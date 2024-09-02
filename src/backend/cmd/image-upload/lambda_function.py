import boto3
import base64
import uuid
import mimetypes

s3 = boto3.client('s3')

def lambda_handler(event, context):
    # Get the base64-encoded image and MIME type from the event
    image_base64 = event['image']
    mime_type = event['mimeType']

    # Decode the image
    image_data = base64.b64decode(image_base64)

    # Determine the file extension based on MIME type
    extension = mimetypes.guess_extension(mime_type)
    if extension is None:
        raise ValueError(f"Unsupported MIME type: {mime_type}")

    # Generate a unique filename with the correct extension
    filename = f"{uuid.uuid4()}{extension}"

    # Specify your S3 bucket name
    bucket_name = 'roadmap-images'

    try:
        # Upload the image to S3
        s3.put_object(Bucket=bucket_name, Key=filename, Body=image_data, ContentType=mime_type)

        # Generate the S3 URL
        s3_url = f"https://{bucket_name}.s3.amazonaws.com/{filename}"

        return {
            'statusCode': 200,
            'body': {
                'message': f'Successfully uploaded {filename} to {bucket_name}',
                'url': s3_url
            }
        }
    except Exception as e:
        return {
            'statusCode': 500,
            'body': {
                'message': f'Error uploading image: {str(e)}',
                'url': None
            }
        }
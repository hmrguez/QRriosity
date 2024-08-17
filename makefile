graph:
	go get github.com/99designs/gqlgen@v0.17.49 && go run github.com/99designs/gqlgen generate

deploy-auth:
	cd src/backend/backend_v2/cmd/auth && \
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go && \
	zip -r auth.zip bootstrap && \
	aws lambda update-function-code --function-name auth-appsync --zip-file fileb://auth.zip

deploy-daily:
	cd src/backend/backend_v2/cmd/daily && \
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go && \
	zip -r daily.zip bootstrap && \
	aws lambda update-function-code --function-name daily-appsync --zip-file fileb://daily.zip

deploy-learning:
	cd src/backend/backend_v2/cmd/learning && \
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go && \
	zip -r learning.zip bootstrap && \
	aws lambda update-function-code --function-name learning-appsync --zip-file fileb://learning.zip
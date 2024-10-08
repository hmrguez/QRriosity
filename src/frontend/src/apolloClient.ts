import {ApolloClient, ApolloLink, HttpLink, InMemoryCache} from '@apollo/client';
import {authLink} from './utils/apolloLink';

const httpLink = new HttpLink({
	uri: 'https://4sxml7ptwfb7xpbyjbu24jbpnu.appsync-api.us-east-2.amazonaws.com/graphql',  // Replace with your AppSync API URL
	headers: {
		'x-api-key': "da2-woozt73tibhh3bglhjlr5flrsy",  // Replace with your AppSync API key if using API key authentication
	},
});

const client = new ApolloClient({
	link: ApolloLink.from([authLink, httpLink]),
	cache: new InMemoryCache({
		typePolicies: {
			Query: {
				fields: {
					getUserByName: {
						merge(existing = {}, incoming) {
							return {...existing, ...incoming};
						}
					}
				}
			}
		}
	}),
	defaultOptions: {
		watchQuery: {
			fetchPolicy: 'no-cache',
		},
		query: {
			fetchPolicy: 'no-cache',
		},
	}
});

export default client;
# go-gqlgen-user-registration
A simple user registration example using gqlgen. 

To generate models run the following 

```
 go run github.com/99designs/gqlgen generate
```

To register a new user in graphql playground

1. execute the app
./go-gqlgen-user-registration
2023/06/10 14:27:53 connect to http://localhost:8080/ for GraphQL playground

or connect with insomnia


```
mutation registerUser {
	registerUser(input: { userId: "1", name: "test"}) {
		id
		name
	}
}
```

You will use the username comeback with a random number between 1 and 100. 

To register a new user in insomnia

Use this as the endpoing, 'query' is the endpoing name. 
http://localhost:8080/query

Then run the following mutation

```
mutation registerUser {
	registerUser(input: { userId: "1", name: "test"}) {
		id
		name
	}
}
```

You will see output like:

```
{
	"data": {
		"registerUser": {
			"id": "1",
			"name": "User26"
		}
	}
}
```

To query a single user

```
query {
	 user(id:1) {
		id,
		name
	}
}
```

To query multiple users

```
query {
	 users {
		id,
		name
	}
}
```




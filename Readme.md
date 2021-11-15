## go-keycloak-auth

Go application for Authentication and authorization with keycloak and JWT - JSON Web Token
The application will be redirect all request (8080) for the keycloak layer (8081)  

**Step 1**: Up docker container
```bash
docker compose up
```

**Step 2**: Run go application
```bash
go mod
go run main.go
```

**Step 3**: Configure the [keycloak](http://localhost:8080/auth/) application in Administration Console:

- Create new Realm: Demo;
- Create new Clients: app with root URL = "http://localhost:8081";
    - dont forget to get the Credential Secret and copy into the application code; 
- Create new User application;


**Step 4**: Login [application](http://localhost:8081/) with the new user

**Step 5**: See all authentication data like access_token, refresh_token and IDTokenin in [JWT](https://jwt.io//) website
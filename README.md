# Simple Movie CRUD API

## Overview
This is a lightweight, single-file application that performs **CRUD (Create, Read, Update, Delete)** operations on movie records. Each movie record includes the following attributes:
- **ID**: Unique identifier for the movie.
- **ISBN**: International Standard Book Number (if applicable).
- **Title**: The title of the movie.
- **Director First Name**: First name of the movie's director.
- **Director Last Name**: Last name of the movie's director.

The application is designed for simplicity, using an in-memory data structure (e.g., a list or array) to store movie records. It is implemented in a single file, making it easy to understand, modify, and maintain. This project is ideal for learning basic CRUD operations or for small-scale use cases without external dependencies.

## Features
- **Create**: Add a new movie record with ID, ISBN, title, and director details.
- **Read**: Retrieve and display all movies or a specific movie by ID.
- **Update**: Modify the details of an existing movie using its ID.
- **Delete**: Remove a movie record from the collection by ID.

#package
- **install**: install the package from the 

```bash
    go get "github.com/gorilla/mux" 

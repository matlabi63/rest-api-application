## mini-cms-api

- REST API Application for managing contents. Written in Go with Echo Framework

## Tech Stack

- Programming language: Go
- Database: MySQL
- Web framework: Echo
- ORM library: GORM
- Request validation: Validator
- Application configuration: Viper

## Notes

- There are two branches in this repository:

- main: REST API application with clean architecture.
- mvc: REST API application with MVC architecture.
- If mvc is chosen, make sure to checkout to the mvc branch with git - checkout mvc

## How to Use

- Clone this repository.

- Copy the .env file for database configuration. Then, configure the database connection inside that file.

- cp .env.example .env
- Create a new database.
- CREATE DATABASE minicmsdb;
- Run the application.
- go run main.go

## Running with Docker

- The application can be run as a Docker container with this command. Make sure to adjust the volume setting inside docker-compose.yml file.

- Run the application.

- docker compose up -d

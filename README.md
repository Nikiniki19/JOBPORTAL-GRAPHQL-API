# Job Portal GraphQL API

A monolithic Job Portal API built with **Golang**, using **GraphQL** for flexible data querying, **GORM** for ORM, and **PostgreSQL** as the database. The project follows **clean architecture** principles and uses **dependency injection**. It allows managing companies and job listings through a single GraphQL endpoint.

## 🚀 Features

- Create and retrieve **Companies**
- Create and retrieve **Jobs**
- View jobs by **Job ID** or **Company ID**
- GraphQL-based API with defined schema
- Monolithic structure (no client/server separation)
- Dependency injection for service layer

## 🧠 GraphQL Operations

### 🔄 Mutations

| Name            | Input Type     | Description          |
|-----------------|----------------|----------------------|
| `createCompany` | `NewCompany`   | Create a new company |
| `createJob`     | `NewJob`       | Create a new job     |

### 🔍 Queries

| Name               | Parameters     | Description                      |
|--------------------|----------------|----------------------------------|
| `viewAllCompanies` | –              | Get all companies                |
| `viewCompanyById`  | `cid: String`  | Get company by ID                |
| `viewAllJobs`      | –              | Get all job listings             |
| `viewJobById`      | `id: String`   | Get job by ID                    |
| `viewJobByCid`     | `cid: String`  | Get jobs posted by a company     |

## 🧪 Tech Stack

- **Go**
- **GraphQL**
- **GORM**
- **PostgreSQL**
- **Clean Architecture**
- **Dependency Injection**

## 📦 Setup Instructions

```bash
# Clone the repo
git clone https://github.com/your-username/jobportal-graphql-api.git
cd jobportal-graphql-api

# Set up environment variables (.env)
cp .env.example .env
# Update .env with your DB credentials

# Run the server
go run main.go

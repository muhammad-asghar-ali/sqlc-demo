### SQLC

generates type-safe code from SQL.

sqlc inspired by PugSQL and HugSQL.

SQLC is code generation tool for Go programming language that allows developers to write SQL queries in a separate file and generate corresponding Go code.

The generated Go code is type-safe, which means it can be checked at compile-time, helping to prevent runtime errors that can occur with dynamic queries.

Database used is postgresql

Demo using managed database which require a sqlc Cloud project and auth token and link is [https://dashboard.sqlc.dev](https://dashboard.sqlc.dev)

### Project Steps:

1. install sqlc
2. create go project
3. create `sqlc.yml` or `sqlc.json`
4. define tables in SQL
5. create quires on the defiend schema
6. run sqlc command to generate the code command: `sqlc generate`

### Why SQLC over RDMBs in GO

- Database/SQL (built-in library)

  1. Works fast (+)
  2. Have to manually map SQL fields to variables (-)
  3. Error-prone as it is not type-safe (-)

- GORM

  1. Queries are already written, takes very short code to execute (+)
  2. Performance overhead (-)
  3. Limited customization (-)

- SQLC
  1. Be precise while not spending too much time (+)
  2. Almost as fast as built-in SQL library (+)
  3. Type-safe queries (+)
  4. Limited database support, for now only supports PostgreSQL, mySQL and SQLite (-)

Install golang migrate tool
```bash 
brew install golang-migrate
```

to run the migration, use the following command
```bash
migrate create -ext sql -dir db/migration -seq init_schema
```
This will create a new migration file in the `db/migration` directory with the name `init_schema.up.sql` and `init_schema.down.sql`. You can then add your SQL statements to these files to create and drop your database schema respectively.

to create sqlc.yaml file, use the following command
```bash
sqlc init
```
This will create a new `sqlc.yaml` file in the root directory of your project.

we are installing postgres driver for testing purpose, you can use any other driver as per xwyour requirement
```bash
go get github.com/lib/pq
```
this will install the postgres driver for your project. You can then use this driver to connect to your postgres database and run your migrations.

ACID properties in database transactions refer to the following:
- Atomicity: This property ensures that a transaction is treated as a single unit of work. If any part of the transaction fails, the entire transaction is rolled back, and the database remains unchanged
- Consistency: This property ensures that a transaction brings the database from one valid state to another valid state. It ensures that any data written to the database must be valid according to all defined rules, including constraints, cascades, and triggers.
- Isolation: This property ensures that concurrent transactions do not interfere with each other. Each transaction should operate as if it is the only transaction in the system, even if other transactions are running concurrently.
- Durability: This property ensures that once a transaction has been committed, it will remain so, even in the event of a system failure. The changes made by the transaction are permanently stored in the database and will not be lost.



Database locking
video 
<!-- add link -->
https://www.youtube.com/watch?v=G2aggv_3Bbg&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=8
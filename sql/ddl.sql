create schema s3explorer
    create table user_book (
        id serial primary key,
        name varchar(50) not null unique,
        last_login timestamp default current_timestamp
    )
    create table token (
        id serial primary key,
        user_id integer references user_book(id),
        token varchar(100)
    );

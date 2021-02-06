CREATE TABLE IF NOT EXISTS users(
    username TEXT NOT NULL,
    password TEXT NOT NULL
);

INSERT INTO users (username,password) 
    VALUES
        ('cesar','cfabrica46'),
        ('arturo','01234');

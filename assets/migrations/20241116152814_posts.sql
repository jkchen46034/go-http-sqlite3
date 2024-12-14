-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  createdAT DATETIME NOT NULL DEFAULT current_timestamp
);

INSERT INTO posts(title, content) values("travel", "moutain biking is a lot of fun");
INSERT INTO posts(title, content) values("travel", "hiking is also a lot of fun");
INSERT INTO posts(title, content) values("gardening", "growing vegetable is a lot of fun");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd

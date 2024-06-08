## ディレクトリ構成
- 参考
  - https://qiita.com/AQUA651/items/7279a4f1f9c977336879

## mysql Config (直書きやめたい)
-User
  - root
- PassWord
  - 'P4ssW0rd!'

## テーブル構成
mysql> SHOW COLUMNS FROM tasks FROM todo;
+------------+-----------------+------+-----+---------+----------------+
| Field      | Type            | Null | Key | Default | Extra          |
+------------+-----------------+------+-----+---------+----------------+
| id         | bigint unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | datetime(3)     | YES  |     | NULL    |                |
| updated_at | datetime(3)     | YES  |     | NULL    |                |
| deleted_at | datetime(3)     | YES  | MUL | NULL    |                |
| title      | longtext        | YES  |     | NULL    |                |
| priority   | longtext        | YES  |     | NULL    |                |
| status     | longtext        | YES  |     | NULL    |                |
+------------+-----------------+------+-----+---------+----------------+

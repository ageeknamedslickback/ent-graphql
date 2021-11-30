BEGIN;
CREATE TABLE `todos`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `text` varchar(255) NOT NULL, `created_at` datetime NOT NULL, `status` varchar(255) NOT NULL DEFAULT 'in_progress', `prority` integer NOT NULL DEFAULT 0, `todo_parent` integer NULL, FOREIGN KEY(`todo_parent`) REFERENCES `todos`(`id`) ON DELETE SET NULL);
COMMIT;

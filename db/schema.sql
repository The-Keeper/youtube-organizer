CREATE TABLE IF NOT EXISTS video (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    channel_id TEXT,
    FOREIGN KEY(channel_id) REFERENCES channel(id)

);

CREATE TABLE  IF NOT EXISTS channel (
  id    TEXT PRIMARY KEY AUTOINCREMENT, 
  title  TEXT
);

CREATE TABLE  IF NOT EXISTS file (
  id    TEXT PRIMARY KEY AUTOINCREMENT, 
  path  TEXT
);

CREATE TABLE  IF NOT EXISTS video_file_map (
    id INTEGER PRIMARY KEY,
    video_id TEXT NOT NULL,
    file_id TEXT NOT NULL,
    FOREIGN KEY(video_id) REFERENCES video(id),
    FOREIGN KEY(file_id) REFERENCES file(id)
)

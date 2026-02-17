# Lum

Open source metadata-driven book manager.

## Core principles

1. **Low hardware requirements**

I want to run the server on my router.

2. **No directory layout requirement**

I want to manage files by their metadata.

3. **All files are read-only**

I don't want to deal with perms or accidental destruction of data.

4. **Use well-known APIs.**

OPDS, Kavita, Stump, whatever makes me avoid re-invent a front-end.

## Roadmap

### Version 0.1

- [ ] Add database
- [ ] Add EPUB metadata parser
- [ ] Add CLI to return metadata
- [ ] Add filesystem watcher
- [ ] Add webserver
- [ ] Add OPDS
- [ ] Add auth

### Version 0.2

- [ ] Add [libmobi](https://github.com/bfabiszewski/libmobi/) support

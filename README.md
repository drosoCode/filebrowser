# Filebrowser onlyoffice fork
[![Build](https://github.com/drosocode/filebrowser/actions/workflows/main.yaml/badge.svg)](https://github.com/drosocode/filebrowser/actions/workflows/main.yaml)

## About this fork

This is a fork of the original [filebrowser repo](https://github.com/filebrowser/filebrowser) to apply [this pull request](https://github.com/filebrowser/filebrowser/pull/1420) that adds an onlyoffice integration.

I will try to keep this fork up to date with the original repo, but please bear in mind that the releases provided here are not as well tested as the ones of the upstream project and may conatains some bugs.

## About filebrowser

filebrowser provides a file managing interface within a specified directory and it can be used to upload, delete, preview, rename and edit your files. It allows the creation of multiple users and each user can have its own directory. It can be used as a standalone app.

## Documentation

For other features not related to this fork, please refer to the original [repo](https://github.com/filebrowser/filebrowser) and [docs](https://filebrowser.org)

## Install

Installation with docker is recommanded, all the images are available in the [github container registry](https://github.com/drosoCode/filebrowser/pkgs/container/filebrowser).

You can use the following example to create your `docker-compose.yml` (using traefik as reverse proxy).

Then login as admin into filebrowser, go in the general settings, enable onlyoffile and set the url to `https://filebrowser.domain.tld/onlyoffice`.

```yaml
version: '3.7'
services:
  filebrowser:
    image: ghcr.io/drosocode/filebrowser:v2
    restart: unless-stopped
    volumes:
      - /path/to/root:/srv
      - /path/filebrowser.db:/database.db
      - /path/.filebrowser.json:/.filebrowser.json
    labels:
        - traefik.http.routers.filebrowser.rule=Host(`filebrowser.domain.tld`)
        - traefik.http.routers.filebrowser.tls=true
        - traefik.http.routers.filebrowser.entrypoints=https
        - traefik.http.routers.filebrowser.service=filebrowser_service
        - traefik.http.services.filebrowser_service.loadbalancer.server.port=80
        - traefik.enable=true

  onlyoffice:
    image: onlyoffice/documentserver
    restart: unless-stopped
    labels:
        - traefik.http.routers.filebrowser_office.rule=Host(`filebrowser.domain.tld`) && PathPrefix(`/onlyoffice`)
        - traefik.http.routers.filebrowser_office.tls=true
        - traefik.http.routers.filebrowser_office.entrypoints=https
        - traefik.http.routers.filebrowser_office.middlewares=filebrowser_office_headers, filebrowser_office_strip
        - traefik.http.middlewares.filebrowser_office_strip.stripprefix.prefixes=/onlyoffice
        - traefik.http.middlewares.filebrowser_office_headers.headers.customrequestheaders.X-Forwarded-Proto=https
        - traefik.http.middlewares.filebrowser_office_headers.headers.customrequestheaders.X-Forwarded-Host=filebrowser.domain.tld/onlyoffice
        - traefik.http.routers.filebrowser_office.service=filebrowser_office_service
        - traefik.http.services.filebrowser_office_service.loadbalancer.server.port=80
        - traefik.enable=true
```
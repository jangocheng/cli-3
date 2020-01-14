package configs

// BackendOnly with Go, Nginx, Certbot (docker-compose.yml)
var BackendOnly string = (`# docker-compose.yml by Create Go App Authors
version: "3.7"

services:
  backend:
    container_name: backend
    build:
      context: ./backend
    networks:
      - cgapp_net
    restart: always

  nginx:
    container_name: nginx
    image: nginx:alpine
    networks:
      - cgapp_net
    volumes:
      - ./nginx/configs/nginx.conf:/etc/nginx/nginx.conf
      - ./certbot/conf:/etc/letsencrypt
      - ./certbot/www:/var/www/certbot
    ports:
      - 80:80
      - 443:443
    restart: unless-stopped
    command: '/bin/sh -c ''while :; do sleep 6h & wait $${!}; nginx -s reload; done & nginx -g "daemon off;"'''
    depends_on:
      - backend

  certbot:
    container_name: certbot
    image: certbot/certbot
    networks:
      - cgapp_net
    volumes:
      - ./certbot/conf:/etc/letsencrypt
      - ./certbot/www:/var/www/certbot
    restart: unless-stopped
    entrypoint: "/bin/sh -c 'trap exit TERM; while :; do certbot renew; sleep 12h & wait $${!}; done;'"

networks:
  cgapp_net:
    name: cgapp_net
`)

// FullStackApp with Go, Node.js, Nginx, Certbot (docker-compose.yml)
var FullStackApp string = (`# docker-compose.yml by Create Go App Authors
version: "3.7"

services:
  frontend:
    container_name: frontend
    build:
      context: ./frontend
    volumes:
      - static:/frontend/build

  backend:
    container_name: backend
    build:
      context: ./backend
    volumes:
      - static:/frontend/build
    networks:
      - cgapp_net
    restart: always
    depends_on:
      - frontend

  nginx:
    container_name: nginx
    image: nginx:alpine
    networks:
      - cgapp_net
    volumes:
      - ./nginx/configs/nginx.conf:/etc/nginx/nginx.conf
      - ./certbot/conf:/etc/letsencrypt
      - ./certbot/www:/var/www/certbot
    ports:
      - 80:80
      - 443:443
    restart: unless-stopped
    command: '/bin/sh -c ''while :; do sleep 6h & wait $${!}; nginx -s reload; done & nginx -g "daemon off;"'''
    depends_on:
      - backend

  certbot:
    container_name: certbot
    image: certbot/certbot
    networks:
      - cgapp_net
    volumes:
      - ./certbot/conf:/etc/letsencrypt
      - ./certbot/www:/var/www/certbot
    restart: unless-stopped
    entrypoint: "/bin/sh -c 'trap exit TERM; while :; do certbot renew; sleep 12h & wait $${!}; done;'"

networks:
  cgapp_net:
    name: cgapp_net

volumes:
  static:
`)

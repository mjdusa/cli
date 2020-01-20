package configs

// BackendService with Go, Nginx, Certbot, Postgres (docker-compose.yml)
var BackendService string = (`# docker-compose.yml by Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)
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
            - ./webserver/configs/default.template.conf:/etc/nginx/conf.d/default.template.conf
            - ./webserver/configs/nginx.conf:/etc/nginx/nginx.conf
            - ./webserver/certbot/conf:/etc/letsencrypt
            - ./webserver/certbot/www:/var/www/certbot
        environment:
            - APP_DOMAIN=localhost
        ports:
            - 80:80
        restart: unless-stopped
        command: /bin/sh -c "envsubst < /etc/nginx/conf.d/default.template.conf > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'"
        depends_on:
            - backend

    certbot:
        container_name: certbot
        image: certbot/certbot
        networks:
            - cgapp_net
        volumes:
            - ./webserver/certbot/conf:/etc/letsencrypt
            - ./webserver/certbot/www:/var/www/certbot
        restart: unless-stopped
        entrypoint: /bin/sh -c "trap exit TERM; while :; do certbot renew; sleep 12h & wait $${!}; done;"
        depends_on:
            - nginx

networks:
    cgapp_net:
        name: cgapp_net
`)

// NginxProdService Nginx config for production (docker-compose.prod.yml)
var NginxProdService string = (`# docker-compose.prod.yml by Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)
version: "3.7"

services:
    nginx:
        environment:
            - APP_DOMAIN=example.com
        ports:
            - 80:80
            - 443:443
`)

// FullstackService docker-compose with Go, Node.js, Nginx, Certbot, Postgres (docker-compose.yml)
var FullstackService string = (`# docker-compose.yml by Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)
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

    nginx:
        container_name: nginx
        image: nginx:alpine
        networks:
            - cgapp_net
        volumes:
            - ./webserver/configs/default.template.conf:/etc/nginx/conf.d/default.template.conf
            - ./webserver/configs/nginx.conf:/etc/nginx/nginx.conf
            - ./webserver/certbot/conf:/etc/letsencrypt
            - ./webserver/certbot/www:/var/www/certbot
        environment:
            - APP_DOMAIN=example.com
        ports:
            - 80:80
        restart: unless-stopped
        command: /bin/sh -c "envsubst < /etc/nginx/conf.d/default.template.conf > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'"
        depends_on:
            - backend

    certbot:
        container_name: certbot
        image: certbot/certbot
        networks:
            - cgapp_net
        volumes:
            - ./webserver/certbot/conf:/etc/letsencrypt
            - ./webserver/certbot/www:/var/www/certbot
		restart: unless-stopped
        entrypoint: /bin/sh -c "trap exit TERM; while :; do certbot renew; sleep 12h & wait $${!}; done;"
        depends_on:
            - nginx

networks:
    cgapp_net:
        name: cgapp_net

volumes:
    static:
`)
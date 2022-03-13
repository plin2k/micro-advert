

build:
	docker build -t plin2k/micro-advert .

prune:
	docker image prune -f

run:
	docker run -p 8080:8080 --name micro-advert --rm --env-file .env plin2k/micro-advert

run-dev:
	docker run -p 8080:8080 --name micro-advert --rm --env-file .env --volume $(PWD)/resources:/app/resources:ro --volume $(PWD)/templates:/app/templates:ro plin2k/micro-advert

stop:
	docker stop micro-advert


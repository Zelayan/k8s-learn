docker-release:
	docker build . -t chenzezeya/hellok8s:v3
	docker push chenzezeya/hellok8s:v3

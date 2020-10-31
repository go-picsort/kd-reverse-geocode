main:
	GOOS=linux GOARCH=amd64       go build -v -o main.sample


.PHONY: clean run

clean:
	go clean
	chmod -R ug+w ../go/
	rm -fr ../go/*
	rm -fv go.sum
	rm -fv *.upx
	find . -name "main.*" -type f -executable -delete


final: main
	ls --color=always -lh main.*
	strip -v -s main.sample
	ls --color=always -lh main.*
	upx --best --color --ultra-brute main.sample
	ls --color=always -lh main.*

run_only:
	./main.sample

run: main run_only
final_run: final run

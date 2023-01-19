module example.com/practice

go 1.19

replace example.com/backend => ../backend

require (
	example.com/backend v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	github.com/joho/godotenv v1.4.0 // indirect
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

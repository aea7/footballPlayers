## Fetch football players for certain teams

To run the application:

go run main.go
-

data/teams.go  ->  array team names (Change to get other players in distinct teams)

models/Response.go  ->  to get response into struct
models/Player.go  ->  to add players and print them

main.go  ->  all logic is here, it is possible to have multiple files and separate some logic.

Another thing that could be done is adding some unit/functional tests to test the endpoint, sorting the player names etc.


# Outcome Checker

This program provides a function (`Solution`) to determine the winner of a game based on a given string representing the game's outcome. The game involves two players, 'X' and 'O,' and the function calculates the winner or declares a tie based on the provided outcomes.


## Usage

```shell
$ go run main.go
```

## Tests

To run the tests, use the following command:

```shell
$ go test -v

go test -v
=== RUN   TestSolution
=== RUN   TestSolution/Player_X_wins_most_of_the_time
=== RUN   TestSolution/Player_O_wins_most_of_the_time
=== RUN   TestSolution/Player_X_wins_most_of_time_in_a_row
=== RUN   TestSolution/Player_O_wins_most_of_time_in_a_row
=== RUN   TestSolution/Both_player_wins_same_number
=== RUN   TestSolution/Both_player_wins_same_consecutive_time_in_a_row
=== RUN   TestSolution/Accept_only_valid_characters
--- PASS: TestSolution (0.00s)
    --- PASS: TestSolution/Player_X_wins_most_of_the_time (0.00s)
    --- PASS: TestSolution/Player_O_wins_most_of_the_time (0.00s)
    --- PASS: TestSolution/Player_X_wins_most_of_time_in_a_row (0.00s)
    --- PASS: TestSolution/Player_O_wins_most_of_time_in_a_row (0.00s)
    --- PASS: TestSolution/Both_player_wins_same_number (0.00s)
    --- PASS: TestSolution/Both_player_wins_same_consecutive_time_in_a_row (0.00s)
    --- PASS: TestSolution/Accept_only_valid_characters (0.00s)
PASS
ok  	app	0.215s
```
package main

import (
	"./data"
	"./models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

func main() {

	teams := data.GetTeams()

	visitedTeams := map[string]bool {}

	for _, elem := range teams {
		visitedTeams[elem] = false
	}

	var wg sync.WaitGroup

	limit := 50
	counter := 0
	players := getTeamsFromBackend(&counter, limit, &wg, visitedTeams)

	keys := make([]string, 0, len(players))
	for k := range players {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, *players[k])
	}


}

func getTeamsFromBackend(counter *int, limit int, wg *sync.WaitGroup, visitedTeams map[string]bool) map[string]*models.Player {
	players := map[string]*models.Player {}
	i := 0
	if limit != 50{
		i = limit/2
	}

	for i < limit {
		i++
		wg.Add(1)
		go callEndpoint(players, counter, visitedTeams, wg, i)
	}
	wg.Wait()
	if *counter != len(visitedTeams){
		getTeamsFromBackend(counter, limit * 2, wg, visitedTeams)
	}
	return players
}

func callEndpoint (players map[string]*models.Player, counter *int, visitedTeams map[string]bool, wg *sync.WaitGroup, next int) {
	defer wg.Done()
	res, err := http.Get("https://vintagemonster.onefootball.com/api/teams/en/" + strconv.Itoa(next) + ".json")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var response = models.Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	teamName := response.Data.Team.Name

	if _, exists := visitedTeams[teamName]; exists {
		*counter += 1

		for _, player := range response.Data.Team.Players{
			playerName := player.Name
			playerAge := player.Age

			if _, exists := players[playerName]; exists {
				players[playerName].Teams = append(players[playerName].Teams, teamName)
			}else {
				players[playerName] = &models.Player{Age: playerAge, Teams: []string{teamName}}
			}
		}

	}

}

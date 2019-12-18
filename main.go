package gamebattles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Player info API
// http://profile.majorleaguegaming.com/Tydra_/
// http://profile.majorleaguegaming.com/api/profile-page-data/Tydra_

// Player stores the information about a player.
type Player struct {
	ID       uint
	UserID   uint
	Username string
	Gamertag string
	Active   bool
}

// Team info API
// https://gamebattles.majorleaguegaming.com/pc/overwatch/team/33834248

// GetTeam gets the players on a team.
func GetTeam(id string) ([]Player, error) {
	r, err := http.Get("https://gb-api.majorleaguegaming.com/api/web/v1/team-members-extended/team/" + id)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	team := struct {
		Errors []struct {
			Code string
		}
		Body []struct {
			TeamMember Player
		}
	}{}
	err = json.Unmarshal(b, &team)
	if err != nil {
		return nil, err
	} else if len(team.Errors) > 0 {
		return nil, fmt.Errorf("error getting team %q: %s", id, team.Errors[0].Code)
	}

	players := []Player{}
	for _, p := range team.Body {
		players = append(players, p.TeamMember)
	}

	return players, nil
}

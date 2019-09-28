package models

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   struct {
		Team struct {
			ID       int    `json:"id"`
			OptaID   int    `json:"optaId"`
			Name     string `json:"name"`
			LogoUrls []struct {
				Size string `json:"size"`
				URL  string `json:"url"`
			} `json:"logoUrls"`
			IsNational   bool `json:"isNational"`
			Competitions []struct {
				CompetitionID   int    `json:"competitionId"`
				CompetitionName string `json:"competitionName"`
			} `json:"competitions"`
			Players []struct {
				ID           string `json:"id"`
				Country      string `json:"country"`
				FirstName    string `json:"firstName"`
				LastName     string `json:"lastName"`
				Name         string `json:"name"`
				Position     string `json:"position"`
				Number       int    `json:"number"`
				BirthDate    string `json:"birthDate"`
				Age          string `json:"age"`
				Height       int    `json:"height"`
				Weight       int    `json:"weight"`
				ThumbnailSrc string `json:"thumbnailSrc"`
			} `json:"players"`
			Officials []struct {
				CountryName string `json:"countryName"`
				ID          string `json:"id"`
				FirstName   string `json:"firstName"`
				LastName    string `json:"lastName"`
				Country     string `json:"country"`
				Position    string `json:"position"`
			} `json:"officials"`
			Colors struct {
				ShirtColorHome string `json:"shirtColorHome"`
				ShirtColorAway string `json:"shirtColorAway"`
				CrestMainColor string `json:"crestMainColor"`
				MainColor      string `json:"mainColor"`
			} `json:"colors"`
		} `json:"team"`
	} `json:"data"`
	Message string `json:"message"`
}

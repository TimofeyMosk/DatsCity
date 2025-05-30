// Package dto provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package dto

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// GameRoundListResponse defines model for game.RoundListResponse.
type GameRoundListResponse struct {
	EventId *string              `json:"eventId,omitempty"`
	Now     *string              `json:"now,omitempty"`
	Rounds  *[]GameRoundResponse `json:"rounds,omitempty"`
}

// GameRoundResponse defines model for game.RoundResponse.
type GameRoundResponse struct {
	Duration *int    `json:"duration,omitempty"`
	EndAt    *string `json:"endAt,omitempty"`
	Name     *string `json:"name,omitempty"`
	Repeat   *int    `json:"repeat,omitempty"`
	StartAt  *string `json:"startAt,omitempty"`
	Status   *string `json:"status,omitempty"`
}

// GamesdkPublicError defines model for gamesdk.PublicError.
type GamesdkPublicError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ModelDoneTowerResponse defines model for model.DoneTowerResponse.
type ModelDoneTowerResponse struct {
	Id    *int     `json:"id,omitempty"`
	Score *float32 `json:"score,omitempty"`
}

// ModelPlayerBuildRequest defines model for model.PlayerBuildRequest.
type ModelPlayerBuildRequest struct {
	Done  *bool                    `json:"done,omitempty"`
	Words *[]ModelTowerWordRequest `json:"words,omitempty"`
}

// ModelPlayerExtendedWordsResponse defines model for model.PlayerExtendedWordsResponse.
type ModelPlayerExtendedWordsResponse struct {
	MapSize     *[]int    `json:"mapSize,omitempty"`
	NextTurnSec *int      `json:"nextTurnSec,omitempty"`
	RoundEndsAt *string   `json:"roundEndsAt,omitempty"`
	ShuffleLeft *int      `json:"shuffleLeft,omitempty"`
	Turn        *int      `json:"turn,omitempty"`
	UsedIndexes *[]int    `json:"usedIndexes,omitempty"`
	Words       *[]string `json:"words,omitempty"`
}

// ModelPlayerResponse defines model for model.PlayerResponse.
type ModelPlayerResponse struct {
	DoneTowers *[]ModelDoneTowerResponse `json:"doneTowers,omitempty"`
	Score      *float32                  `json:"score,omitempty"`
	Tower      *ModelPlayerTowerResponse `json:"tower,omitempty"`
}

// ModelPlayerTowerResponse defines model for model.PlayerTowerResponse.
type ModelPlayerTowerResponse struct {
	Score *float32           `json:"score,omitempty"`
	Words *[]ModelPlayerWord `json:"words,omitempty"`
}

// ModelPlayerWord defines model for model.PlayerWord.
type ModelPlayerWord struct {
	Dir  *int    `json:"dir,omitempty"`
	Pos  *[]int  `json:"pos,omitempty"`
	Text *string `json:"text,omitempty"`
}

// ModelPlayerWordsResponse defines model for model.PlayerWordsResponse.
type ModelPlayerWordsResponse struct {
	ShuffleLeft *int      `json:"shuffleLeft,omitempty"`
	Words       *[]string `json:"words,omitempty"`
}

// ModelTowerWordRequest defines model for model.TowerWordRequest.
type ModelTowerWordRequest struct {
	Dir *int   `json:"dir,omitempty"`
	Id  *int   `json:"id,omitempty"`
	Pos *[]int `json:"pos,omitempty"`
}

// PostApiBuildJSONRequestBody defines body for PostApiBuild for application/json ContentType.
type PostApiBuildJSONRequestBody = ModelPlayerBuildRequest

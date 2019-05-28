package houndify // import github.com/soundhound/houndify-sdk-go/houndify

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ParseWrittenResponse will take final server response JSON (as a string)
// and parse out the human readable text to be displayed or spoken the end user.
// If the string is invalid JSON, the server had an error, or there was nothing
// to reply with, an error is returned.
func ParseWrittenResponse(serverResponseJSON string) (string, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(serverResponseJSON), &result)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("failed to decode json")
	}
	if !strings.EqualFold(result["Status"].(string), "OK") {
		return "", errors.New(result["ErrorMessage"].(string))
	}
	if result["NumToReturn"].(float64) < 1 {
		return "", errors.New("no results to return")
	}
	return result["AllResults"].([]interface{})[0].(map[string]interface{})["WrittenResponseLong"].(string), nil
}

func parseConversationState(serverResponseJSON string) (interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(serverResponseJSON), &result)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("failed to decode json")
	}
	if !strings.EqualFold(result["Status"].(string), "OK") {
		return nil, errors.New(result["ErrorMessage"].(string))
	}
	if result["NumToReturn"].(float64) < 1 {
		return nil, errors.New("no results to return")
	}
	return result["AllResults"].([]interface{})[0].(map[string]interface{})["ConversationState"], nil
}
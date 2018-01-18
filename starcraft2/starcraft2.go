/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 10:35:54
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-17 19:56:33
 */

// Package starcraft2 is a client library to use Blizzard Starcraft 2 API calls.
package starcraft2

import (
	"strconv"

	"github.com/FuzzyStatic/blizzard"
)

// Starcraft2 regional API URLs, locale, access token, api key
type Starcraft2 struct {
	Auth         blizzard.Auth
	Locale       string
	DataURL      string
	CommunityURL string
}

// New create new Starcraft2 structure
func New(auth blizzard.Auth, region blizzard.Region) *Starcraft2 {
	var s = Starcraft2{Auth: auth}

	switch region {
	case blizzard.EU:
		s.Locale = "en_GB"
		s.DataURL = "https://eu.api.battle.net/sc2/data"
		s.CommunityURL = "https://eu.api.battle.net/sc2"
	case blizzard.KR:
		s.Locale = "ko_KR"
		s.DataURL = "https://kr.api.battle.net/sc2/data"
		s.CommunityURL = "https://kr.api.battle.net/sc2"
	case blizzard.SEA:
		s.Locale = "en_US"
		s.DataURL = "https://sea.api.battle.net/sc2/data"
		s.CommunityURL = "https://sea.api.battle.net/sc2"
	case blizzard.TW:
		s.Locale = "zh_TW"
		s.DataURL = "https://tw.api.battle.net/sc2/data"
		s.CommunityURL = "https://tw.api.battle.net/sc2"
	case blizzard.US:
		s.Locale = "en_US"
		s.DataURL = "https://us.api.battle.net/sc2/data"
		s.CommunityURL = "https://us.api.battle.net/sc2"
	default: // USA! USA!
		s.Locale = "en_US"
		s.DataURL = "https://us.api.battle.net/sc2/data"
		s.CommunityURL = "https://us.api.battle.net/sc2"
	}

	return &s
}

// GetProfileJSON gets profile JSON information
func (s *Starcraft2) GetProfileJSON(id, region int, name string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = s.CommunityURL + profilePath + "/" + strconv.Itoa(id) + "/" + strconv.Itoa(region) +
		"/" + name + "/?" + localeQuery + s.Locale + "&" + apiKeyQuery + s.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetProfile puts profile info into Profile structure
func (s *Starcraft2) GetProfile(id, region int, name string) (*Profile, error) {
	var (
		profile Profile
		json    *[]byte
		err     error
	)

	json, err = s.GetProfileJSON(id, region, name)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// GetProfileLaddersJSON gets profile ladders JSON information
func (s *Starcraft2) GetProfileLaddersJSON(id, region int, name string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = s.CommunityURL + profilePath + "/" + strconv.Itoa(id) + "/" + strconv.Itoa(region) +
		"/" + name + "/" + laddersPath + "?" + localeQuery + s.Locale + "&" + apiKeyQuery +
		s.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetProfileLadders puts profile ladders info into ProfileLadders structure
func (s *Starcraft2) GetProfileLadders(id, region int, name string) (*ProfileLadders, error) {
	var (
		profileLadders ProfileLadders
		json           *[]byte
		err            error
	)

	json, err = s.GetProfileLaddersJSON(id, region, name)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &profileLadders)
	if err != nil {
		return nil, err
	}

	return &profileLadders, nil
}

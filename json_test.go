package gpjson

import (
	"os"

	"testing"
)

func Test_json(t *testing.T) {

	rawjson, err := os.Open("test_json.json")
	if err != nil {
		t.Fatalf("Error getting json: %s", err)
	}

	j, err := NewJsonFromReader(rawjson)
	if err != nil {
		t.Fatalf("Error converting json: %s", err)
	}

	//t.Logf("Got JSON: %# v", pretty.Formatter(j))

	kid, err := j.Get("killID").Int64()
	if err != nil {
		t.Errorf("Error reading killid integer: %s", err)
	}
	if kid != 38863346 {
		t.Errorf("got kill id %d but expected 38863346", kid)
	}

	topDamageName, err := j.Get("attackers").Idx(0).Get("character").Get("name").String()
	if err != nil {
		t.Errorf("Error reading name: %s", err)
	}
	if topDamageName != "Hellcharm Bloodsmoke" {
		t.Errorf("got %s as top damage, expected 'Hellcharm Bloodsmoke'", topDamageName)
	}

	secondDamageName, err := j.Get("attackers").Idx(1).Get("character").Get("name").String()
	if err != nil {
		t.Fatalf("Error reading name: %s", err)
	}
	if secondDamageName != "Paladin Fett" {
		t.Errorf("got %s as second damage, expected 'Paladin Fett'", secondDamageName)
	}

	victimName, err := j.Get("victim").Get("character").Get("name").String()
	if err != nil {
		t.Errorf("error reading victim name: %s", err)
	}
	if victimName != "gh0stryd3r" {
		t.Errorf("unexpected victim name, expected 'gh0stryd3r'")
	}

	_, err = j.Get("victim").Idx(0).String()
	if err == nil {
		t.Errorf("expected error, got none")
	}

	_, err = j.Get("attackers").Idx(-1).String()
	if err == nil {
		t.Errorf("expected error, got none")
	}

	_, err = j.Get("attackers").Idx(43534).String()
	if err == nil {
		t.Errorf("expected error, got none")
	}

	_, err = j.Get("attackers").Get("buttes").String()
	if err == nil {
		t.Errorf("expected error, got none")
	}

	_, err = j.Get("buttes").String()
	if err == nil {
		t.Errorf("expected error, got none")
	}
}

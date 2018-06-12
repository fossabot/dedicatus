package models

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Personality struct {
	KGID          string
	CanonicalName string
	Nickname      []string
}

func (p Personality) ToString() string {
	return fmt.Sprintf("kg:%s %s (%s)", p.KGID, p.CanonicalName, strings.Join(p.Nickname, ", "))
}

func CreatePersonality(ctx context.Context, KGID string, name string) (*Personality, error) {
	key := datastore.NewIncompleteKey(ctx, personalityEntityKind, nil)

	p := &Personality{
		KGID:          KGID,
		CanonicalName: name,
		Nickname:      []string{name},
	}

	_, err := datastore.Put(ctx, key, p)
	return p, err
}

func GetPersonality(ctx context.Context, key *datastore.Key) (*Personality, error) {
	p := new(Personality)
	err := datastore.Get(ctx, key, p)
	return p, err
}

func GetPersonalityByName(ctx context.Context, name string) (*datastore.Key, *Personality, error) {
	var ps []*Personality
	keys, err := datastore.NewQuery(personalityEntityKind).Filter("CanonicalName = ", name).Limit(1).GetAll(ctx, &ps)
	if len(keys) == 1 {
		return keys[0], ps[0], nil
	}
	return nil, nil, err
}

func GetPersonalityByKGID(ctx context.Context, KGID string) (*Personality, error) {
	var ps []*Personality
	keys, err := datastore.NewQuery(personalityEntityKind).Filter("KGID = ", KGID).Limit(1).GetAll(ctx, &ps)
	if len(keys) == 1 {
		return ps[0], nil
	}
	return nil, err
}

func TryFindPersonality(ctx context.Context, query string) (*datastore.Key, error) {
	keys, err := datastore.NewQuery(personalityEntityKind).Filter("Nickname = ", query).Limit(1).KeysOnly().GetAll(ctx, nil)
	if len(keys) == 1 {
		return keys[0], nil
	}
	return nil, err
}

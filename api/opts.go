package api

import (
	"github.com/dghubble/sling"
)

type Opt func(s *sling.Sling) *sling.Sling

func Authorize(token string) func(s *sling.Sling) *sling.Sling {
	return func(s *sling.Sling) *sling.Sling {
		return s.Set("Authorization", "bearer "+token)
	}
}

func Query(q interface{}) func(s *sling.Sling) *sling.Sling {
	return func(s *sling.Sling) *sling.Sling {
		return s.QueryStruct(q)
	}
}
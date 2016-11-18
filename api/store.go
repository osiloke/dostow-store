package api

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type StoreService struct {
	sling *sling.Sling
}

func newStoreService(sling *sling.Sling) *StoreService {
	return &StoreService{
		sling: sling.Path("store"),
	}
}

func (s *StoreService) List(store string, params ...interface{}) (*json.RawMessage, *http.Response, error) {
	var rows *json.RawMessage = &json.RawMessage{}
	apiError := new(APIError)
	_s := s.sling.New().Path("store/" + store)
	for _, params := range params {
		_s.QueryStruct(params)
	}
	resp, err := _s.Receive(rows, apiError)
	return rows, resp, relevantError(err, apiError)
}
func (s *StoreService) Get(store, id string, data interface{}) error {
	apiError := new(APIError)
	_s := s.sling.New().Get("store/" + store + "/" + id)
	_, err := _s.Receive(data, apiError)
	return relevantError(err, apiError)
}
func (s *StoreService) Create(store string, data interface{}) (*json.RawMessage, error) {
	var result *json.RawMessage = &json.RawMessage{}
	apiError := new(APIError)
	_s := s.sling.New().Post("store/" + store).BodyJSON(data)
	_, err := _s.Receive(result, apiError)
	return result, relevantError(err, apiError)
}
func (s *StoreService) Update(store, id string, data interface{}) (*json.RawMessage, error) {
	var result *json.RawMessage = &json.RawMessage{}
	apiError := new(APIError)
	_s := s.sling.New().Put("store/" + store + "/" + id).BodyJSON(data)
	_, err := _s.Receive(result, apiError)
	return result, relevantError(err, apiError)
}
func (s *StoreService) Remove(store, id string) (*json.RawMessage, error) {
	var result *json.RawMessage = &json.RawMessage{}
	apiError := new(APIError)
	_s := s.sling.New().Delete("store/" + store + "/" + id)
	_, err := _s.Receive(result, apiError)
	return result, relevantError(err, apiError)
}
package semrush

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
)

type LocationsServiceOp struct {
	client *Client
}

type LocationsService interface {
	GetLocations(context.Context, BrowseLocation) (*LocationsResponse, error)
	GetLocation(context.Context, ReadLocation) (*LocationResponse, error)
	UpdateLocations(context.Context, EditLocations) (*LocationsResponse, error)
	UpdateLocation(context.Context, EditLocation) (*LocationResponse, error)
}

func (s *LocationsServiceOp) GetLocation(ctx context.Context, req ReadLocation) (*LocationResponse, error) {
	s.client.TokenAccess(ctx)

	var reqResponse []byte

	purchasingURL := url + `listing-management/v1/external/locations/` + req.LocationID

	errRequest := s.client.Request("GET", purchasingURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response LocationResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *LocationsServiceOp) UpdateLocation(ctx context.Context, req EditLocation) (*LocationResponse, error) {
	s.client.TokenAccess(ctx)

	var reqResponse []byte

	productURL := url + `listing-management/v1/external/locations/` + req.LocationID

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("PUT", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response LocationResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *LocationsServiceOp) GetLocations(ctx context.Context, req BrowseLocation) (*LocationsResponse, error) {
	s.client.TokenAccess(ctx)

	var reqResponse []byte
	urlBuild := []string{}

	if req.Page != nil {
		urlBuild = append(urlBuild, "page="+*req.Page)
	}

	if req.Size != nil {
		urlBuild = append(urlBuild, "size="+*req.Size)
	}

	purchasingURL := url + `listing-management/v1/external/locations?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", purchasingURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response LocationsResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *LocationsServiceOp) UpdateLocations(ctx context.Context, req EditLocations) (*LocationsResponse, error) {
	s.client.TokenAccess(ctx)

	var reqResponse []byte

	productURL := url + `listing-management/v1/external/locations`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("PUT", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response LocationsResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

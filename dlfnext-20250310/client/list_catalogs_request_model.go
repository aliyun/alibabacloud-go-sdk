// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogNamePattern(v string) *ListCatalogsRequest
	GetCatalogNamePattern() *string
	SetMaxResults(v int32) *ListCatalogsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListCatalogsRequest
	GetPageToken() *string
}

type ListCatalogsRequest struct {
	// The catalog name pattern.
	//
	// example:
	//
	// mi
	CatalogNamePattern *string `json:"catalogNamePattern,omitempty" xml:"catalogNamePattern,omitempty"`
	// The maximum number of records to retrieve at a time.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token used to retrieve the next page of results. If the response does not include a token, pass an empty string ("") or an empty character (\\"\\").
	//
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogsRequest) GetCatalogNamePattern() *string {
	return s.CatalogNamePattern
}

func (s *ListCatalogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCatalogsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListCatalogsRequest) SetCatalogNamePattern(v string) *ListCatalogsRequest {
	s.CatalogNamePattern = &v
	return s
}

func (s *ListCatalogsRequest) SetMaxResults(v int32) *ListCatalogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCatalogsRequest) SetPageToken(v string) *ListCatalogsRequest {
	s.PageToken = &v
	return s
}

func (s *ListCatalogsRequest) Validate() error {
	return dara.Validate(s)
}

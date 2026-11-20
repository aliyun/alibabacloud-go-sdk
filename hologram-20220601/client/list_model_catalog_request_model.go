// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListModelCatalogRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListModelCatalogRequest
	GetNextToken() *string
	SetProvider(v string) *ListModelCatalogRequest
	GetProvider() *string
}

type ListModelCatalogRequest struct {
	// The maximum number of entries per page.
	//
	// Default value:
	//
	// 	20
	//
	// Maximum value:
	//
	// 	100
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Set this parameter to the value of NextToken returned from the previous API call.
	//
	// example:
	//
	// uat-assembly-cut-3d-bbig
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The model provider.
	//
	// example:
	//
	// holo_bailian
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListModelCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelCatalogRequest) GoString() string {
	return s.String()
}

func (s *ListModelCatalogRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelCatalogRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelCatalogRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListModelCatalogRequest) SetMaxResults(v int32) *ListModelCatalogRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelCatalogRequest) SetNextToken(v string) *ListModelCatalogRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelCatalogRequest) SetProvider(v string) *ListModelCatalogRequest {
	s.Provider = &v
	return s
}

func (s *ListModelCatalogRequest) Validate() error {
	return dara.Validate(s)
}

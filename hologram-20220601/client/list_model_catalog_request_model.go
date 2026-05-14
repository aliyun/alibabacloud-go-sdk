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
}

type ListModelCatalogRequest struct {
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// uat-assembly-cut-3d-bbig
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

func (s *ListModelCatalogRequest) SetMaxResults(v int32) *ListModelCatalogRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelCatalogRequest) SetNextToken(v string) *ListModelCatalogRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelCatalogRequest) Validate() error {
	return dara.Validate(s)
}

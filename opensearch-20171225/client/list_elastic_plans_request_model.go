// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListElasticPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *ListElasticPlansRequest
	GetEnabled() *bool
	SetMaxResults(v int32) *ListElasticPlansRequest
	GetMaxResults() *int32
	SetName(v string) *ListElasticPlansRequest
	GetName() *string
	SetNextToken(v string) *ListElasticPlansRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListElasticPlansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListElasticPlansRequest
	GetPageSize() *int32
}

type ListElasticPlansRequest struct {
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// "test"
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 20
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListElasticPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListElasticPlansRequest) GoString() string {
	return s.String()
}

func (s *ListElasticPlansRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListElasticPlansRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListElasticPlansRequest) GetName() *string {
	return s.Name
}

func (s *ListElasticPlansRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListElasticPlansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListElasticPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListElasticPlansRequest) SetEnabled(v bool) *ListElasticPlansRequest {
	s.Enabled = &v
	return s
}

func (s *ListElasticPlansRequest) SetMaxResults(v int32) *ListElasticPlansRequest {
	s.MaxResults = &v
	return s
}

func (s *ListElasticPlansRequest) SetName(v string) *ListElasticPlansRequest {
	s.Name = &v
	return s
}

func (s *ListElasticPlansRequest) SetNextToken(v string) *ListElasticPlansRequest {
	s.NextToken = &v
	return s
}

func (s *ListElasticPlansRequest) SetPageNumber(v int32) *ListElasticPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *ListElasticPlansRequest) SetPageSize(v int32) *ListElasticPlansRequest {
	s.PageSize = &v
	return s
}

func (s *ListElasticPlansRequest) Validate() error {
	return dara.Validate(s)
}

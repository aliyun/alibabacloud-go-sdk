// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateTTSVoicesCustomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPrivateTTSVoicesCustomRequest
	GetMaxResults() *int32
	SetName(v string) *ListPrivateTTSVoicesCustomRequest
	GetName() *string
	SetNextToken(v string) *ListPrivateTTSVoicesCustomRequest
	GetNextToken() *string
	SetPageIndex(v int32) *ListPrivateTTSVoicesCustomRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListPrivateTTSVoicesCustomRequest
	GetPageSize() *int32
}

type ListPrivateTTSVoicesCustomRequest struct {
	// example:
	//
	// 100
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Q45algIClks=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPrivateTTSVoicesCustomRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateTTSVoicesCustomRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateTTSVoicesCustomRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrivateTTSVoicesCustomRequest) GetName() *string {
	return s.Name
}

func (s *ListPrivateTTSVoicesCustomRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrivateTTSVoicesCustomRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListPrivateTTSVoicesCustomRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPrivateTTSVoicesCustomRequest) SetMaxResults(v int32) *ListPrivateTTSVoicesCustomRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomRequest) SetName(v string) *ListPrivateTTSVoicesCustomRequest {
	s.Name = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomRequest) SetNextToken(v string) *ListPrivateTTSVoicesCustomRequest {
	s.NextToken = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomRequest) SetPageIndex(v int32) *ListPrivateTTSVoicesCustomRequest {
	s.PageIndex = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomRequest) SetPageSize(v int32) *ListPrivateTTSVoicesCustomRequest {
	s.PageSize = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomRequest) Validate() error {
	return dara.Validate(s)
}

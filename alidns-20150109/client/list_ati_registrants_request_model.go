// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiRegistrantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListAtiRegistrantsRequest
	GetClientToken() *string
	SetMaxResults(v int32) *ListAtiRegistrantsRequest
	GetMaxResults() *int32
	SetName(v string) *ListAtiRegistrantsRequest
	GetName() *string
	SetNextToken(v string) *ListAtiRegistrantsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListAtiRegistrantsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAtiRegistrantsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListAtiRegistrantsRequest
	GetStatus() *string
}

type ListAtiRegistrantsRequest struct {
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 张xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 审核通过
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAtiRegistrantsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAtiRegistrantsRequest) GoString() string {
	return s.String()
}

func (s *ListAtiRegistrantsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListAtiRegistrantsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAtiRegistrantsRequest) GetName() *string {
	return s.Name
}

func (s *ListAtiRegistrantsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAtiRegistrantsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAtiRegistrantsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAtiRegistrantsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAtiRegistrantsRequest) SetClientToken(v string) *ListAtiRegistrantsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAtiRegistrantsRequest) SetMaxResults(v int32) *ListAtiRegistrantsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAtiRegistrantsRequest) SetName(v string) *ListAtiRegistrantsRequest {
	s.Name = &v
	return s
}

func (s *ListAtiRegistrantsRequest) SetNextToken(v string) *ListAtiRegistrantsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAtiRegistrantsRequest) SetPageNumber(v int32) *ListAtiRegistrantsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAtiRegistrantsRequest) SetPageSize(v int32) *ListAtiRegistrantsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAtiRegistrantsRequest) SetStatus(v string) *ListAtiRegistrantsRequest {
	s.Status = &v
	return s
}

func (s *ListAtiRegistrantsRequest) Validate() error {
	return dara.Validate(s)
}

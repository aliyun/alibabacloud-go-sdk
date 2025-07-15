// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnassignedNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUnassignedNumbersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUnassignedNumbersRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListUnassignedNumbersRequest
	GetSearchPattern() *string
}

type ListUnassignedNumbersRequest struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0833
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListUnassignedNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUnassignedNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUnassignedNumbersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnassignedNumbersRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListUnassignedNumbersRequest) SetPageNumber(v int32) *ListUnassignedNumbersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUnassignedNumbersRequest) SetPageSize(v int32) *ListUnassignedNumbersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUnassignedNumbersRequest) SetSearchPattern(v string) *ListUnassignedNumbersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListUnassignedNumbersRequest) Validate() error {
	return dara.Validate(s)
}

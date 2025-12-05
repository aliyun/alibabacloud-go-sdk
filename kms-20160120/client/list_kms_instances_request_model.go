// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKmsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v string) *ListKmsInstancesRequest
	GetFilters() *string
	SetPageNumber(v int32) *ListKmsInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKmsInstancesRequest
	GetPageSize() *int32
}

type ListKmsInstancesRequest struct {
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListKmsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKmsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesRequest) GetFilters() *string {
	return s.Filters
}

func (s *ListKmsInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKmsInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKmsInstancesRequest) SetFilters(v string) *ListKmsInstancesRequest {
	s.Filters = &v
	return s
}

func (s *ListKmsInstancesRequest) SetPageNumber(v int32) *ListKmsInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKmsInstancesRequest) SetPageSize(v int32) *ListKmsInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListKmsInstancesRequest) Validate() error {
	return dara.Validate(s)
}

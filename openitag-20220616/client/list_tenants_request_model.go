// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTenantsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTenantsRequest
	GetPageSize() *int32
}

type ListTenantsRequest struct {
	// Page number of the tenant list. The starting value is 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of tenants displayed per page in a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTenantsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTenantsRequest) GoString() string {
	return s.String()
}

func (s *ListTenantsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTenantsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTenantsRequest) SetPageNumber(v int32) *ListTenantsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTenantsRequest) SetPageSize(v int32) *ListTenantsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTenantsRequest) Validate() error {
	return dara.Validate(s)
}

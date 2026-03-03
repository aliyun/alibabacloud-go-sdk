// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderBy(v string) *ListPdpImageRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpImageRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpImageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpImageRequest
	GetPageSize() *int32
	SetServiceGroupId(v int64) *ListPdpImageRequest
	GetServiceGroupId() *int64
}

type ListPdpImageRequest struct {
	// example:
	//
	// gmtCreated
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// desc
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s ListPdpImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpImageRequest) GoString() string {
	return s.String()
}

func (s *ListPdpImageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpImageRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpImageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpImageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpImageRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListPdpImageRequest) SetOrderBy(v string) *ListPdpImageRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpImageRequest) SetOrderDirection(v string) *ListPdpImageRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpImageRequest) SetPageNumber(v int32) *ListPdpImageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpImageRequest) SetPageSize(v int32) *ListPdpImageRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpImageRequest) SetServiceGroupId(v int64) *ListPdpImageRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListPdpImageRequest) Validate() error {
	return dara.Validate(s)
}

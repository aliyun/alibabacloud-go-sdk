// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderBy(v string) *ListPdpConfigsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpConfigsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpConfigsRequest
	GetPageSize() *int32
	SetServiceGroupId(v int64) *ListPdpConfigsRequest
	GetServiceGroupId() *int64
	SetType(v string) *ListPdpConfigsRequest
	GetType() *string
}

type ListPdpConfigsRequest struct {
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPdpConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListPdpConfigsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpConfigsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpConfigsRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListPdpConfigsRequest) GetType() *string {
	return s.Type
}

func (s *ListPdpConfigsRequest) SetOrderBy(v string) *ListPdpConfigsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpConfigsRequest) SetOrderDirection(v string) *ListPdpConfigsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpConfigsRequest) SetPageNumber(v int32) *ListPdpConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpConfigsRequest) SetPageSize(v int32) *ListPdpConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpConfigsRequest) SetServiceGroupId(v int64) *ListPdpConfigsRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListPdpConfigsRequest) SetType(v string) *ListPdpConfigsRequest {
	s.Type = &v
	return s
}

func (s *ListPdpConfigsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMicroServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderBy(v string) *ListMicroServiceRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMicroServiceRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMicroServiceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMicroServiceRequest
	GetPageSize() *int32
	SetPbcId(v int64) *ListMicroServiceRequest
	GetPbcId() *int64
	SetServiceIds(v string) *ListMicroServiceRequest
	GetServiceIds() *string
}

type ListMicroServiceRequest struct {
	// example:
	//
	// gmtCreated
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// desc
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// 1
	ServiceIds *string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty"`
}

func (s ListMicroServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMicroServiceRequest) GoString() string {
	return s.String()
}

func (s *ListMicroServiceRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMicroServiceRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMicroServiceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMicroServiceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMicroServiceRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListMicroServiceRequest) GetServiceIds() *string {
	return s.ServiceIds
}

func (s *ListMicroServiceRequest) SetOrderBy(v string) *ListMicroServiceRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMicroServiceRequest) SetOrderDirection(v string) *ListMicroServiceRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMicroServiceRequest) SetPageNumber(v int32) *ListMicroServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMicroServiceRequest) SetPageSize(v int32) *ListMicroServiceRequest {
	s.PageSize = &v
	return s
}

func (s *ListMicroServiceRequest) SetPbcId(v int64) *ListMicroServiceRequest {
	s.PbcId = &v
	return s
}

func (s *ListMicroServiceRequest) SetServiceIds(v string) *ListMicroServiceRequest {
	s.ServiceIds = &v
	return s
}

func (s *ListMicroServiceRequest) Validate() error {
	return dara.Validate(s)
}

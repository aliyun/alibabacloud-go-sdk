// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRuntimeTokensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ListRuntimeTokensRequest
	GetEnterpriseId() *int64
	SetOrderBy(v string) *ListRuntimeTokensRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListRuntimeTokensRequest
	GetOrderDirection() *string
	SetPageNumber(v string) *ListRuntimeTokensRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListRuntimeTokensRequest
	GetPageSize() *string
	SetPbcId(v int64) *ListRuntimeTokensRequest
	GetPbcId() *int64
	SetServiceGroupId(v int64) *ListRuntimeTokensRequest
	GetServiceGroupId() *int64
}

type ListRuntimeTokensRequest struct {
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
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
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// 1
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s ListRuntimeTokensRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRuntimeTokensRequest) GoString() string {
	return s.String()
}

func (s *ListRuntimeTokensRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListRuntimeTokensRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListRuntimeTokensRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListRuntimeTokensRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListRuntimeTokensRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListRuntimeTokensRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListRuntimeTokensRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListRuntimeTokensRequest) SetEnterpriseId(v int64) *ListRuntimeTokensRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListRuntimeTokensRequest) SetOrderBy(v string) *ListRuntimeTokensRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRuntimeTokensRequest) SetOrderDirection(v string) *ListRuntimeTokensRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListRuntimeTokensRequest) SetPageNumber(v string) *ListRuntimeTokensRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRuntimeTokensRequest) SetPageSize(v string) *ListRuntimeTokensRequest {
	s.PageSize = &v
	return s
}

func (s *ListRuntimeTokensRequest) SetPbcId(v int64) *ListRuntimeTokensRequest {
	s.PbcId = &v
	return s
}

func (s *ListRuntimeTokensRequest) SetServiceGroupId(v int64) *ListRuntimeTokensRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListRuntimeTokensRequest) Validate() error {
	return dara.Validate(s)
}

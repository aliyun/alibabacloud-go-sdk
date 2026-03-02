// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpLanesForServiceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListPdpLanesForServiceGroupShrinkRequest
	GetCompanyId() *int64
	SetEnv(v string) *ListPdpLanesForServiceGroupShrinkRequest
	GetEnv() *string
	SetKeyword(v string) *ListPdpLanesForServiceGroupShrinkRequest
	GetKeyword() *string
	SetLaneIdsShrink(v string) *ListPdpLanesForServiceGroupShrinkRequest
	GetLaneIdsShrink() *string
	SetOperator(v string) *ListPdpLanesForServiceGroupShrinkRequest
	GetOperator() *string
	SetOrderBy(v string) *ListPdpLanesForServiceGroupShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpLanesForServiceGroupShrinkRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpLanesForServiceGroupShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpLanesForServiceGroupShrinkRequest
	GetPageSize() *int32
	SetServiceGroupId(v int64) *ListPdpLanesForServiceGroupShrinkRequest
	GetServiceGroupId() *int64
	SetServiceId(v int64) *ListPdpLanesForServiceGroupShrinkRequest
	GetServiceId() *int64
}

type ListPdpLanesForServiceGroupShrinkRequest struct {
	// example:
	//
	// 40
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// yunmall
	Keyword       *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LaneIdsShrink *string `json:"laneIds,omitempty" xml:"laneIds,omitempty"`
	// example:
	//
	// filter
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// gmtModified
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 199
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// example:
	//
	// 100
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s ListPdpLanesForServiceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpLanesForServiceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetLaneIdsShrink() *string {
	return s.LaneIdsShrink
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetOperator() *string {
	return s.Operator
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetCompanyId(v int64) *ListPdpLanesForServiceGroupShrinkRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetEnv(v string) *ListPdpLanesForServiceGroupShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetKeyword(v string) *ListPdpLanesForServiceGroupShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetLaneIdsShrink(v string) *ListPdpLanesForServiceGroupShrinkRequest {
	s.LaneIdsShrink = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetOperator(v string) *ListPdpLanesForServiceGroupShrinkRequest {
	s.Operator = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetOrderBy(v string) *ListPdpLanesForServiceGroupShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetOrderDirection(v string) *ListPdpLanesForServiceGroupShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetPageNumber(v int32) *ListPdpLanesForServiceGroupShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetPageSize(v int32) *ListPdpLanesForServiceGroupShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetServiceGroupId(v int64) *ListPdpLanesForServiceGroupShrinkRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) SetServiceId(v int64) *ListPdpLanesForServiceGroupShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *ListPdpLanesForServiceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

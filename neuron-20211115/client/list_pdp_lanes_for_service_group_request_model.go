// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpLanesForServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListPdpLanesForServiceGroupRequest
	GetCompanyId() *int64
	SetEnv(v string) *ListPdpLanesForServiceGroupRequest
	GetEnv() *string
	SetKeyword(v string) *ListPdpLanesForServiceGroupRequest
	GetKeyword() *string
	SetLaneIds(v []*int64) *ListPdpLanesForServiceGroupRequest
	GetLaneIds() []*int64
	SetOperator(v string) *ListPdpLanesForServiceGroupRequest
	GetOperator() *string
	SetOrderBy(v string) *ListPdpLanesForServiceGroupRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpLanesForServiceGroupRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpLanesForServiceGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpLanesForServiceGroupRequest
	GetPageSize() *int32
	SetServiceGroupId(v int64) *ListPdpLanesForServiceGroupRequest
	GetServiceGroupId() *int64
	SetServiceId(v int64) *ListPdpLanesForServiceGroupRequest
	GetServiceId() *int64
}

type ListPdpLanesForServiceGroupRequest struct {
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
	Keyword *string  `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LaneIds []*int64 `json:"laneIds,omitempty" xml:"laneIds,omitempty" type:"Repeated"`
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

func (s ListPdpLanesForServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpLanesForServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListPdpLanesForServiceGroupRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPdpLanesForServiceGroupRequest) GetEnv() *string {
	return s.Env
}

func (s *ListPdpLanesForServiceGroupRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListPdpLanesForServiceGroupRequest) GetLaneIds() []*int64 {
	return s.LaneIds
}

func (s *ListPdpLanesForServiceGroupRequest) GetOperator() *string {
	return s.Operator
}

func (s *ListPdpLanesForServiceGroupRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpLanesForServiceGroupRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpLanesForServiceGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpLanesForServiceGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpLanesForServiceGroupRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListPdpLanesForServiceGroupRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *ListPdpLanesForServiceGroupRequest) SetCompanyId(v int64) *ListPdpLanesForServiceGroupRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetEnv(v string) *ListPdpLanesForServiceGroupRequest {
	s.Env = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetKeyword(v string) *ListPdpLanesForServiceGroupRequest {
	s.Keyword = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetLaneIds(v []*int64) *ListPdpLanesForServiceGroupRequest {
	s.LaneIds = v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetOperator(v string) *ListPdpLanesForServiceGroupRequest {
	s.Operator = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetOrderBy(v string) *ListPdpLanesForServiceGroupRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetOrderDirection(v string) *ListPdpLanesForServiceGroupRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetPageNumber(v int32) *ListPdpLanesForServiceGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetPageSize(v int32) *ListPdpLanesForServiceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetServiceGroupId(v int64) *ListPdpLanesForServiceGroupRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) SetServiceId(v int64) *ListPdpLanesForServiceGroupRequest {
	s.ServiceId = &v
	return s
}

func (s *ListPdpLanesForServiceGroupRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentServiceStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetAgentServiceStatusShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetAgentServiceStatusShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetAgentServiceStatusShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetAgentServiceStatusShrinkRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetAgentServiceStatusShrinkRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetAgentServiceStatusShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetInstanceId(v string) *GetAgentServiceStatusShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentServiceStatusShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetAgentServiceStatusShrinkRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetAgentServiceStatusShrinkRequest
	GetTimeLatitudeType() *string
}

type GetAgentServiceStatusShrinkRequest struct {
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// example:
	//
	// 1
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// true
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// example:
	//
	// true
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetAgentServiceStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentServiceStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetAgentServiceStatusShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentServiceStatusShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetAgentServiceStatusShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAgentServiceStatusShrinkRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetAgentServiceStatusShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetAgentServiceStatusShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentServiceStatusShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentServiceStatusShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAgentServiceStatusShrinkRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetAgentServiceStatusShrinkRequest) SetAgentIdsShrink(v string) *GetAgentServiceStatusShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetCurrentPage(v int32) *GetAgentServiceStatusShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetDepIdsShrink(v string) *GetAgentServiceStatusShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetEndDate(v int64) *GetAgentServiceStatusShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetExistAgentGrouping(v bool) *GetAgentServiceStatusShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetExistDepartmentGrouping(v bool) *GetAgentServiceStatusShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetInstanceId(v string) *GetAgentServiceStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetPageSize(v int32) *GetAgentServiceStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetStartDate(v int64) *GetAgentServiceStatusShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetTimeLatitudeType(v string) *GetAgentServiceStatusShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}

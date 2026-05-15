// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupStatusTotalShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetSkillGroupStatusTotalShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetSkillGroupStatusTotalShrinkRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetSkillGroupStatusTotalShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupStatusTotalShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupStatusTotalShrinkRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetSkillGroupStatusTotalShrinkRequest
	GetTimeLatitudeType() *string
}

type GetSkillGroupStatusTotalShrinkRequest struct {
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
	// false
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// example:
	//
	// fasle
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// example:
	//
	// fasle
	ExistSkillGroupGrouping *bool   `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	GroupIdsShrink          *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
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

func (s GetSkillGroupStatusTotalShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupStatusTotalShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupStatusTotalShrinkRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetAgentIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupStatusTotalShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetEndDate(v int64) *GetSkillGroupStatusTotalShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetExistAgentGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetInstanceId(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetPageSize(v int32) *GetSkillGroupStatusTotalShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetStartDate(v int64) *GetSkillGroupStatusTotalShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetTimeLatitudeType(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupServiceStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetSkillGroupServiceStatusShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetSkillGroupServiceStatusShrinkRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest
	GetExistAgentGrouping() *bool
	SetExistChannelInstanceGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest
	GetExistChannelInstanceGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistRobotInstanceGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest
	GetExistRobotInstanceGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetSkillGroupServiceStatusShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupServiceStatusShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupServiceStatusShrinkRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetSkillGroupServiceStatusShrinkRequest
	GetTimeLatitudeType() *string
}

type GetSkillGroupServiceStatusShrinkRequest struct {
	// List of agent IDs.
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// The current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// End date UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Specifies whether to query by skill group. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// Specifies whether to query by Channel instance group. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistChannelInstanceGrouping *bool `json:"ExistChannelInstanceGrouping,omitempty" xml:"ExistChannelInstanceGrouping,omitempty"`
	// Specifies whether to query by department group. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// Specifies whether to query by robot instance group. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistRobotInstanceGrouping *bool `json:"ExistRobotInstanceGrouping,omitempty" xml:"ExistRobotInstanceGrouping,omitempty"`
	// Specifies whether to query by skill group. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	// List of skill group IDs.
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// AICCS instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page size. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start date UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Time latitude type. Valid values:
	//
	// - **minute**: Minute.
	//
	// - **hour**: Hour.
	//
	// - **day**: Day.
	//
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetSkillGroupServiceStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetExistChannelInstanceGrouping() *bool {
	return s.ExistChannelInstanceGrouping
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetExistRobotInstanceGrouping() *bool {
	return s.ExistRobotInstanceGrouping
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupServiceStatusShrinkRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetAgentIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupServiceStatusShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetEndDate(v int64) *GetSkillGroupServiceStatusShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistAgentGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistChannelInstanceGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistChannelInstanceGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistRobotInstanceGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistRobotInstanceGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetInstanceId(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetPageSize(v int32) *GetSkillGroupServiceStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetStartDate(v int64) *GetSkillGroupServiceStatusShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetTimeLatitudeType(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}

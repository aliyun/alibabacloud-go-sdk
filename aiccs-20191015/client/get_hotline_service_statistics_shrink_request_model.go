// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineServiceStatisticsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetHotlineServiceStatisticsShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetHotlineServiceStatisticsShrinkRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetHotlineServiceStatisticsShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetHotlineServiceStatisticsShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetHotlineServiceStatisticsShrinkRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetHotlineServiceStatisticsShrinkRequest
	GetTimeLatitudeType() *string
}

type GetHotlineServiceStatisticsShrinkRequest struct {
	// List of agent IDs.
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// Current page. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// End Datetime Variable as a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Specifies whether to query by agent group. Valid values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// true
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// Specifies whether to query by department group. Valid values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// true
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// Indicates whether to query by skill group grouping. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
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
	// Time latitude type for grouped queries. Valid values:
	//
	// - **minute**: minute
	//
	// - **hour**: hour
	//
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetHotlineServiceStatisticsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineServiceStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetHotlineServiceStatisticsShrinkRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetAgentIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetCurrentPage(v int32) *GetHotlineServiceStatisticsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetDepIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetEndDate(v int64) *GetHotlineServiceStatisticsShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetExistAgentGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetExistDepartmentGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetGroupIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetInstanceId(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetPageSize(v int32) *GetHotlineServiceStatisticsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetStartDate(v int64) *GetHotlineServiceStatisticsShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetTimeLatitudeType(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

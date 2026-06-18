// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupAndAgentStatusSummaryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupAndAgentStatusSummaryShrinkRequest
	GetStartDate() *int64
}

type GetSkillGroupAndAgentStatusSummaryShrinkRequest struct {
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
	// Specifies whether to query by department grouping. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// Specifies whether to query by skill group grouping. Default value: **false**. Valid values:
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
	// You can obtain it in the **Instance Management*	- section of the left-side navigation pane in the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start date UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetSkillGroupAndAgentStatusSummaryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetEndDate(v int64) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetInstanceId(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetPageSize(v int32) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetStartDate(v int64) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) Validate() error {
	return dara.Validate(s)
}

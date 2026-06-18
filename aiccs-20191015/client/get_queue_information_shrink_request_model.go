// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueInformationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetQueueInformationShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetQueueInformationShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetQueueInformationShrinkRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetQueueInformationShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetQueueInformationShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetQueueInformationShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetQueueInformationShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetQueueInformationShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetQueueInformationShrinkRequest
	GetStartDate() *int64
}

type GetQueueInformationShrinkRequest struct {
	// The current page. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// End UNIX timestamp. Unit: milliseconds.
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
	// Specifies whether to query data grouped by skill group. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	// The list of skill group IDs.
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetQueueInformationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueInformationShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetQueueInformationShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetQueueInformationShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetQueueInformationShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetQueueInformationShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetQueueInformationShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetQueueInformationShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetQueueInformationShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueueInformationShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueueInformationShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetQueueInformationShrinkRequest) SetCurrentPage(v int32) *GetQueueInformationShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetDepIdsShrink(v string) *GetQueueInformationShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetEndDate(v int64) *GetQueueInformationShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetExistDepartmentGrouping(v bool) *GetQueueInformationShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetQueueInformationShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetGroupIdsShrink(v string) *GetQueueInformationShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetInstanceId(v string) *GetQueueInformationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetPageSize(v int32) *GetQueueInformationShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetStartDate(v int64) *GetQueueInformationShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) Validate() error {
	return dara.Validate(s)
}

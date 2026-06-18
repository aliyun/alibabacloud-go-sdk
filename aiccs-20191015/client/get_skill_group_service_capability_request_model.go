// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupServiceCapabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSkillGroupServiceCapabilityRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetSkillGroupServiceCapabilityRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetSkillGroupServiceCapabilityRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceCapabilityRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceCapabilityRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetSkillGroupServiceCapabilityRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetSkillGroupServiceCapabilityRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupServiceCapabilityRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupServiceCapabilityRequest
	GetStartDate() *int64
}

type GetSkillGroupServiceCapabilityRequest struct {
	// Current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
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
	// true
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// Specifies whether to query by skill group grouping. Default value: **false**. Valid values:
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
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
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
	// Start UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetSkillGroupServiceCapabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupServiceCapabilityRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetSkillGroupServiceCapabilityRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupServiceCapabilityRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupServiceCapabilityRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupServiceCapabilityRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetSkillGroupServiceCapabilityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupServiceCapabilityRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupServiceCapabilityRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupServiceCapabilityRequest) SetCurrentPage(v int32) *GetSkillGroupServiceCapabilityRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetDepIds(v []*int64) *GetSkillGroupServiceCapabilityRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetEndDate(v int64) *GetSkillGroupServiceCapabilityRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceCapabilityRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceCapabilityRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetGroupIds(v []*int64) *GetSkillGroupServiceCapabilityRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetInstanceId(v string) *GetSkillGroupServiceCapabilityRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetPageSize(v int32) *GetSkillGroupServiceCapabilityRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetStartDate(v int64) *GetSkillGroupServiceCapabilityRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) Validate() error {
	return dara.Validate(s)
}

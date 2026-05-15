// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupServiceCapabilityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSkillGroupServiceCapabilityShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetSkillGroupServiceCapabilityShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetSkillGroupServiceCapabilityShrinkRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceCapabilityShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceCapabilityShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetSkillGroupServiceCapabilityShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetSkillGroupServiceCapabilityShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupServiceCapabilityShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupServiceCapabilityShrinkRequest
	GetStartDate() *int64
}

type GetSkillGroupServiceCapabilityShrinkRequest struct {
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
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// example:
	//
	// true
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetSkillGroupServiceCapabilityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetEndDate(v int64) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetInstanceId(v string) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetPageSize(v int32) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetStartDate(v int64) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) Validate() error {
	return dara.Validate(s)
}

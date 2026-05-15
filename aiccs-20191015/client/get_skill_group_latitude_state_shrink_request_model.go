// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupLatitudeStateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSkillGroupLatitudeStateShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetSkillGroupLatitudeStateShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetSkillGroupLatitudeStateShrinkRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSkillGroupLatitudeStateShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupLatitudeStateShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetSkillGroupLatitudeStateShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetSkillGroupLatitudeStateShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupLatitudeStateShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupLatitudeStateShrinkRequest
	GetStartDate() *int64
}

type GetSkillGroupLatitudeStateShrinkRequest struct {
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetSkillGroupLatitudeStateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupLatitudeStateShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupLatitudeStateShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupLatitudeStateShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetEndDate(v int64) *GetSkillGroupLatitudeStateShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupLatitudeStateShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupLatitudeStateShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupLatitudeStateShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetInstanceId(v string) *GetSkillGroupLatitudeStateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetPageSize(v int32) *GetSkillGroupLatitudeStateShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetStartDate(v int64) *GetSkillGroupLatitudeStateShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) Validate() error {
	return dara.Validate(s)
}

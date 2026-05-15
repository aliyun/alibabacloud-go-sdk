// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupLatitudeStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSkillGroupLatitudeStateRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetSkillGroupLatitudeStateRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetSkillGroupLatitudeStateRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSkillGroupLatitudeStateRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupLatitudeStateRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetSkillGroupLatitudeStateRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetSkillGroupLatitudeStateRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupLatitudeStateRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupLatitudeStateRequest
	GetStartDate() *int64
}

type GetSkillGroupLatitudeStateRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIds      []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
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
	ExistSkillGroupGrouping *bool    `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	GroupIds                []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
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

func (s GetSkillGroupLatitudeStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupLatitudeStateRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupLatitudeStateRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetSkillGroupLatitudeStateRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupLatitudeStateRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupLatitudeStateRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupLatitudeStateRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetSkillGroupLatitudeStateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupLatitudeStateRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupLatitudeStateRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupLatitudeStateRequest) SetCurrentPage(v int32) *GetSkillGroupLatitudeStateRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetDepIds(v []*int64) *GetSkillGroupLatitudeStateRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetEndDate(v int64) *GetSkillGroupLatitudeStateRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupLatitudeStateRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupLatitudeStateRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetGroupIds(v []*int64) *GetSkillGroupLatitudeStateRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetInstanceId(v string) *GetSkillGroupLatitudeStateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetPageSize(v int32) *GetSkillGroupLatitudeStateRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetStartDate(v int64) *GetSkillGroupLatitudeStateRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) Validate() error {
	return dara.Validate(s)
}

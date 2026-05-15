// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupAndAgentStatusSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSkillGroupAndAgentStatusSummaryRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetSkillGroupAndAgentStatusSummaryRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetSkillGroupAndAgentStatusSummaryRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetSkillGroupAndAgentStatusSummaryRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetSkillGroupAndAgentStatusSummaryRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupAndAgentStatusSummaryRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupAndAgentStatusSummaryRequest
	GetStartDate() *int64
}

type GetSkillGroupAndAgentStatusSummaryRequest struct {
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
	// false
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// example:
	//
	// false
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

func (s GetSkillGroupAndAgentStatusSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetCurrentPage(v int32) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetDepIds(v []*int64) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetEndDate(v int64) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetGroupIds(v []*int64) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetInstanceId(v string) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetPageSize(v int32) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetStartDate(v int64) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) Validate() error {
	return dara.Validate(s)
}

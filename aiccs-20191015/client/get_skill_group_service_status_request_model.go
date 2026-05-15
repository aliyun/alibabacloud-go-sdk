// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetSkillGroupServiceStatusRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetSkillGroupServiceStatusRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetSkillGroupServiceStatusRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetSkillGroupServiceStatusRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetSkillGroupServiceStatusRequest
	GetExistAgentGrouping() *bool
	SetExistChannelInstanceGrouping(v bool) *GetSkillGroupServiceStatusRequest
	GetExistChannelInstanceGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceStatusRequest
	GetExistDepartmentGrouping() *bool
	SetExistRobotInstanceGrouping(v bool) *GetSkillGroupServiceStatusRequest
	GetExistRobotInstanceGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceStatusRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetSkillGroupServiceStatusRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetSkillGroupServiceStatusRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupServiceStatusRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupServiceStatusRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetSkillGroupServiceStatusRequest
	GetTimeLatitudeType() *string
}

type GetSkillGroupServiceStatusRequest struct {
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
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
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// example:
	//
	// false
	ExistChannelInstanceGrouping *bool `json:"ExistChannelInstanceGrouping,omitempty" xml:"ExistChannelInstanceGrouping,omitempty"`
	// example:
	//
	// false
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// example:
	//
	// false
	ExistRobotInstanceGrouping *bool `json:"ExistRobotInstanceGrouping,omitempty" xml:"ExistRobotInstanceGrouping,omitempty"`
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
	// 10
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

func (s GetSkillGroupServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetSkillGroupServiceStatusRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupServiceStatusRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetSkillGroupServiceStatusRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupServiceStatusRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetSkillGroupServiceStatusRequest) GetExistChannelInstanceGrouping() *bool {
	return s.ExistChannelInstanceGrouping
}

func (s *GetSkillGroupServiceStatusRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupServiceStatusRequest) GetExistRobotInstanceGrouping() *bool {
	return s.ExistRobotInstanceGrouping
}

func (s *GetSkillGroupServiceStatusRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupServiceStatusRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetSkillGroupServiceStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupServiceStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupServiceStatusRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupServiceStatusRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetSkillGroupServiceStatusRequest) SetAgentIds(v []*int64) *GetSkillGroupServiceStatusRequest {
	s.AgentIds = v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetCurrentPage(v int32) *GetSkillGroupServiceStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetDepIds(v []*int64) *GetSkillGroupServiceStatusRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetEndDate(v int64) *GetSkillGroupServiceStatusRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistAgentGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistChannelInstanceGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistChannelInstanceGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistRobotInstanceGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistRobotInstanceGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetGroupIds(v []*int64) *GetSkillGroupServiceStatusRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetInstanceId(v string) *GetSkillGroupServiceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetPageSize(v int32) *GetSkillGroupServiceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetStartDate(v int64) *GetSkillGroupServiceStatusRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetTimeLatitudeType(v string) *GetSkillGroupServiceStatusRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}

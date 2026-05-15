// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupStatusTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetSkillGroupStatusTotalRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetSkillGroupStatusTotalRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetSkillGroupStatusTotalRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetSkillGroupStatusTotalRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetSkillGroupStatusTotalRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetSkillGroupStatusTotalRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupStatusTotalRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetSkillGroupStatusTotalRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetSkillGroupStatusTotalRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupStatusTotalRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupStatusTotalRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetSkillGroupStatusTotalRequest
	GetTimeLatitudeType() *string
}

type GetSkillGroupStatusTotalRequest struct {
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
	// fasle
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// example:
	//
	// fasle
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
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetSkillGroupStatusTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupStatusTotalRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetSkillGroupStatusTotalRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupStatusTotalRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetSkillGroupStatusTotalRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupStatusTotalRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetSkillGroupStatusTotalRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupStatusTotalRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupStatusTotalRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetSkillGroupStatusTotalRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupStatusTotalRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupStatusTotalRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupStatusTotalRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetSkillGroupStatusTotalRequest) SetAgentIds(v []*int64) *GetSkillGroupStatusTotalRequest {
	s.AgentIds = v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetCurrentPage(v int32) *GetSkillGroupStatusTotalRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetDepIds(v []*int64) *GetSkillGroupStatusTotalRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetEndDate(v int64) *GetSkillGroupStatusTotalRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetExistAgentGrouping(v bool) *GetSkillGroupStatusTotalRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupStatusTotalRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupStatusTotalRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetGroupIds(v []*int64) *GetSkillGroupStatusTotalRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetInstanceId(v string) *GetSkillGroupStatusTotalRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetPageSize(v int32) *GetSkillGroupStatusTotalRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetStartDate(v int64) *GetSkillGroupStatusTotalRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetTimeLatitudeType(v string) *GetSkillGroupStatusTotalRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineServiceVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetOnlineServiceVolumeRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetOnlineServiceVolumeRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetOnlineServiceVolumeRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetOnlineServiceVolumeRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetOnlineServiceVolumeRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetOnlineServiceVolumeRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetOnlineServiceVolumeRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetOnlineServiceVolumeRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetOnlineServiceVolumeRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetOnlineServiceVolumeRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetOnlineServiceVolumeRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetOnlineServiceVolumeRequest
	GetTimeLatitudeType() *string
}

type GetOnlineServiceVolumeRequest struct {
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
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetOnlineServiceVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineServiceVolumeRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetOnlineServiceVolumeRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOnlineServiceVolumeRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetOnlineServiceVolumeRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetOnlineServiceVolumeRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetOnlineServiceVolumeRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetOnlineServiceVolumeRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetOnlineServiceVolumeRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetOnlineServiceVolumeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOnlineServiceVolumeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOnlineServiceVolumeRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetOnlineServiceVolumeRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetOnlineServiceVolumeRequest) SetAgentIds(v []*int64) *GetOnlineServiceVolumeRequest {
	s.AgentIds = v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetCurrentPage(v int32) *GetOnlineServiceVolumeRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetDepIds(v []*int64) *GetOnlineServiceVolumeRequest {
	s.DepIds = v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetEndDate(v int64) *GetOnlineServiceVolumeRequest {
	s.EndDate = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetExistAgentGrouping(v bool) *GetOnlineServiceVolumeRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetExistDepartmentGrouping(v bool) *GetOnlineServiceVolumeRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetExistSkillGroupGrouping(v bool) *GetOnlineServiceVolumeRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetGroupIds(v []*int64) *GetOnlineServiceVolumeRequest {
	s.GroupIds = v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetInstanceId(v string) *GetOnlineServiceVolumeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetPageSize(v int32) *GetOnlineServiceVolumeRequest {
	s.PageSize = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetStartDate(v int64) *GetOnlineServiceVolumeRequest {
	s.StartDate = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetTimeLatitudeType(v string) *GetOnlineServiceVolumeRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) Validate() error {
	return dara.Validate(s)
}

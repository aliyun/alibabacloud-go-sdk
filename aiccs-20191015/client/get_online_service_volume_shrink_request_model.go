// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineServiceVolumeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetOnlineServiceVolumeShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetOnlineServiceVolumeShrinkRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetOnlineServiceVolumeShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetOnlineServiceVolumeShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetOnlineServiceVolumeShrinkRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetOnlineServiceVolumeShrinkRequest
	GetTimeLatitudeType() *string
}

type GetOnlineServiceVolumeShrinkRequest struct {
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
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
	// false
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// example:
	//
	// false
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// example:
	//
	// false
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
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetOnlineServiceVolumeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineServiceVolumeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetOnlineServiceVolumeShrinkRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetAgentIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetCurrentPage(v int32) *GetOnlineServiceVolumeShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetDepIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetEndDate(v int64) *GetOnlineServiceVolumeShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetExistAgentGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetExistDepartmentGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetGroupIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetInstanceId(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetPageSize(v int32) *GetOnlineServiceVolumeShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetStartDate(v int64) *GetOnlineServiceVolumeShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetTimeLatitudeType(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) Validate() error {
	return dara.Validate(s)
}

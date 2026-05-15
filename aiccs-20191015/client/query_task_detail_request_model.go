// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAni(v string) *QueryTaskDetailRequest
	GetAni() *string
	SetCurrentPage(v int32) *QueryTaskDetailRequest
	GetCurrentPage() *int32
	SetDepartmentIdList(v string) *QueryTaskDetailRequest
	GetDepartmentIdList() *string
	SetDnis(v string) *QueryTaskDetailRequest
	GetDnis() *string
	SetEndReasonList(v string) *QueryTaskDetailRequest
	GetEndReasonList() *string
	SetInstanceId(v string) *QueryTaskDetailRequest
	GetInstanceId() *string
	SetOutboundTaskId(v string) *QueryTaskDetailRequest
	GetOutboundTaskId() *string
	SetPageSize(v int32) *QueryTaskDetailRequest
	GetPageSize() *int32
	SetPriorityList(v string) *QueryTaskDetailRequest
	GetPriorityList() *string
	SetServicerId(v string) *QueryTaskDetailRequest
	GetServicerId() *string
	SetServicerName(v string) *QueryTaskDetailRequest
	GetServicerName() *string
	SetSid(v string) *QueryTaskDetailRequest
	GetSid() *string
	SetSkillGroup(v string) *QueryTaskDetailRequest
	GetSkillGroup() *string
	SetStatusList(v string) *QueryTaskDetailRequest
	GetStatusList() *string
	SetTaskId(v int64) *QueryTaskDetailRequest
	GetTaskId() *int64
}

type QueryTaskDetailRequest struct {
	Ani              *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	CurrentPage      *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepartmentIdList *string `json:"DepartmentIdList,omitempty" xml:"DepartmentIdList,omitempty"`
	Dnis             *string `json:"Dnis,omitempty" xml:"Dnis,omitempty"`
	EndReasonList    *string `json:"EndReasonList,omitempty" xml:"EndReasonList,omitempty"`
	// This parameter is required.
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundTaskId *string `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PriorityList   *string `json:"PriorityList,omitempty" xml:"PriorityList,omitempty"`
	ServicerId     *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	ServicerName   *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	Sid            *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	SkillGroup     *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	StatusList     *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	TaskId         *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailRequest) GetAni() *string {
	return s.Ani
}

func (s *QueryTaskDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryTaskDetailRequest) GetDepartmentIdList() *string {
	return s.DepartmentIdList
}

func (s *QueryTaskDetailRequest) GetDnis() *string {
	return s.Dnis
}

func (s *QueryTaskDetailRequest) GetEndReasonList() *string {
	return s.EndReasonList
}

func (s *QueryTaskDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTaskDetailRequest) GetOutboundTaskId() *string {
	return s.OutboundTaskId
}

func (s *QueryTaskDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskDetailRequest) GetPriorityList() *string {
	return s.PriorityList
}

func (s *QueryTaskDetailRequest) GetServicerId() *string {
	return s.ServicerId
}

func (s *QueryTaskDetailRequest) GetServicerName() *string {
	return s.ServicerName
}

func (s *QueryTaskDetailRequest) GetSid() *string {
	return s.Sid
}

func (s *QueryTaskDetailRequest) GetSkillGroup() *string {
	return s.SkillGroup
}

func (s *QueryTaskDetailRequest) GetStatusList() *string {
	return s.StatusList
}

func (s *QueryTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryTaskDetailRequest) SetAni(v string) *QueryTaskDetailRequest {
	s.Ani = &v
	return s
}

func (s *QueryTaskDetailRequest) SetCurrentPage(v int32) *QueryTaskDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTaskDetailRequest) SetDepartmentIdList(v string) *QueryTaskDetailRequest {
	s.DepartmentIdList = &v
	return s
}

func (s *QueryTaskDetailRequest) SetDnis(v string) *QueryTaskDetailRequest {
	s.Dnis = &v
	return s
}

func (s *QueryTaskDetailRequest) SetEndReasonList(v string) *QueryTaskDetailRequest {
	s.EndReasonList = &v
	return s
}

func (s *QueryTaskDetailRequest) SetInstanceId(v string) *QueryTaskDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailRequest) SetOutboundTaskId(v string) *QueryTaskDetailRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *QueryTaskDetailRequest) SetPageSize(v int32) *QueryTaskDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailRequest) SetPriorityList(v string) *QueryTaskDetailRequest {
	s.PriorityList = &v
	return s
}

func (s *QueryTaskDetailRequest) SetServicerId(v string) *QueryTaskDetailRequest {
	s.ServicerId = &v
	return s
}

func (s *QueryTaskDetailRequest) SetServicerName(v string) *QueryTaskDetailRequest {
	s.ServicerName = &v
	return s
}

func (s *QueryTaskDetailRequest) SetSid(v string) *QueryTaskDetailRequest {
	s.Sid = &v
	return s
}

func (s *QueryTaskDetailRequest) SetSkillGroup(v string) *QueryTaskDetailRequest {
	s.SkillGroup = &v
	return s
}

func (s *QueryTaskDetailRequest) SetStatusList(v string) *QueryTaskDetailRequest {
	s.StatusList = &v
	return s
}

func (s *QueryTaskDetailRequest) SetTaskId(v int64) *QueryTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}

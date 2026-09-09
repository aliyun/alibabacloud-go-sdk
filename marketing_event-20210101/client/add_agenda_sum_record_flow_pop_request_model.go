// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgendaSumRecordFlowPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveNum(v int32) *AddAgendaSumRecordFlowPopRequest
	GetActiveNum() *int32
	SetAgendaId(v int64) *AddAgendaSumRecordFlowPopRequest
	GetAgendaId() *int64
	SetAttendancePercent(v string) *AddAgendaSumRecordFlowPopRequest
	GetAttendancePercent() *string
	SetFlowTime(v int64) *AddAgendaSumRecordFlowPopRequest
	GetFlowTime() *int64
	SetSessionName(v string) *AddAgendaSumRecordFlowPopRequest
	GetSessionName() *string
	SetTotalPv(v int32) *AddAgendaSumRecordFlowPopRequest
	GetTotalPv() *int32
	SetTotalUv(v int32) *AddAgendaSumRecordFlowPopRequest
	GetTotalUv() *int32
}

type AddAgendaSumRecordFlowPopRequest struct {
	// example:
	//
	// 10
	ActiveNum *int32 `json:"ActiveNum,omitempty" xml:"ActiveNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	AgendaId *int64 `json:"AgendaId,omitempty" xml:"AgendaId,omitempty"`
	// example:
	//
	// 90.81
	AttendancePercent *string `json:"AttendancePercent,omitempty" xml:"AttendancePercent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1784443333333
	FlowTime *int64 `json:"FlowTime,omitempty" xml:"FlowTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 云栖大会主论坛
	SessionName *string `json:"SessionName,omitempty" xml:"SessionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	TotalPv *int32 `json:"TotalPv,omitempty" xml:"TotalPv,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8
	TotalUv *int32 `json:"TotalUv,omitempty" xml:"TotalUv,omitempty"`
}

func (s AddAgendaSumRecordFlowPopRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAgendaSumRecordFlowPopRequest) GoString() string {
	return s.String()
}

func (s *AddAgendaSumRecordFlowPopRequest) GetActiveNum() *int32 {
	return s.ActiveNum
}

func (s *AddAgendaSumRecordFlowPopRequest) GetAgendaId() *int64 {
	return s.AgendaId
}

func (s *AddAgendaSumRecordFlowPopRequest) GetAttendancePercent() *string {
	return s.AttendancePercent
}

func (s *AddAgendaSumRecordFlowPopRequest) GetFlowTime() *int64 {
	return s.FlowTime
}

func (s *AddAgendaSumRecordFlowPopRequest) GetSessionName() *string {
	return s.SessionName
}

func (s *AddAgendaSumRecordFlowPopRequest) GetTotalPv() *int32 {
	return s.TotalPv
}

func (s *AddAgendaSumRecordFlowPopRequest) GetTotalUv() *int32 {
	return s.TotalUv
}

func (s *AddAgendaSumRecordFlowPopRequest) SetActiveNum(v int32) *AddAgendaSumRecordFlowPopRequest {
	s.ActiveNum = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopRequest) SetAgendaId(v int64) *AddAgendaSumRecordFlowPopRequest {
	s.AgendaId = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopRequest) SetAttendancePercent(v string) *AddAgendaSumRecordFlowPopRequest {
	s.AttendancePercent = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopRequest) SetFlowTime(v int64) *AddAgendaSumRecordFlowPopRequest {
	s.FlowTime = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopRequest) SetSessionName(v string) *AddAgendaSumRecordFlowPopRequest {
	s.SessionName = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopRequest) SetTotalPv(v int32) *AddAgendaSumRecordFlowPopRequest {
	s.TotalPv = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopRequest) SetTotalUv(v int32) *AddAgendaSumRecordFlowPopRequest {
	s.TotalUv = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopRequest) Validate() error {
	return dara.Validate(s)
}

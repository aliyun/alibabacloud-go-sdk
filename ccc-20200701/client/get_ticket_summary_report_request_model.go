// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketSummaryReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignee(v string) *GetTicketSummaryReportRequest
	GetAssignee() *string
	SetAssigneeType(v string) *GetTicketSummaryReportRequest
	GetAssigneeType() *string
	SetCategoryId(v string) *GetTicketSummaryReportRequest
	GetCategoryId() *string
	SetCreator(v string) *GetTicketSummaryReportRequest
	GetCreator() *string
	SetEndTime(v int64) *GetTicketSummaryReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetTicketSummaryReportRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *GetTicketSummaryReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetTicketSummaryReportRequest
	GetPageSize() *int32
	SetParticipant(v string) *GetTicketSummaryReportRequest
	GetParticipant() *string
	SetStartTime(v int64) *GetTicketSummaryReportRequest
	GetStartTime() *int64
	SetState(v string) *GetTicketSummaryReportRequest
	GetState() *string
}

type GetTicketSummaryReportRequest struct {
	// Assignee ID. This can be an agent ID or a skill group ID.
	//
	// example:
	//
	// assignee@ccc-test
	Assignee *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	// Assignee type.
	//
	// - Agent
	//
	// - SkillGroup
	//
	// example:
	//
	// Agent
	AssigneeType *string `json:"AssigneeType,omitempty" xml:"AssigneeType,omitempty"`
	// Ticket category ID.
	//
	// example:
	//
	// 43c2671b-***-***-86d0-6bd187905cc8
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Creator ID.
	//
	// example:
	//
	// creator@ccc-test
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// End time. Filter tickets by creation time.
	//
	// example:
	//
	// 1719590399999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Participant ID.
	//
	// example:
	//
	// participant@ccc-test
	Participant *string `json:"Participant,omitempty" xml:"Participant,omitempty"`
	// Start time. Filter tickets by creation time.
	//
	// example:
	//
	// 1716998400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Ticket state.
	//
	// - Processing
	//
	// - Withdrawal
	//
	// - Rejected
	//
	// - Closed
	//
	// example:
	//
	// 无
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetTicketSummaryReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTicketSummaryReportRequest) GoString() string {
	return s.String()
}

func (s *GetTicketSummaryReportRequest) GetAssignee() *string {
	return s.Assignee
}

func (s *GetTicketSummaryReportRequest) GetAssigneeType() *string {
	return s.AssigneeType
}

func (s *GetTicketSummaryReportRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetTicketSummaryReportRequest) GetCreator() *string {
	return s.Creator
}

func (s *GetTicketSummaryReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetTicketSummaryReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTicketSummaryReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetTicketSummaryReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTicketSummaryReportRequest) GetParticipant() *string {
	return s.Participant
}

func (s *GetTicketSummaryReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetTicketSummaryReportRequest) GetState() *string {
	return s.State
}

func (s *GetTicketSummaryReportRequest) SetAssignee(v string) *GetTicketSummaryReportRequest {
	s.Assignee = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetAssigneeType(v string) *GetTicketSummaryReportRequest {
	s.AssigneeType = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetCategoryId(v string) *GetTicketSummaryReportRequest {
	s.CategoryId = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetCreator(v string) *GetTicketSummaryReportRequest {
	s.Creator = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetEndTime(v int64) *GetTicketSummaryReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetInstanceId(v string) *GetTicketSummaryReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetPageNumber(v int32) *GetTicketSummaryReportRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetPageSize(v int32) *GetTicketSummaryReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetParticipant(v string) *GetTicketSummaryReportRequest {
	s.Participant = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetStartTime(v int64) *GetTicketSummaryReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetTicketSummaryReportRequest) SetState(v string) *GetTicketSummaryReportRequest {
	s.State = &v
	return s
}

func (s *GetTicketSummaryReportRequest) Validate() error {
	return dara.Validate(s)
}

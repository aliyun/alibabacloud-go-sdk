// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignee(v string) *ListTicketsRequest
	GetAssignee() *string
	SetAssigneeType(v string) *ListTicketsRequest
	GetAssigneeType() *string
	SetCategoryId(v string) *ListTicketsRequest
	GetCategoryId() *string
	SetCreator(v string) *ListTicketsRequest
	GetCreator() *string
	SetCustomerId(v string) *ListTicketsRequest
	GetCustomerId() *string
	SetEndTime(v int64) *ListTicketsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListTicketsRequest
	GetInstanceId() *string
	SetJobIdList(v string) *ListTicketsRequest
	GetJobIdList() *string
	SetPageNumber(v int64) *ListTicketsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTicketsRequest
	GetPageSize() *int64
	SetParticipant(v string) *ListTicketsRequest
	GetParticipant() *string
	SetStartTime(v int64) *ListTicketsRequest
	GetStartTime() *int64
	SetState(v string) *ListTicketsRequest
	GetState() *string
	SetTicketId(v string) *ListTicketsRequest
	GetTicketId() *string
	SetTitle(v string) *ListTicketsRequest
	GetTitle() *string
}

type ListTicketsRequest struct {
	// example:
	//
	// assignee@cccV2-kmz
	Assignee *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	// example:
	//
	// Agent
	AssigneeType *string `json:"AssigneeType,omitempty" xml:"AssigneeType,omitempty"`
	// example:
	//
	// 43c2671b-****-4223-86d0-6bd187905cc8
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// creator@cccV2-kmz
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 51e155ce-3747-*****-b402-13c69597b920
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// example:
	//
	// 1646928000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["job-12******","job-23****"]
	JobIdList *string `json:"JobIdList,omitempty" xml:"JobIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// participant@cccV2-kmz
	Participant *string `json:"Participant,omitempty" xml:"Participant,omitempty"`
	// example:
	//
	// 1646841600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Processing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 5491d3b4-14ee-4341-b5f1-db2c78beddeb
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListTicketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsRequest) GetAssignee() *string {
	return s.Assignee
}

func (s *ListTicketsRequest) GetAssigneeType() *string {
	return s.AssigneeType
}

func (s *ListTicketsRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListTicketsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListTicketsRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListTicketsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListTicketsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTicketsRequest) GetJobIdList() *string {
	return s.JobIdList
}

func (s *ListTicketsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTicketsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTicketsRequest) GetParticipant() *string {
	return s.Participant
}

func (s *ListTicketsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListTicketsRequest) GetState() *string {
	return s.State
}

func (s *ListTicketsRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketsRequest) GetTitle() *string {
	return s.Title
}

func (s *ListTicketsRequest) SetAssignee(v string) *ListTicketsRequest {
	s.Assignee = &v
	return s
}

func (s *ListTicketsRequest) SetAssigneeType(v string) *ListTicketsRequest {
	s.AssigneeType = &v
	return s
}

func (s *ListTicketsRequest) SetCategoryId(v string) *ListTicketsRequest {
	s.CategoryId = &v
	return s
}

func (s *ListTicketsRequest) SetCreator(v string) *ListTicketsRequest {
	s.Creator = &v
	return s
}

func (s *ListTicketsRequest) SetCustomerId(v string) *ListTicketsRequest {
	s.CustomerId = &v
	return s
}

func (s *ListTicketsRequest) SetEndTime(v int64) *ListTicketsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTicketsRequest) SetInstanceId(v string) *ListTicketsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTicketsRequest) SetJobIdList(v string) *ListTicketsRequest {
	s.JobIdList = &v
	return s
}

func (s *ListTicketsRequest) SetPageNumber(v int64) *ListTicketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsRequest) SetPageSize(v int64) *ListTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketsRequest) SetParticipant(v string) *ListTicketsRequest {
	s.Participant = &v
	return s
}

func (s *ListTicketsRequest) SetStartTime(v int64) *ListTicketsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTicketsRequest) SetState(v string) *ListTicketsRequest {
	s.State = &v
	return s
}

func (s *ListTicketsRequest) SetTicketId(v string) *ListTicketsRequest {
	s.TicketId = &v
	return s
}

func (s *ListTicketsRequest) SetTitle(v string) *ListTicketsRequest {
	s.Title = &v
	return s
}

func (s *ListTicketsRequest) Validate() error {
	return dara.Validate(s)
}

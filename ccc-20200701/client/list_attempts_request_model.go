// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttemptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListAttemptsRequest
	GetAgentId() *string
	SetAttemptId(v string) *ListAttemptsRequest
	GetAttemptId() *string
	SetCallee(v string) *ListAttemptsRequest
	GetCallee() *string
	SetCaller(v string) *ListAttemptsRequest
	GetCaller() *string
	SetCampaignId(v string) *ListAttemptsRequest
	GetCampaignId() *string
	SetCaseId(v string) *ListAttemptsRequest
	GetCaseId() *string
	SetContactId(v string) *ListAttemptsRequest
	GetContactId() *string
	SetCriteria(v string) *ListAttemptsRequest
	GetCriteria() *string
	SetEndTime(v int64) *ListAttemptsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListAttemptsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListAttemptsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAttemptsRequest
	GetPageSize() *int32
	SetQueueId(v string) *ListAttemptsRequest
	GetQueueId() *string
	SetStartTime(v int64) *ListAttemptsRequest
	GetStartTime() *int64
}

type ListAttemptsRequest struct {
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// job-16976964500325****
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// example:
	//
	// 1888888****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 05711234****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6badb397-a8b5-40b6-21019d382a09
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60ecb1a2-4480-4d01-bede-c5b7655bfadf
	CaseId *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// example:
	//
	// job-16976964500325****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Criteria  *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// example:
	//
	// 1634115698291
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// example:
	//
	// 1634115688291
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAttemptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAttemptsRequest) GoString() string {
	return s.String()
}

func (s *ListAttemptsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAttemptsRequest) GetAttemptId() *string {
	return s.AttemptId
}

func (s *ListAttemptsRequest) GetCallee() *string {
	return s.Callee
}

func (s *ListAttemptsRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListAttemptsRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListAttemptsRequest) GetCaseId() *string {
	return s.CaseId
}

func (s *ListAttemptsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListAttemptsRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ListAttemptsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAttemptsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAttemptsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAttemptsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAttemptsRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *ListAttemptsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAttemptsRequest) SetAgentId(v string) *ListAttemptsRequest {
	s.AgentId = &v
	return s
}

func (s *ListAttemptsRequest) SetAttemptId(v string) *ListAttemptsRequest {
	s.AttemptId = &v
	return s
}

func (s *ListAttemptsRequest) SetCallee(v string) *ListAttemptsRequest {
	s.Callee = &v
	return s
}

func (s *ListAttemptsRequest) SetCaller(v string) *ListAttemptsRequest {
	s.Caller = &v
	return s
}

func (s *ListAttemptsRequest) SetCampaignId(v string) *ListAttemptsRequest {
	s.CampaignId = &v
	return s
}

func (s *ListAttemptsRequest) SetCaseId(v string) *ListAttemptsRequest {
	s.CaseId = &v
	return s
}

func (s *ListAttemptsRequest) SetContactId(v string) *ListAttemptsRequest {
	s.ContactId = &v
	return s
}

func (s *ListAttemptsRequest) SetCriteria(v string) *ListAttemptsRequest {
	s.Criteria = &v
	return s
}

func (s *ListAttemptsRequest) SetEndTime(v int64) *ListAttemptsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAttemptsRequest) SetInstanceId(v string) *ListAttemptsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAttemptsRequest) SetPageNumber(v int32) *ListAttemptsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAttemptsRequest) SetPageSize(v int32) *ListAttemptsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAttemptsRequest) SetQueueId(v string) *ListAttemptsRequest {
	s.QueueId = &v
	return s
}

func (s *ListAttemptsRequest) SetStartTime(v int64) *ListAttemptsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAttemptsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttemptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAttemptsResponseBody
	GetCode() *string
	SetData(v *ListAttemptsResponseBodyData) *ListAttemptsResponseBody
	GetData() *ListAttemptsResponseBodyData
	SetHttpStatusCode(v int32) *ListAttemptsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAttemptsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAttemptsResponseBody
	GetRequestId() *string
}

type ListAttemptsResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAttemptsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7CC6523B-0E51-1B62-8DA5-6A9831CAE315
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAttemptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAttemptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAttemptsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAttemptsResponseBody) GetData() *ListAttemptsResponseBodyData {
	return s.Data
}

func (s *ListAttemptsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAttemptsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAttemptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAttemptsResponseBody) SetCode(v string) *ListAttemptsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAttemptsResponseBody) SetData(v *ListAttemptsResponseBodyData) *ListAttemptsResponseBody {
	s.Data = v
	return s
}

func (s *ListAttemptsResponseBody) SetHttpStatusCode(v int32) *ListAttemptsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAttemptsResponseBody) SetMessage(v string) *ListAttemptsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAttemptsResponseBody) SetRequestId(v string) *ListAttemptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAttemptsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAttemptsResponseBodyData struct {
	List []*ListAttemptsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAttemptsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAttemptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAttemptsResponseBodyData) GetList() []*ListAttemptsResponseBodyDataList {
	return s.List
}

func (s *ListAttemptsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAttemptsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAttemptsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAttemptsResponseBodyData) SetList(v []*ListAttemptsResponseBodyDataList) *ListAttemptsResponseBodyData {
	s.List = v
	return s
}

func (s *ListAttemptsResponseBodyData) SetPageNumber(v int32) *ListAttemptsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAttemptsResponseBodyData) SetPageSize(v int32) *ListAttemptsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAttemptsResponseBodyData) SetTotalCount(v int32) *ListAttemptsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAttemptsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAttemptsResponseBodyDataList struct {
	// example:
	//
	// 1632883592732
	AgentEstablishedTime *int64 `json:"AgentEstablishedTime,omitempty" xml:"AgentEstablishedTime,omitempty"`
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 23
	AgentRingDuration *int64 `json:"AgentRingDuration,omitempty" xml:"AgentRingDuration,omitempty"`
	// example:
	//
	// 1634196287869
	AssignAgentTime *int64 `json:"AssignAgentTime,omitempty" xml:"AssignAgentTime,omitempty"`
	// example:
	//
	// job-1704342174816****
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// example:
	//
	// 1888888****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 05711234****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// 083046e3-5822-4cda-9b84-04f2a02eb605
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// example:
	//
	// 21d194a7-60b7-4824-932b-48ed03a83704
	CaseId *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// example:
	//
	// job-1704342174816****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// 1634196286708
	CustomerEstablishedTime *int64 `json:"CustomerEstablishedTime,omitempty" xml:"CustomerEstablishedTime,omitempty"`
	// example:
	//
	// 1634196317888
	CustomerReleasedTime *int64 `json:"CustomerReleasedTime,omitempty" xml:"CustomerReleasedTime,omitempty"`
	// example:
	//
	// 2734
	DialDuration *int64 `json:"DialDuration,omitempty" xml:"DialDuration,omitempty"`
	// example:
	//
	// 1634196283974
	DialTime *int64 `json:"DialTime,omitempty" xml:"DialTime,omitempty"`
	// example:
	//
	// 1634196287789
	EnqueueTime *int64 `json:"EnqueueTime,omitempty" xml:"EnqueueTime,omitempty"`
	// example:
	//
	// 1634196286740
	EnterIvrTime *int64 `json:"EnterIvrTime,omitempty" xml:"EnterIvrTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1049
	IvrDuration *int64 `json:"IvrDuration,omitempty" xml:"IvrDuration,omitempty"`
	// example:
	//
	// 80
	QueueDuration *int64 `json:"QueueDuration,omitempty" xml:"QueueDuration,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
}

func (s ListAttemptsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListAttemptsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListAttemptsResponseBodyDataList) GetAgentEstablishedTime() *int64 {
	return s.AgentEstablishedTime
}

func (s *ListAttemptsResponseBodyDataList) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAttemptsResponseBodyDataList) GetAgentRingDuration() *int64 {
	return s.AgentRingDuration
}

func (s *ListAttemptsResponseBodyDataList) GetAssignAgentTime() *int64 {
	return s.AssignAgentTime
}

func (s *ListAttemptsResponseBodyDataList) GetAttemptId() *string {
	return s.AttemptId
}

func (s *ListAttemptsResponseBodyDataList) GetCallee() *string {
	return s.Callee
}

func (s *ListAttemptsResponseBodyDataList) GetCaller() *string {
	return s.Caller
}

func (s *ListAttemptsResponseBodyDataList) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListAttemptsResponseBodyDataList) GetCaseId() *string {
	return s.CaseId
}

func (s *ListAttemptsResponseBodyDataList) GetContactId() *string {
	return s.ContactId
}

func (s *ListAttemptsResponseBodyDataList) GetCustomerEstablishedTime() *int64 {
	return s.CustomerEstablishedTime
}

func (s *ListAttemptsResponseBodyDataList) GetCustomerReleasedTime() *int64 {
	return s.CustomerReleasedTime
}

func (s *ListAttemptsResponseBodyDataList) GetDialDuration() *int64 {
	return s.DialDuration
}

func (s *ListAttemptsResponseBodyDataList) GetDialTime() *int64 {
	return s.DialTime
}

func (s *ListAttemptsResponseBodyDataList) GetEnqueueTime() *int64 {
	return s.EnqueueTime
}

func (s *ListAttemptsResponseBodyDataList) GetEnterIvrTime() *int64 {
	return s.EnterIvrTime
}

func (s *ListAttemptsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAttemptsResponseBodyDataList) GetIvrDuration() *int64 {
	return s.IvrDuration
}

func (s *ListAttemptsResponseBodyDataList) GetQueueDuration() *int64 {
	return s.QueueDuration
}

func (s *ListAttemptsResponseBodyDataList) GetQueueId() *string {
	return s.QueueId
}

func (s *ListAttemptsResponseBodyDataList) SetAgentEstablishedTime(v int64) *ListAttemptsResponseBodyDataList {
	s.AgentEstablishedTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetAgentId(v string) *ListAttemptsResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetAgentRingDuration(v int64) *ListAttemptsResponseBodyDataList {
	s.AgentRingDuration = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetAssignAgentTime(v int64) *ListAttemptsResponseBodyDataList {
	s.AssignAgentTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetAttemptId(v string) *ListAttemptsResponseBodyDataList {
	s.AttemptId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCallee(v string) *ListAttemptsResponseBodyDataList {
	s.Callee = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCaller(v string) *ListAttemptsResponseBodyDataList {
	s.Caller = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCampaignId(v string) *ListAttemptsResponseBodyDataList {
	s.CampaignId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCaseId(v string) *ListAttemptsResponseBodyDataList {
	s.CaseId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetContactId(v string) *ListAttemptsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCustomerEstablishedTime(v int64) *ListAttemptsResponseBodyDataList {
	s.CustomerEstablishedTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCustomerReleasedTime(v int64) *ListAttemptsResponseBodyDataList {
	s.CustomerReleasedTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetDialDuration(v int64) *ListAttemptsResponseBodyDataList {
	s.DialDuration = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetDialTime(v int64) *ListAttemptsResponseBodyDataList {
	s.DialTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetEnqueueTime(v int64) *ListAttemptsResponseBodyDataList {
	s.EnqueueTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetEnterIvrTime(v int64) *ListAttemptsResponseBodyDataList {
	s.EnterIvrTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetInstanceId(v string) *ListAttemptsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetIvrDuration(v int64) *ListAttemptsResponseBodyDataList {
	s.IvrDuration = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetQueueDuration(v int64) *ListAttemptsResponseBodyDataList {
	s.QueueDuration = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetQueueId(v string) *ListAttemptsResponseBodyDataList {
	s.QueueId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

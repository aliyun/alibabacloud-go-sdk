// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCampaignsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCampaignsResponseBody
	GetCode() *string
	SetData(v *ListCampaignsResponseBodyData) *ListCampaignsResponseBody
	GetData() *ListCampaignsResponseBodyData
	SetHttpStatusCode(v int64) *ListCampaignsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListCampaignsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCampaignsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCampaignsResponseBody
	GetSuccess() *bool
}

type ListCampaignsResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCampaignsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6CCEF32F-8614-535F-A1D9-D85B8C0DC4F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCampaignsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCampaignsResponseBody) GetData() *ListCampaignsResponseBodyData {
	return s.Data
}

func (s *ListCampaignsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListCampaignsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCampaignsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCampaignsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCampaignsResponseBody) SetCode(v string) *ListCampaignsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCampaignsResponseBody) SetData(v *ListCampaignsResponseBodyData) *ListCampaignsResponseBody {
	s.Data = v
	return s
}

func (s *ListCampaignsResponseBody) SetHttpStatusCode(v int64) *ListCampaignsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCampaignsResponseBody) SetMessage(v string) *ListCampaignsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCampaignsResponseBody) SetRequestId(v string) *ListCampaignsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCampaignsResponseBody) SetSuccess(v bool) *ListCampaignsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCampaignsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCampaignsResponseBodyData struct {
	List []*ListCampaignsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCampaignsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBodyData) GetList() []*ListCampaignsResponseBodyDataList {
	return s.List
}

func (s *ListCampaignsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCampaignsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCampaignsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCampaignsResponseBodyData) SetList(v []*ListCampaignsResponseBodyDataList) *ListCampaignsResponseBodyData {
	s.List = v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageNumber(v int64) *ListCampaignsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageSize(v int64) *ListCampaignsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetTotalCount(v int64) *ListCampaignsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCampaignsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCampaignsResponseBodyDataList struct {
	// example:
	//
	// 1634008800000
	ActualEndTime *int64 `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	// example:
	//
	// 1634000460000
	ActualStartTime *int64 `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// example:
	//
	// 6badb397-a8b5-40b6-21019d382a09
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// example:
	//
	// 0
	CasesAborted *int64 `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	// example:
	//
	// 40
	CasesConnected *int64 `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	// example:
	//
	// 0
	CasesUncompleted *int64   `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	CompletionRate   *float32 `json:"CompletionRate,omitempty" xml:"CompletionRate,omitempty"`
	ContactFlowId    *string  `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// example:
	//
	// 1
	MaxAttemptCount *int64 `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	// example:
	//
	// 1
	MinAttemptInterval *int64 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// example:
	//
	// test-campaign
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1634054400000
	PlanedEndTime *int64 `json:"PlanedEndTime,omitempty" xml:"PlanedEndTime,omitempty"`
	// example:
	//
	// 1633968000000
	PlanedStartTime *int64 `json:"PlanedStartTime,omitempty" xml:"PlanedStartTime,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	QueueId   *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// false
	Simulation *bool `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	// example:
	//
	// Completed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// {"ratio":1}
	StrategyParameters *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	// example:
	//
	// PACING
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// example:
	//
	// 100
	TotalCases *int64 `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
}

func (s ListCampaignsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBodyDataList) GetActualEndTime() *int64 {
	return s.ActualEndTime
}

func (s *ListCampaignsResponseBodyDataList) GetActualStartTime() *int64 {
	return s.ActualStartTime
}

func (s *ListCampaignsResponseBodyDataList) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListCampaignsResponseBodyDataList) GetCasesAborted() *int64 {
	return s.CasesAborted
}

func (s *ListCampaignsResponseBodyDataList) GetCasesConnected() *int64 {
	return s.CasesConnected
}

func (s *ListCampaignsResponseBodyDataList) GetCasesUncompleted() *int64 {
	return s.CasesUncompleted
}

func (s *ListCampaignsResponseBodyDataList) GetCompletionRate() *float32 {
	return s.CompletionRate
}

func (s *ListCampaignsResponseBodyDataList) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *ListCampaignsResponseBodyDataList) GetMaxAttemptCount() *int64 {
	return s.MaxAttemptCount
}

func (s *ListCampaignsResponseBodyDataList) GetMinAttemptInterval() *int64 {
	return s.MinAttemptInterval
}

func (s *ListCampaignsResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListCampaignsResponseBodyDataList) GetPlanedEndTime() *int64 {
	return s.PlanedEndTime
}

func (s *ListCampaignsResponseBodyDataList) GetPlanedStartTime() *int64 {
	return s.PlanedStartTime
}

func (s *ListCampaignsResponseBodyDataList) GetQueueId() *string {
	return s.QueueId
}

func (s *ListCampaignsResponseBodyDataList) GetQueueName() *string {
	return s.QueueName
}

func (s *ListCampaignsResponseBodyDataList) GetSimulation() *bool {
	return s.Simulation
}

func (s *ListCampaignsResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListCampaignsResponseBodyDataList) GetStrategyParameters() *string {
	return s.StrategyParameters
}

func (s *ListCampaignsResponseBodyDataList) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListCampaignsResponseBodyDataList) GetTotalCases() *int64 {
	return s.TotalCases
}

func (s *ListCampaignsResponseBodyDataList) SetActualEndTime(v int64) *ListCampaignsResponseBodyDataList {
	s.ActualEndTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetActualStartTime(v int64) *ListCampaignsResponseBodyDataList {
	s.ActualStartTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCampaignId(v string) *ListCampaignsResponseBodyDataList {
	s.CampaignId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesAborted(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesAborted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesConnected(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesConnected = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesUncompleted(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesUncompleted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCompletionRate(v float32) *ListCampaignsResponseBodyDataList {
	s.CompletionRate = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetContactFlowId(v string) *ListCampaignsResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetMaxAttemptCount(v int64) *ListCampaignsResponseBodyDataList {
	s.MaxAttemptCount = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetMinAttemptInterval(v int64) *ListCampaignsResponseBodyDataList {
	s.MinAttemptInterval = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetName(v string) *ListCampaignsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetPlanedEndTime(v int64) *ListCampaignsResponseBodyDataList {
	s.PlanedEndTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetPlanedStartTime(v int64) *ListCampaignsResponseBodyDataList {
	s.PlanedStartTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetQueueId(v string) *ListCampaignsResponseBodyDataList {
	s.QueueId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetQueueName(v string) *ListCampaignsResponseBodyDataList {
	s.QueueName = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetSimulation(v bool) *ListCampaignsResponseBodyDataList {
	s.Simulation = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetState(v string) *ListCampaignsResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetStrategyParameters(v string) *ListCampaignsResponseBodyDataList {
	s.StrategyParameters = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetStrategyType(v string) *ListCampaignsResponseBodyDataList {
	s.StrategyType = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetTotalCases(v int64) *ListCampaignsResponseBodyDataList {
	s.TotalCases = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

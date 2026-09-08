// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCampaignResponseBody
	GetCode() *string
	SetData(v *GetCampaignResponseBodyData) *GetCampaignResponseBody
	GetData() *GetCampaignResponseBodyData
	SetHttpStatusCode(v int64) *GetCampaignResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *GetCampaignResponseBody
	GetRequestId() *string
}

type GetCampaignResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetCampaignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6CCEF32F-8614-535F-A1D9-D85B8C0DC4F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *GetCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCampaignResponseBody) GetData() *GetCampaignResponseBodyData {
	return s.Data
}

func (s *GetCampaignResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCampaignResponseBody) SetCode(v string) *GetCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *GetCampaignResponseBody) SetData(v *GetCampaignResponseBodyData) *GetCampaignResponseBody {
	s.Data = v
	return s
}

func (s *GetCampaignResponseBody) SetHttpStatusCode(v int64) *GetCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCampaignResponseBody) SetRequestId(v string) *GetCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCampaignResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCampaignResponseBodyData struct {
	// The actual end time of the predictive outbound campaign. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1634008800000
	ActualEndTime *int64 `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	// The actual start time of the predictive outbound campaign. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1634000460000
	ActualStartTime *int64 `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// The ID of the predictive outbound campaign.
	//
	// example:
	//
	// 6badb397-a8b5-40b6-21019d382a09
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The number of aborted contacts.
	//
	// example:
	//
	// 0
	CasesAborted *int64 `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	// The number of connected contacts.
	//
	// example:
	//
	// 40
	CasesConnected *int64 `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	// The number of uncompleted contacts.
	//
	// example:
	//
	// 0
	CasesUncompleted             *int64  `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	CasesUncompletedAfterAttempt *string `json:"CasesUncompletedAfterAttempt,omitempty" xml:"CasesUncompletedAfterAttempt,omitempty"`
	// The number of attempted but unconnected contacts that are still eligible for redial.
	//
	// example:
	//
	// 1
	CasesUncompletedAfterAttempted *int64 `json:"CasesUncompletedAfterAttempted,omitempty" xml:"CasesUncompletedAfterAttempted,omitempty"`
	// The ratio of connected contacts to the total number of contacts.
	CompletionRate *float32 `json:"CompletionRate,omitempty" xml:"CompletionRate,omitempty"`
	// The ID of the associated Contact Flow.
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// The maximum number of call attempts for each contact. If an attempt fails, the contact is redialed until this limit is reached.
	//
	// example:
	//
	// 1
	MaxAttemptCount *int64 `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	// The minimum interval, in seconds, to wait before redialing a failed call.
	//
	// example:
	//
	// 1
	MinAttemptInterval *int64 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// The name of the predictive outbound campaign.
	//
	// example:
	//
	// test-campaign
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The planned end time of the predictive outbound campaign. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1634054400000
	PlanedEndTime *int64 `json:"PlanedEndTime,omitempty" xml:"PlanedEndTime,omitempty"`
	// The planned start time of the predictive outbound campaign. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1633968000000
	PlanedStartTime *int64 `json:"PlanedStartTime,omitempty" xml:"PlanedStartTime,omitempty"`
	// The ID of the associated skill group.
	//
	// example:
	//
	// skillgroup@ccc-test
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The name of the associated skill group.
	//
	// example:
	//
	// 测试技能组
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// Indicates whether the campaign is a test simulation. This parameter is not intended for production use.
	//
	// example:
	//
	// false
	Simulation *bool `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	// The parameters for a test simulation. This parameter is not intended for production use.
	//
	// example:
	//
	// 无
	SimulationParameters *string `json:"SimulationParameters,omitempty" xml:"SimulationParameters,omitempty"`
	// The state of the predictive outbound campaign.
	//
	// example:
	//
	// Completed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The strategy parameters for the predictive outbound campaign, in JSON format. For a `PID` strategy, the format is `{"abandonRate":"5","historicalConnectedRate":"35"}`. For a `PACING` strategy, the format is `{"ratio":1}`. `abandonRate` specifies the target abandon rate, `historicalConnectedRate` specifies the historical connection rate for reference, and `ratio` specifies the fixed dialing ratio.
	//
	// example:
	//
	// {"ratio":1}
	StrategyParameters *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	// The dialing strategy for the predictive outbound campaign.
	//
	// example:
	//
	// PACING
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// The total number of contacts.
	//
	// example:
	//
	// 100
	TotalCases *int64 `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
}

func (s GetCampaignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCampaignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCampaignResponseBodyData) GetActualEndTime() *int64 {
	return s.ActualEndTime
}

func (s *GetCampaignResponseBodyData) GetActualStartTime() *int64 {
	return s.ActualStartTime
}

func (s *GetCampaignResponseBodyData) GetCampaignId() *string {
	return s.CampaignId
}

func (s *GetCampaignResponseBodyData) GetCasesAborted() *int64 {
	return s.CasesAborted
}

func (s *GetCampaignResponseBodyData) GetCasesConnected() *int64 {
	return s.CasesConnected
}

func (s *GetCampaignResponseBodyData) GetCasesUncompleted() *int64 {
	return s.CasesUncompleted
}

func (s *GetCampaignResponseBodyData) GetCasesUncompletedAfterAttempt() *string {
	return s.CasesUncompletedAfterAttempt
}

func (s *GetCampaignResponseBodyData) GetCasesUncompletedAfterAttempted() *int64 {
	return s.CasesUncompletedAfterAttempted
}

func (s *GetCampaignResponseBodyData) GetCompletionRate() *float32 {
	return s.CompletionRate
}

func (s *GetCampaignResponseBodyData) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *GetCampaignResponseBodyData) GetMaxAttemptCount() *int64 {
	return s.MaxAttemptCount
}

func (s *GetCampaignResponseBodyData) GetMinAttemptInterval() *int64 {
	return s.MinAttemptInterval
}

func (s *GetCampaignResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetCampaignResponseBodyData) GetPlanedEndTime() *int64 {
	return s.PlanedEndTime
}

func (s *GetCampaignResponseBodyData) GetPlanedStartTime() *int64 {
	return s.PlanedStartTime
}

func (s *GetCampaignResponseBodyData) GetQueueId() *string {
	return s.QueueId
}

func (s *GetCampaignResponseBodyData) GetQueueName() *string {
	return s.QueueName
}

func (s *GetCampaignResponseBodyData) GetSimulation() *bool {
	return s.Simulation
}

func (s *GetCampaignResponseBodyData) GetSimulationParameters() *string {
	return s.SimulationParameters
}

func (s *GetCampaignResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetCampaignResponseBodyData) GetStrategyParameters() *string {
	return s.StrategyParameters
}

func (s *GetCampaignResponseBodyData) GetStrategyType() *string {
	return s.StrategyType
}

func (s *GetCampaignResponseBodyData) GetTotalCases() *int64 {
	return s.TotalCases
}

func (s *GetCampaignResponseBodyData) SetActualEndTime(v int64) *GetCampaignResponseBodyData {
	s.ActualEndTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetActualStartTime(v int64) *GetCampaignResponseBodyData {
	s.ActualStartTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCampaignId(v string) *GetCampaignResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesAborted(v int64) *GetCampaignResponseBodyData {
	s.CasesAborted = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesConnected(v int64) *GetCampaignResponseBodyData {
	s.CasesConnected = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesUncompleted(v int64) *GetCampaignResponseBodyData {
	s.CasesUncompleted = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesUncompletedAfterAttempt(v string) *GetCampaignResponseBodyData {
	s.CasesUncompletedAfterAttempt = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesUncompletedAfterAttempted(v int64) *GetCampaignResponseBodyData {
	s.CasesUncompletedAfterAttempted = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCompletionRate(v float32) *GetCampaignResponseBodyData {
	s.CompletionRate = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetContactFlowId(v string) *GetCampaignResponseBodyData {
	s.ContactFlowId = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetMaxAttemptCount(v int64) *GetCampaignResponseBodyData {
	s.MaxAttemptCount = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetMinAttemptInterval(v int64) *GetCampaignResponseBodyData {
	s.MinAttemptInterval = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetName(v string) *GetCampaignResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetPlanedEndTime(v int64) *GetCampaignResponseBodyData {
	s.PlanedEndTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetPlanedStartTime(v int64) *GetCampaignResponseBodyData {
	s.PlanedStartTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetQueueId(v string) *GetCampaignResponseBodyData {
	s.QueueId = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetQueueName(v string) *GetCampaignResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetSimulation(v bool) *GetCampaignResponseBodyData {
	s.Simulation = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetSimulationParameters(v string) *GetCampaignResponseBodyData {
	s.SimulationParameters = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetState(v string) *GetCampaignResponseBodyData {
	s.State = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetStrategyParameters(v string) *GetCampaignResponseBodyData {
	s.StrategyParameters = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetStrategyType(v string) *GetCampaignResponseBodyData {
	s.StrategyType = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetTotalCases(v int64) *GetCampaignResponseBodyData {
	s.TotalCases = &v
	return s
}

func (s *GetCampaignResponseBodyData) Validate() error {
	return dara.Validate(s)
}

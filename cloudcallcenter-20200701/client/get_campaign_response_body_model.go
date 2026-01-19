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
	SetSuccess(v bool) *GetCampaignResponseBody
	GetSuccess() *bool
}

type GetCampaignResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetCampaignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *GetCampaignResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *GetCampaignResponseBody) SetSuccess(v bool) *GetCampaignResponseBody {
	s.Success = &v
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
	ActualEndTime        *int64  `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	ActualStartTime      *int64  `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	CampaignId           *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CasesAborted         *int64  `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	CasesConnected       *int64  `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	CasesUncompleted     *int64  `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	CompletedRate        *int64  `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	MaxAttemptCount      *int64  `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	MinAttemptInterval   *int64  `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanedEndTime        *int64  `json:"PlanedEndTime,omitempty" xml:"PlanedEndTime,omitempty"`
	PlanedStartTime      *int64  `json:"PlanedStartTime,omitempty" xml:"PlanedStartTime,omitempty"`
	QueueId              *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	QueueName            *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Simulation           *bool   `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	SimulationParameters *string `json:"SimulationParameters,omitempty" xml:"SimulationParameters,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
	StrategyParameters   *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	StrategyType         *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	TotalCases           *int64  `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
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

func (s *GetCampaignResponseBodyData) GetCompletedRate() *int64 {
	return s.CompletedRate
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

func (s *GetCampaignResponseBodyData) SetCompletedRate(v int64) *GetCampaignResponseBodyData {
	s.CompletedRate = &v
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCampaignShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableTime(v string) *CreateCampaignShrinkRequest
	GetCallableTime() *string
	SetCaseFileKey(v string) *CreateCampaignShrinkRequest
	GetCaseFileKey() *string
	SetCaseListShrink(v string) *CreateCampaignShrinkRequest
	GetCaseListShrink() *string
	SetContactFlowId(v string) *CreateCampaignShrinkRequest
	GetContactFlowId() *string
	SetEndTime(v string) *CreateCampaignShrinkRequest
	GetEndTime() *string
	SetExecutingUntilTimeout(v bool) *CreateCampaignShrinkRequest
	GetExecutingUntilTimeout() *bool
	SetInstanceId(v string) *CreateCampaignShrinkRequest
	GetInstanceId() *string
	SetMaxAttemptCount(v int64) *CreateCampaignShrinkRequest
	GetMaxAttemptCount() *int64
	SetMinAttemptInterval(v int64) *CreateCampaignShrinkRequest
	GetMinAttemptInterval() *int64
	SetName(v string) *CreateCampaignShrinkRequest
	GetName() *string
	SetQueueId(v string) *CreateCampaignShrinkRequest
	GetQueueId() *string
	SetSimulation(v bool) *CreateCampaignShrinkRequest
	GetSimulation() *bool
	SetSimulationParameters(v string) *CreateCampaignShrinkRequest
	GetSimulationParameters() *string
	SetStartTime(v string) *CreateCampaignShrinkRequest
	GetStartTime() *string
	SetStrategyParameters(v string) *CreateCampaignShrinkRequest
	GetStrategyParameters() *string
	SetStrategyType(v string) *CreateCampaignShrinkRequest
	GetStrategyType() *string
}

type CreateCampaignShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"beginTime":"00:00:00","endTime":"23:00:00" }]
	CallableTime *string `json:"CallableTime,omitempty" xml:"CallableTime,omitempty"`
	// example:
	//
	// ccc-test/namelist.csv
	CaseFileKey    *string `json:"CaseFileKey,omitempty" xml:"CaseFileKey,omitempty"`
	CaseListShrink *string `json:"CaseList,omitempty" xml:"CaseList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c1f2bc75-422e-43c7-9c9d9d95633a
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1634313600000
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutingUntilTimeout *bool   `json:"ExecutingUntilTimeout,omitempty" xml:"ExecutingUntilTimeout,omitempty"`
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
	MaxAttemptCount *int64 `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MinAttemptInterval *int64 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-campaign
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	QueueId              *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	Simulation           *bool   `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	SimulationParameters *string `json:"SimulationParameters,omitempty" xml:"SimulationParameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1634140800000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ratio":1}
	StrategyParameters *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACING
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s CreateCampaignShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCampaignShrinkRequest) GetCallableTime() *string {
	return s.CallableTime
}

func (s *CreateCampaignShrinkRequest) GetCaseFileKey() *string {
	return s.CaseFileKey
}

func (s *CreateCampaignShrinkRequest) GetCaseListShrink() *string {
	return s.CaseListShrink
}

func (s *CreateCampaignShrinkRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *CreateCampaignShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateCampaignShrinkRequest) GetExecutingUntilTimeout() *bool {
	return s.ExecutingUntilTimeout
}

func (s *CreateCampaignShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCampaignShrinkRequest) GetMaxAttemptCount() *int64 {
	return s.MaxAttemptCount
}

func (s *CreateCampaignShrinkRequest) GetMinAttemptInterval() *int64 {
	return s.MinAttemptInterval
}

func (s *CreateCampaignShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCampaignShrinkRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *CreateCampaignShrinkRequest) GetSimulation() *bool {
	return s.Simulation
}

func (s *CreateCampaignShrinkRequest) GetSimulationParameters() *string {
	return s.SimulationParameters
}

func (s *CreateCampaignShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateCampaignShrinkRequest) GetStrategyParameters() *string {
	return s.StrategyParameters
}

func (s *CreateCampaignShrinkRequest) GetStrategyType() *string {
	return s.StrategyType
}

func (s *CreateCampaignShrinkRequest) SetCallableTime(v string) *CreateCampaignShrinkRequest {
	s.CallableTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetCaseFileKey(v string) *CreateCampaignShrinkRequest {
	s.CaseFileKey = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetCaseListShrink(v string) *CreateCampaignShrinkRequest {
	s.CaseListShrink = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetContactFlowId(v string) *CreateCampaignShrinkRequest {
	s.ContactFlowId = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetEndTime(v string) *CreateCampaignShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetExecutingUntilTimeout(v bool) *CreateCampaignShrinkRequest {
	s.ExecutingUntilTimeout = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetInstanceId(v string) *CreateCampaignShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetMaxAttemptCount(v int64) *CreateCampaignShrinkRequest {
	s.MaxAttemptCount = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetMinAttemptInterval(v int64) *CreateCampaignShrinkRequest {
	s.MinAttemptInterval = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetName(v string) *CreateCampaignShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetQueueId(v string) *CreateCampaignShrinkRequest {
	s.QueueId = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetSimulation(v bool) *CreateCampaignShrinkRequest {
	s.Simulation = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetSimulationParameters(v string) *CreateCampaignShrinkRequest {
	s.SimulationParameters = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetStartTime(v string) *CreateCampaignShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetStrategyParameters(v string) *CreateCampaignShrinkRequest {
	s.StrategyParameters = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetStrategyType(v string) *CreateCampaignShrinkRequest {
	s.StrategyType = &v
	return s
}

func (s *CreateCampaignShrinkRequest) Validate() error {
	return dara.Validate(s)
}

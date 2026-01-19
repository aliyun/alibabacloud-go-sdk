// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableTime(v string) *CreateCampaignRequest
	GetCallableTime() *string
	SetCaseFileKey(v string) *CreateCampaignRequest
	GetCaseFileKey() *string
	SetCaseList(v []*CreateCampaignRequestCaseList) *CreateCampaignRequest
	GetCaseList() []*CreateCampaignRequestCaseList
	SetContactFlowId(v string) *CreateCampaignRequest
	GetContactFlowId() *string
	SetEndTime(v string) *CreateCampaignRequest
	GetEndTime() *string
	SetExecutingUntilTimeout(v bool) *CreateCampaignRequest
	GetExecutingUntilTimeout() *bool
	SetFlashSmsParameters(v string) *CreateCampaignRequest
	GetFlashSmsParameters() *string
	SetInstGroupId(v string) *CreateCampaignRequest
	GetInstGroupId() *string
	SetInstanceId(v string) *CreateCampaignRequest
	GetInstanceId() *string
	SetMaxAttemptCount(v int64) *CreateCampaignRequest
	GetMaxAttemptCount() *int64
	SetMinAttemptInterval(v int64) *CreateCampaignRequest
	GetMinAttemptInterval() *int64
	SetName(v string) *CreateCampaignRequest
	GetName() *string
	SetNumberList(v []*string) *CreateCampaignRequest
	GetNumberList() []*string
	SetQueueId(v string) *CreateCampaignRequest
	GetQueueId() *string
	SetSimulation(v bool) *CreateCampaignRequest
	GetSimulation() *bool
	SetSimulationParameters(v string) *CreateCampaignRequest
	GetSimulationParameters() *string
	SetStartTime(v string) *CreateCampaignRequest
	GetStartTime() *string
	SetStrategyParameters(v string) *CreateCampaignRequest
	GetStrategyParameters() *string
	SetStrategyType(v string) *CreateCampaignRequest
	GetStrategyType() *string
}

type CreateCampaignRequest struct {
	CallableTime          *string                          `json:"CallableTime,omitempty" xml:"CallableTime,omitempty"`
	CaseFileKey           *string                          `json:"CaseFileKey,omitempty" xml:"CaseFileKey,omitempty"`
	CaseList              []*CreateCampaignRequestCaseList `json:"CaseList,omitempty" xml:"CaseList,omitempty" type:"Repeated"`
	ContactFlowId         *string                          `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	EndTime               *string                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutingUntilTimeout *bool                            `json:"ExecutingUntilTimeout,omitempty" xml:"ExecutingUntilTimeout,omitempty"`
	FlashSmsParameters    *string                          `json:"FlashSmsParameters,omitempty" xml:"FlashSmsParameters,omitempty"`
	InstGroupId           *string                          `json:"InstGroupId,omitempty" xml:"InstGroupId,omitempty"`
	InstanceId            *string                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxAttemptCount       *int64                           `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	MinAttemptInterval    *int64                           `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	Name                  *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	NumberList            []*string                        `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
	QueueId               *string                          `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	Simulation            *bool                            `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	SimulationParameters  *string                          `json:"SimulationParameters,omitempty" xml:"SimulationParameters,omitempty"`
	StartTime             *string                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StrategyParameters    *string                          `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	StrategyType          *string                          `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s CreateCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignRequest) GoString() string {
	return s.String()
}

func (s *CreateCampaignRequest) GetCallableTime() *string {
	return s.CallableTime
}

func (s *CreateCampaignRequest) GetCaseFileKey() *string {
	return s.CaseFileKey
}

func (s *CreateCampaignRequest) GetCaseList() []*CreateCampaignRequestCaseList {
	return s.CaseList
}

func (s *CreateCampaignRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *CreateCampaignRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateCampaignRequest) GetExecutingUntilTimeout() *bool {
	return s.ExecutingUntilTimeout
}

func (s *CreateCampaignRequest) GetFlashSmsParameters() *string {
	return s.FlashSmsParameters
}

func (s *CreateCampaignRequest) GetInstGroupId() *string {
	return s.InstGroupId
}

func (s *CreateCampaignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCampaignRequest) GetMaxAttemptCount() *int64 {
	return s.MaxAttemptCount
}

func (s *CreateCampaignRequest) GetMinAttemptInterval() *int64 {
	return s.MinAttemptInterval
}

func (s *CreateCampaignRequest) GetName() *string {
	return s.Name
}

func (s *CreateCampaignRequest) GetNumberList() []*string {
	return s.NumberList
}

func (s *CreateCampaignRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *CreateCampaignRequest) GetSimulation() *bool {
	return s.Simulation
}

func (s *CreateCampaignRequest) GetSimulationParameters() *string {
	return s.SimulationParameters
}

func (s *CreateCampaignRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateCampaignRequest) GetStrategyParameters() *string {
	return s.StrategyParameters
}

func (s *CreateCampaignRequest) GetStrategyType() *string {
	return s.StrategyType
}

func (s *CreateCampaignRequest) SetCallableTime(v string) *CreateCampaignRequest {
	s.CallableTime = &v
	return s
}

func (s *CreateCampaignRequest) SetCaseFileKey(v string) *CreateCampaignRequest {
	s.CaseFileKey = &v
	return s
}

func (s *CreateCampaignRequest) SetCaseList(v []*CreateCampaignRequestCaseList) *CreateCampaignRequest {
	s.CaseList = v
	return s
}

func (s *CreateCampaignRequest) SetContactFlowId(v string) *CreateCampaignRequest {
	s.ContactFlowId = &v
	return s
}

func (s *CreateCampaignRequest) SetEndTime(v string) *CreateCampaignRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCampaignRequest) SetExecutingUntilTimeout(v bool) *CreateCampaignRequest {
	s.ExecutingUntilTimeout = &v
	return s
}

func (s *CreateCampaignRequest) SetFlashSmsParameters(v string) *CreateCampaignRequest {
	s.FlashSmsParameters = &v
	return s
}

func (s *CreateCampaignRequest) SetInstGroupId(v string) *CreateCampaignRequest {
	s.InstGroupId = &v
	return s
}

func (s *CreateCampaignRequest) SetInstanceId(v string) *CreateCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCampaignRequest) SetMaxAttemptCount(v int64) *CreateCampaignRequest {
	s.MaxAttemptCount = &v
	return s
}

func (s *CreateCampaignRequest) SetMinAttemptInterval(v int64) *CreateCampaignRequest {
	s.MinAttemptInterval = &v
	return s
}

func (s *CreateCampaignRequest) SetName(v string) *CreateCampaignRequest {
	s.Name = &v
	return s
}

func (s *CreateCampaignRequest) SetNumberList(v []*string) *CreateCampaignRequest {
	s.NumberList = v
	return s
}

func (s *CreateCampaignRequest) SetQueueId(v string) *CreateCampaignRequest {
	s.QueueId = &v
	return s
}

func (s *CreateCampaignRequest) SetSimulation(v bool) *CreateCampaignRequest {
	s.Simulation = &v
	return s
}

func (s *CreateCampaignRequest) SetSimulationParameters(v string) *CreateCampaignRequest {
	s.SimulationParameters = &v
	return s
}

func (s *CreateCampaignRequest) SetStartTime(v string) *CreateCampaignRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCampaignRequest) SetStrategyParameters(v string) *CreateCampaignRequest {
	s.StrategyParameters = &v
	return s
}

func (s *CreateCampaignRequest) SetStrategyType(v string) *CreateCampaignRequest {
	s.StrategyType = &v
	return s
}

func (s *CreateCampaignRequest) Validate() error {
	if s.CaseList != nil {
		for _, item := range s.CaseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCampaignRequestCaseList struct {
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	PhoneNumber     *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ReferenceId     *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
}

func (s CreateCampaignRequestCaseList) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignRequestCaseList) GoString() string {
	return s.String()
}

func (s *CreateCampaignRequestCaseList) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *CreateCampaignRequestCaseList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateCampaignRequestCaseList) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *CreateCampaignRequestCaseList) SetCustomVariables(v string) *CreateCampaignRequestCaseList {
	s.CustomVariables = &v
	return s
}

func (s *CreateCampaignRequestCaseList) SetPhoneNumber(v string) *CreateCampaignRequestCaseList {
	s.PhoneNumber = &v
	return s
}

func (s *CreateCampaignRequestCaseList) SetReferenceId(v string) *CreateCampaignRequestCaseList {
	s.ReferenceId = &v
	return s
}

func (s *CreateCampaignRequestCaseList) Validate() error {
	return dara.Validate(s)
}

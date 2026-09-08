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
	SetFlashSmsParameters(v string) *CreateCampaignShrinkRequest
	GetFlashSmsParameters() *string
	SetInstGroupId(v string) *CreateCampaignShrinkRequest
	GetInstGroupId() *string
	SetInstanceId(v string) *CreateCampaignShrinkRequest
	GetInstanceId() *string
	SetMaxAttemptCount(v int64) *CreateCampaignShrinkRequest
	GetMaxAttemptCount() *int64
	SetMinAttemptInterval(v int64) *CreateCampaignShrinkRequest
	GetMinAttemptInterval() *int64
	SetName(v string) *CreateCampaignShrinkRequest
	GetName() *string
	SetNumberListShrink(v string) *CreateCampaignShrinkRequest
	GetNumberListShrink() *string
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
	// The callable time window for the predictive outbound dialing activity, formatted as a JSON object containing two properties: beginTime and endTime. Example: [{"beginTime":"00:00:00","endTime":"23:00:00"}].
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"beginTime":"00:00:00","endTime":"23:00:00" }]
	CallableTime *string `json:"CallableTime,omitempty" xml:"CallableTime,omitempty"`
	// Predictive outbound dialing contact file, specified as the key of an OSS object. Obtain this key by calling the GetCaseFileUploadUrl API.
	//
	// example:
	//
	// ccc-test/namelist.csv
	CaseFileKey *string `json:"CaseFileKey,omitempty" xml:"CaseFileKey,omitempty"`
	// List of predictive outbound dialing contacts. This parameter cannot be used together with CaseFileKey (import from file). You must choose either file import or list import.
	CaseListShrink *string `json:"CaseList,omitempty" xml:"CaseList,omitempty"`
	// The contact flow ID associated with the predictive outbound dialing activity.
	//
	// This parameter is required.
	//
	// example:
	//
	// c1f2bc75-422e-43c7-9c9d9d95633a
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// The end time of the predictive outbound calling activity, formatted as a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1634313600000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Whether to keep the activity in the executing state until it expires. The default value is false. If false, the activity automatically transitions to the completed state after all contacts have been called. If true, the activity remains in the executing state even after all contacts have been called, allowing you to append additional contacts and continue dialing until the activity expires or is manually stopped.
	//
	// example:
	//
	// false
	ExecutingUntilTimeout *bool `json:"ExecutingUntilTimeout,omitempty" xml:"ExecutingUntilTimeout,omitempty"`
	// Flash SMS parameters
	//
	// example:
	//
	// {"applicationId":"08e6b63a-****-****-****-689a288cdbb5","templateId":"325"}
	FlashSmsParameters *string `json:"FlashSmsParameters,omitempty" xml:"FlashSmsParameters,omitempty"`
	// Phone number collection ID
	//
	// example:
	//
	// 0d368091-2c70-4d26-979a-6997ddc9c34f
	InstGroupId *string `json:"InstGroupId,omitempty" xml:"InstGroupId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of attempts for the predictive outbound calling activity. This specifies how many times a number can be redialed if the initial call fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxAttemptCount *int64 `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	// The minimum redial interval for the predictive outbound calling activity, which specifies the minimum time interval between redial attempts after a failed call, in minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MinAttemptInterval *int64 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// Name of the predictive outbound dialing activity.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-campaign
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of caller numbers
	NumberListShrink *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	// The skill group ID associated with the predictive outbound dialing activity.
	//
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// Indicates whether this is a simulation activity used for testing. Regular customers do not need to concern themselves with this.
	//
	// example:
	//
	// 无
	Simulation *bool `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	// Simulation parameters used for testing. Regular customers do not need to concern themselves with this.
	//
	// example:
	//
	// 无
	SimulationParameters *string `json:"SimulationParameters,omitempty" xml:"SimulationParameters,omitempty"`
	// The start time of the predictive outbound dialing activity, in Unix timestamp format with millisecond precision.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1634140800000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Strategy parameters for the predictive outbound dialing activity. For PID strategy, an example format is: {"abandonRate":"5","historicalConnectedRate":"35"}. For PACING strategy, an example format is: {"ratio":1}. abandonRate represents the desired abandonment rate, historicalConnectedRate represents the historical reference connection rate, and ratio represents the fixed dialing ratio.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ratio":1}
	StrategyParameters *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	// The strategy pattern for the predictive outbound calling activity.
	//
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

func (s *CreateCampaignShrinkRequest) GetFlashSmsParameters() *string {
	return s.FlashSmsParameters
}

func (s *CreateCampaignShrinkRequest) GetInstGroupId() *string {
	return s.InstGroupId
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

func (s *CreateCampaignShrinkRequest) GetNumberListShrink() *string {
	return s.NumberListShrink
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

func (s *CreateCampaignShrinkRequest) SetFlashSmsParameters(v string) *CreateCampaignShrinkRequest {
	s.FlashSmsParameters = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetInstGroupId(v string) *CreateCampaignShrinkRequest {
	s.InstGroupId = &v
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

func (s *CreateCampaignShrinkRequest) SetNumberListShrink(v string) *CreateCampaignShrinkRequest {
	s.NumberListShrink = &v
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

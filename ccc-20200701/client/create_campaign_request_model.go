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
	CaseList []*CreateCampaignRequestCaseList `json:"CaseList,omitempty" xml:"CaseList,omitempty" type:"Repeated"`
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
	NumberList []*string `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
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
	// Customer-defined custom variables in JSON object format. The object can contain up to 10 properties, each with a name and value defined by the customer.
	//
	// example:
	//
	// {"name":"customer","客户标签":"tag"}
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// Contact phone number.
	//
	// example:
	//
	// 1888888888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Business ID, an identifier from the Customer\\"s Operational System, used in integration scenarios.
	//
	// example:
	//
	// 01
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
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

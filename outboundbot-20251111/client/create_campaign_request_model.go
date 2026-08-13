// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttemptOrder(v string) *CreateCampaignRequest
	GetAttemptOrder() *string
	SetCallableTime(v string) *CreateCampaignRequest
	GetCallableTime() *string
	SetCaseFileKey(v string) *CreateCampaignRequest
	GetCaseFileKey() *string
	SetCases(v []*CreateCampaignRequestCases) *CreateCampaignRequest
	GetCases() []*CreateCampaignRequestCases
	SetDialingTimeoutSeconds(v int32) *CreateCampaignRequest
	GetDialingTimeoutSeconds() *int32
	SetEndTime(v int64) *CreateCampaignRequest
	GetEndTime() *int64
	SetFixedQuota(v int32) *CreateCampaignRequest
	GetFixedQuota() *int32
	SetFlashSmsParameters(v string) *CreateCampaignRequest
	GetFlashSmsParameters() *string
	SetHolidayRestricted(v bool) *CreateCampaignRequest
	GetHolidayRestricted() *bool
	SetInstanceId(v string) *CreateCampaignRequest
	GetInstanceId() *string
	SetMaxAttemptCount(v int32) *CreateCampaignRequest
	GetMaxAttemptCount() *int32
	SetMinAttemptInterval(v int32) *CreateCampaignRequest
	GetMinAttemptInterval() *int32
	SetName(v string) *CreateCampaignRequest
	GetName() *string
	SetNumbers(v []*string) *CreateCampaignRequest
	GetNumbers() []*string
	SetRedialRestrictions(v string) *CreateCampaignRequest
	GetRedialRestrictions() *string
	SetRunUntilEndTime(v bool) *CreateCampaignRequest
	GetRunUntilEndTime() *bool
	SetScriptId(v string) *CreateCampaignRequest
	GetScriptId() *string
	SetStartTime(v int64) *CreateCampaignRequest
	GetStartTime() *int64
	SetWeight(v int32) *CreateCampaignRequest
	GetWeight() *int32
}

type CreateCampaignRequest struct {
	// The call execution order. Default value: MIN_ATTEMPT_FIRST. Valid values:
	//
	// - PRIORITY_FIRST: priority first.
	//
	// - MIN_ATTEMPT_FIRST: minimum attempt count first.
	//
	// example:
	//
	// MIN_ATTEMPT_FIRST
	AttemptOrder *string `json:"AttemptOrder,omitempty" xml:"AttemptOrder,omitempty"`
	// The callable time range for the task. The value is a JSON object that contains two properties: beginTime and EndTime.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"beginTime":"09:00:00","endTime":"18:00:00" }]。
	CallableTime *string `json:"CallableTime,omitempty" xml:"CallableTime,omitempty"`
	// The task contact list, which is an OSS object key obtained through the GenerateFileUploadParams operation. You can also leave this parameter empty and append contacts later through the AppendCases operation.
	//
	// example:
	//
	// cases/customer.csv
	CaseFileKey *string `json:"CaseFileKey,omitempty" xml:"CaseFileKey,omitempty"`
	// The contact list. You can also leave this parameter empty and append contacts later through the AppendCases operation.
	Cases []*CreateCampaignRequestCases `json:"Cases,omitempty" xml:"Cases,omitempty" type:"Repeated"`
	// The dialing timeout period, in seconds. Default value: 25.
	//
	// example:
	//
	// 25
	DialingTimeoutSeconds *int32 `json:"DialingTimeoutSeconds,omitempty" xml:"DialingTimeoutSeconds,omitempty"`
	// The task end time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1579965079000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The minimum concurrency for the task. A value of 0 indicates no guaranteed minimum, and resources are allocated by weight.
	//
	// If multiple tasks have a minimum concurrency configured:
	//
	// - If the total concurrency is less than the instance total concurrency, the minimum concurrency of each task is satisfied first, and the remaining resources are allocated proportionally by weight.
	//
	// - If the total concurrency exceeds the instance total concurrency, the minimum concurrency no longer serves as a guaranteed minimum but is used as a weight factor in the calculation.
	//
	// example:
	//
	// 0
	FixedQuota *int32 `json:"FixedQuota,omitempty" xml:"FixedQuota,omitempty"`
	// The flash SMS parameters.
	//
	// example:
	//
	// {}
	FlashSmsParameters *string `json:"FlashSmsParameters,omitempty" xml:"FlashSmsParameters,omitempty"`
	// Specifies whether to prohibit outbound calls on holidays.
	//
	// example:
	//
	// false
	HolidayRestricted *bool `json:"HolidayRestricted,omitempty" xml:"HolidayRestricted,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of attempts. This specifies the maximum number of times a number is called when the call fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MaxAttemptCount *int32 `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	// The interval between attempts.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	MinAttemptInterval *int32 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Updated_task_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of caller numbers for the outbound task.
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	// The list of redial restriction conditions. If this parameter is not specified, no restrictions are applied. Valid values:
	//
	// - CALLEE_NOT_EXISTS: Do not call nonexistent numbers.
	//
	// - OUT_OF_SERVICE: Do not call numbers that are out of service.
	//
	// example:
	//
	// None
	RedialRestrictions *string `json:"RedialRestrictions,omitempty" xml:"RedialRestrictions,omitempty"`
	// Specifies whether to keep the scheduling state until the task end time after all contacts are called. Default value: false. Valid values:
	//
	// - true: The task remains in the scheduling state, and you can continue to append contacts.
	//
	// - false: The task changes to completed, and you cannot append contacts.
	//
	// example:
	//
	// false
	RunUntilEndTime *bool `json:"RunUntilEndTime,omitempty" xml:"RunUntilEndTime,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aa279896-64a6-4182-864c-4f2b04ec8d17
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The task start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1578965079000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The weight. The value is an integer in the range of 0 to 100. A larger value indicates more concurrency allocated during scheduling.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignRequest) GoString() string {
	return s.String()
}

func (s *CreateCampaignRequest) GetAttemptOrder() *string {
	return s.AttemptOrder
}

func (s *CreateCampaignRequest) GetCallableTime() *string {
	return s.CallableTime
}

func (s *CreateCampaignRequest) GetCaseFileKey() *string {
	return s.CaseFileKey
}

func (s *CreateCampaignRequest) GetCases() []*CreateCampaignRequestCases {
	return s.Cases
}

func (s *CreateCampaignRequest) GetDialingTimeoutSeconds() *int32 {
	return s.DialingTimeoutSeconds
}

func (s *CreateCampaignRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateCampaignRequest) GetFixedQuota() *int32 {
	return s.FixedQuota
}

func (s *CreateCampaignRequest) GetFlashSmsParameters() *string {
	return s.FlashSmsParameters
}

func (s *CreateCampaignRequest) GetHolidayRestricted() *bool {
	return s.HolidayRestricted
}

func (s *CreateCampaignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCampaignRequest) GetMaxAttemptCount() *int32 {
	return s.MaxAttemptCount
}

func (s *CreateCampaignRequest) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *CreateCampaignRequest) GetName() *string {
	return s.Name
}

func (s *CreateCampaignRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *CreateCampaignRequest) GetRedialRestrictions() *string {
	return s.RedialRestrictions
}

func (s *CreateCampaignRequest) GetRunUntilEndTime() *bool {
	return s.RunUntilEndTime
}

func (s *CreateCampaignRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateCampaignRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateCampaignRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateCampaignRequest) SetAttemptOrder(v string) *CreateCampaignRequest {
	s.AttemptOrder = &v
	return s
}

func (s *CreateCampaignRequest) SetCallableTime(v string) *CreateCampaignRequest {
	s.CallableTime = &v
	return s
}

func (s *CreateCampaignRequest) SetCaseFileKey(v string) *CreateCampaignRequest {
	s.CaseFileKey = &v
	return s
}

func (s *CreateCampaignRequest) SetCases(v []*CreateCampaignRequestCases) *CreateCampaignRequest {
	s.Cases = v
	return s
}

func (s *CreateCampaignRequest) SetDialingTimeoutSeconds(v int32) *CreateCampaignRequest {
	s.DialingTimeoutSeconds = &v
	return s
}

func (s *CreateCampaignRequest) SetEndTime(v int64) *CreateCampaignRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCampaignRequest) SetFixedQuota(v int32) *CreateCampaignRequest {
	s.FixedQuota = &v
	return s
}

func (s *CreateCampaignRequest) SetFlashSmsParameters(v string) *CreateCampaignRequest {
	s.FlashSmsParameters = &v
	return s
}

func (s *CreateCampaignRequest) SetHolidayRestricted(v bool) *CreateCampaignRequest {
	s.HolidayRestricted = &v
	return s
}

func (s *CreateCampaignRequest) SetInstanceId(v string) *CreateCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCampaignRequest) SetMaxAttemptCount(v int32) *CreateCampaignRequest {
	s.MaxAttemptCount = &v
	return s
}

func (s *CreateCampaignRequest) SetMinAttemptInterval(v int32) *CreateCampaignRequest {
	s.MinAttemptInterval = &v
	return s
}

func (s *CreateCampaignRequest) SetName(v string) *CreateCampaignRequest {
	s.Name = &v
	return s
}

func (s *CreateCampaignRequest) SetNumbers(v []*string) *CreateCampaignRequest {
	s.Numbers = v
	return s
}

func (s *CreateCampaignRequest) SetRedialRestrictions(v string) *CreateCampaignRequest {
	s.RedialRestrictions = &v
	return s
}

func (s *CreateCampaignRequest) SetRunUntilEndTime(v bool) *CreateCampaignRequest {
	s.RunUntilEndTime = &v
	return s
}

func (s *CreateCampaignRequest) SetScriptId(v string) *CreateCampaignRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateCampaignRequest) SetStartTime(v int64) *CreateCampaignRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCampaignRequest) SetWeight(v int32) *CreateCampaignRequest {
	s.Weight = &v
	return s
}

func (s *CreateCampaignRequest) Validate() error {
	if s.Cases != nil {
		for _, item := range s.Cases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCampaignRequestCases struct {
	// The custom variables defined by the customer. The value is a JSON object that contains up to 10 properties. The name and value of each property are defined by the customer.
	//
	// example:
	//
	// {"key1":"value1"}
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// 133********
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The priority of the contact. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The business ID of the contact.
	//
	// example:
	//
	// bizId-1
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
}

func (s CreateCampaignRequestCases) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignRequestCases) GoString() string {
	return s.String()
}

func (s *CreateCampaignRequestCases) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *CreateCampaignRequestCases) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateCampaignRequestCases) GetPriority() *string {
	return s.Priority
}

func (s *CreateCampaignRequestCases) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *CreateCampaignRequestCases) SetCustomVariables(v string) *CreateCampaignRequestCases {
	s.CustomVariables = &v
	return s
}

func (s *CreateCampaignRequestCases) SetPhoneNumber(v string) *CreateCampaignRequestCases {
	s.PhoneNumber = &v
	return s
}

func (s *CreateCampaignRequestCases) SetPriority(v string) *CreateCampaignRequestCases {
	s.Priority = &v
	return s
}

func (s *CreateCampaignRequestCases) SetReferenceId(v string) *CreateCampaignRequestCases {
	s.ReferenceId = &v
	return s
}

func (s *CreateCampaignRequestCases) Validate() error {
	return dara.Validate(s)
}

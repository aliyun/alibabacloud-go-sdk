// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCampaignShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttemptOrder(v string) *CreateCampaignShrinkRequest
	GetAttemptOrder() *string
	SetCallableTime(v string) *CreateCampaignShrinkRequest
	GetCallableTime() *string
	SetCaseFileKey(v string) *CreateCampaignShrinkRequest
	GetCaseFileKey() *string
	SetCasesShrink(v string) *CreateCampaignShrinkRequest
	GetCasesShrink() *string
	SetDialingTimeoutSeconds(v int32) *CreateCampaignShrinkRequest
	GetDialingTimeoutSeconds() *int32
	SetEndTime(v int64) *CreateCampaignShrinkRequest
	GetEndTime() *int64
	SetFixedQuota(v int32) *CreateCampaignShrinkRequest
	GetFixedQuota() *int32
	SetFlashSmsParameters(v string) *CreateCampaignShrinkRequest
	GetFlashSmsParameters() *string
	SetHolidayRestricted(v bool) *CreateCampaignShrinkRequest
	GetHolidayRestricted() *bool
	SetInstanceId(v string) *CreateCampaignShrinkRequest
	GetInstanceId() *string
	SetMaxAttemptCount(v int32) *CreateCampaignShrinkRequest
	GetMaxAttemptCount() *int32
	SetMinAttemptInterval(v int32) *CreateCampaignShrinkRequest
	GetMinAttemptInterval() *int32
	SetName(v string) *CreateCampaignShrinkRequest
	GetName() *string
	SetNumbersShrink(v string) *CreateCampaignShrinkRequest
	GetNumbersShrink() *string
	SetRedialRestrictions(v string) *CreateCampaignShrinkRequest
	GetRedialRestrictions() *string
	SetRunUntilEndTime(v bool) *CreateCampaignShrinkRequest
	GetRunUntilEndTime() *bool
	SetScriptId(v string) *CreateCampaignShrinkRequest
	GetScriptId() *string
	SetStartTime(v int64) *CreateCampaignShrinkRequest
	GetStartTime() *int64
	SetWeight(v int32) *CreateCampaignShrinkRequest
	GetWeight() *int32
}

type CreateCampaignShrinkRequest struct {
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
	CasesShrink *string `json:"Cases,omitempty" xml:"Cases,omitempty"`
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
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
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

func (s CreateCampaignShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCampaignShrinkRequest) GetAttemptOrder() *string {
	return s.AttemptOrder
}

func (s *CreateCampaignShrinkRequest) GetCallableTime() *string {
	return s.CallableTime
}

func (s *CreateCampaignShrinkRequest) GetCaseFileKey() *string {
	return s.CaseFileKey
}

func (s *CreateCampaignShrinkRequest) GetCasesShrink() *string {
	return s.CasesShrink
}

func (s *CreateCampaignShrinkRequest) GetDialingTimeoutSeconds() *int32 {
	return s.DialingTimeoutSeconds
}

func (s *CreateCampaignShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateCampaignShrinkRequest) GetFixedQuota() *int32 {
	return s.FixedQuota
}

func (s *CreateCampaignShrinkRequest) GetFlashSmsParameters() *string {
	return s.FlashSmsParameters
}

func (s *CreateCampaignShrinkRequest) GetHolidayRestricted() *bool {
	return s.HolidayRestricted
}

func (s *CreateCampaignShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCampaignShrinkRequest) GetMaxAttemptCount() *int32 {
	return s.MaxAttemptCount
}

func (s *CreateCampaignShrinkRequest) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *CreateCampaignShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCampaignShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *CreateCampaignShrinkRequest) GetRedialRestrictions() *string {
	return s.RedialRestrictions
}

func (s *CreateCampaignShrinkRequest) GetRunUntilEndTime() *bool {
	return s.RunUntilEndTime
}

func (s *CreateCampaignShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateCampaignShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateCampaignShrinkRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateCampaignShrinkRequest) SetAttemptOrder(v string) *CreateCampaignShrinkRequest {
	s.AttemptOrder = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetCallableTime(v string) *CreateCampaignShrinkRequest {
	s.CallableTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetCaseFileKey(v string) *CreateCampaignShrinkRequest {
	s.CaseFileKey = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetCasesShrink(v string) *CreateCampaignShrinkRequest {
	s.CasesShrink = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetDialingTimeoutSeconds(v int32) *CreateCampaignShrinkRequest {
	s.DialingTimeoutSeconds = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetEndTime(v int64) *CreateCampaignShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetFixedQuota(v int32) *CreateCampaignShrinkRequest {
	s.FixedQuota = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetFlashSmsParameters(v string) *CreateCampaignShrinkRequest {
	s.FlashSmsParameters = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetHolidayRestricted(v bool) *CreateCampaignShrinkRequest {
	s.HolidayRestricted = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetInstanceId(v string) *CreateCampaignShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetMaxAttemptCount(v int32) *CreateCampaignShrinkRequest {
	s.MaxAttemptCount = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetMinAttemptInterval(v int32) *CreateCampaignShrinkRequest {
	s.MinAttemptInterval = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetName(v string) *CreateCampaignShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetNumbersShrink(v string) *CreateCampaignShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetRedialRestrictions(v string) *CreateCampaignShrinkRequest {
	s.RedialRestrictions = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetRunUntilEndTime(v bool) *CreateCampaignShrinkRequest {
	s.RunUntilEndTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetScriptId(v string) *CreateCampaignShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetStartTime(v int64) *CreateCampaignShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetWeight(v int32) *CreateCampaignShrinkRequest {
	s.Weight = &v
	return s
}

func (s *CreateCampaignShrinkRequest) Validate() error {
	return dara.Validate(s)
}

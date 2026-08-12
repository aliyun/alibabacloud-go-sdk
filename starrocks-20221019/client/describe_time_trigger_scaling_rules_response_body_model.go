// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTimeTriggerScalingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeTimeTriggerScalingRulesResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*DescribeTimeTriggerScalingRulesResponseBodyData) *DescribeTimeTriggerScalingRulesResponseBody
	GetData() []*DescribeTimeTriggerScalingRulesResponseBodyData
	SetErrCode(v string) *DescribeTimeTriggerScalingRulesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeTimeTriggerScalingRulesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeTimeTriggerScalingRulesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeTimeTriggerScalingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTimeTriggerScalingRulesResponseBody
	GetSuccess() *bool
}

type DescribeTimeTriggerScalingRulesResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data []*DescribeTimeTriggerScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code of the request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTimeTriggerScalingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimeTriggerScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) GetData() []*DescribeTimeTriggerScalingRulesResponseBodyData {
	return s.Data
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) SetAccessDeniedDetail(v string) *DescribeTimeTriggerScalingRulesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) SetData(v []*DescribeTimeTriggerScalingRulesResponseBodyData) *DescribeTimeTriggerScalingRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) SetErrCode(v string) *DescribeTimeTriggerScalingRulesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) SetErrMessage(v string) *DescribeTimeTriggerScalingRulesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) SetHttpStatusCode(v int32) *DescribeTimeTriggerScalingRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) SetRequestId(v string) *DescribeTimeTriggerScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) SetSuccess(v bool) *DescribeTimeTriggerScalingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTimeTriggerScalingRulesResponseBodyData struct {
	// The number of elastic nodes.
	//
	// example:
	//
	// 3
	NodeNumber *string `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	// The scale-in rule.
	ScalingInRule *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule `json:"ScalingInRule,omitempty" xml:"ScalingInRule,omitempty" type:"Struct"`
	// The scale-out rule.
	ScalingOutRule *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule `json:"ScalingOutRule,omitempty" xml:"ScalingOutRule,omitempty" type:"Struct"`
	// The ID of the scaling rule.
	//
	// example:
	//
	// r-d1775b776110****
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	// The name of the Auto Scaling policy.
	//
	// example:
	//
	// scale-test1
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// The current status of the rule. Valid values:
	//
	// - INACTIVE: The rule is not triggered.
	//
	// - ACTIVE: The rule is triggered.
	//
	// - DISABLED: The rule is disabled due to an overdue payment.
	//
	// example:
	//
	// INACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTimeTriggerScalingRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimeTriggerScalingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) GetNodeNumber() *string {
	return s.NodeNumber
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) GetScalingInRule() *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	return s.ScalingInRule
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) GetScalingOutRule() *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	return s.ScalingOutRule
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) GetScalingRuleId() *string {
	return s.ScalingRuleId
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) SetNodeNumber(v string) *DescribeTimeTriggerScalingRulesResponseBodyData {
	s.NodeNumber = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) SetScalingInRule(v *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) *DescribeTimeTriggerScalingRulesResponseBodyData {
	s.ScalingInRule = v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) SetScalingOutRule(v *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) *DescribeTimeTriggerScalingRulesResponseBodyData {
	s.ScalingOutRule = v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) SetScalingRuleId(v string) *DescribeTimeTriggerScalingRulesResponseBodyData {
	s.ScalingRuleId = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) SetScalingRuleName(v string) *DescribeTimeTriggerScalingRulesResponseBodyData {
	s.ScalingRuleName = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) SetStatus(v string) *DescribeTimeTriggerScalingRulesResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyData) Validate() error {
	if s.ScalingInRule != nil {
		if err := s.ScalingInRule.Validate(); err != nil {
			return err
		}
	}
	if s.ScalingOutRule != nil {
		if err := s.ScalingOutRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule struct {
	// The day of the execution time.
	//
	// example:
	//
	// 1
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// The hour of the execution time.
	//
	// example:
	//
	// 12
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// The minute of the execution time.
	//
	// example:
	//
	// 24
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// The month of the execution time.
	//
	// example:
	//
	// 3
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// The trigger interval. Its meaning depends on the RecurrenceType value:
	//
	// - DAILY: The interval in days.
	//
	// - WEEKLY: The interval in weeks.
	//
	// - MONTHLY: The interval in months.
	//
	// example:
	//
	// WEEKLY
	RecurrenceInterval *int32 `json:"RecurrenceInterval,omitempty" xml:"RecurrenceInterval,omitempty"`
	// The trigger method. Valid values:
	//
	// - ONCE: One-time
	//
	// - DAILY: Daily
	//
	// - WEEKLY: Weekly
	//
	// - MONTHLY: Monthly
	//
	// example:
	//
	// ONCE
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The trigger values. Their meaning depends on the RecurrenceType value:
	//
	// -WEEKLY: The days of the week for execution.
	//
	// -MONTHLY: The days of the month for execution.
	RecurrenceValues []*string `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// The second of the execution time.
	//
	// example:
	//
	// 30
	Second *int32 `json:"Second,omitempty" xml:"Second,omitempty"`
	// The year of the execution time.
	//
	// example:
	//
	// 2025
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GoString() string {
	return s.String()
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetDay() *int32 {
	return s.Day
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetHour() *int32 {
	return s.Hour
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetMinute() *int32 {
	return s.Minute
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetMonth() *int32 {
	return s.Month
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetRecurrenceInterval() *int32 {
	return s.RecurrenceInterval
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetRecurrenceValues() []*string {
	return s.RecurrenceValues
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetSecond() *int32 {
	return s.Second
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) GetYear() *int32 {
	return s.Year
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetDay(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.Day = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetHour(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.Hour = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetMinute(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.Minute = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetMonth(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.Month = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetRecurrenceInterval(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.RecurrenceInterval = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetRecurrenceType(v string) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetRecurrenceValues(v []*string) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.RecurrenceValues = v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetSecond(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.Second = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) SetYear(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule {
	s.Year = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingInRule) Validate() error {
	return dara.Validate(s)
}

type DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule struct {
	// The day of the execution time.
	//
	// example:
	//
	// 10
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// The hour of the execution time.
	//
	// example:
	//
	// 3
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// The minute of the execution time.
	//
	// example:
	//
	// 30
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// The month of the execution time.
	//
	// example:
	//
	// 12
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// The trigger interval. Its meaning depends on the RecurrenceType value:
	//
	// - DAILY: The interval in days.
	//
	// - WEEKLY: The interval in weeks.
	//
	// - MONTHLY: The interval in months.
	//
	// example:
	//
	// 1
	RecurrenceInterval *int32 `json:"RecurrenceInterval,omitempty" xml:"RecurrenceInterval,omitempty"`
	// The trigger method. Valid values:
	//
	// - ONCE: One-time
	//
	// - DAILY: Daily
	//
	// - WEEKLY: Weekly
	//
	// - MONTHLY: Monthly
	//
	// example:
	//
	// ONCE
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The trigger values. Their meaning depends on the RecurrenceType value:
	//
	// -WEEKLY: The days of the week for execution.
	//
	// -MONTHLY: The days of the month for execution.
	RecurrenceValues []*string `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// The second of the execution time.
	//
	// example:
	//
	// 20
	Second *int32 `json:"Second,omitempty" xml:"Second,omitempty"`
	// The year of the execution time.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GoString() string {
	return s.String()
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetDay() *int32 {
	return s.Day
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetHour() *int32 {
	return s.Hour
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetMinute() *int32 {
	return s.Minute
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetMonth() *int32 {
	return s.Month
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetRecurrenceInterval() *int32 {
	return s.RecurrenceInterval
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetRecurrenceValues() []*string {
	return s.RecurrenceValues
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetSecond() *int32 {
	return s.Second
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) GetYear() *int32 {
	return s.Year
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetDay(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.Day = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetHour(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.Hour = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetMinute(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.Minute = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetMonth(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.Month = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetRecurrenceInterval(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.RecurrenceInterval = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetRecurrenceType(v string) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetRecurrenceValues(v []*string) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.RecurrenceValues = v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetSecond(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.Second = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) SetYear(v int32) *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule {
	s.Year = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponseBodyDataScalingOutRule) Validate() error {
	return dara.Validate(s)
}

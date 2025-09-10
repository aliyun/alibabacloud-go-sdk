// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaAlarmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListQuotaAlarmsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaAlarmsResponseBody
	GetNextToken() *string
	SetQuotaAlarms(v []*ListQuotaAlarmsResponseBodyQuotaAlarms) *ListQuotaAlarmsResponseBody
	GetQuotaAlarms() []*ListQuotaAlarmsResponseBodyQuotaAlarms
	SetRequestId(v string) *ListQuotaAlarmsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListQuotaAlarmsResponseBody
	GetTotalCount() *int32
}

type ListQuotaAlarmsResponseBody struct {
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends.
	//
	// > An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details about the quota alert rules.
	QuotaAlarms []*ListQuotaAlarmsResponseBodyQuotaAlarms `json:"QuotaAlarms,omitempty" xml:"QuotaAlarms,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 87F3B755-3BD2-4C76-B36A-93247002918C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of quota alerts.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaAlarmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaAlarmsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaAlarmsResponseBody) GetQuotaAlarms() []*ListQuotaAlarmsResponseBodyQuotaAlarms {
	return s.QuotaAlarms
}

func (s *ListQuotaAlarmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaAlarmsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotaAlarmsResponseBody) SetMaxResults(v int32) *ListQuotaAlarmsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetNextToken(v string) *ListQuotaAlarmsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetQuotaAlarms(v []*ListQuotaAlarmsResponseBodyQuotaAlarms) *ListQuotaAlarmsResponseBody {
	s.QuotaAlarms = v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetRequestId(v string) *ListQuotaAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetTotalCount(v int32) *ListQuotaAlarmsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQuotaAlarmsResponseBodyQuotaAlarms struct {
	// The ID of the alert.
	//
	// example:
	//
	// a2efa7fc-832f-47bb-8054-39e28012****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the alert event.
	//
	// example:
	//
	// rules
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert was created.
	//
	// example:
	//
	// 2020-11-27T07:23:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the alert threshold was reached. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	ExceedThreshold *bool `json:"ExceedThreshold,omitempty" xml:"ExceedThreshold,omitempty"`
	// The alert notification methods.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The alert contact.
	//
	// >  Valid value: accountContact. Only the account contact is supported.
	//
	// example:
	//
	// accountContact
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// config
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_hvnoqv
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimensions.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	QuotaDimensions map[string]interface{} `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 73
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The value of the quota.
	//
	// example:
	//
	// 200
	QuotaValue *float32 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The numeric value of the alert threshold.
	//
	// example:
	//
	// 160
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold.
	//
	// example:
	//
	// 80
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// 	- used: The alert is created for the used quota.
	//
	// 	- usable: The alert is created for the available quota.
	//
	// example:
	//
	// used
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The webhook URL.
	//
	// example:
	//
	// https://www.aliyun.com/webhook
	WebHook *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s ListQuotaAlarmsResponseBodyQuotaAlarms) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaAlarmsResponseBodyQuotaAlarms) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetAlarmId() *string {
	return s.AlarmId
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetAlarmName() *string {
	return s.AlarmName
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetExceedThreshold() *bool {
	return s.ExceedThreshold
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetNotifyTarget() *string {
	return s.NotifyTarget
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetQuotaDimensions() map[string]interface{} {
	return s.QuotaDimensions
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetQuotaUsage() *float32 {
	return s.QuotaUsage
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetQuotaValue() *float32 {
	return s.QuotaValue
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetThresholdPercent() *float32 {
	return s.ThresholdPercent
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetThresholdType() *string {
	return s.ThresholdType
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) GetWebHook() *string {
	return s.WebHook
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetAlarmId(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.AlarmId = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetAlarmName(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.AlarmName = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetCreateTime(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.CreateTime = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetExceedThreshold(v bool) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ExceedThreshold = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetNotifyChannels(v []*string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.NotifyChannels = v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetNotifyTarget(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.NotifyTarget = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetProductCode(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaActionCode(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaDimensions(v map[string]interface{}) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaDimensions = v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaUsage(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaUsage = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaValue(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaValue = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetThreshold(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.Threshold = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetThresholdPercent(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ThresholdPercent = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetThresholdType(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ThresholdType = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetWebHook(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.WebHook = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) Validate() error {
	return dara.Validate(s)
}

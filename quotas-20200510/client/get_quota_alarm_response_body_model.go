// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaAlarm(v *GetQuotaAlarmResponseBodyQuotaAlarm) *GetQuotaAlarmResponseBody
	GetQuotaAlarm() *GetQuotaAlarmResponseBodyQuotaAlarm
	SetRequestId(v string) *GetQuotaAlarmResponseBody
	GetRequestId() *string
}

type GetQuotaAlarmResponseBody struct {
	// The details of the quota alert.
	QuotaAlarm *GetQuotaAlarmResponseBodyQuotaAlarm `json:"QuotaAlarm,omitempty" xml:"QuotaAlarm,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 81B9D511-F3DD-43B1-9A81-1795DDB52ADF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQuotaAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmResponseBody) GetQuotaAlarm() *GetQuotaAlarmResponseBodyQuotaAlarm {
	return s.QuotaAlarm
}

func (s *GetQuotaAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaAlarmResponseBody) SetQuotaAlarm(v *GetQuotaAlarmResponseBodyQuotaAlarm) *GetQuotaAlarmResponseBody {
	s.QuotaAlarm = v
	return s
}

func (s *GetQuotaAlarmResponseBody) SetRequestId(v string) *GetQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQuotaAlarmResponseBodyQuotaAlarm struct {
	// The ID of the quota alert.
	//
	// example:
	//
	// 78d7e436-4b25-4897-84b5-d7b656bb****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the quota alert.
	//
	// example:
	//
	// tf-testacccn-hangzhouquotasquotaalarm81611
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert was created.
	//
	// example:
	//
	// 2021-01-21T03:47:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The alert notification methods.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The alert contact.
	//
	// example:
	//
	// accountContact
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimension.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	QuotaDimension map[string]interface{} `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 28
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The quota value.
	//
	// example:
	//
	// 804
	QuotaValue *float32 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The numeric value of the alert threshold.
	//
	// example:
	//
	// 29
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold.
	//
	// example:
	//
	// 50
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
	// The webhook URL. Quota Center sends alert notifications to the specified URL by using HTTP POST requests.
	//
	// example:
	//
	// https://alert.aliyun.com/callback
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s GetQuotaAlarmResponseBodyQuotaAlarm) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaAlarmResponseBodyQuotaAlarm) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetAlarmId() *string {
	return s.AlarmId
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetAlarmName() *string {
	return s.AlarmName
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetNotifyTarget() *string {
	return s.NotifyTarget
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetQuotaDimension() map[string]interface{} {
	return s.QuotaDimension
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetQuotaUsage() *float32 {
	return s.QuotaUsage
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetQuotaValue() *float32 {
	return s.QuotaValue
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetThreshold() *float32 {
	return s.Threshold
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetThresholdPercent() *float32 {
	return s.ThresholdPercent
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetThresholdType() *string {
	return s.ThresholdType
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) GetWebhook() *string {
	return s.Webhook
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetAlarmId(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.AlarmId = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetAlarmName(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.AlarmName = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetCreateTime(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetNotifyChannels(v []*string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.NotifyChannels = v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetNotifyTarget(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.NotifyTarget = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetProductCode(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.ProductCode = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaActionCode(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaActionCode = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaDimension(v map[string]interface{}) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaDimension = v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaUsage(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaUsage = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaValue(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaValue = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetThreshold(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.Threshold = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetThresholdPercent(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.ThresholdPercent = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetThresholdType(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.ThresholdType = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetWebhook(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.Webhook = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) Validate() error {
	return dara.Validate(s)
}

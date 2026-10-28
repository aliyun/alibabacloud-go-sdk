// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v string) *UpdateQuotaAlarmRequest
	GetAlarmId() *string
	SetAlarmName(v string) *UpdateQuotaAlarmRequest
	GetAlarmName() *string
	SetThreshold(v float32) *UpdateQuotaAlarmRequest
	GetThreshold() *float32
	SetThresholdPercent(v float32) *UpdateQuotaAlarmRequest
	GetThresholdPercent() *float32
	SetThresholdType(v string) *UpdateQuotaAlarmRequest
	GetThresholdType() *string
	SetWebHook(v string) *UpdateQuotaAlarmRequest
	GetWebHook() *string
}

type UpdateQuotaAlarmRequest struct {
	// The ID of the quota alert.
	//
	// > For more information about how to obtain the ID of a quota alert, see [ListQuotaAlarms](https://help.aliyun.com/document_detail/440561.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// a2efa7fc-832f-47bb-8054-39e28012****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the quota alert.
	//
	// > For more information about how to obtain the name of a quota alert, see [ListQuotaAlarms](https://help.aliyun.com/document_detail/440561.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rules
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The numeric value of the alert threshold. Valid values:
	//
	// - If you set `ThresholdType` to `used`, an alert is triggered when the used quota is greater than or equal to this value. The alert threshold must be greater than the current used quota.
	//
	// - If you set `ThresholdType` to `usable`, an alert is triggered when the usable quota is less than or equal to this value. The alert threshold must be less than the current usable quota.
	//
	// > You must set either Threshold or ThresholdPercent.
	//
	// example:
	//
	// 160
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold. Valid values:
	//
	// - If you set `ThresholdType` to `used`, an alert is triggered when the percentage of the used quota is greater than or equal to this value. The value must be in the range of (50%, 100%].
	//
	// - If you set `ThresholdType` to `usable`, an alert is triggered when the percentage of the usable quota is less than or equal to this value. The value must be in the range of (0%, 50%].
	//
	// > You must set either Threshold or ThresholdPercent.
	//
	// example:
	//
	// 51
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// - used (default): The alert is triggered based on the used quota.
	//
	// - usable: The alert is triggered based on the usable quota.
	//
	// example:
	//
	// usable
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The webhook URL. Quota Center sends alert notifications to this URL using an HTTP POST request.
	//
	// example:
	//
	// https://alert.aliyun.com/callback
	WebHook *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s UpdateQuotaAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaAlarmRequest) GetAlarmId() *string {
	return s.AlarmId
}

func (s *UpdateQuotaAlarmRequest) GetAlarmName() *string {
	return s.AlarmName
}

func (s *UpdateQuotaAlarmRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateQuotaAlarmRequest) GetThresholdPercent() *float32 {
	return s.ThresholdPercent
}

func (s *UpdateQuotaAlarmRequest) GetThresholdType() *string {
	return s.ThresholdType
}

func (s *UpdateQuotaAlarmRequest) GetWebHook() *string {
	return s.WebHook
}

func (s *UpdateQuotaAlarmRequest) SetAlarmId(v string) *UpdateQuotaAlarmRequest {
	s.AlarmId = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetAlarmName(v string) *UpdateQuotaAlarmRequest {
	s.AlarmName = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetThreshold(v float32) *UpdateQuotaAlarmRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetThresholdPercent(v float32) *UpdateQuotaAlarmRequest {
	s.ThresholdPercent = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetThresholdType(v string) *UpdateQuotaAlarmRequest {
	s.ThresholdType = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetWebHook(v string) *UpdateQuotaAlarmRequest {
	s.WebHook = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) Validate() error {
	return dara.Validate(s)
}

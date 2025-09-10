// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmName(v string) *CreateQuotaAlarmRequest
	GetAlarmName() *string
	SetOriginalContext(v string) *CreateQuotaAlarmRequest
	GetOriginalContext() *string
	SetProductCode(v string) *CreateQuotaAlarmRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *CreateQuotaAlarmRequest
	GetQuotaActionCode() *string
	SetQuotaDimensions(v []*CreateQuotaAlarmRequestQuotaDimensions) *CreateQuotaAlarmRequest
	GetQuotaDimensions() []*CreateQuotaAlarmRequestQuotaDimensions
	SetThreshold(v float32) *CreateQuotaAlarmRequest
	GetThreshold() *float32
	SetThresholdPercent(v float32) *CreateQuotaAlarmRequest
	GetThresholdPercent() *float32
	SetThresholdType(v string) *CreateQuotaAlarmRequest
	GetThresholdType() *string
	SetWebHook(v string) *CreateQuotaAlarmRequest
	GetWebHook() *string
}

type CreateQuotaAlarmRequest struct {
	// Quota alarm name
	//
	// This parameter is required.
	//
	// example:
	//
	// q_344t4 alarm
	AlarmName       *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	OriginalContext *string `json:"OriginalContext,omitempty" xml:"OriginalContext,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, check the `ProductCode` parameter that is described in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// config
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >  To obtain the quota ID of an Alibaba Cloud service, call the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation and check the value of `QuotaActionCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// q_hvnoqv
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimensions. A quota item is uniquely identified by the values of `Dimensions` and `QuotaActionCode`.
	//
	// >  This parameter is required for specific Alibaba Cloud services. You can call the [ListProductQuotaDimensions](https://help.aliyun.com/document_detail/440553.html) operation to query the quota dimensions that are supported by an Alibaba Cloud service. The value of `Requisite` in the response indicates whether a dimension is required.
	QuotaDimensions []*CreateQuotaAlarmRequestQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
	// The numerical value of the quota alarm threshold. Value settings:
	//
	// - When `ThresholdType` is set to `used`, an alarm notification will be sent if the used amount of the quota is greater than or equal to the preset value. The quota alarm threshold must be greater than the used amount of the quota.
	//
	// - When `ThresholdType` is set to `usable`, an alarm notification will be sent if the remaining available amount of the quota is less than or equal to the preset value. The quota alarm threshold must be less than the remaining available amount of the quota.
	//
	// > One of this parameter and ThresholdPercent must be set.
	//
	// example:
	//
	// 150
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the quota alert threshold. Values:
	//
	// - When `ThresholdType` is set to `used`, you will receive an alert notification if the used amount of the quota is greater than or equal to the preset percentage. The value range is (50%, 100%].
	//
	// - When `ThresholdType` is set to `usable`, you will receive an alert notification if the remaining available amount of the quota is less than or equal to the preset percentage. The value range is (0%, 50%].
	//
	// > One of this parameter and Threshold must be set.
	//
	// example:
	//
	// 50
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// 	- used (default): The alert is created for the used quota.
	//
	// 	- usable: The alert is created for the available quota.
	//
	// example:
	//
	// used
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The quota center sends alert information to the specified public URL address via a POST request using the HTTP protocol.
	//
	// example:
	//
	// https://alert.aliyun.com/callback
	WebHook *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s CreateQuotaAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmRequest) GetAlarmName() *string {
	return s.AlarmName
}

func (s *CreateQuotaAlarmRequest) GetOriginalContext() *string {
	return s.OriginalContext
}

func (s *CreateQuotaAlarmRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateQuotaAlarmRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *CreateQuotaAlarmRequest) GetQuotaDimensions() []*CreateQuotaAlarmRequestQuotaDimensions {
	return s.QuotaDimensions
}

func (s *CreateQuotaAlarmRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateQuotaAlarmRequest) GetThresholdPercent() *float32 {
	return s.ThresholdPercent
}

func (s *CreateQuotaAlarmRequest) GetThresholdType() *string {
	return s.ThresholdType
}

func (s *CreateQuotaAlarmRequest) GetWebHook() *string {
	return s.WebHook
}

func (s *CreateQuotaAlarmRequest) SetAlarmName(v string) *CreateQuotaAlarmRequest {
	s.AlarmName = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetOriginalContext(v string) *CreateQuotaAlarmRequest {
	s.OriginalContext = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetProductCode(v string) *CreateQuotaAlarmRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetQuotaActionCode(v string) *CreateQuotaAlarmRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetQuotaDimensions(v []*CreateQuotaAlarmRequestQuotaDimensions) *CreateQuotaAlarmRequest {
	s.QuotaDimensions = v
	return s
}

func (s *CreateQuotaAlarmRequest) SetThreshold(v float32) *CreateQuotaAlarmRequest {
	s.Threshold = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetThresholdPercent(v float32) *CreateQuotaAlarmRequest {
	s.ThresholdPercent = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetThresholdType(v string) *CreateQuotaAlarmRequest {
	s.ThresholdType = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetWebHook(v string) *CreateQuotaAlarmRequest {
	s.WebHook = &v
	return s
}

func (s *CreateQuotaAlarmRequest) Validate() error {
	return dara.Validate(s)
}

type CreateQuotaAlarmRequestQuotaDimensions struct {
	// The key of the dimension.
	//
	// >  You must configure `Dimensions.N.Key` and `Dimensions.N.Value` at the same time. The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service. You can call the [ListProductQuotaDimensions](https://help.aliyun.com/document_detail/440553.html) operation to query the dimensions that are supported by an Alibaba Cloud service. The number of items in the returned array is N.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >  You must configure `Dimensions.N.Key` and `Dimensions.N.Value` at the same time. The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service. You can call the [ListProductQuotaDimensions](https://help.aliyun.com/document_detail/440553.html) operation to query the dimensions that are supported by an Alibaba Cloud service. The number of items in the returned array is N.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQuotaAlarmRequestQuotaDimensions) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaAlarmRequestQuotaDimensions) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmRequestQuotaDimensions) GetKey() *string {
	return s.Key
}

func (s *CreateQuotaAlarmRequestQuotaDimensions) GetValue() *string {
	return s.Value
}

func (s *CreateQuotaAlarmRequestQuotaDimensions) SetKey(v string) *CreateQuotaAlarmRequestQuotaDimensions {
	s.Key = &v
	return s
}

func (s *CreateQuotaAlarmRequestQuotaDimensions) SetValue(v string) *CreateQuotaAlarmRequestQuotaDimensions {
	s.Value = &v
	return s
}

func (s *CreateQuotaAlarmRequestQuotaDimensions) Validate() error {
	return dara.Validate(s)
}

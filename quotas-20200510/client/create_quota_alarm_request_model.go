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
	// The name of the quota alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// q_hvnoqv alarm
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The original context information.
	//
	// example:
	//
	// {}
	OriginalContext *string `json:"OriginalContext,omitempty" xml:"OriginalContext,omitempty"`
	// The abbreviation of the cloud service.
	//
	// > For more information about how to obtain the abbreviation of a cloud service, see ProductCode in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// config
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// > For more information about how to obtain the quota ID of a cloud service, see QuotaActionCode in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// q_hvnoqv
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimensions.
	//
	// The dimensions are properties of a quota item. A quota item is uniquely identified by Dimensions and QuotaActionCode.
	//
	// > Some dimensions are required for specific cloud services. You can call the [ListProductQuotaDimensions](https://help.aliyun.com/document_detail/440553.html) operation to query the supported quota dimensions of a cloud service. The Requisite response parameter indicates whether a dimension is required.
	QuotaDimensions []*CreateQuotaAlarmRequestQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
	// The numeric value of the quota alert threshold. Valid values:
	//
	// - If ThresholdType is set to used, you receive an alert notification when the used quota is greater than or equal to the specified value. The quota alert threshold must be greater than the used quota.
	//
	// - If ThresholdType is set to usable, you receive an alert notification when the available quota is less than or equal to the specified value. The quota alert threshold must be less than the available quota.
	//
	// > You must specify either this parameter or ThresholdPercent.
	//
	// example:
	//
	// 150
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the quota alert threshold. Valid values:
	//
	// - If ThresholdType is set to used, you receive an alert notification when the used quota is greater than or equal to the specified percentage. Valid values: (50%, 100%].
	//
	// - If ThresholdType is set to usable, you receive an alert notification when the available quota is less than or equal to the specified percentage. Valid values: (0%, 50%].
	//
	// > You must specify either this parameter or Threshold.
	//
	// example:
	//
	// 50
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The quota alert type. Valid values:
	//
	// - used: alerts on the used quota. This is the default value.
	//
	// - usable: alerts on the remaining available quota.
	//
	// example:
	//
	// used
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The public URL to which Quota Center sends alert notifications by using HTTP POST requests.
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
	if s.QuotaDimensions != nil {
		for _, item := range s.QuotaDimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateQuotaAlarmRequestQuotaDimensions struct {
	// The key of the quota dimension.
	//
	// > You must specify both Dimensions.N.Key and Dimensions.N.Value. The value range of N depends on the number of dimensions supported by the cloud service. You can call the [ListProductQuotaDimensions](https://help.aliyun.com/document_detail/440553.html) operation to query the supported quota dimensions. The length of the returned array is the maximum value of N.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension.
	//
	//  > You must specify both Dimensions.N.Key and Dimensions.N.Value. The value range of N depends on the number of dimensions supported by the cloud service. You can call the [ListProductQuotaDimensions](https://help.aliyun.com/document_detail/440553.html) operation to query the supported quota dimensions. The length of the returned array is the maximum value of N.
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

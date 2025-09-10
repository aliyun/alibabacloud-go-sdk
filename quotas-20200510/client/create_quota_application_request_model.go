// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditMode(v string) *CreateQuotaApplicationRequest
	GetAuditMode() *string
	SetDesireValue(v float32) *CreateQuotaApplicationRequest
	GetDesireValue() *float32
	SetDimensions(v []*CreateQuotaApplicationRequestDimensions) *CreateQuotaApplicationRequest
	GetDimensions() []*CreateQuotaApplicationRequestDimensions
	SetEffectiveTime(v string) *CreateQuotaApplicationRequest
	GetEffectiveTime() *string
	SetEnvLanguage(v string) *CreateQuotaApplicationRequest
	GetEnvLanguage() *string
	SetExpireTime(v string) *CreateQuotaApplicationRequest
	GetExpireTime() *string
	SetNoticeType(v int32) *CreateQuotaApplicationRequest
	GetNoticeType() *int32
	SetProductCode(v string) *CreateQuotaApplicationRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *CreateQuotaApplicationRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *CreateQuotaApplicationRequest
	GetQuotaCategory() *string
	SetReason(v string) *CreateQuotaApplicationRequest
	GetReason() *string
}

type CreateQuotaApplicationRequest struct {
	// >  This parameter is deprecated and is not recommended.
	//
	// The mode in which you want the application to be reviewed.
	//
	// Valid values:
	//
	// 	- Async
	//
	// 	- Sync
	//
	// example:
	//
	// Sync
	AuditMode *string `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	// The requested value of the quota.
	//
	// >
	//
	// 	- You can specify the DesireValue parameter based on the values of the `TotalUsage` and `ApplicableRange` parameters that are returned by the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation.
	//
	// 	- Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and detailed reasons when you submit an application to increase the value of the quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// 804
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions. A quota item is uniquely determined by the values of Dimensions and QuotaActionCode.
	//
	// >  Some dimensions are required. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the quota dimensions that are supported by an Alibaba Cloud service. The value of `Requisite` in the response indicates whether a dimension is required.
	Dimensions []*CreateQuotaApplicationRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// >  If you do not specify an end time, the default end time is 99 years after the quota application is submitted.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification.
	//
	// Valid values:
	//
	// 	- en: English
	//
	// 	- zh: Chinese
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// >  If you do not specify a start time, the default start time is the time when the quota application is submitted.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// 	- 0 (default): no
	//
	// 	- 3: sends a notification.
	//
	// example:
	//
	// 0
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, check the `ProductCode` parameter that is described in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >  To query the quota ID of an Alibaba Cloud service, check the `QuotaActionCode` parameter that is described in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// Default value: CommonQuota.
	//
	// Valid values:
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: whitelist quota
	//
	// 	- CommonQuota: general quota
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the application.
	//
	// > Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and detailed reasons when you submit an application to increase the value of the quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateQuotaApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationRequest) GetAuditMode() *string {
	return s.AuditMode
}

func (s *CreateQuotaApplicationRequest) GetDesireValue() *float32 {
	return s.DesireValue
}

func (s *CreateQuotaApplicationRequest) GetDimensions() []*CreateQuotaApplicationRequestDimensions {
	return s.Dimensions
}

func (s *CreateQuotaApplicationRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *CreateQuotaApplicationRequest) GetEnvLanguage() *string {
	return s.EnvLanguage
}

func (s *CreateQuotaApplicationRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateQuotaApplicationRequest) GetNoticeType() *int32 {
	return s.NoticeType
}

func (s *CreateQuotaApplicationRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateQuotaApplicationRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *CreateQuotaApplicationRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *CreateQuotaApplicationRequest) GetReason() *string {
	return s.Reason
}

func (s *CreateQuotaApplicationRequest) SetAuditMode(v string) *CreateQuotaApplicationRequest {
	s.AuditMode = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetDesireValue(v float32) *CreateQuotaApplicationRequest {
	s.DesireValue = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetDimensions(v []*CreateQuotaApplicationRequestDimensions) *CreateQuotaApplicationRequest {
	s.Dimensions = v
	return s
}

func (s *CreateQuotaApplicationRequest) SetEffectiveTime(v string) *CreateQuotaApplicationRequest {
	s.EffectiveTime = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetEnvLanguage(v string) *CreateQuotaApplicationRequest {
	s.EnvLanguage = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetExpireTime(v string) *CreateQuotaApplicationRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetNoticeType(v int32) *CreateQuotaApplicationRequest {
	s.NoticeType = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetProductCode(v string) *CreateQuotaApplicationRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetQuotaActionCode(v string) *CreateQuotaApplicationRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetQuotaCategory(v string) *CreateQuotaApplicationRequest {
	s.QuotaCategory = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetReason(v string) *CreateQuotaApplicationRequest {
	s.Reason = &v
	return s
}

func (s *CreateQuotaApplicationRequest) Validate() error {
	return dara.Validate(s)
}

type CreateQuotaApplicationRequestDimensions struct {
	// The key of the dimension.
	//
	// >  You must configure `Dimensions.N.Key` and `Dimensions.N.Value` at the same time. The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the dimensions that are supported by an Alibaba Cloud service. The number of elements in the returned array is N.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >  You must configure `Dimensions.N.Key` and `Dimensions.N.Value` at the same time. The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the dimensions that are supported by an Alibaba Cloud service. The number of elements in the returned array is N.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQuotaApplicationRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationRequestDimensions) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *CreateQuotaApplicationRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *CreateQuotaApplicationRequestDimensions) SetKey(v string) *CreateQuotaApplicationRequestDimensions {
	s.Key = &v
	return s
}

func (s *CreateQuotaApplicationRequestDimensions) SetValue(v string) *CreateQuotaApplicationRequestDimensions {
	s.Value = &v
	return s
}

func (s *CreateQuotaApplicationRequestDimensions) Validate() error {
	return dara.Validate(s)
}

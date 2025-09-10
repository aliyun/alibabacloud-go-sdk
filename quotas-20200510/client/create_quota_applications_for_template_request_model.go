// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaApplicationsForTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUids(v []*string) *CreateQuotaApplicationsForTemplateRequest
	GetAliyunUids() []*string
	SetDesireValue(v float64) *CreateQuotaApplicationsForTemplateRequest
	GetDesireValue() *float64
	SetDimensions(v []*CreateQuotaApplicationsForTemplateRequestDimensions) *CreateQuotaApplicationsForTemplateRequest
	GetDimensions() []*CreateQuotaApplicationsForTemplateRequestDimensions
	SetEffectiveTime(v string) *CreateQuotaApplicationsForTemplateRequest
	GetEffectiveTime() *string
	SetEnvLanguage(v string) *CreateQuotaApplicationsForTemplateRequest
	GetEnvLanguage() *string
	SetExpireTime(v string) *CreateQuotaApplicationsForTemplateRequest
	GetExpireTime() *string
	SetNoticeType(v int32) *CreateQuotaApplicationsForTemplateRequest
	GetNoticeType() *int32
	SetProductCode(v string) *CreateQuotaApplicationsForTemplateRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *CreateQuotaApplicationsForTemplateRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *CreateQuotaApplicationsForTemplateRequest
	GetQuotaCategory() *string
	SetReason(v string) *CreateQuotaApplicationsForTemplateRequest
	GetReason() *string
}

type CreateQuotaApplicationsForTemplateRequest struct {
	// The Alibaba Cloud accounts that correspond to the resource directory members for which the quotas are applied.
	//
	// >  You can submit a quota increase application for a maximum of 50 members at a time. For more information about the members of a resource directory, see [ListAccounts](https://help.aliyun.com/document_detail/604207.html).
	//
	// This parameter is required.
	AliyunUids []*string `json:"AliyunUids,omitempty" xml:"AliyunUids,omitempty" type:"Repeated"`
	// The requested value of the quota.
	//
	// >
	//
	// 	- You can specify DesireValue based on the values of `TotalUsage` and `ApplicableRange` in the response to the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation.
	//
	// 	- Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, specify a reasonable quota value and a detailed reason.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	DesireValue *float64 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateQuotaApplicationsForTemplateRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// >  If you do not specify a start time, the value is the time when the quota application is submitted.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the notification about the application result. Valid values:
	//
	// 	- zh (default): Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// >  If you do not specify an end time, the value is 99 years after the start time of the validity period.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// 	- 0 (default): no
	//
	// 	- 3: yes
	//
	// example:
	//
	// 0
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >  To query the quota ID of an Alibaba Cloud service, call the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) and check the value of `QuotaActionCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// This parameter is required.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the quota application.
	//
	// >  Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and detailed reasons when you submit the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetAliyunUids() []*string {
	return s.AliyunUids
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetDesireValue() *float64 {
	return s.DesireValue
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetDimensions() []*CreateQuotaApplicationsForTemplateRequestDimensions {
	return s.Dimensions
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetEnvLanguage() *string {
	return s.EnvLanguage
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetNoticeType() *int32 {
	return s.NoticeType
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *CreateQuotaApplicationsForTemplateRequest) GetReason() *string {
	return s.Reason
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetAliyunUids(v []*string) *CreateQuotaApplicationsForTemplateRequest {
	s.AliyunUids = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetDesireValue(v float64) *CreateQuotaApplicationsForTemplateRequest {
	s.DesireValue = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetDimensions(v []*CreateQuotaApplicationsForTemplateRequestDimensions) *CreateQuotaApplicationsForTemplateRequest {
	s.Dimensions = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetEffectiveTime(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.EffectiveTime = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetEnvLanguage(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.EnvLanguage = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetExpireTime(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetNoticeType(v int32) *CreateQuotaApplicationsForTemplateRequest {
	s.NoticeType = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetProductCode(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetQuotaActionCode(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetQuotaCategory(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.QuotaCategory = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetReason(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.Reason = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type CreateQuotaApplicationsForTemplateRequestDimensions struct {
	// The key of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateRequestDimensions) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *CreateQuotaApplicationsForTemplateRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *CreateQuotaApplicationsForTemplateRequestDimensions) SetKey(v string) *CreateQuotaApplicationsForTemplateRequestDimensions {
	s.Key = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequestDimensions) SetValue(v string) *CreateQuotaApplicationsForTemplateRequestDimensions {
	s.Value = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequestDimensions) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateQuotaItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesireValue(v float32) *CreateTemplateQuotaItemRequest
	GetDesireValue() *float32
	SetDimensions(v []*CreateTemplateQuotaItemRequestDimensions) *CreateTemplateQuotaItemRequest
	GetDimensions() []*CreateTemplateQuotaItemRequestDimensions
	SetEffectiveTime(v string) *CreateTemplateQuotaItemRequest
	GetEffectiveTime() *string
	SetEnvLanguage(v string) *CreateTemplateQuotaItemRequest
	GetEnvLanguage() *string
	SetExpireTime(v string) *CreateTemplateQuotaItemRequest
	GetExpireTime() *string
	SetNoticeType(v int64) *CreateTemplateQuotaItemRequest
	GetNoticeType() *int64
	SetProductCode(v string) *CreateTemplateQuotaItemRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *CreateTemplateQuotaItemRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *CreateTemplateQuotaItemRequest
	GetQuotaCategory() *string
}

type CreateTemplateQuotaItemRequest struct {
	// The requested value of the quota.
	//
	// >
	//
	// 	- You can specify DesireValue based on the values of `TotalUsage` and `ApplicableRange` in the response to the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation.
	//
	// 	- Applications are reviewed by the technical support team for each cloud service. To increase the success rate of your application, specify a reasonable quota value and a detailed reason.
	//
	// This parameter is required.
	//
	// example:
	//
	// 804
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateTemplateQuotaItemRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If you leave this parameter empty, the quota takes effect immediately.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// 	- zh (default value): Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If you leave this parameter empty, no end time is specified.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// 	- 0 (default value): no
	//
	// 	- 3: yes
	//
	// example:
	//
	// 0
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >  To obtain the quota ID of an Alibaba Cloud service, call the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation and check the value of `QuotaActionCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- WhiteListLabel: privilege
	//
	// 	- FlowControl: API rate limit
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s CreateTemplateQuotaItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateQuotaItemRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateQuotaItemRequest) GetDesireValue() *float32 {
	return s.DesireValue
}

func (s *CreateTemplateQuotaItemRequest) GetDimensions() []*CreateTemplateQuotaItemRequestDimensions {
	return s.Dimensions
}

func (s *CreateTemplateQuotaItemRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *CreateTemplateQuotaItemRequest) GetEnvLanguage() *string {
	return s.EnvLanguage
}

func (s *CreateTemplateQuotaItemRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateTemplateQuotaItemRequest) GetNoticeType() *int64 {
	return s.NoticeType
}

func (s *CreateTemplateQuotaItemRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateTemplateQuotaItemRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *CreateTemplateQuotaItemRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *CreateTemplateQuotaItemRequest) SetDesireValue(v float32) *CreateTemplateQuotaItemRequest {
	s.DesireValue = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetDimensions(v []*CreateTemplateQuotaItemRequestDimensions) *CreateTemplateQuotaItemRequest {
	s.Dimensions = v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetEffectiveTime(v string) *CreateTemplateQuotaItemRequest {
	s.EffectiveTime = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetEnvLanguage(v string) *CreateTemplateQuotaItemRequest {
	s.EnvLanguage = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetExpireTime(v string) *CreateTemplateQuotaItemRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetNoticeType(v int64) *CreateTemplateQuotaItemRequest {
	s.NoticeType = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetProductCode(v string) *CreateTemplateQuotaItemRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetQuotaActionCode(v string) *CreateTemplateQuotaItemRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetQuotaCategory(v string) *CreateTemplateQuotaItemRequest {
	s.QuotaCategory = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateQuotaItemRequestDimensions struct {
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

func (s CreateTemplateQuotaItemRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateQuotaItemRequestDimensions) GoString() string {
	return s.String()
}

func (s *CreateTemplateQuotaItemRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *CreateTemplateQuotaItemRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *CreateTemplateQuotaItemRequestDimensions) SetKey(v string) *CreateTemplateQuotaItemRequestDimensions {
	s.Key = &v
	return s
}

func (s *CreateTemplateQuotaItemRequestDimensions) SetValue(v string) *CreateTemplateQuotaItemRequestDimensions {
	s.Value = &v
	return s
}

func (s *CreateTemplateQuotaItemRequestDimensions) Validate() error {
	return dara.Validate(s)
}

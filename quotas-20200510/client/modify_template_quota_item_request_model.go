// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateQuotaItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesireValue(v float32) *ModifyTemplateQuotaItemRequest
	GetDesireValue() *float32
	SetDimensions(v []*ModifyTemplateQuotaItemRequestDimensions) *ModifyTemplateQuotaItemRequest
	GetDimensions() []*ModifyTemplateQuotaItemRequestDimensions
	SetEffectiveTime(v string) *ModifyTemplateQuotaItemRequest
	GetEffectiveTime() *string
	SetEnvLanguage(v string) *ModifyTemplateQuotaItemRequest
	GetEnvLanguage() *string
	SetExpireTime(v string) *ModifyTemplateQuotaItemRequest
	GetExpireTime() *string
	SetId(v string) *ModifyTemplateQuotaItemRequest
	GetId() *string
	SetNoticeType(v int64) *ModifyTemplateQuotaItemRequest
	GetNoticeType() *int64
	SetProductCode(v string) *ModifyTemplateQuotaItemRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ModifyTemplateQuotaItemRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *ModifyTemplateQuotaItemRequest
	GetQuotaCategory() *string
}

type ModifyTemplateQuotaItemRequest struct {
	// The value of the quota that you want to apply for.
	//
	// > - You can set this parameter based on the `TotalUsage` (quota usage) and `ApplicableRange` (the range of values for the quota) parameters of the target quota in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// >
	//
	// > - The quota application is reviewed by the technical support of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and a detailed reason when you submit the application.
	//
	// example:
	//
	// 804
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*ModifyTemplateQuotaItemRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The UTC time when the quota takes effect.
	//
	// This parameter is applicable only to privileges (WhiteListLabel).
	//
	// > If you do not specify this parameter, the quota takes effect immediately.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the notification for a quota alert. Valid values:
	//
	// - zh (default): Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The UTC time when the quota expires.
	//
	// This parameter is applicable only to privileges (WhiteListLabel).
	//
	// > If you do not specify this parameter, the quota does not expire.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the quota template.
	//
	// For more information about how to obtain the ID of a quota template, see [ListQuotaApplicationTemplates](https://help.aliyun.com/document_detail/450403.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// - 0 (default): no.
	//
	// - 3: yes.
	//
	// example:
	//
	// 0
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information about how to obtain the abbreviation of an Alibaba Cloud service name, see the `ProductCode` parameter in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// > For more information about how to obtain the quota ID of an Alibaba Cloud service, see the `QuotaActionCode` parameter in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// - CommonQuota (default): general quota.
	//
	// - WhiteListLabel: privilege.
	//
	// - FlowControl: API rate limit.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ModifyTemplateQuotaItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateQuotaItemRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemRequest) GetDesireValue() *float32 {
	return s.DesireValue
}

func (s *ModifyTemplateQuotaItemRequest) GetDimensions() []*ModifyTemplateQuotaItemRequestDimensions {
	return s.Dimensions
}

func (s *ModifyTemplateQuotaItemRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyTemplateQuotaItemRequest) GetEnvLanguage() *string {
	return s.EnvLanguage
}

func (s *ModifyTemplateQuotaItemRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModifyTemplateQuotaItemRequest) GetId() *string {
	return s.Id
}

func (s *ModifyTemplateQuotaItemRequest) GetNoticeType() *int64 {
	return s.NoticeType
}

func (s *ModifyTemplateQuotaItemRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ModifyTemplateQuotaItemRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ModifyTemplateQuotaItemRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ModifyTemplateQuotaItemRequest) SetDesireValue(v float32) *ModifyTemplateQuotaItemRequest {
	s.DesireValue = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetDimensions(v []*ModifyTemplateQuotaItemRequestDimensions) *ModifyTemplateQuotaItemRequest {
	s.Dimensions = v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetEffectiveTime(v string) *ModifyTemplateQuotaItemRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetEnvLanguage(v string) *ModifyTemplateQuotaItemRequest {
	s.EnvLanguage = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetExpireTime(v string) *ModifyTemplateQuotaItemRequest {
	s.ExpireTime = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetId(v string) *ModifyTemplateQuotaItemRequest {
	s.Id = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetNoticeType(v int64) *ModifyTemplateQuotaItemRequest {
	s.NoticeType = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetProductCode(v string) *ModifyTemplateQuotaItemRequest {
	s.ProductCode = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetQuotaActionCode(v string) *ModifyTemplateQuotaItemRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetQuotaCategory(v string) *ModifyTemplateQuotaItemRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyTemplateQuotaItemRequestDimensions struct {
	// The key of the quota dimension.
	//
	// > - The value of N depends on the number of dimensions that are supported by the Alibaba Cloud service.
	//
	// >
	//
	// > - If you set `ProductCode` to a value such as `ecs`, `ecs-spec`, `actiontrail`, or `ess`, you must specify both the key and value of the quota dimension.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension.
	//
	// > - The value of N depends on the number of dimensions that are supported by the Alibaba Cloud service.
	//
	// >
	//
	// > - If you set `ProductCode` to a value such as `ecs`, `ecs-spec`, `actiontrail`, or `ess`, you must specify both the key and value of the quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyTemplateQuotaItemRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateQuotaItemRequestDimensions) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *ModifyTemplateQuotaItemRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *ModifyTemplateQuotaItemRequestDimensions) SetKey(v string) *ModifyTemplateQuotaItemRequestDimensions {
	s.Key = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequestDimensions) SetValue(v string) *ModifyTemplateQuotaItemRequestDimensions {
	s.Value = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequestDimensions) Validate() error {
	return dara.Validate(s)
}

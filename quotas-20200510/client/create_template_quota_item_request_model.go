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
	// The requested quota value.
	//
	// > - Set the requested quota value based on the `TotalUsage` (quota usage) and `ApplicableRange` (applicable range for the quota item) of the target quota in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// > - Quota applications are reviewed by the technical support team of each cloud service. To increase the approval rate, specify a reasonable requested value and provide a detailed reason when you apply for a quota increase.
	//
	// This parameter is required.
	//
	// example:
	//
	// 804
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateTemplateQuotaItemRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The UTC time when the quota takes effect.
	//
	// This parameter applies only to privileges (WhiteListLabel).
	//
	// > An empty value indicates that the quota takes effect immediately.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
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
	// This parameter applies only to privileges (WhiteListLabel).
	//
	// > An empty value indicates that no UTC expiration time is specified.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification for the quota application result. Valid values:
	//
	// - 0 (default): No.
	//
	// - 3: Yes.
	//
	// example:
	//
	// 0
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the cloud service.
	//
	// > For information about how to obtain the abbreviation of a cloud service, see `ProductCode` in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// > For information about how to obtain the quota ID of a cloud service, see `QuotaActionCode` in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota category. Valid values:
	//
	// - CommonQuota: general quota.
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

type CreateTemplateQuotaItemRequestDimensions struct {
	// The key of the quota dimension.
	//
	// > - The value range of N depends on the number of dimensions supported by the cloud service.
	//
	// > - If `ProductCode` is `ecs`, `ecs-spec`, `actiontrail`, or `ess`, you must specify both the key and the value of the quota dimension.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension.
	//
	// > - The value range of N depends on the number of dimensions supported by the cloud service.
	//
	// > - If `ProductCode` is `ecs`, `ecs-spec`, `actiontrail`, or `ess`, you must specify both the key and the value of the quota dimension.
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

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
	// The list of Alibaba Cloud accounts that belong to the resource directory members for which you want to apply for quotas.
	//
	// > You can apply for quota increases for a maximum of 50 members at a time. For more information about the members of a resource directory, see [ListAccounts](https://help.aliyun.com/document_detail/604207.html).
	//
	// This parameter is required.
	AliyunUids []*string `json:"AliyunUids,omitempty" xml:"AliyunUids,omitempty" type:"Repeated"`
	// The requested quota value.
	//
	// > - Set the requested quota value based on the `TotalUsage` (quota usage) and `ApplicableRange` (applicable range for quota application) of the target quota in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// > - Quota applications are reviewed and approved by the technical support team of each cloud service. To increase the approval rate, specify a reasonable requested value and provide a detailed reason when you submit a quota application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	DesireValue *float64 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateQuotaApplicationsForTemplateRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The UTC time when the quota takes effect. This parameter applies only to privileges (WhiteListLabel).
	//
	//
	//
	// > If you do not specify an effective period for the current account, the commit time is used by default.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the notification for the quota application result. Valid values:
	//
	// - zh (default): Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The UTC time when the quota expires. This parameter applies only to privileges (WhiteListLabel).
	//
	// > If you do not specify an expiration time for the current account, the expiration time is set to 99 years after the effective period of the current quota.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send notifications for quota application results. Valid values:
	//
	// - 0 (default): no.
	//
	// - 3: yes.
	//
	// example:
	//
	// 0
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviated name of the cloud service.
	//
	// > For more information about how to obtain the abbreviated name of a cloud service, see `ProductCode` in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// > For more information about how to obtain the quota ID of a cloud service, see `QuotaActionCode` in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// - CommonQuota: general quota.
	//
	// - FlowControl: API rate limit.
	//
	// - WhiteListLabel: privilege.
	//
	// This parameter is required.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the quota application.
	//
	// > Quota applications are reviewed and approved by the technical support team of each cloud service. To increase the approval rate, specify a reasonable requested value and provide a detailed reason when you submit a quota application.
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

type CreateQuotaApplicationsForTemplateRequestDimensions struct {
	// The key of the quota dimension.
	//
	// > - The value range of N depends on the number of dimensions supported by the cloud service.
	//
	// > - When `ProductCode` is set to `ecs`, `ecs-spec`, `actiontrail`, or `ess`, you must specify both the key and the value of the quota dimension.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension.
	//
	// > - The value range of N depends on the number of dimensions supported by the cloud service.
	//
	// > - When `ProductCode` is set to `ecs`, `ecs-spec`, `actiontrail`, or `ess`, you must specify both the key and the value of the quota dimension.
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

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
	// 	Notice: This parameter is deprecated and is not recommended.
	//
	// The quota review mode.
	//
	// example:
	//
	// Sync
	AuditMode *string `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	// The requested quota value.
	//
	// > - Set the requested quota value based on the `TotalUsage` (quota usage) and `ApplicableRange` (applicable range of the quota) of the target quota in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// > - Quota increase requests are reviewed by the technical support of each cloud service. To increase the chances of approval, provide a reasonable requested value and a detailed reason when you submit the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 804
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	//
	// The dimensions are attributes of the quota. The combination of dimensions (Dimensions) and quota ID (QuotaActionCode) uniquely identifies a quota.
	//
	// > Some cloud services require dimensions. You can call [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) to query the quota dimensions supported by a cloud service. The `Requisite` response parameter indicates whether a dimension is required.
	Dimensions []*CreateQuotaApplicationRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The UTC time when the quota takes effect.
	//
	// This parameter is applicable only to privileges (WhiteListLabel).
	//
	// > If you do not specify an effective period, the default value is 99 years from the start time of the current quota.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification.
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The UTC time when the quota expires.
	//
	// This parameter is applicable only to privileges (WhiteListLabel).
	//
	// > If you do not specify an expiration time, the default value is the submission time.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the request result. Valid values:
	//
	// - 0 (default): does not send a notification.
	//
	// - 3: sends a notification.
	//
	// example:
	//
	// 0
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// > For information about how to obtain the abbreviation of a cloud service name, see `ProductCode` in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
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
	// The quota type.
	//
	// Default value: CommonQuota.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the quota increase request.
	//
	// > Quota increase requests are reviewed by the technical support of each cloud service. To increase the chances of approval, provide a reasonable requested value and a detailed reason when you submit the request.
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

type CreateQuotaApplicationRequestDimensions struct {
	// The key of the quota dimension.
	//
	// > `Dimensions.N.Key` and `Dimensions.N.Value` must be specified together. The valid values of N depend on the number of dimensions supported by the cloud service. You can call [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) to query the quota dimensions supported by a cloud service. The length of the returned array is the value of N.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension.
	//
	// > `Dimensions.N.Key` and `Dimensions.N.Value` must be specified together. The valid values of N depend on the number of dimensions supported by the cloud service. You can call [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) to query the quota dimensions supported by a cloud service. The length of the returned array is the value of N.
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

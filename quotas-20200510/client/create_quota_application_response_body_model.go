// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateQuotaApplicationResponseBody
	GetApplicationId() *string
	SetApplyTime(v string) *CreateQuotaApplicationResponseBody
	GetApplyTime() *string
	SetApproveValue(v float32) *CreateQuotaApplicationResponseBody
	GetApproveValue() *float32
	SetAuditReason(v string) *CreateQuotaApplicationResponseBody
	GetAuditReason() *string
	SetDesireValue(v int32) *CreateQuotaApplicationResponseBody
	GetDesireValue() *int32
	SetDimension(v map[string]interface{}) *CreateQuotaApplicationResponseBody
	GetDimension() map[string]interface{}
	SetEffectiveTime(v string) *CreateQuotaApplicationResponseBody
	GetEffectiveTime() *string
	SetExpireTime(v string) *CreateQuotaApplicationResponseBody
	GetExpireTime() *string
	SetNoticeType(v int64) *CreateQuotaApplicationResponseBody
	GetNoticeType() *int64
	SetProductCode(v string) *CreateQuotaApplicationResponseBody
	GetProductCode() *string
	SetQuotaActionCode(v string) *CreateQuotaApplicationResponseBody
	GetQuotaActionCode() *string
	SetQuotaArn(v string) *CreateQuotaApplicationResponseBody
	GetQuotaArn() *string
	SetQuotaDescription(v string) *CreateQuotaApplicationResponseBody
	GetQuotaDescription() *string
	SetQuotaName(v string) *CreateQuotaApplicationResponseBody
	GetQuotaName() *string
	SetQuotaUnit(v string) *CreateQuotaApplicationResponseBody
	GetQuotaUnit() *string
	SetReason(v string) *CreateQuotaApplicationResponseBody
	GetReason() *string
	SetRequestId(v string) *CreateQuotaApplicationResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateQuotaApplicationResponseBody
	GetStatus() *string
}

type CreateQuotaApplicationResponseBody struct {
	// The ID of the application.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was submitted.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The quota value that is approved.
	//
	// example:
	//
	// 804
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The requested value of the quota.
	//
	// example:
	//
	// 12
	DesireValue *int32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimension.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the new quota value takes effect.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the new quota expires.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the notification about the application result is sent. Valid values:
	//
	// 	- 0: The notification is not sent.
	//
	// 	- 3: The notification is sent.
	//
	// example:
	//
	// 3
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.c5.large
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:*:quota/ecs/ecs.m2.medium/prepaid/classic/instancetype/cn-hangzhou-b
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// ecs.c5.large
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// ecs.c5.large
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// AMOUNT
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The approval state of the quota increase application.
	//
	// Valid values:
	//
	// 	- Cancel: The application is canceled.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is in review.
	//
	// 	- Disagree: The application is rejected.
	//
	// example:
	//
	// Process
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateQuotaApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateQuotaApplicationResponseBody) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *CreateQuotaApplicationResponseBody) GetApproveValue() *float32 {
	return s.ApproveValue
}

func (s *CreateQuotaApplicationResponseBody) GetAuditReason() *string {
	return s.AuditReason
}

func (s *CreateQuotaApplicationResponseBody) GetDesireValue() *int32 {
	return s.DesireValue
}

func (s *CreateQuotaApplicationResponseBody) GetDimension() map[string]interface{} {
	return s.Dimension
}

func (s *CreateQuotaApplicationResponseBody) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *CreateQuotaApplicationResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateQuotaApplicationResponseBody) GetNoticeType() *int64 {
	return s.NoticeType
}

func (s *CreateQuotaApplicationResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateQuotaApplicationResponseBody) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *CreateQuotaApplicationResponseBody) GetQuotaArn() *string {
	return s.QuotaArn
}

func (s *CreateQuotaApplicationResponseBody) GetQuotaDescription() *string {
	return s.QuotaDescription
}

func (s *CreateQuotaApplicationResponseBody) GetQuotaName() *string {
	return s.QuotaName
}

func (s *CreateQuotaApplicationResponseBody) GetQuotaUnit() *string {
	return s.QuotaUnit
}

func (s *CreateQuotaApplicationResponseBody) GetReason() *string {
	return s.Reason
}

func (s *CreateQuotaApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQuotaApplicationResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateQuotaApplicationResponseBody) SetApplicationId(v string) *CreateQuotaApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetApplyTime(v string) *CreateQuotaApplicationResponseBody {
	s.ApplyTime = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetApproveValue(v float32) *CreateQuotaApplicationResponseBody {
	s.ApproveValue = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetAuditReason(v string) *CreateQuotaApplicationResponseBody {
	s.AuditReason = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetDesireValue(v int32) *CreateQuotaApplicationResponseBody {
	s.DesireValue = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetDimension(v map[string]interface{}) *CreateQuotaApplicationResponseBody {
	s.Dimension = v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetEffectiveTime(v string) *CreateQuotaApplicationResponseBody {
	s.EffectiveTime = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetExpireTime(v string) *CreateQuotaApplicationResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetNoticeType(v int64) *CreateQuotaApplicationResponseBody {
	s.NoticeType = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetProductCode(v string) *CreateQuotaApplicationResponseBody {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaActionCode(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaArn(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaArn = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaDescription(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaDescription = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaName(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaName = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaUnit(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaUnit = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetReason(v string) *CreateQuotaApplicationResponseBody {
	s.Reason = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetRequestId(v string) *CreateQuotaApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetStatus(v string) *CreateQuotaApplicationResponseBody {
	s.Status = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

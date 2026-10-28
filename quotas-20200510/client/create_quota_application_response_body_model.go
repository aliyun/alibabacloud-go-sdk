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
	// The ID of the quota increase request.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the quota increase request was submitted. The time is in ISO 8601 format and in UTC. The value ends with Z.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The approved quota value.
	//
	// example:
	//
	// 804
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The review comments on the quota increase request.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The requested quota value.
	//
	// example:
	//
	// 804
	DesireValue *int32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The effective period of the quota. The time is in ISO 8601 format and in UTC. The value ends with Z.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the quota expires. The time is in ISO 8601 format and in UTC. The value ends with Z.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether a notification is sent about the request result. Valid values:
	//
	// - 0: no notification is sent.
	//
	// - 3: a notification is sent.
	//
	// example:
	//
	// 3
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas::141339776561****:quota/ecs/q_security-groups
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of security groups that can be owned by the current account.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Security Groups
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// AMOUNT
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the quota increase request.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The approval status of the quota increase request.
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

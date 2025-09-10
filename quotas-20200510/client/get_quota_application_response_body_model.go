// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaApplication(v *GetQuotaApplicationResponseBodyQuotaApplication) *GetQuotaApplicationResponseBody
	GetQuotaApplication() *GetQuotaApplicationResponseBodyQuotaApplication
	SetRequestId(v string) *GetQuotaApplicationResponseBody
	GetRequestId() *string
}

type GetQuotaApplicationResponseBody struct {
	// The details of the quota application.
	QuotaApplication *GetQuotaApplicationResponseBodyQuotaApplication `json:"QuotaApplication,omitempty" xml:"QuotaApplication,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7BBD1D37-094C-4485-8B7D-64682F82BC18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQuotaApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationResponseBody) GetQuotaApplication() *GetQuotaApplicationResponseBodyQuotaApplication {
	return s.QuotaApplication
}

func (s *GetQuotaApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaApplicationResponseBody) SetQuotaApplication(v *GetQuotaApplicationResponseBodyQuotaApplication) *GetQuotaApplicationResponseBody {
	s.QuotaApplication = v
	return s
}

func (s *GetQuotaApplicationResponseBody) SetRequestId(v string) *GetQuotaApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQuotaApplicationResponseBodyQuotaApplication struct {
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
	// The approved quota value.
	//
	// example:
	//
	// 10
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The expected value of the quota.
	//
	// example:
	//
	// 804
	DesireValue *int32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The dimension.
	//
	// Format: `{"regionId":"Region"}`.
	//
	// example:
	//
	// ["cn-shanghai","cn-hangzhou"]
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the new quota value takes effect.
	//
	// example:
	//
	// 2021-01-19 15:30:00
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the new quota expires.
	//
	// example:
	//
	// 2023-06-29 15:30:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The method of that is used to send alert notifications. Valid values:
	//
	// 	- 0: Quota Center does not send a notification.
	//
	// 	- 1: Quota Center sends an email notification.
	//
	// 	- 2: Quota Center sends an SMS notification.
	//
	// example:
	//
	// 0
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:120886317861****:quota/ecs/q_security-groups/
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota.
	//
	// 	- FlowControl: API rate limit.
	//
	// 	- WhiteListLabel: whitelist quota.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of security groups that can be owned by the current account
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Security Groups
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the new quota value.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is being reviewed.
	//
	// 	- Cancel: The application is closed.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetQuotaApplicationResponseBodyQuotaApplication) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaApplicationResponseBodyQuotaApplication) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetApproveValue() *float32 {
	return s.ApproveValue
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetAuditReason() *string {
	return s.AuditReason
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetDesireValue() *int32 {
	return s.DesireValue
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetDimension() map[string]interface{} {
	return s.Dimension
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetNoticeType() *int64 {
	return s.NoticeType
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetQuotaArn() *string {
	return s.QuotaArn
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetQuotaDescription() *string {
	return s.QuotaDescription
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetQuotaName() *string {
	return s.QuotaName
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetQuotaUnit() *string {
	return s.QuotaUnit
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetReason() *string {
	return s.Reason
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetApplicationId(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ApplicationId = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetApplyTime(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ApplyTime = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetApproveValue(v float32) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ApproveValue = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetAuditReason(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.AuditReason = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetDesireValue(v int32) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.DesireValue = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetDimension(v map[string]interface{}) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.Dimension = v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetEffectiveTime(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.EffectiveTime = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetExpireTime(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ExpireTime = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetNoticeType(v int64) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.NoticeType = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetProductCode(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ProductCode = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaActionCode(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaActionCode = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaArn(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaArn = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaCategory(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaCategory = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaDescription(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaDescription = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaName(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaName = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaUnit(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaUnit = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetReason(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.Reason = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetStatus(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.Status = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) Validate() error {
	return dara.Validate(s)
}

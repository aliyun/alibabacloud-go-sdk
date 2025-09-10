// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsDetailForTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListQuotaApplicationsDetailForTemplateResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaApplicationsDetailForTemplateResponseBody
	GetNextToken() *string
	SetQuotaApplications(v []*ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) *ListQuotaApplicationsDetailForTemplateResponseBody
	GetQuotaApplications() []*ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications
	SetRequestId(v string) *ListQuotaApplicationsDetailForTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListQuotaApplicationsDetailForTemplateResponseBody
	GetTotalCount() *int32
}

type ListQuotaApplicationsDetailForTemplateResponseBody struct {
	// The maximum number of records that can be returned for the query.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details of the quota increase application.
	QuotaApplications []*ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications `json:"QuotaApplications,omitempty" xml:"QuotaApplications,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 9
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) GetQuotaApplications() []*ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	return s.QuotaApplications
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetMaxResults(v int32) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetNextToken(v string) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetQuotaApplications(v []*ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.QuotaApplications = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetRequestId(v string) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetTotalCount(v int32) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 175376458581****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The ID of the quota increase application.
	//
	// example:
	//
	// b926571d-cc09-4711-b547-58a615f0****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the quota increase application was submitted. The value is displayed in UTC.
	//
	// example:
	//
	// 2021-01-15T09:13:53Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The quota value that is approved.
	//
	// example:
	//
	// 101
	ApproveValue *float64 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The approval result of the quota increase application.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The ID of the quota application batch.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The requested value of the quota.
	//
	// example:
	//
	// 60
	DesireValue *float64 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2021-01-15T11:46:25Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2021-01-17T11:46:25Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// 	- 0: no
	//
	// 	- 3: yes
	//
	// example:
	//
	// 3
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The calculation cycle of the quota.
	Period *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.n1.large
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:*:120886317861****:quota/csk/q_i5uzm3
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota (default): general quota
	//
	// 	- WhiteListLabel: privilege
	//
	// 	- FlowControl: API rate limit
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of nodes in a cluster
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The quota dimensions.
	QuotaDimension map[string]*string `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty"`
	// The quota name.
	//
	// example:
	//
	// Maximum Number of Nodes
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// GiB
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the quota increase application.
	//
	// example:
	//
	// Business expansion
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The approval status of the quota increase application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is pending approval.
	//
	// 	- Cancel: The application is canceled.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetApproveValue() *float64 {
	return s.ApproveValue
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetAuditReason() *string {
	return s.AuditReason
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetBatchQuotaApplicationId() *string {
	return s.BatchQuotaApplicationId
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetDesireValue() *float64 {
	return s.DesireValue
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetEnvLanguage() *string {
	return s.EnvLanguage
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetNoticeType() *int32 {
	return s.NoticeType
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetPeriod() *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod {
	return s.Period
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetQuotaArn() *string {
	return s.QuotaArn
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetQuotaDescription() *string {
	return s.QuotaDescription
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetQuotaDimension() map[string]*string {
	return s.QuotaDimension
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetQuotaUnit() *string {
	return s.QuotaUnit
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetReason() *string {
	return s.Reason
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GetStatus() *string {
	return s.Status
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetAliyunUid(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.AliyunUid = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetApplicationId(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetApplyTime(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ApplyTime = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetApproveValue(v float64) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ApproveValue = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetAuditReason(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.AuditReason = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetDesireValue(v float64) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetEffectiveTime(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetEnvLanguage(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.EnvLanguage = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetExpireTime(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetNoticeType(v int32) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.NoticeType = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetPeriod(v *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.Period = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetProductCode(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaActionCode(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaArn(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaArn = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaCategory(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaDescription(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaDescription = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaDimension(v map[string]*string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaDimension = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaName(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaName = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaUnit(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaUnit = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetReason(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.Reason = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetStatus(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.Status = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod struct {
	// The unit of the calculation cycle of the quota.
	//
	// example:
	//
	// second
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle of the quota.
	//
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) GetPeriodValue() *int32 {
	return s.PeriodValue
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) SetPeriodUnit(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) SetPeriodValue(v int32) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod {
	s.PeriodValue = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) Validate() error {
	return dara.Validate(s)
}

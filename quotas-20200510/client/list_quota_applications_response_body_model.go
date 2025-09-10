// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListQuotaApplicationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaApplicationsResponseBody
	GetNextToken() *string
	SetQuotaApplications(v []*ListQuotaApplicationsResponseBodyQuotaApplications) *ListQuotaApplicationsResponseBody
	GetQuotaApplications() []*ListQuotaApplicationsResponseBodyQuotaApplications
	SetRequestId(v string) *ListQuotaApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListQuotaApplicationsResponseBody
	GetTotalCount() *int32
}

type ListQuotaApplicationsResponseBody struct {
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details of the quota increase applications.
	QuotaApplications []*ListQuotaApplicationsResponseBodyQuotaApplications `json:"QuotaApplications,omitempty" xml:"QuotaApplications,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 730925FB-0BD0-40AC-AF3A-A6C6E9716879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of applications.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaApplicationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaApplicationsResponseBody) GetQuotaApplications() []*ListQuotaApplicationsResponseBodyQuotaApplications {
	return s.QuotaApplications
}

func (s *ListQuotaApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaApplicationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotaApplicationsResponseBody) SetMaxResults(v int32) *ListQuotaApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetNextToken(v string) *ListQuotaApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetQuotaApplications(v []*ListQuotaApplicationsResponseBodyQuotaApplications) *ListQuotaApplicationsResponseBody {
	s.QuotaApplications = v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetRequestId(v string) *ListQuotaApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetTotalCount(v int32) *ListQuotaApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationsResponseBodyQuotaApplications struct {
	// The ID of the application.
	//
	// example:
	//
	// b926571d-cc09-4711-b547-58a615f0****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was submitted.
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
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The remarks of the application.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The quota value that is approved.
	//
	// example:
	//
	// 101
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The dimension of the application.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the application took effect.
	//
	// example:
	//
	// 2021-01-15T11:46:25Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the application expired.
	//
	// example:
	//
	// 2021-01-17T11:46:25Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// 	- 0: A notification about the application result is not sent.
	//
	// 	- 3: A notification about the application result is sent.
	//
	// example:
	//
	// 3
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The calculation cycle of the quota.
	Period *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// csk
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_i5uzm3
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:*:120886317861****:quota/csk/q_i5uzm3
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota (default): general quota.
	//
	// 	- WhiteListLabel: whitelist quota.
	//
	// 	- FlowControl: API rate limit.
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
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Nodes
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the new quota value.
	//
	// example:
	//
	// Node
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	//
	// example:
	//
	// Business expansion
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is being reviewed.
	//
	// 	- Cancel: The application is canceled.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsResponseBodyQuotaApplications) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsResponseBodyQuotaApplications) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetApproveValue() *float32 {
	return s.ApproveValue
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetAuditReason() *string {
	return s.AuditReason
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetComment() *string {
	return s.Comment
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetDesireValue() *float32 {
	return s.DesireValue
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetDimension() map[string]interface{} {
	return s.Dimension
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetNoticeType() *int32 {
	return s.NoticeType
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetPeriod() *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod {
	return s.Period
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetQuotaArn() *string {
	return s.QuotaArn
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetQuotaDescription() *string {
	return s.QuotaDescription
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetQuotaUnit() *string {
	return s.QuotaUnit
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetReason() *string {
	return s.Reason
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) GetStatus() *string {
	return s.Status
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApplicationId(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApplyTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApplyTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApproveValue(v float32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApproveValue = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetAuditReason(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.AuditReason = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetComment(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Comment = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetDesireValue(v float32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetDimension(v map[string]interface{}) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Dimension = v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetEffectiveTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetExpireTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetNoticeType(v int32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.NoticeType = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetPeriod(v *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Period = v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetProductCode(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaActionCode(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaArn(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaArn = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaCategory(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaDescription(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaDescription = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaName(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaName = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaUnit(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaUnit = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetReason(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Reason = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetStatus(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Status = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod struct {
	// The unit of the calculation cycle. Valid values:
	//
	// 	- second
	//
	// 	- minute
	//
	// 	- hour
	//
	// 	- day
	//
	// 	- week
	//
	// example:
	//
	// second
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle.
	//
	// example:
	//
	// 1
	PeriodValue *int64 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) GetPeriodValue() *int64 {
	return s.PeriodValue
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) SetPeriodUnit(v string) *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) SetPeriodValue(v int64) *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod {
	s.PeriodValue = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) Validate() error {
	return dara.Validate(s)
}

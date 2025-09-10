// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsForTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListQuotaApplicationsForTemplateResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaApplicationsForTemplateResponseBody
	GetNextToken() *string
	SetQuotaBatchApplications(v []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) *ListQuotaApplicationsForTemplateResponseBody
	GetQuotaBatchApplications() []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications
	SetRequestId(v string) *ListQuotaApplicationsForTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListQuotaApplicationsForTemplateResponseBody
	GetTotalCount() *int32
}

type ListQuotaApplicationsForTemplateResponseBody struct {
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
	// The queried quota application records.
	QuotaBatchApplications []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications `json:"QuotaBatchApplications,omitempty" xml:"QuotaBatchApplications,omitempty" type:"Repeated"`
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
	// 67
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaApplicationsForTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaApplicationsForTemplateResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaApplicationsForTemplateResponseBody) GetQuotaBatchApplications() []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	return s.QuotaBatchApplications
}

func (s *ListQuotaApplicationsForTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaApplicationsForTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetMaxResults(v int32) *ListQuotaApplicationsForTemplateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetNextToken(v string) *ListQuotaApplicationsForTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetQuotaBatchApplications(v []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) *ListQuotaApplicationsForTemplateResponseBody {
	s.QuotaBatchApplications = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetRequestId(v string) *ListQuotaApplicationsForTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetTotalCount(v int32) *ListQuotaApplicationsForTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications struct {
	// The Alibaba Cloud accounts that correspond to the resource directory members for which the quotas are applied.
	AliyunUids []*string `json:"AliyunUids,omitempty" xml:"AliyunUids,omitempty" type:"Repeated"`
	// The time when the quota increase application was submitted. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-23T02:38:18Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The number of applications in different approval states.
	AuditStatusVos []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos `json:"AuditStatusVos,omitempty" xml:"AuditStatusVos,omitempty" type:"Repeated"`
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
	// 105
	DesireValue *float64 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	//
	// Format example: {"regionId":"cn-hangzhou"}.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2023-05-22T16:00:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2024-05-14T02:08:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// vpc
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// q_fhoz4k
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the quota increase application.
	//
	// example:
	//
	// Business expansion
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetAliyunUids() []*string {
	return s.AliyunUids
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetAuditStatusVos() []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos {
	return s.AuditStatusVos
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetBatchQuotaApplicationId() *string {
	return s.BatchQuotaApplicationId
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetDesireValue() *float64 {
	return s.DesireValue
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetDimensions() map[string]interface{} {
	return s.Dimensions
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GetReason() *string {
	return s.Reason
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetAliyunUids(v []*string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.AliyunUids = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetApplyTime(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.ApplyTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetAuditStatusVos(v []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.AuditStatusVos = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetDesireValue(v float64) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetDimensions(v map[string]interface{}) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.Dimensions = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetEffectiveTime(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetExpireTime(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetProductCode(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetQuotaActionCode(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetQuotaCategory(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetReason(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.Reason = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos struct {
	// The number of approval tickets.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The approval state of the quota increase application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is in review.
	//
	// 	- Cancel: The application is canceled.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) GetCount() *int32 {
	return s.Count
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) GetStatus() *string {
	return s.Status
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) SetCount(v int32) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos {
	s.Count = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) SetStatus(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos {
	s.Status = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) Validate() error {
	return dara.Validate(s)
}

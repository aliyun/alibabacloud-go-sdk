// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsForTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyEndTime(v string) *ListQuotaApplicationsForTemplateRequest
	GetApplyEndTime() *string
	SetApplyStartTime(v string) *ListQuotaApplicationsForTemplateRequest
	GetApplyStartTime() *string
	SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsForTemplateRequest
	GetBatchQuotaApplicationId() *string
	SetMaxResults(v int32) *ListQuotaApplicationsForTemplateRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaApplicationsForTemplateRequest
	GetNextToken() *string
	SetProductCode(v string) *ListQuotaApplicationsForTemplateRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ListQuotaApplicationsForTemplateRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *ListQuotaApplicationsForTemplateRequest
	GetQuotaCategory() *string
}

type ListQuotaApplicationsForTemplateRequest struct {
	// The end of the time range to query. The value is in UTC.
	//
	// > The interval between `ApplyStartTime` and `ApplyEndTime` must be 7 days or less. If you do not set these two parameters, quota applications submitted within the last 7 days are queried.
	//
	// example:
	//
	// 2023-05-22T16:00:00Z
	ApplyEndTime *string `json:"ApplyEndTime,omitempty" xml:"ApplyEndTime,omitempty"`
	// The beginning of the time range to query. The value is in UTC.
	//
	// > The interval between `ApplyStartTime` and `ApplyEndTime` must be 7 days or less. If you do not set these two parameters, quota applications submitted within the last 7 days are queried.
	//
	// example:
	//
	// 2023-05-22T16:00:00Z
	ApplyStartTime *string `json:"ApplyStartTime,omitempty" xml:"ApplyStartTime,omitempty"`
	// The ID of the batch quota application.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The maximum number of results to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the start of the data to read.
	//
	// > If you do not set this parameter, the query starts from the beginning.
	//
	// example:
	//
	// 4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service.
	//
	// > For more information about how to obtain the abbreviation of an Alibaba Cloud service, see the `ProductCode` parameter in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// - CommonQuota: general quota.
	//
	// - FlowControl: API rate limit.
	//
	// - WhiteListLabel: privilege.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ListQuotaApplicationsForTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateRequest) GetApplyEndTime() *string {
	return s.ApplyEndTime
}

func (s *ListQuotaApplicationsForTemplateRequest) GetApplyStartTime() *string {
	return s.ApplyStartTime
}

func (s *ListQuotaApplicationsForTemplateRequest) GetBatchQuotaApplicationId() *string {
	return s.BatchQuotaApplicationId
}

func (s *ListQuotaApplicationsForTemplateRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaApplicationsForTemplateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaApplicationsForTemplateRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaApplicationsForTemplateRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaApplicationsForTemplateRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaApplicationsForTemplateRequest) SetApplyEndTime(v string) *ListQuotaApplicationsForTemplateRequest {
	s.ApplyEndTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetApplyStartTime(v string) *ListQuotaApplicationsForTemplateRequest {
	s.ApplyStartTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsForTemplateRequest {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetMaxResults(v int32) *ListQuotaApplicationsForTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetNextToken(v string) *ListQuotaApplicationsForTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetProductCode(v string) *ListQuotaApplicationsForTemplateRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetQuotaActionCode(v string) *ListQuotaApplicationsForTemplateRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetQuotaCategory(v string) *ListQuotaApplicationsForTemplateRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsDetailForTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUid(v string) *ListQuotaApplicationsDetailForTemplateRequest
	GetAliyunUid() *string
	SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsDetailForTemplateRequest
	GetBatchQuotaApplicationId() *string
	SetMaxResults(v int32) *ListQuotaApplicationsDetailForTemplateRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaApplicationsDetailForTemplateRequest
	GetNextToken() *string
	SetProductCode(v string) *ListQuotaApplicationsDetailForTemplateRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ListQuotaApplicationsDetailForTemplateRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *ListQuotaApplicationsDetailForTemplateRequest
	GetQuotaCategory() *string
	SetStatus(v string) *ListQuotaApplicationsDetailForTemplateRequest
	GetStatus() *string
}

type ListQuotaApplicationsDetailForTemplateRequest struct {
	// The Alibaba Cloud account that is used to submit the quota increase application.
	//
	// example:
	//
	// 135048337611****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The ID of the quota application batch.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// >  An empty value indicates that the query starts from the beginning.
	//
	// example:
	//
	// 4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.c5.large
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
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

func (s ListQuotaApplicationsDetailForTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) GetBatchQuotaApplicationId() *string {
	return s.BatchQuotaApplicationId
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) GetStatus() *string {
	return s.Status
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetAliyunUid(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.AliyunUid = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetMaxResults(v int32) *ListQuotaApplicationsDetailForTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetNextToken(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetProductCode(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetQuotaActionCode(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetQuotaCategory(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetStatus(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.Status = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) Validate() error {
	return dara.Validate(s)
}

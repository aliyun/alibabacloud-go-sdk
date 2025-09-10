// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductQuotaDimensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductQuotaDimensionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductQuotaDimensionsRequest
	GetNextToken() *string
	SetProductCode(v string) *ListProductQuotaDimensionsRequest
	GetProductCode() *string
	SetQuotaCategory(v string) *ListProductQuotaDimensionsRequest
	GetQuotaCategory() *string
}

type ListProductQuotaDimensionsRequest struct {
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- FlowControl: API rate limit.
	//
	// 	- CommonQuota: general quota. This is the default value.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ListProductQuotaDimensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotaDimensionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductQuotaDimensionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductQuotaDimensionsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductQuotaDimensionsRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListProductQuotaDimensionsRequest) SetMaxResults(v int32) *ListProductQuotaDimensionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) SetNextToken(v string) *ListProductQuotaDimensionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) SetProductCode(v string) *ListProductQuotaDimensionsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) SetQuotaCategory(v string) *ListProductQuotaDimensionsRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) Validate() error {
	return dara.Validate(s)
}

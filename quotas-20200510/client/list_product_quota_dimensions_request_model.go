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
	// The maximum number of entries to return for this query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that specifies the starting position of the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviated name of the cloud service.
	//
	// > For more information about how to obtain the abbreviated name of a cloud service, see `ProductCode` in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// disk
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota category. Valid values:
	//
	// - FlowControl: API rate limits.
	//
	// - CommonQuota: general quotas. This is the default value.
	//
	// > Privileges (WhiteListLabel) do not have quota dimensions.
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

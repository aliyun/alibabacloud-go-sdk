// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductsResponseBody
	GetNextToken() *string
	SetProductInfo(v []*ListProductsResponseBodyProductInfo) *ListProductsResponseBody
	GetProductInfo() []*ListProductsResponseBodyProductInfo
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductsResponseBody
	GetTotalCount() *int32
}

type ListProductsResponseBody struct {
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the Alibaba Cloud services.
	ProductInfo []*ListProductsResponseBodyProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1DA9C136-11BC-4C39-ADC6-B86276128072
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsResponseBody) GetProductInfo() []*ListProductsResponseBodyProductInfo {
	return s.ProductInfo
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductsResponseBody) SetMaxResults(v int32) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetProductInfo(v []*ListProductsResponseBodyProductInfo) *ListProductsResponseBody {
	s.ProductInfo = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotalCount(v int32) *ListProductsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProductsResponseBodyProductInfo struct {
	// Indicates whether the Alibaba Cloud service supports general quotas. Valid values:
	//
	// 	- support: The Alibaba Cloud service supports general quotas.
	//
	// 	- unsupport: The Alibaba Cloud service does not support general quotas.
	//
	// example:
	//
	// support
	CommonQuotaSupport *string `json:"CommonQuotaSupport,omitempty" xml:"CommonQuotaSupport,omitempty"`
	// >  This parameter is discontinued and is not recommended.
	//
	// Indicates whether the Alibaba Cloud service supports dynamic quota adjustment. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Dynamic *bool `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	// Indicates whether the Alibaba Cloud service supports API rate limits. Valid values:
	//
	// 	- support: The Alibaba Cloud service supports API rate limits.
	//
	// 	- unsupport: The Alibaba Cloud service does not support API rate limits.
	//
	// example:
	//
	// unsupport
	FlowControlSupport *string `json:"FlowControlSupport,omitempty" xml:"FlowControlSupport,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the Alibaba Cloud service.
	//
	// example:
	//
	// Elastic Compute Service (ECS)
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The English name of the Alibaba Cloud service.
	//
	// example:
	//
	// Elastic Compute Service
	ProductNameEn *string `json:"ProductNameEn,omitempty" xml:"ProductNameEn,omitempty"`
	// The ID of the service category.
	//
	// example:
	//
	// 5
	SecondCategoryId *int64 `json:"SecondCategoryId,omitempty" xml:"SecondCategoryId,omitempty"`
	// The name of the service category.
	//
	// example:
	//
	// Elastic Compute
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
	// The English name of the service category.
	//
	// example:
	//
	// Elastic Compute
	SecondCategoryNameEn *string `json:"SecondCategoryNameEn,omitempty" xml:"SecondCategoryNameEn,omitempty"`
	// Indicates whether the Alibaba Cloud service supports whitelist quotas. Valid values:
	//
	// 	- support: The Alibaba Cloud service supports whitelist quotas.
	//
	// 	- unsupport: The Alibaba Cloud service does not support whitelist quotas.
	//
	// example:
	//
	// support
	WhiteListLabelQuotaSupport *string `json:"WhiteListLabelQuotaSupport,omitempty" xml:"WhiteListLabelQuotaSupport,omitempty"`
}

func (s ListProductsResponseBodyProductInfo) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyProductInfo) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProductInfo) GetCommonQuotaSupport() *string {
	return s.CommonQuotaSupport
}

func (s *ListProductsResponseBodyProductInfo) GetDynamic() *bool {
	return s.Dynamic
}

func (s *ListProductsResponseBodyProductInfo) GetFlowControlSupport() *string {
	return s.FlowControlSupport
}

func (s *ListProductsResponseBodyProductInfo) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductsResponseBodyProductInfo) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsResponseBodyProductInfo) GetProductNameEn() *string {
	return s.ProductNameEn
}

func (s *ListProductsResponseBodyProductInfo) GetSecondCategoryId() *int64 {
	return s.SecondCategoryId
}

func (s *ListProductsResponseBodyProductInfo) GetSecondCategoryName() *string {
	return s.SecondCategoryName
}

func (s *ListProductsResponseBodyProductInfo) GetSecondCategoryNameEn() *string {
	return s.SecondCategoryNameEn
}

func (s *ListProductsResponseBodyProductInfo) GetWhiteListLabelQuotaSupport() *string {
	return s.WhiteListLabelQuotaSupport
}

func (s *ListProductsResponseBodyProductInfo) SetCommonQuotaSupport(v string) *ListProductsResponseBodyProductInfo {
	s.CommonQuotaSupport = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetDynamic(v bool) *ListProductsResponseBodyProductInfo {
	s.Dynamic = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetFlowControlSupport(v string) *ListProductsResponseBodyProductInfo {
	s.FlowControlSupport = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetProductCode(v string) *ListProductsResponseBodyProductInfo {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetProductName(v string) *ListProductsResponseBodyProductInfo {
	s.ProductName = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetProductNameEn(v string) *ListProductsResponseBodyProductInfo {
	s.ProductNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryId(v int64) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryId = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryName(v string) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryName = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryNameEn(v string) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetWhiteListLabelQuotaSupport(v string) *ListProductsResponseBodyProductInfo {
	s.WhiteListLabelQuotaSupport = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) Validate() error {
	return dara.Validate(s)
}

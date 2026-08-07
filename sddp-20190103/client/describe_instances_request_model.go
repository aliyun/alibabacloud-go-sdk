// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInstancesRequest
	GetCurrentPage() *int32
	SetFeatureType(v int32) *DescribeInstancesRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeInstancesRequest
	GetLang() *string
	SetName(v string) *DescribeInstancesRequest
	GetName() *string
	SetPageSize(v int32) *DescribeInstancesRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeInstancesRequest
	GetProductCode() *string
	SetProductId(v int64) *DescribeInstancesRequest
	GetProductId() *int64
	SetRiskLevelId(v int64) *DescribeInstancesRequest
	GetRiskLevelId() *int64
	SetRuleId(v int64) *DescribeInstancesRequest
	GetRuleId() *int64
	SetServiceRegionId(v string) *DescribeInstancesRequest
	GetServiceRegionId() *string
}

type DescribeInstancesRequest struct {
	// The page number of the current page in a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the request and response. Default value: **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese (Simplified).
	//
	// - **en_us**: English (US).
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The search keyword. Fuzzy match is supported. For example, if you enter data, all data entries that contain data in the search item are returned.
	//
	// example:
	//
	// data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of data asset instances to return on each page in a paged query. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the product to which the data asset instance belongs, such as MaxCompute, OSS, or RDS. For supported product names, see [Data types from which sensitive data can be detected](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the product to which the data asset instance belongs. You can call the [DescribeDataAssets](~~DescribeDataAssets~~) operation to obtain the product ID.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The risk level ID of the data asset instance. A higher risk level ID indicates more sensitive data is detected. Valid values:
	//
	// - **1**: No sensitive data is detected. No risk.
	//
	// - **2**: Sensitive data risk at level 1.
	//
	// - **3**: Sensitive data risk at level 2.
	//
	// - **4**: Sensitive data risk at level 3.
	//
	// - **5**: Sensitive data risk at level 4.
	//
	// - **6**: Sensitive data risk at level 5.
	//
	// - **7**: Sensitive data risk at level 6.
	//
	// - **8**: Sensitive data risk at level 7.
	//
	// - **9**: Sensitive data risk at level 8.
	//
	// - **10**: Sensitive data risk at level 9.
	//
	// - **11**: Sensitive data risk at level 10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the data asset instance hits. You can call the [DescribeRules](~~DescribeRules~~) operation and obtain the rule ID from the **Id*	- response parameter.
	//
	// example:
	//
	// 1111111
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The region of the data asset instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/214257.html).
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstancesRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeInstancesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInstancesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstancesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeInstancesRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeInstancesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeInstancesRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *DescribeInstancesRequest) SetCurrentPage(v int32) *DescribeInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesRequest) SetFeatureType(v int32) *DescribeInstancesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeInstancesRequest) SetLang(v string) *DescribeInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstancesRequest) SetName(v string) *DescribeInstancesRequest {
	s.Name = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetProductCode(v string) *DescribeInstancesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstancesRequest) SetProductId(v int64) *DescribeInstancesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeInstancesRequest) SetRiskLevelId(v int64) *DescribeInstancesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeInstancesRequest) SetRuleId(v int64) *DescribeInstancesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeInstancesRequest) SetServiceRegionId(v string) *DescribeInstancesRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	return dara.Validate(s)
}

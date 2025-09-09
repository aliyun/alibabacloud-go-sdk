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
	// The number of the page to return. Default value: **1**.
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
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The keyword that is used to search for data assets. DSC searches for data assets based on the keyword that you specify in fuzzy match mode. For example, if you specify data, all data assets whose names contain data are queried.
	//
	// example:
	//
	// data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the data asset belongs, such as MaxCompute, OSS, and ApsaraDB RDS. For more information about the types of data assets from which DSC can scan for sensitive data, see [Supported data assets](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs. You can call the [DescribeDataAssets](~~DescribeDataAssets~~) operation to query the ID of the service.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level ID of the data asset. A higher sensitivity level indicates that the identified data is more sensitive. Valid values:
	//
	// 	- **1**: No sensitive data is identified.
	//
	// 	- **2**: sensitive data at level 1.
	//
	// 	- **3**: sensitive data at level 2.
	//
	// 	- **4**: sensitive data at level 3
	//
	// 	- **5**: sensitive data at level 4.
	//
	// 	- **6**: sensitive data at level 5.
	//
	// 	- **7**: sensitive data at level 6.
	//
	// 	- **8**: sensitive data at level 7.
	//
	// 	- **9**: sensitive data at level 8.
	//
	// 	- **10**: sensitive data at level 9.
	//
	// 	- **11**: sensitive data at level 10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the data asset hits. You can call the [DescribeRules](~~DescribeRules~~) operation and obtain the ID of the sensitive data detection rule from the **Id*	- response parameter.
	//
	// example:
	//
	// 1111111
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The region where the data asset resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/214257.html).
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

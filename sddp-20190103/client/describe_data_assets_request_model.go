// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataAssetsRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeDataAssetsRequest
	GetLang() *string
	SetName(v string) *DescribeDataAssetsRequest
	GetName() *string
	SetPageSize(v int32) *DescribeDataAssetsRequest
	GetPageSize() *int32
	SetRangeId(v int32) *DescribeDataAssetsRequest
	GetRangeId() *int32
	SetRiskLevels(v string) *DescribeDataAssetsRequest
	GetRiskLevels() *string
	SetRuleId(v int64) *DescribeDataAssetsRequest
	GetRuleId() *int64
}

type DescribeDataAssetsRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
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
	// The keyword that is used to search for data assets. Fuzzy search is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the data asset that you want to query. Valid values:
	//
	// 	- **1**: MaxCompute project
	//
	// 	- **2**: MaxCompute table
	//
	// 	- **3**: MaxCompute package
	//
	// 	- **11**: AnalyticDB for MySQL database
	//
	// 	- **12**: AnalyticDB for MySQL table
	//
	// 	- **21**: Object Storage Service (OSS) bucket
	//
	// 	- **22**: OSS object
	//
	// 	- **31**: Tablestore instance
	//
	// 	- **32**: Tablestore table
	//
	// 	- **51**: ApsaraDB RDS database
	//
	// 	- **52**: ApsaraDB RDS table
	//
	// 	- **61**: self-managed database hosted on an Elastic Compute Service (ECS) instance
	//
	// 	- **62**: self-managed table hosted on an ECS instance
	//
	// 	- **71**: PolarDB-X database
	//
	// 	- **72**: PolarDB-X table
	//
	// 	- **81**: PolarDB database
	//
	// 	- **82**: PolarDB table
	//
	// 	- **91**: AnalyticDB for PostgreSQL database
	//
	// 	- **92**: AnalyticDB for PostgreSQL table
	//
	// example:
	//
	// 1
	RangeId *int32 `json:"RangeId,omitempty" xml:"RangeId,omitempty"`
	// The sensitivity level of the data asset. Separate multiple sensitivity levels with commas (,). Valid values:
	//
	// 	- **2**: S1, indicating the low sensitivity level
	//
	// 	- **3**: S2, indicating the medium sensitivity level
	//
	// 	- **4**: S3, indicating the high sensitivity level
	//
	// 	- **5**: S4, indicating the highest sensitivity level
	//
	// example:
	//
	// 2
	RiskLevels *string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty"`
	// The unique ID of the sensitive data detection rule that the data assets to be queried hit.
	//
	// > If you query sensitive data detection results based on the sensitive data detection rule that the data assets hit, you can call the [DescribeRules](~~DescribeRules~~) operation to query the ID of the sensitive data detection rule.
	//
	// example:
	//
	// 11122200
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeDataAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataAssetsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataAssetsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeDataAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataAssetsRequest) GetRangeId() *int32 {
	return s.RangeId
}

func (s *DescribeDataAssetsRequest) GetRiskLevels() *string {
	return s.RiskLevels
}

func (s *DescribeDataAssetsRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDataAssetsRequest) SetCurrentPage(v int32) *DescribeDataAssetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetLang(v string) *DescribeDataAssetsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetName(v string) *DescribeDataAssetsRequest {
	s.Name = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetPageSize(v int32) *DescribeDataAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetRangeId(v int32) *DescribeDataAssetsRequest {
	s.RangeId = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetRiskLevels(v string) *DescribeDataAssetsRequest {
	s.RiskLevels = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetRuleId(v int64) *DescribeDataAssetsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDataAssetsRequest) Validate() error {
	return dara.Validate(s)
}

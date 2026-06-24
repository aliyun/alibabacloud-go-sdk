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
	// The page number to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the request and response. The default value is **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese (Simplified)
	//
	// - **en_us**: English (US)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The keyword for a fuzzy search of data assets.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. The default value is **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of data asset to query. Valid values:
	//
	// - **1**: MaxCompute project
	//
	// - **2**: MaxCompute table
	//
	// - **3**: MaxCompute package
	//
	// - **11**: AnalyticDB for MySQL database
	//
	// - **12**: AnalyticDB for MySQL table
	//
	// - **21**: OSS bucket
	//
	// - **22**: OSS object
	//
	// - **31**: Tablestore instance
	//
	// - **32**: Tablestore table
	//
	// - **51**: RDS database
	//
	// - **52**: RDS table
	//
	// - **61**: Self-managed database on an ECS instance
	//
	// - **62**: Self-managed table on an ECS instance
	//
	// - **71**: DRDS database
	//
	// - **72**: DRDS table
	//
	// - **81**: PolarDB database
	//
	// - **82**: PolarDB table
	//
	// - **91**: GPDB database
	//
	// - **92**: GPDB table
	//
	// example:
	//
	// 1
	RangeId *int32 `json:"RangeId,omitempty" xml:"RangeId,omitempty"`
	// The risk levels of the data assets to query. Separate multiple risk levels with commas (,).
	//
	// - **2**: S1, low risk level
	//
	// - **3**: S2, medium risk level
	//
	// - **4**: S3, high risk level
	//
	// - **5**: S4, highest risk level
	//
	// example:
	//
	// 2
	RiskLevels *string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty"`
	// The ID of the sensitive data detection rule that the data asset matches.
	//
	// > To find data assets based on the sensitive data detection rules they match, call the [DescribeRules](~~DescribeRules~~) operation to get the rule IDs.
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

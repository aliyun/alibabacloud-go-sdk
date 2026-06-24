// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePackagesRequest
	GetCurrentPage() *int32
	SetInstanceId(v int64) *DescribePackagesRequest
	GetInstanceId() *int64
	SetLang(v string) *DescribePackagesRequest
	GetLang() *string
	SetName(v string) *DescribePackagesRequest
	GetName() *string
	SetPageSize(v int32) *DescribePackagesRequest
	GetPageSize() *int32
	SetProductId(v int64) *DescribePackagesRequest
	GetProductId() *int64
	SetRiskLevelId(v int64) *DescribePackagesRequest
	GetRiskLevelId() *int64
	SetRuleId(v int64) *DescribePackagesRequest
	GetRuleId() *int64
}

type DescribePackagesRequest struct {
	// The page number to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the asset instance to which the data asset package belongs.
	//
	// > To query the list of MaxCompute data asset packages that are authorized for an SDPP connection by instance ID, call the **DescribeInstances*	- operation to obtain the instance ID.
	//
	// example:
	//
	// 12321
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the request and response. The default value is **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The keyword for the search. Fuzzy matching is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the product to which the data asset package belongs.
	//
	// > To query the list of MaxCompute data asset packages that are authorized for an SDPP connection by product ID, call the **DescribeDataAssets*	- operation to obtain the product ID.
	//
	// example:
	//
	// 2566600
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the risk level for the data asset package.
	//
	// - **1**: N/A: No sensitive data is detected.
	//
	// - **2**: S1: Level 1 sensitive data.
	//
	// - **3**: S2: Level 2 sensitive data.
	//
	// - **4**: S3: Level 3 sensitive data.
	//
	// - **5**: S4: Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the data asset package matches.
	//
	// > To query the list of MaxCompute data asset packages that are authorized for an SDPP connection by the ID of a matching sensitive data detection rule, call the **DescribeRules*	- operation to obtain the rule ID.
	//
	// example:
	//
	// 266666
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribePackagesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePackagesRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribePackagesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePackagesRequest) GetName() *string {
	return s.Name
}

func (s *DescribePackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePackagesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribePackagesRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribePackagesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribePackagesRequest) SetCurrentPage(v int32) *DescribePackagesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackagesRequest) SetInstanceId(v int64) *DescribePackagesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackagesRequest) SetLang(v string) *DescribePackagesRequest {
	s.Lang = &v
	return s
}

func (s *DescribePackagesRequest) SetName(v string) *DescribePackagesRequest {
	s.Name = &v
	return s
}

func (s *DescribePackagesRequest) SetPageSize(v int32) *DescribePackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackagesRequest) SetProductId(v int64) *DescribePackagesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribePackagesRequest) SetRiskLevelId(v int64) *DescribePackagesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribePackagesRequest) SetRuleId(v int64) *DescribePackagesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribePackagesRequest) Validate() error {
	return dara.Validate(s)
}

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
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance to which the package belongs.
	//
	// > You can call the **DescribeInstances*	- operation to query the ID of the instance.
	//
	// example:
	//
	// 12321
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the service to which the package belongs.
	//
	// > You can call the **DescribeDataAssets*	- operation to query the ID of the service.
	//
	// example:
	//
	// 2566600
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the package. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the package hits.
	//
	// > You can call the **DescribeRules*	- operation to query the ID of the sensitive data detection rule.
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

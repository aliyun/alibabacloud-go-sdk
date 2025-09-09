// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeTablesRequest
	GetCurrentPage() *int32
	SetInstanceId(v int64) *DescribeTablesRequest
	GetInstanceId() *int64
	SetLang(v string) *DescribeTablesRequest
	GetLang() *string
	SetName(v string) *DescribeTablesRequest
	GetName() *string
	SetPackageId(v int64) *DescribeTablesRequest
	GetPackageId() *int64
	SetPageSize(v int32) *DescribeTablesRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeTablesRequest
	GetProductCode() *string
	SetProductId(v int64) *DescribeTablesRequest
	GetProductId() *int64
	SetRiskLevelId(v int64) *DescribeTablesRequest
	GetRiskLevelId() *int64
	SetRuleId(v int64) *DescribeTablesRequest
	GetRuleId() *int64
	SetServiceRegionId(v string) *DescribeTablesRequest
	GetServiceRegionId() *string
	SetTemplateId(v int64) *DescribeTablesRequest
	GetTemplateId() *int64
}

type DescribeTablesRequest struct {
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the data asset to which the table belongs. You can call the [DescribeInstances](~~DescribeInstances~~) operation to obtain the ID of the data asset.
	//
	// example:
	//
	// 1
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
	// The search keyword. Fuzzy match is supported. For example, if you specify test, all tables whose names contain test are retrieved.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the package to which the table belongs. You can call the [DescribePackages](~~DescribePackages~~) operation to obtain the ID of the package.
	//
	// example:
	//
	// 555555
	PackageId *int64 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the table belongs, such as MaxCompute, OSS, and ApsaraDB RDS. For more information about the types of data assets from which Data Security Center (DSC) can scan for sensitive data, see [Supported data assets](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the table belongs. You can call the [DescribeDataAssets](~~DescribeDataAssets~~) operation to obtain the ID of the service.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the table. Each sensitivity level ID corresponds to a sensitivity level name. Valid values:
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
	// The ID of the sensitive data detection rule that the table hits. You can call the [DescribeRules](~~DescribeRules~~) operation to obtain the ID of the sensitive data detection rule.
	//
	// example:
	//
	// 333322
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The region in which DSC is activated. For more information, see [Supported regions](https://help.aliyun.com/document_detail/214257.html).
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific rule template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTablesRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeTablesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTablesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeTablesRequest) GetPackageId() *int64 {
	return s.PackageId
}

func (s *DescribeTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTablesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeTablesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeTablesRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeTablesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeTablesRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *DescribeTablesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeTablesRequest) SetCurrentPage(v int32) *DescribeTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesRequest) SetInstanceId(v int64) *DescribeTablesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTablesRequest) SetLang(v string) *DescribeTablesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTablesRequest) SetName(v string) *DescribeTablesRequest {
	s.Name = &v
	return s
}

func (s *DescribeTablesRequest) SetPackageId(v int64) *DescribeTablesRequest {
	s.PackageId = &v
	return s
}

func (s *DescribeTablesRequest) SetPageSize(v int32) *DescribeTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesRequest) SetProductCode(v string) *DescribeTablesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeTablesRequest) SetProductId(v int64) *DescribeTablesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeTablesRequest) SetRiskLevelId(v int64) *DescribeTablesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeTablesRequest) SetRuleId(v int64) *DescribeTablesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeTablesRequest) SetServiceRegionId(v string) *DescribeTablesRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeTablesRequest) SetTemplateId(v int64) *DescribeTablesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeTablesRequest) Validate() error {
	return dara.Validate(s)
}

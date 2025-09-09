// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeColumnsRequest
	GetCurrentPage() *int32
	SetEngineType(v string) *DescribeColumnsRequest
	GetEngineType() *string
	SetInstanceId(v int64) *DescribeColumnsRequest
	GetInstanceId() *int64
	SetInstanceName(v string) *DescribeColumnsRequest
	GetInstanceName() *string
	SetLang(v string) *DescribeColumnsRequest
	GetLang() *string
	SetModelTagId(v string) *DescribeColumnsRequest
	GetModelTagId() *string
	SetName(v string) *DescribeColumnsRequest
	GetName() *string
	SetPageSize(v int32) *DescribeColumnsRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeColumnsRequest
	GetProductCode() *string
	SetProductId(v string) *DescribeColumnsRequest
	GetProductId() *string
	SetRiskLevelId(v int64) *DescribeColumnsRequest
	GetRiskLevelId() *int64
	SetRuleId(v int64) *DescribeColumnsRequest
	GetRuleId() *int64
	SetRuleName(v string) *DescribeColumnsRequest
	GetRuleName() *string
	SetSensLevelName(v string) *DescribeColumnsRequest
	GetSensLevelName() *string
	SetTableId(v int64) *DescribeColumnsRequest
	GetTableId() *int64
	SetTableName(v string) *DescribeColumnsRequest
	GetTableName() *string
	SetTemplateId(v string) *DescribeColumnsRequest
	GetTemplateId() *string
	SetTemplateRuleId(v string) *DescribeColumnsRequest
	GetTemplateRuleId() *string
}

type DescribeColumnsRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **MariaDB**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The ID of the instance to which data in the column of the table belongs.
	//
	// > You can call the [DescribeInstances](~~DescribeRules~~) operation to query the IDs of instances.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance to which data in the column of the table belongs.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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
	// The data tag.
	//
	// 	- 101: personal sensitive information
	//
	// 	- 102: personal information
	//
	// example:
	//
	// 101
	ModelTagId *string `json:"ModelTagId,omitempty" xml:"ModelTagId,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// For example, if you enter **test**, all columns whose names contain **test*	- are retrieved.
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
	// The name of the service to which data in the column of the table belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data object belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: Object Storage Service (OSS)
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore (OTS)
	//
	// 	- **5**: ApsaraDB RDS
	//
	// 	- **6**: self-managed database
	//
	// 	- **7**: PolarDB for Xscale (PolarDB-X)
	//
	// 	- **8**: PolarDB
	//
	// 	- **9**: AnalyticDB for PostgreSQL
	//
	// 	- **10**: ApsaraDB for OceanBase
	//
	// 	- **11**: ApsaraDB for MongoDB
	//
	// 	- **25**: ApsaraDB for Redis
	//
	// example:
	//
	// 5
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A
	//
	// 	- **2**: S1
	//
	// 	- **3**: S2
	//
	// 	- **4**: S3
	//
	// 	- **5**: S4
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that data in the column of the table hits.
	//
	// > You can call the [DescribeRules](~~DescribeRules~~) operation to query the IDs of sensitive data detection rules.
	//
	// example:
	//
	// 11111
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the sensitive data detection rule that data in the column of the table hits.
	//
	// example:
	//
	// ID card number (the Chinese mainland)
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The name of the sensitivity level of the data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **N/A**: No sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S2
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	// The ID of the table to which the column belongs.
	//
	// > You can call the [DescribeTables](~~DescribeTables~~) operation to query the IDs of tables.
	//
	// example:
	//
	// 11132334
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the industry-specific classification template.
	//
	// >  You can call the [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html) operation to obtain the IDs of industry-specific classification templates.
	//
	// example:
	//
	// 5
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the template rule that is hit.
	//
	// >  You can call the [DescribeCategoryTemplateRuleList](https://help.aliyun.com/document_detail/410143.html) operation to obtain the IDs of hit template rules.
	//
	// example:
	//
	// 1542
	TemplateRuleId *string `json:"TemplateRuleId,omitempty" xml:"TemplateRuleId,omitempty"`
}

func (s DescribeColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeColumnsRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeColumnsRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeColumnsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeColumnsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeColumnsRequest) GetModelTagId() *string {
	return s.ModelTagId
}

func (s *DescribeColumnsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeColumnsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeColumnsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeColumnsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeColumnsRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeColumnsRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeColumnsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeColumnsRequest) GetSensLevelName() *string {
	return s.SensLevelName
}

func (s *DescribeColumnsRequest) GetTableId() *int64 {
	return s.TableId
}

func (s *DescribeColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeColumnsRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeColumnsRequest) GetTemplateRuleId() *string {
	return s.TemplateRuleId
}

func (s *DescribeColumnsRequest) SetCurrentPage(v int32) *DescribeColumnsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeColumnsRequest) SetEngineType(v string) *DescribeColumnsRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeColumnsRequest) SetInstanceId(v int64) *DescribeColumnsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeColumnsRequest) SetInstanceName(v string) *DescribeColumnsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnsRequest) SetLang(v string) *DescribeColumnsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeColumnsRequest) SetModelTagId(v string) *DescribeColumnsRequest {
	s.ModelTagId = &v
	return s
}

func (s *DescribeColumnsRequest) SetName(v string) *DescribeColumnsRequest {
	s.Name = &v
	return s
}

func (s *DescribeColumnsRequest) SetPageSize(v int32) *DescribeColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsRequest) SetProductCode(v string) *DescribeColumnsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsRequest) SetProductId(v string) *DescribeColumnsRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeColumnsRequest) SetRiskLevelId(v int64) *DescribeColumnsRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeColumnsRequest) SetRuleId(v int64) *DescribeColumnsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeColumnsRequest) SetRuleName(v string) *DescribeColumnsRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeColumnsRequest) SetSensLevelName(v string) *DescribeColumnsRequest {
	s.SensLevelName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableId(v int64) *DescribeColumnsRequest {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableName(v string) *DescribeColumnsRequest {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTemplateId(v string) *DescribeColumnsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeColumnsRequest) SetTemplateRuleId(v string) *DescribeColumnsRequest {
	s.TemplateRuleId = &v
	return s
}

func (s *DescribeColumnsRequest) Validate() error {
	return dara.Validate(s)
}

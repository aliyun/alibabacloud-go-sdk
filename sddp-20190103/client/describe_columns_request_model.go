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
	SetServiceRegionId(v string) *DescribeColumnsRequest
	GetServiceRegionId() *string
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
	// The page number for paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Engine type. Valid values:
	//
	// - **MySQL**.
	//
	// - **MariaDB**.
	//
	// - **Oracle**.
	//
	// - **PostgreSQL**.
	//
	// - **SQLServer**.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The ID of the asset instance to which the column data in the data asset table belongs.
	//
	// > Query column data in data asset tables authorized to connect to Data Security Center using the ID of the asset instance to which the column data in the data asset table belongs. Obtain the asset instance ID by calling the [DescribeInstances](~~DescribeRules~~) API.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset instance to which the column data in the data asset table belongs.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language type for requests and responses. The default value is **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Data tag.
	//
	// - 101: Personal sensitive information
	//
	// - 102: Personal information
	//
	// example:
	//
	// 101
	ModelTagId *string `json:"ModelTagId,omitempty" xml:"ModelTagId,omitempty"`
	// The keyword for search. Supports fuzzy match.
	//
	// For example, entering **test*	- returns all data containing **test**.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of data entries displayed per page in the list.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product name to which the column data in the data asset table belongs. Valid values: **MaxCompute, OSS, ADS, OTS, RDS**, and others.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID corresponding to the product name to which the data object belongs. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// example:
	//
	// 5
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The risk level ID of the sensitive data detection rule. Valid values:
	//
	// - **1**: N/A.
	//
	// - **2**: S1.
	//
	// - **3**: S2.
	//
	// - **4**: S3.
	//
	// - **5**: S4.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The unique ID of the sensitive data detection rule hit by the column data in the asset table.
	//
	// > Query column data in data asset tables authorized to connect to Data Security Center using the ID of the sensitive data detection rule hit by the column data in the asset table. Obtain the sensitive data detection rule ID by calling the [DescribeRules](~~DescribeRules~~) API.
	//
	// example:
	//
	// 11111
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the sensitive data detection rule hit by the column data in the data asset table.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Sensitivity level name. Valid values:
	//
	// - **N/A**: No sensitive data detected.
	//
	// - **S1**: Level 1 sensitive data.
	//
	// - **S2**: Level 2 sensitive data.
	//
	// - **S3**: Level 3 sensitive data.
	//
	// - **S4**: Level 4 sensitive data.
	//
	// example:
	//
	// S2
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The unique ID of the asset table that contains the columns in data asset tables such as MaxCompute and RDS.
	//
	// > Query column data in data asset tables authorized to connect to Data Security Center using the asset table ID. Obtain the asset table ID by calling the [DescribeTables](~~DescribeTables~~) API.
	//
	// example:
	//
	// 11132334
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the data asset table.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// Industry template ID.
	//
	// > Obtain the industry template ID by calling [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html).
	//
	// example:
	//
	// 5
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the hit template rule.
	//
	// > Obtain the hit template rule ID by calling [DescribeCategoryTemplateRuleList](https://help.aliyun.com/document_detail/410143.html).
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

func (s *DescribeColumnsRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
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

func (s *DescribeColumnsRequest) SetServiceRegionId(v string) *DescribeColumnsRequest {
	s.ServiceRegionId = &v
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

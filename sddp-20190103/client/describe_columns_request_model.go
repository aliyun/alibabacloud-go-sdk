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
	// The page number for a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The database engine type. Valid values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The ID of the asset instance to which the column data in the data asset table belongs.
	//
	// > Queries column data in data asset tables authorized for connection by Data Security Center based on the asset instance ID. You can call the [DescribeInstances](~~DescribeRules~~) operation to obtain the instance ID.
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
	// The language of the request and response. Default value: **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The data tag.
	//
	// - 101: personal sensitive information
	//
	// - 102: personal information
	//
	// example:
	//
	// 101
	ModelTagId *string `json:"ModelTagId,omitempty" xml:"ModelTagId,omitempty"`
	// The keyword to search for. Fuzzy match is supported.
	//
	// For example, if you enter **test**, all data entries that contain **test*	- in the search fields are returned.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the product to which the column data in the data asset table belongs. Valid values: **MaxCompute, OSS, ADS, OTS, RDS**, and others.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID that corresponds to the product name to which the data object belongs. Valid values:
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
	// The unique ID of the sensitive data detection rule that the column data matches.
	//
	// > Queries column data in data asset tables authorized for connection by Data Security Center based on the ID of the sensitive data detection rule that the column data matches. You can call the [DescribeRules](~~DescribeRules~~) operation to obtain the rule ID.
	//
	// example:
	//
	// 11111
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the sensitive data detection rule that the column data in the data asset table matches.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The sensitivity level name. Valid values:
	//
	// - **N/A**: No sensitive data is detected.
	//
	// - **S1**: Level-1 sensitive data.
	//
	// - **S2**: Level-2 sensitive data.
	//
	// - **S3**: Level-3 sensitive data.
	//
	// - **S4**: Level-4 sensitive data.
	//
	// example:
	//
	// S2
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The unique ID of the data asset table to which the columns belong in MaxCompute, ApsaraDB RDS, or other assets.
	//
	// > Queries column data in data asset tables authorized for connection by Data Security Center based on the table ID. You can call the [DescribeTables](~~DescribeTables~~) operation to obtain the table ID.
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
	// The industry template ID.
	//
	// > You can call the [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html) operation to obtain the industry template ID.
	//
	// example:
	//
	// 5
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the matched template rule.
	//
	// > You can call the [DescribeCategoryTemplateRuleList](https://help.aliyun.com/document_detail/410143.html) operation to obtain the matched template rule ID.
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

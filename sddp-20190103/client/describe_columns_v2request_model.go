// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnsV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeColumnsV2Request
	GetCurrentPage() *int32
	SetEngineType(v string) *DescribeColumnsV2Request
	GetEngineType() *string
	SetInstanceId(v int64) *DescribeColumnsV2Request
	GetInstanceId() *int64
	SetInstanceName(v string) *DescribeColumnsV2Request
	GetInstanceName() *string
	SetLang(v string) *DescribeColumnsV2Request
	GetLang() *string
	SetName(v string) *DescribeColumnsV2Request
	GetName() *string
	SetPageSize(v int32) *DescribeColumnsV2Request
	GetPageSize() *int32
	SetProductCode(v string) *DescribeColumnsV2Request
	GetProductCode() *string
	SetRiskLevelId(v int64) *DescribeColumnsV2Request
	GetRiskLevelId() *int64
	SetRuleId(v int64) *DescribeColumnsV2Request
	GetRuleId() *int64
	SetRuleName(v string) *DescribeColumnsV2Request
	GetRuleName() *string
	SetSensLevelName(v string) *DescribeColumnsV2Request
	GetSensLevelName() *string
	SetTableId(v string) *DescribeColumnsV2Request
	GetTableId() *string
	SetTableName(v string) *DescribeColumnsV2Request
	GetTableName() *string
}

type DescribeColumnsV2Request struct {
	// When performing a paginated query, sets the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Engine type. Values:
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
	// ID of the asset instance to which the column data in the data asset table belongs.
	//
	// > Query the data in the columns of the data assets authorized by the Data Security Center based on the ID of the asset instance to which the column data in the data asset table belongs. The asset instance ID can be obtained by calling the [DescribeInstances](https://help.aliyun.com/document_detail/141708.html) interface.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Name of the asset instance to which the column data in the data asset table belongs.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Sets the language type for requests and received messages, default is **zh_cn**.
	//
	// Values:
	//
	// - **zh_cn**: Simplified Chinese
	//
	// - **en_us**: English (United States)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Search keyword, supports fuzzy matching.
	//
	// For example, entering **test*	- will search for all data information containing **test*	- in the search items.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// When performing a paginated query, sets the maximum number of data asset instances displayed per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Product name to which the column data in the data asset table belongs. Values: **MaxCompute, OSS, ADS, OTS, RDS**, etc.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Risk level ID of the sensitive data recognition rule. Values:
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
	// Unique identifier ID of the sensitive data recognition rule hit by the column data in the asset table.
	//
	// > Query the data in the columns of the data assets authorized by the Data Security Center based on the ID of the sensitive data recognition rule hit by the column data in the asset table. The sensitive data recognition rule ID can be obtained by calling the [DescribeRules](https://help.aliyun.com/document_detail/141389.html) interface.
	//
	// example:
	//
	// 11122200
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Name of the sensitive data recognition rule hit by the column data in the data asset table.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Sensitive level name. Values:
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
	// Unique identifier ID of the asset table to which the column in MaxCompute, RDS, etc., belongs.
	//
	// > Query the data in the columns of the data assets authorized by the Data Security Center based on the ID of the asset table. The asset table ID can be obtained by calling the [DescribeTables](https://help.aliyun.com/document_detail/141709.html) interface.
	//
	// example:
	//
	// 11132334
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// Name of the data asset table.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsV2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeColumnsV2Request) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeColumnsV2Request) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeColumnsV2Request) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeColumnsV2Request) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeColumnsV2Request) GetLang() *string {
	return s.Lang
}

func (s *DescribeColumnsV2Request) GetName() *string {
	return s.Name
}

func (s *DescribeColumnsV2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeColumnsV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeColumnsV2Request) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeColumnsV2Request) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeColumnsV2Request) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeColumnsV2Request) GetSensLevelName() *string {
	return s.SensLevelName
}

func (s *DescribeColumnsV2Request) GetTableId() *string {
	return s.TableId
}

func (s *DescribeColumnsV2Request) GetTableName() *string {
	return s.TableName
}

func (s *DescribeColumnsV2Request) SetCurrentPage(v int32) *DescribeColumnsV2Request {
	s.CurrentPage = &v
	return s
}

func (s *DescribeColumnsV2Request) SetEngineType(v string) *DescribeColumnsV2Request {
	s.EngineType = &v
	return s
}

func (s *DescribeColumnsV2Request) SetInstanceId(v int64) *DescribeColumnsV2Request {
	s.InstanceId = &v
	return s
}

func (s *DescribeColumnsV2Request) SetInstanceName(v string) *DescribeColumnsV2Request {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnsV2Request) SetLang(v string) *DescribeColumnsV2Request {
	s.Lang = &v
	return s
}

func (s *DescribeColumnsV2Request) SetName(v string) *DescribeColumnsV2Request {
	s.Name = &v
	return s
}

func (s *DescribeColumnsV2Request) SetPageSize(v int32) *DescribeColumnsV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsV2Request) SetProductCode(v string) *DescribeColumnsV2Request {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsV2Request) SetRiskLevelId(v int64) *DescribeColumnsV2Request {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeColumnsV2Request) SetRuleId(v int64) *DescribeColumnsV2Request {
	s.RuleId = &v
	return s
}

func (s *DescribeColumnsV2Request) SetRuleName(v string) *DescribeColumnsV2Request {
	s.RuleName = &v
	return s
}

func (s *DescribeColumnsV2Request) SetSensLevelName(v string) *DescribeColumnsV2Request {
	s.SensLevelName = &v
	return s
}

func (s *DescribeColumnsV2Request) SetTableId(v string) *DescribeColumnsV2Request {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsV2Request) SetTableName(v string) *DescribeColumnsV2Request {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsV2Request) Validate() error {
	return dara.Validate(s)
}

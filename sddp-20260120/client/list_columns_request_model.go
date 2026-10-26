// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListColumnsRequest
	GetCurrentPage() *int32
	SetDataAssetSourceId(v string) *ListColumnsRequest
	GetDataAssetSourceId() *string
	SetDataSourceName(v string) *ListColumnsRequest
	GetDataSourceName() *string
	SetEngineType(v string) *ListColumnsRequest
	GetEngineType() *string
	SetInstanceName(v string) *ListColumnsRequest
	GetInstanceName() *string
	SetLang(v string) *ListColumnsRequest
	GetLang() *string
	SetName(v string) *ListColumnsRequest
	GetName() *string
	SetPageSize(v int32) *ListColumnsRequest
	GetPageSize() *int32
	SetProductCode(v string) *ListColumnsRequest
	GetProductCode() *string
	SetRiskLevelId(v int64) *ListColumnsRequest
	GetRiskLevelId() *int64
	SetRuleId(v int64) *ListColumnsRequest
	GetRuleId() *int64
	SetTableName(v string) *ListColumnsRequest
	GetTableName() *string
	SetTemplateId(v int64) *ListColumnsRequest
	GetTemplateId() *int64
}

type ListColumnsRequest struct {
	// The page number of the current query. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The cloud database instance identity. Set this parameter to the InstanceId in the instance list.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	// The name of the database. Set this parameter to the DbName returned by ListDataSources.
	//
	// example:
	//
	// business_db
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The type of the database engine. This parameter must be consistent with the EngineType of the target instance, such as MySQL or PostgreSQL.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The compatible parameter for the instance identity. DataAssetSourceId takes precedence. If both are specified, DataAssetSourceId is used.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language of the response description. Valid values: zh: Chinese; en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The filter condition for the column name.
	//
	// example:
	//
	// phone
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Keep other filter conditions consistent when querying the next page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the cloud service. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two parameters must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The identifier of the sensitivity level. You can obtain this value from the return list of ListRiskLevels.
	//
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The identifier of the detection rule or model. Set this parameter to the ID of the corresponding detection rule in the target template.
	//
	// example:
	//
	// 1001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the data table. In PostgreSQL or Oracle scenarios, retain the schema.table format obtained from the table list.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The identifier of the data classification and grading template. You can obtain this value from Items[].Id returned by ListTemplates. Keep this parameter consistent with the selected sensitivity level and detection model.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsRequest) GoString() string {
	return s.String()
}

func (s *ListColumnsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListColumnsRequest) GetDataAssetSourceId() *string {
	return s.DataAssetSourceId
}

func (s *ListColumnsRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListColumnsRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ListColumnsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListColumnsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListColumnsRequest) GetName() *string {
	return s.Name
}

func (s *ListColumnsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListColumnsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListColumnsRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *ListColumnsRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListColumnsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListColumnsRequest) SetCurrentPage(v int32) *ListColumnsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListColumnsRequest) SetDataAssetSourceId(v string) *ListColumnsRequest {
	s.DataAssetSourceId = &v
	return s
}

func (s *ListColumnsRequest) SetDataSourceName(v string) *ListColumnsRequest {
	s.DataSourceName = &v
	return s
}

func (s *ListColumnsRequest) SetEngineType(v string) *ListColumnsRequest {
	s.EngineType = &v
	return s
}

func (s *ListColumnsRequest) SetInstanceName(v string) *ListColumnsRequest {
	s.InstanceName = &v
	return s
}

func (s *ListColumnsRequest) SetLang(v string) *ListColumnsRequest {
	s.Lang = &v
	return s
}

func (s *ListColumnsRequest) SetName(v string) *ListColumnsRequest {
	s.Name = &v
	return s
}

func (s *ListColumnsRequest) SetPageSize(v int32) *ListColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *ListColumnsRequest) SetProductCode(v string) *ListColumnsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListColumnsRequest) SetRiskLevelId(v int64) *ListColumnsRequest {
	s.RiskLevelId = &v
	return s
}

func (s *ListColumnsRequest) SetRuleId(v int64) *ListColumnsRequest {
	s.RuleId = &v
	return s
}

func (s *ListColumnsRequest) SetTableName(v string) *ListColumnsRequest {
	s.TableName = &v
	return s
}

func (s *ListColumnsRequest) SetTemplateId(v int64) *ListColumnsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListColumnsRequest) Validate() error {
	return dara.Validate(s)
}

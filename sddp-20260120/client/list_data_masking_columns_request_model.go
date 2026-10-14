// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ListDataMaskingColumnsRequest
	GetColumnName() *string
	SetCurrentPage(v int32) *ListDataMaskingColumnsRequest
	GetCurrentPage() *int32
	SetDbName(v string) *ListDataMaskingColumnsRequest
	GetDbName() *string
	SetEngineType(v string) *ListDataMaskingColumnsRequest
	GetEngineType() *string
	SetInstanceId(v string) *ListDataMaskingColumnsRequest
	GetInstanceId() *string
	SetLang(v string) *ListDataMaskingColumnsRequest
	GetLang() *string
	SetMaskingStatus(v string) *ListDataMaskingColumnsRequest
	GetMaskingStatus() *string
	SetPageSize(v int32) *ListDataMaskingColumnsRequest
	GetPageSize() *int32
	SetProductCode(v string) *ListDataMaskingColumnsRequest
	GetProductCode() *string
	SetProductId(v int64) *ListDataMaskingColumnsRequest
	GetProductId() *int64
	SetProductIds(v string) *ListDataMaskingColumnsRequest
	GetProductIds() *string
	SetRiskLeveLId(v int64) *ListDataMaskingColumnsRequest
	GetRiskLeveLId() *int64
	SetRiskLevelIds(v string) *ListDataMaskingColumnsRequest
	GetRiskLevelIds() *string
	SetTableName(v string) *ListDataMaskingColumnsRequest
	GetTableName() *string
	SetTemplateId(v int64) *ListDataMaskingColumnsRequest
	GetTemplateId() *int64
	SetTemplateRuleIds(v string) *ListDataMaskingColumnsRequest
	GetTemplateRuleIds() *string
}

type ListDataMaskingColumnsRequest struct {
	// The column name filter.
	//
	// example:
	//
	// phone
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The page number of the current query. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The database name.
	//
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The database engine type. This value must match the EngineType of the target instance, for example, MySQL or PostgreSQL.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The instance identity of the database. You can obtain this value from Items[].InstanceId in the response of the ListDataMaskingInstances operation.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The column encryption status filter. Valid values:
	//
	// - NotEncrypted: Not encrypted.
	//
	// - Success: Encryption succeeded.
	//
	// - Failed: Encryption failed.
	//
	// example:
	//
	// NotEncrypted
	MaskingStatus *string `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// The number of records per page. Keep other filter conditions consistent when querying the next page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The cloud product code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two values must correspond.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud product type ID. Valid values:
	//
	// - 5: RDS.
	//
	// - 8: POLARDB.
	//
	// - 27: PolarDBX2.
	//
	// This value must correspond to ProductCode if both are specified.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The compatible product ID list parameter. For column queries, use a single ProductId or ProductCode. Call the operation separately for each product when querying multiple products.
	//
	// example:
	//
	// 5
	ProductIds *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
	// The ID of a sensitivity level. You can obtain this value from the ListRiskLevels operation. The uppercase L in the parameter name must be preserved.
	//
	// example:
	//
	// 1
	RiskLeveLId *int64 `json:"RiskLeveLId,omitempty" xml:"RiskLeveLId,omitempty"`
	// The list of sensitivity level IDs, separated by commas. Select values from the ListRiskLevels response that corresponds to the same TemplateId.
	//
	// example:
	//
	// 1,2
	RiskLevelIds *string `json:"RiskLevelIds,omitempty" xml:"RiskLevelIds,omitempty"`
	// The table name. For PostgreSQL or Oracle scenarios, preserve the schema.table format obtained from the table list.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the data classification template. You can obtain this value from Items[].Id in the response of the ListTemplates operation. This value must be consistent with the selected sensitivity level and detection model.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The list of detection model IDs, separated by commas. Select values from the ListIdentifyModels response that corresponds to the same TemplateId.
	//
	// example:
	//
	// 1001,1002
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
}

func (s ListDataMaskingColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingColumnsRequest) GoString() string {
	return s.String()
}

func (s *ListDataMaskingColumnsRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDataMaskingColumnsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataMaskingColumnsRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataMaskingColumnsRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ListDataMaskingColumnsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataMaskingColumnsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataMaskingColumnsRequest) GetMaskingStatus() *string {
	return s.MaskingStatus
}

func (s *ListDataMaskingColumnsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataMaskingColumnsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataMaskingColumnsRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListDataMaskingColumnsRequest) GetProductIds() *string {
	return s.ProductIds
}

func (s *ListDataMaskingColumnsRequest) GetRiskLeveLId() *int64 {
	return s.RiskLeveLId
}

func (s *ListDataMaskingColumnsRequest) GetRiskLevelIds() *string {
	return s.RiskLevelIds
}

func (s *ListDataMaskingColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDataMaskingColumnsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListDataMaskingColumnsRequest) GetTemplateRuleIds() *string {
	return s.TemplateRuleIds
}

func (s *ListDataMaskingColumnsRequest) SetColumnName(v string) *ListDataMaskingColumnsRequest {
	s.ColumnName = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetCurrentPage(v int32) *ListDataMaskingColumnsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetDbName(v string) *ListDataMaskingColumnsRequest {
	s.DbName = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetEngineType(v string) *ListDataMaskingColumnsRequest {
	s.EngineType = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetInstanceId(v string) *ListDataMaskingColumnsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetLang(v string) *ListDataMaskingColumnsRequest {
	s.Lang = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetMaskingStatus(v string) *ListDataMaskingColumnsRequest {
	s.MaskingStatus = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetPageSize(v int32) *ListDataMaskingColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetProductCode(v string) *ListDataMaskingColumnsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetProductId(v int64) *ListDataMaskingColumnsRequest {
	s.ProductId = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetProductIds(v string) *ListDataMaskingColumnsRequest {
	s.ProductIds = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetRiskLeveLId(v int64) *ListDataMaskingColumnsRequest {
	s.RiskLeveLId = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetRiskLevelIds(v string) *ListDataMaskingColumnsRequest {
	s.RiskLevelIds = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetTableName(v string) *ListDataMaskingColumnsRequest {
	s.TableName = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetTemplateId(v int64) *ListDataMaskingColumnsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) SetTemplateRuleIds(v string) *ListDataMaskingColumnsRequest {
	s.TemplateRuleIds = &v
	return s
}

func (s *ListDataMaskingColumnsRequest) Validate() error {
	return dara.Validate(s)
}

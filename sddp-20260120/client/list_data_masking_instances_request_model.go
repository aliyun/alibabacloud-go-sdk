// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ListDataMaskingInstancesRequest
	GetColumnName() *string
	SetCurrentPage(v int32) *ListDataMaskingInstancesRequest
	GetCurrentPage() *int32
	SetDbName(v string) *ListDataMaskingInstancesRequest
	GetDbName() *string
	SetEngineType(v string) *ListDataMaskingInstancesRequest
	GetEngineType() *string
	SetInstanceId(v string) *ListDataMaskingInstancesRequest
	GetInstanceId() *string
	SetLang(v string) *ListDataMaskingInstancesRequest
	GetLang() *string
	SetMaskingStatus(v string) *ListDataMaskingInstancesRequest
	GetMaskingStatus() *string
	SetModelTagId(v string) *ListDataMaskingInstancesRequest
	GetModelTagId() *string
	SetPageSize(v int32) *ListDataMaskingInstancesRequest
	GetPageSize() *int32
	SetProductCode(v string) *ListDataMaskingInstancesRequest
	GetProductCode() *string
	SetProductId(v int64) *ListDataMaskingInstancesRequest
	GetProductId() *int64
	SetProductIds(v string) *ListDataMaskingInstancesRequest
	GetProductIds() *string
	SetRiskLevelId(v int64) *ListDataMaskingInstancesRequest
	GetRiskLevelId() *int64
	SetRiskLevelIds(v string) *ListDataMaskingInstancesRequest
	GetRiskLevelIds() *string
	SetTableName(v string) *ListDataMaskingInstancesRequest
	GetTableName() *string
	SetTemplateId(v int64) *ListDataMaskingInstancesRequest
	GetTemplateId() *int64
	SetTemplateRuleIds(v string) *ListDataMaskingInstancesRequest
	GetTemplateRuleIds() *string
}

type ListDataMaskingInstancesRequest struct {
	// The filter condition for the column name.
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
	// The name of the database.
	//
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The type of the database engine. The value must be consistent with the EngineType of the target instance, such as MySQL or PostgreSQL.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The database instance identity. You can obtain the identity from the Items[].InstanceId parameter in the response of the ListDataMaskingInstances operation.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response description. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The column encryption status filter. Valid values: NotEncrypted: not encrypted. Success: encrypted. Failed: encryption failed.
	//
	// example:
	//
	// NotEncrypted
	MaskingStatus *string `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// The ID of the model tag. This parameter is used to filter data that matches the specified model tag.
	//
	// example:
	//
	// 101
	ModelTagId *string `json:"ModelTagId,omitempty" xml:"ModelTagId,omitempty"`
	// The number of entries to return on each page. Keep other filter conditions consistent when querying the next page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the cloud service. Column encryption supports RDS, POLARDB, and PolarDBX2. If you specify this parameter and ProductId at the same time, the values must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the cloud service type. Valid values: 5: RDS. 8: POLARDB. 27: PolarDBX2. If you specify this parameter and ProductCode at the same time, the values must match.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The list of cloud service type IDs. Separate multiple values with commas (,). In column encryption scenarios, the valid values are 5 (RDS), 8 (POLARDB), and 27 (PolarDBX2).
	//
	// example:
	//
	// 5
	ProductIds *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
	// The ID of the sensitivity level. You can obtain the ID from the response of the ListRiskLevels operation.
	//
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The list of sensitivity level IDs. Separate multiple values with commas (,). Select the values from the response of the ListRiskLevels operation that corresponds to the same TemplateId.
	//
	// example:
	//
	// 1,2
	RiskLevelIds *string `json:"RiskLevelIds,omitempty" xml:"RiskLevelIds,omitempty"`
	// The name of the table. In PostgreSQL or Oracle scenarios, retain the schema.table format obtained from the table list.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the data classification template. You can obtain the ID from the Items[].Id parameter in the response of the ListTemplates operation. The ID must be consistent with the selected sensitivity level and identification model.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The list of identification model IDs. Separate multiple values with commas (,). Select the values from the response of the ListIdentifyModels operation that corresponds to the same TemplateId.
	//
	// example:
	//
	// 1001,1002
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
}

func (s ListDataMaskingInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDataMaskingInstancesRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDataMaskingInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataMaskingInstancesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataMaskingInstancesRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ListDataMaskingInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataMaskingInstancesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataMaskingInstancesRequest) GetMaskingStatus() *string {
	return s.MaskingStatus
}

func (s *ListDataMaskingInstancesRequest) GetModelTagId() *string {
	return s.ModelTagId
}

func (s *ListDataMaskingInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataMaskingInstancesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataMaskingInstancesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListDataMaskingInstancesRequest) GetProductIds() *string {
	return s.ProductIds
}

func (s *ListDataMaskingInstancesRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *ListDataMaskingInstancesRequest) GetRiskLevelIds() *string {
	return s.RiskLevelIds
}

func (s *ListDataMaskingInstancesRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDataMaskingInstancesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListDataMaskingInstancesRequest) GetTemplateRuleIds() *string {
	return s.TemplateRuleIds
}

func (s *ListDataMaskingInstancesRequest) SetColumnName(v string) *ListDataMaskingInstancesRequest {
	s.ColumnName = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetCurrentPage(v int32) *ListDataMaskingInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetDbName(v string) *ListDataMaskingInstancesRequest {
	s.DbName = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetEngineType(v string) *ListDataMaskingInstancesRequest {
	s.EngineType = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetInstanceId(v string) *ListDataMaskingInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetLang(v string) *ListDataMaskingInstancesRequest {
	s.Lang = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetMaskingStatus(v string) *ListDataMaskingInstancesRequest {
	s.MaskingStatus = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetModelTagId(v string) *ListDataMaskingInstancesRequest {
	s.ModelTagId = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetPageSize(v int32) *ListDataMaskingInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetProductCode(v string) *ListDataMaskingInstancesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetProductId(v int64) *ListDataMaskingInstancesRequest {
	s.ProductId = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetProductIds(v string) *ListDataMaskingInstancesRequest {
	s.ProductIds = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetRiskLevelId(v int64) *ListDataMaskingInstancesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetRiskLevelIds(v string) *ListDataMaskingInstancesRequest {
	s.RiskLevelIds = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetTableName(v string) *ListDataMaskingInstancesRequest {
	s.TableName = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetTemplateId(v int64) *ListDataMaskingInstancesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) SetTemplateRuleIds(v string) *ListDataMaskingInstancesRequest {
	s.TemplateRuleIds = &v
	return s
}

func (s *ListDataMaskingInstancesRequest) Validate() error {
	return dara.Validate(s)
}

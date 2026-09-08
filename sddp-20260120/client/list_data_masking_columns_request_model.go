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
	// example:
	//
	// phone
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// NotEncrypted
	MaskingStatus *string `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 5
	ProductIds *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
	// example:
	//
	// 1
	RiskLeveLId *int64 `json:"RiskLeveLId,omitempty" xml:"RiskLeveLId,omitempty"`
	// example:
	//
	// 1,2
	RiskLevelIds *string `json:"RiskLevelIds,omitempty" xml:"RiskLevelIds,omitempty"`
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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

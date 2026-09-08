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
	// 101
	ModelTagId *string `json:"ModelTagId,omitempty" xml:"ModelTagId,omitempty"`
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
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
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

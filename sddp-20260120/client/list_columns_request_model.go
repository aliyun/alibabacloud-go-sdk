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
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	DataSourceName    *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	EngineType        *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RiskLevelId       *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleId            *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TemplateId        *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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

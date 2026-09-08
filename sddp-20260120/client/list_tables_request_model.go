// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListTablesRequest
	GetCurrentPage() *int32
	SetDataAssetSourceId(v string) *ListTablesRequest
	GetDataAssetSourceId() *string
	SetDataSourceName(v string) *ListTablesRequest
	GetDataSourceName() *string
	SetLang(v string) *ListTablesRequest
	GetLang() *string
	SetMarker(v int64) *ListTablesRequest
	GetMarker() *int64
	SetName(v string) *ListTablesRequest
	GetName() *string
	SetPageSize(v int32) *ListTablesRequest
	GetPageSize() *int32
	SetProductCode(v string) *ListTablesRequest
	GetProductCode() *string
	SetProductId(v int64) *ListTablesRequest
	GetProductId() *int64
	SetRiskLevelId(v int64) *ListTablesRequest
	GetRiskLevelId() *int64
	SetRuleId(v int64) *ListTablesRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *ListTablesRequest
	GetTemplateId() *int64
}

type ListTablesRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	// example:
	//
	// business_db
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 0
	Marker *int64 `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// example:
	//
	// customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// example:
	//
	// 1001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTablesRequest) GetDataAssetSourceId() *string {
	return s.DataAssetSourceId
}

func (s *ListTablesRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListTablesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListTablesRequest) GetMarker() *int64 {
	return s.Marker
}

func (s *ListTablesRequest) GetName() *string {
	return s.Name
}

func (s *ListTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListTablesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListTablesRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *ListTablesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListTablesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListTablesRequest) SetCurrentPage(v int32) *ListTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTablesRequest) SetDataAssetSourceId(v string) *ListTablesRequest {
	s.DataAssetSourceId = &v
	return s
}

func (s *ListTablesRequest) SetDataSourceName(v string) *ListTablesRequest {
	s.DataSourceName = &v
	return s
}

func (s *ListTablesRequest) SetLang(v string) *ListTablesRequest {
	s.Lang = &v
	return s
}

func (s *ListTablesRequest) SetMarker(v int64) *ListTablesRequest {
	s.Marker = &v
	return s
}

func (s *ListTablesRequest) SetName(v string) *ListTablesRequest {
	s.Name = &v
	return s
}

func (s *ListTablesRequest) SetPageSize(v int32) *ListTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablesRequest) SetProductCode(v string) *ListTablesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListTablesRequest) SetProductId(v int64) *ListTablesRequest {
	s.ProductId = &v
	return s
}

func (s *ListTablesRequest) SetRiskLevelId(v int64) *ListTablesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *ListTablesRequest) SetRuleId(v int64) *ListTablesRequest {
	s.RuleId = &v
	return s
}

func (s *ListTablesRequest) SetTemplateId(v int64) *ListTablesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}

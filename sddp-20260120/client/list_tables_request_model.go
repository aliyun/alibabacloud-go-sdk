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
	// The page number of the current query. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance identity of the cloud database instance. Use the InstanceId from the instance list.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	// The database name. Use the DbName returned by ListDataSources.
	//
	// example:
	//
	// business_db
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
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
	// The starting position for cursor-based pagination. Use the NextMarker returned in the previous request for subsequent requests. Keep it consistent with the original filter conditions.
	//
	// example:
	//
	// 0
	Marker *int64 `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The name filter condition.
	//
	// example:
	//
	// customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Keep other filter conditions consistent when querying the next page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The cloud service code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also passed, the two values must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud service type ID. Valid values:
	//
	// - 5: RDS.
	//
	// - 8: POLARDB.
	//
	// - 27: PolarDBX2.
	//
	// If specified together with ProductCode, the values must match.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level ID. You can obtain the value from the returned list of ListRiskLevels.
	//
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The detection rule or model ID. Use the ID of the corresponding detection rule in the target template.
	//
	// example:
	//
	// 1001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The data classification template ID. You can obtain the value from Items[].Id in the response of ListTemplates. Keep it consistent with the selected sensitivity level and detection model.
	//
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

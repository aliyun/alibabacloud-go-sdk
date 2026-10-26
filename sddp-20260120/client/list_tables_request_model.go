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
	// The page number of the current page. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance identity of the cloud database instance. Set the value to the InstanceId value in the instance list.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	// The name of the database. Set the value to the DbName value returned by the ListDataSources operation.
	//
	// example:
	//
	// business_db
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The language of the response. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start position of the cursor-based pagination. Use the NextMarker value returned in the previous request for subsequent requests. Keep the value consistent with the original filter conditions.
	//
	// example:
	//
	// 0
	Marker *int64 `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The name filter.
	//
	// example:
	//
	// customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Keep other filter conditions consistent when you query the next page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the cloud service. Column encryption supports ApsaraDB RDS, PolarDB, and PolarDB-X 2.0. If you specify this parameter and ProductId at the same time, the values must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the cloud service type. Valid values: 5: ApsaraDB RDS. 8: PolarDB. 27: PolarDB-X 2.0. If you specify this parameter and ProductCode at the same time, the values must match.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the sensitivity level. You can obtain the ID from the response of the ListRiskLevels operation.
	//
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the detection rule or model. Set the value to the ID of the corresponding detection rule in the target template.
	//
	// example:
	//
	// 1001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the data classification template. You can obtain the ID from the Items[].Id parameter in the response of the ListTemplates operation. The value must be consistent with the selected sensitivity level and detection model.
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

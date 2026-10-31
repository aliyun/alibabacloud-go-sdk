// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListQualityTemplatesRequestListQuery) *ListQualityTemplatesRequest
	GetListQuery() *ListQualityTemplatesRequestListQuery
	SetOpTenantId(v int64) *ListQualityTemplatesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListQualityTemplatesRequest
	GetOpUserId() *string
}

type ListQualityTemplatesRequest struct {
	// The paged query conditions.
	ListQuery *ListQualityTemplatesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListQualityTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesRequest) GetListQuery() *ListQualityTemplatesRequestListQuery {
	return s.ListQuery
}

func (s *ListQualityTemplatesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityTemplatesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListQualityTemplatesRequest) SetListQuery(v *ListQualityTemplatesRequestListQuery) *ListQualityTemplatesRequest {
	s.ListQuery = v
	return s
}

func (s *ListQualityTemplatesRequest) SetOpTenantId(v int64) *ListQualityTemplatesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityTemplatesRequest) SetOpUserId(v string) *ListQualityTemplatesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListQualityTemplatesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityTemplatesRequestListQuery struct {
	// The rule types. Valid values:
	//
	// - CONSISTENT: Consistency.
	//
	// - EFFECTIVE: Effectiveness.
	//
	// - TIMELINESE: Timeliness.
	//
	// - ACCURATE: Accuracy.
	//
	// - UNIQUENESS: Uniqueness.
	//
	// - COMPLETENESS: Completeness.
	//
	// - STABILITY: Stability.
	//
	// - CUSTOM: Custom.
	CatalogList []*string `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	// Specifies whether the template is owned by the current user.
	CurrentUserOwned *bool `json:"CurrentUserOwned,omitempty" xml:"CurrentUserOwned,omitempty"`
	// The search keyword filter, which supports template names.
	//
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The supported data source types, such as MAX_COMPUTE, MYSQL, and HIVE.
	SupportDataSourceTypeList []*string `json:"SupportDataSourceTypeList,omitempty" xml:"SupportDataSourceTypeList,omitempty" type:"Repeated"`
	// The template owners.
	TemplateOwnerList []*string `json:"TemplateOwnerList,omitempty" xml:"TemplateOwnerList,omitempty" type:"Repeated"`
	// The template sources. Valid values:
	//
	// - SYSTEM: System template.
	//
	// - CUSTOM: Custom template.
	TemplateSourceList []*string `json:"TemplateSourceList,omitempty" xml:"TemplateSourceList,omitempty" type:"Repeated"`
	// The templatetype. Valid values:
	//
	// - FIELD_NULL_VALUE_VALIDATE: Field null value check.
	//
	// - FIELD_EMPTY_STRING_VALIDATE: Field empty character check.
	//
	// - FIELD_UNIQUE_VALIDATE: Field uniqueness check.
	//
	// - FIELD_GROUP_COUNT_VALIDATE: Field unique value count check.
	//
	// - FIELD_DUPLICATE_VALUE_COUNT_VALIDATE: Field duplicate value count check.
	//
	// - FUNCTION_TIME_COMPARE: Time function comparison.
	//
	// - SINGLE_TABLE_TIME_COMPARE: Non-partitioned table time field comparison.
	//
	// - DOUBLE_TABLE_TIME_COMPARE: Two-table time field comparison.
	//
	// - FIELD_FORMAT_VALIDATE: Field format check.
	//
	// - FIELD_LENGTH_VALIDATE: Field length check.
	//
	// - FIELD_VALUE_RANGE_VALIDATE: Field value range check.
	//
	// - CODE_TABLE_COMPARE: Lookup table reference comparison.
	//
	// - STANDARD_CODE_TABLE_COMPARE: Data standard lookup table reference comparison.
	//
	// - SINGLE_TABLE_FIELD_VALUE_COMPARE: Non-partitioned table field value consistency comparison.
	//
	// - SINGLE_TABLE_FIELD_STATISTICAL_COMPARE: Non-partitioned table field statistical value consistency comparison.
	//
	// - SINGLE_TABLE_FIELD_EXP_COMPARE: Non-partitioned table field business logic consistency comparison.
	//
	// - DOUBLE_TABLE_FIELD_VALUE_COMPARE: Two-table field value consistency comparison.
	//
	// - DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: Two-table field statistical value consistency comparison.
	//
	// - CROSS_DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: Cross-source two-table field statistical value consistency comparison.
	//
	// - DOUBLE_TABLE_FIELD_EXP_COMPARE: Two-table field business logic consistency comparison.
	//
	// - TABLE_STABILITY_VALIDATE: Table stability check.
	//
	// - TABLE_FLUCTUATION_VALIDATE: Table fluctuation check.
	//
	// - FIELD_STABILITY_VALIDATE: Field stability check.
	//
	// - FIELD_FLUCTUATION_VALIDATE: Field fluctuation check.
	//
	// - CUSTOM_STATISTICAL_VALIDATE: Custom statistical metric check.
	//
	// - CUSTOM_DATA_DETAILS_VALIDATE: Custom data details check.
	//
	// - DATASOURCE_AVAILABLE_CHECK: Data source connectivity monitoring.
	//
	// - TABLE_SCHEMA_CHECK: Table schema change monitoring.
	//
	// - REAL_TIME_OFFLINE_COMPARE: Real-time and offline comparison.
	//
	// - REAL_TIME_STATISTICAL_VALIDATE: Real-time statistical value monitoring.
	//
	// - REAL_TIME_MULTI_CHAIN_COMPARE: Real-time multi-link comparison.
	TemplateTypeList []*string `json:"TemplateTypeList,omitempty" xml:"TemplateTypeList,omitempty" type:"Repeated"`
	// The monitored object types. Valid values:
	//
	// - TABLE: Dataphin table.
	//
	// - DATASOURCE_TABLE: Global table.
	//
	// - DATASOURCE: Data source.
	//
	// - INDEX: Metric.
	//
	// - REALTIME_LOGICAL_TABLE: Real-time metadata table.
	WatchTypeList []*string `json:"WatchTypeList,omitempty" xml:"WatchTypeList,omitempty" type:"Repeated"`
}

func (s ListQualityTemplatesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesRequestListQuery) GetCatalogList() []*string {
	return s.CatalogList
}

func (s *ListQualityTemplatesRequestListQuery) GetCurrentUserOwned() *bool {
	return s.CurrentUserOwned
}

func (s *ListQualityTemplatesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListQualityTemplatesRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListQualityTemplatesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityTemplatesRequestListQuery) GetSupportDataSourceTypeList() []*string {
	return s.SupportDataSourceTypeList
}

func (s *ListQualityTemplatesRequestListQuery) GetTemplateOwnerList() []*string {
	return s.TemplateOwnerList
}

func (s *ListQualityTemplatesRequestListQuery) GetTemplateSourceList() []*string {
	return s.TemplateSourceList
}

func (s *ListQualityTemplatesRequestListQuery) GetTemplateTypeList() []*string {
	return s.TemplateTypeList
}

func (s *ListQualityTemplatesRequestListQuery) GetWatchTypeList() []*string {
	return s.WatchTypeList
}

func (s *ListQualityTemplatesRequestListQuery) SetCatalogList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.CatalogList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetCurrentUserOwned(v bool) *ListQualityTemplatesRequestListQuery {
	s.CurrentUserOwned = &v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetKeyword(v string) *ListQualityTemplatesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetPageNo(v int32) *ListQualityTemplatesRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetPageSize(v int32) *ListQualityTemplatesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetSupportDataSourceTypeList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.SupportDataSourceTypeList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetTemplateOwnerList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.TemplateOwnerList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetTemplateSourceList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.TemplateSourceList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetTemplateTypeList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.TemplateTypeList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetWatchTypeList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.WatchTypeList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) Validate() error {
	return dara.Validate(s)
}

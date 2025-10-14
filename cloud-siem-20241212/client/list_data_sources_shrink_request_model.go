// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceFrom(v string) *ListDataSourcesShrinkRequest
	GetDataSourceFrom() *string
	SetDataSourceIdsShrink(v string) *ListDataSourcesShrinkRequest
	GetDataSourceIdsShrink() *string
	SetDataSourceName(v string) *ListDataSourcesShrinkRequest
	GetDataSourceName() *string
	SetDataSourceStatus(v string) *ListDataSourcesShrinkRequest
	GetDataSourceStatus() *string
	SetDataSourceStoreStatus(v string) *ListDataSourcesShrinkRequest
	GetDataSourceStoreStatus() *string
	SetDataSourceTemplateIdsShrink(v string) *ListDataSourcesShrinkRequest
	GetDataSourceTemplateIdsShrink() *string
	SetDataSourceType(v string) *ListDataSourcesShrinkRequest
	GetDataSourceType() *string
	SetLang(v string) *ListDataSourcesShrinkRequest
	GetLang() *string
	SetLogProjectName(v string) *ListDataSourcesShrinkRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *ListDataSourcesShrinkRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *ListDataSourcesShrinkRequest
	GetLogStoreName() *string
	SetLogUserIdsShrink(v string) *ListDataSourcesShrinkRequest
	GetLogUserIdsShrink() *string
	SetMaxResults(v int32) *ListDataSourcesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSourcesShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *ListDataSourcesShrinkRequest
	GetOrder() *string
	SetOrderField(v string) *ListDataSourcesShrinkRequest
	GetOrderField() *string
	SetPageNumber(v string) *ListDataSourcesShrinkRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataSourcesShrinkRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDataSourcesShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataSourcesShrinkRequest
	GetRoleFor() *int64
}

type ListDataSourcesShrinkRequest struct {
	// example:
	//
	// center。
	DataSourceFrom      *string `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	DataSourceIdsShrink *string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty"`
	// example:
	//
	// alibaba_cloud_waf_alert_log。
	DataSourceName              *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceStatus            *string `json:"DataSourceStatus,omitempty" xml:"DataSourceStatus,omitempty"`
	DataSourceStoreStatus       *string `json:"DataSourceStoreStatus,omitempty" xml:"DataSourceStoreStatus,omitempty"`
	DataSourceTemplateIdsShrink *string `json:"DataSourceTemplateIds,omitempty" xml:"DataSourceTemplateIds,omitempty"`
	// example:
	//
	// custom。
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou。
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// example:
	//
	// audit-activity。
	LogStoreName     *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	LogUserIdsShrink *string `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// desc。
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// UpdateTime。
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// example:
	//
	// 1。
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5。
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDataSourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesShrinkRequest) GetDataSourceFrom() *string {
	return s.DataSourceFrom
}

func (s *ListDataSourcesShrinkRequest) GetDataSourceIdsShrink() *string {
	return s.DataSourceIdsShrink
}

func (s *ListDataSourcesShrinkRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListDataSourcesShrinkRequest) GetDataSourceStatus() *string {
	return s.DataSourceStatus
}

func (s *ListDataSourcesShrinkRequest) GetDataSourceStoreStatus() *string {
	return s.DataSourceStoreStatus
}

func (s *ListDataSourcesShrinkRequest) GetDataSourceTemplateIdsShrink() *string {
	return s.DataSourceTemplateIdsShrink
}

func (s *ListDataSourcesShrinkRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDataSourcesShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSourcesShrinkRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *ListDataSourcesShrinkRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ListDataSourcesShrinkRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ListDataSourcesShrinkRequest) GetLogUserIdsShrink() *string {
	return s.LogUserIdsShrink
}

func (s *ListDataSourcesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSourcesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSourcesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDataSourcesShrinkRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListDataSourcesShrinkRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataSourcesShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataSourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourcesShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataSourcesShrinkRequest) SetDataSourceFrom(v string) *ListDataSourcesShrinkRequest {
	s.DataSourceFrom = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetDataSourceIdsShrink(v string) *ListDataSourcesShrinkRequest {
	s.DataSourceIdsShrink = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetDataSourceName(v string) *ListDataSourcesShrinkRequest {
	s.DataSourceName = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetDataSourceStatus(v string) *ListDataSourcesShrinkRequest {
	s.DataSourceStatus = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetDataSourceStoreStatus(v string) *ListDataSourcesShrinkRequest {
	s.DataSourceStoreStatus = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetDataSourceTemplateIdsShrink(v string) *ListDataSourcesShrinkRequest {
	s.DataSourceTemplateIdsShrink = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetDataSourceType(v string) *ListDataSourcesShrinkRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetLang(v string) *ListDataSourcesShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetLogProjectName(v string) *ListDataSourcesShrinkRequest {
	s.LogProjectName = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetLogRegionId(v string) *ListDataSourcesShrinkRequest {
	s.LogRegionId = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetLogStoreName(v string) *ListDataSourcesShrinkRequest {
	s.LogStoreName = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetLogUserIdsShrink(v string) *ListDataSourcesShrinkRequest {
	s.LogUserIdsShrink = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetMaxResults(v int32) *ListDataSourcesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetNextToken(v string) *ListDataSourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetOrder(v string) *ListDataSourcesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetOrderField(v string) *ListDataSourcesShrinkRequest {
	s.OrderField = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetPageNumber(v string) *ListDataSourcesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetPageSize(v string) *ListDataSourcesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetRegionId(v string) *ListDataSourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetRoleFor(v int64) *ListDataSourcesShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

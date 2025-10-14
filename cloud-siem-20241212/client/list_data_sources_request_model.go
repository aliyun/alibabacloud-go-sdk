// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceFrom(v string) *ListDataSourcesRequest
	GetDataSourceFrom() *string
	SetDataSourceIds(v []*string) *ListDataSourcesRequest
	GetDataSourceIds() []*string
	SetDataSourceName(v string) *ListDataSourcesRequest
	GetDataSourceName() *string
	SetDataSourceStatus(v string) *ListDataSourcesRequest
	GetDataSourceStatus() *string
	SetDataSourceStoreStatus(v string) *ListDataSourcesRequest
	GetDataSourceStoreStatus() *string
	SetDataSourceTemplateIds(v []*string) *ListDataSourcesRequest
	GetDataSourceTemplateIds() []*string
	SetDataSourceType(v string) *ListDataSourcesRequest
	GetDataSourceType() *string
	SetLang(v string) *ListDataSourcesRequest
	GetLang() *string
	SetLogProjectName(v string) *ListDataSourcesRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *ListDataSourcesRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *ListDataSourcesRequest
	GetLogStoreName() *string
	SetLogUserIds(v []*int64) *ListDataSourcesRequest
	GetLogUserIds() []*int64
	SetMaxResults(v int32) *ListDataSourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSourcesRequest
	GetNextToken() *string
	SetOrder(v string) *ListDataSourcesRequest
	GetOrder() *string
	SetOrderField(v string) *ListDataSourcesRequest
	GetOrderField() *string
	SetPageNumber(v string) *ListDataSourcesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataSourcesRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDataSourcesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataSourcesRequest
	GetRoleFor() *int64
}

type ListDataSourcesRequest struct {
	// example:
	//
	// center。
	DataSourceFrom *string   `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	DataSourceIds  []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// alibaba_cloud_waf_alert_log。
	DataSourceName        *string   `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceStatus      *string   `json:"DataSourceStatus,omitempty" xml:"DataSourceStatus,omitempty"`
	DataSourceStoreStatus *string   `json:"DataSourceStoreStatus,omitempty" xml:"DataSourceStoreStatus,omitempty"`
	DataSourceTemplateIds []*string `json:"DataSourceTemplateIds,omitempty" xml:"DataSourceTemplateIds,omitempty" type:"Repeated"`
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
	LogStoreName *string  `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	LogUserIds   []*int64 `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty" type:"Repeated"`
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

func (s ListDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) GetDataSourceFrom() *string {
	return s.DataSourceFrom
}

func (s *ListDataSourcesRequest) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *ListDataSourcesRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListDataSourcesRequest) GetDataSourceStatus() *string {
	return s.DataSourceStatus
}

func (s *ListDataSourcesRequest) GetDataSourceStoreStatus() *string {
	return s.DataSourceStoreStatus
}

func (s *ListDataSourcesRequest) GetDataSourceTemplateIds() []*string {
	return s.DataSourceTemplateIds
}

func (s *ListDataSourcesRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDataSourcesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSourcesRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *ListDataSourcesRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ListDataSourcesRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ListDataSourcesRequest) GetLogUserIds() []*int64 {
	return s.LogUserIds
}

func (s *ListDataSourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSourcesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDataSourcesRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListDataSourcesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataSourcesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataSourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourcesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataSourcesRequest) SetDataSourceFrom(v string) *ListDataSourcesRequest {
	s.DataSourceFrom = &v
	return s
}

func (s *ListDataSourcesRequest) SetDataSourceIds(v []*string) *ListDataSourcesRequest {
	s.DataSourceIds = v
	return s
}

func (s *ListDataSourcesRequest) SetDataSourceName(v string) *ListDataSourcesRequest {
	s.DataSourceName = &v
	return s
}

func (s *ListDataSourcesRequest) SetDataSourceStatus(v string) *ListDataSourcesRequest {
	s.DataSourceStatus = &v
	return s
}

func (s *ListDataSourcesRequest) SetDataSourceStoreStatus(v string) *ListDataSourcesRequest {
	s.DataSourceStoreStatus = &v
	return s
}

func (s *ListDataSourcesRequest) SetDataSourceTemplateIds(v []*string) *ListDataSourcesRequest {
	s.DataSourceTemplateIds = v
	return s
}

func (s *ListDataSourcesRequest) SetDataSourceType(v string) *ListDataSourcesRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListDataSourcesRequest) SetLang(v string) *ListDataSourcesRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSourcesRequest) SetLogProjectName(v string) *ListDataSourcesRequest {
	s.LogProjectName = &v
	return s
}

func (s *ListDataSourcesRequest) SetLogRegionId(v string) *ListDataSourcesRequest {
	s.LogRegionId = &v
	return s
}

func (s *ListDataSourcesRequest) SetLogStoreName(v string) *ListDataSourcesRequest {
	s.LogStoreName = &v
	return s
}

func (s *ListDataSourcesRequest) SetLogUserIds(v []*int64) *ListDataSourcesRequest {
	s.LogUserIds = v
	return s
}

func (s *ListDataSourcesRequest) SetMaxResults(v int32) *ListDataSourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataSourcesRequest) SetNextToken(v string) *ListDataSourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataSourcesRequest) SetOrder(v string) *ListDataSourcesRequest {
	s.Order = &v
	return s
}

func (s *ListDataSourcesRequest) SetOrderField(v string) *ListDataSourcesRequest {
	s.OrderField = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageNumber(v string) *ListDataSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageSize(v string) *ListDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesRequest) SetRegionId(v string) *ListDataSourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourcesRequest) SetRoleFor(v int64) *ListDataSourcesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}

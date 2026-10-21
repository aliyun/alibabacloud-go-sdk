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
	// The data source origin. Valid values:
	//
	// - center
	//
	// - custom
	//
	// example:
	//
	// center
	DataSourceFrom *string `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	// The list of data source IDs.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// The data source name.
	//
	// example:
	//
	// alibaba_cloud_waf_alert_log
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The data source status. Valid values:
	//
	// - unconfigured
	//
	// - normal
	//
	// - abnormal
	//
	// example:
	//
	// unconfigured
	DataSourceStatus *string `json:"DataSourceStatus,omitempty" xml:"DataSourceStatus,omitempty"`
	// The Logstore status. Valid values:
	//
	// - normal
	//
	// - abnormal
	//
	// example:
	//
	// normal
	DataSourceStoreStatus *string `json:"DataSourceStoreStatus,omitempty" xml:"DataSourceStoreStatus,omitempty"`
	// The list of data source template IDs.
	DataSourceTemplateIds []*string `json:"DataSourceTemplateIds,omitempty" xml:"DataSourceTemplateIds,omitempty" type:"Repeated"`
	// The data source type. Valid values:
	//
	//  - preset
	//
	//  - custom
	//
	// example:
	//
	// custom
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required. If you do not specify this parameter, the service returns a 400 error: "The specified Lang can not be empty." Explicitly pass a language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The Simple Log Service project name.
	//
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// The log storage region ID.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The Simple Log Service Logstore name.
	//
	// example:
	//
	// audit-activity
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The list of data ingestion user IDs.
	LogUserIds []*int64 `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return in this request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Valid values:
	//
	// - desc
	//
	// - asc
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field used for sorting.
	//
	// example:
	//
	// UpdateTime
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the threat analysis data management center is located. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: The assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 173326*******
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

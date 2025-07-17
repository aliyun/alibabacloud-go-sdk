// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDataSourcesResponseBodyPagingInfo) *ListDataSourcesResponseBody
	GetPagingInfo() *ListDataSourcesResponseBodyPagingInfo
	SetRequestId(v string) *ListDataSourcesResponseBody
	GetRequestId() *string
}

type ListDataSourcesResponseBody struct {
	// The pagination information.
	PagingInfo *ListDataSourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7BE1433F-6D55-5D86-9344-CA6F7DD19B13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) GetPagingInfo() *ListDataSourcesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourcesResponseBody) SetPagingInfo(v *ListDataSourcesResponseBodyPagingInfo) *ListDataSourcesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataSourcesResponseBodyPagingInfo struct {
	// The data source groups. Each element in the array indicates a data source group. Each data source group contains data sources in the development environment (if any) and the production environment.
	DataSources []*ListDataSourcesResponseBodyPagingInfoDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 131
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfo) GetDataSources() []*ListDataSourcesResponseBodyPagingInfoDataSources {
	return s.DataSources
}

func (s *ListDataSourcesResponseBodyPagingInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataSourcesResponseBodyPagingInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataSourcesResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetDataSources(v []*ListDataSourcesResponseBodyPagingInfoDataSources) *ListDataSourcesResponseBodyPagingInfo {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetPageNumber(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetPageSize(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetTotalCount(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDataSourcesResponseBodyPagingInfoDataSources struct {
	// The data sources. Each element is the information of a single data source with a unique data source ID.
	DataSource []*ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Repeated"`
	// The name of the data source.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// mysql
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfoDataSources) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfoDataSources) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) GetDataSource() []*ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	return s.DataSource
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) GetName() *string {
	return s.Name
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) GetType() *string {
	return s.Type
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetDataSource(v []*ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.DataSource = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetName(v string) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.Name = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetType(v string) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.Type = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) Validate() error {
	return dara.Validate(s)
}

type ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource struct {
	// The connection configurations of the data source, including the connection address, access identity, and environment information. The envType parameter specifies the environment in which the data source is used. Valid values of the envType parameter:
	//
	// 	- Dev: development environment
	//
	// 	- Prod: production environment
	//
	// The parameters that you need to configure for the data source vary based on the mode in which the data source is added. For more information, see [Data source connection information (ConnectionProperties)](https://help.aliyun.com/document_detail/2852465.html).
	//
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The mode in which the data source is added. The mode varies based on the data source type. Valid values:
	//
	// 	- InstanceMode: instance mode
	//
	// 	- UrlMode: connection string mode
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The time when the data source was added. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1648711113000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who adds the data source.
	//
	// example:
	//
	// 1624387842781448
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// 16035
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the data source was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1648711113000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the user who modifies the data source.
	//
	// example:
	//
	// 1624387842781448
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The unique business key of the data source. For example, the unique business key of a Hologres data source is in the `${tenantOwnerId}:${regionId}:${type}:${instanceId}:${database}` format.
	//
	// example:
	//
	// 1648711121000:cn-beijing:odps:yongxunQA_beijing_standard
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetConnectionProperties() interface{} {
	return s.ConnectionProperties
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetDescription() *string {
	return s.Description
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetId() *int64 {
	return s.Id
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetConnectionProperties(v interface{}) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ConnectionProperties = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetConnectionPropertiesMode(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetCreateTime(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetCreateUser(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.CreateUser = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetDescription(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.Description = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetId(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.Id = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetModifyTime(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ModifyTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetModifyUser(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ModifyUser = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetQualifiedName(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.QualifiedName = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) Validate() error {
	return dara.Validate(s)
}

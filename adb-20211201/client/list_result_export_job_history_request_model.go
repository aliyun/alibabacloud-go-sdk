// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResultExportJobHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListResultExportJobHistoryRequest
	GetDBClusterId() *string
	SetDatabaseUser(v string) *ListResultExportJobHistoryRequest
	GetDatabaseUser() *string
	SetEndTime(v string) *ListResultExportJobHistoryRequest
	GetEndTime() *string
	SetOrder(v *ListResultExportJobHistoryRequestOrder) *ListResultExportJobHistoryRequest
	GetOrder() *ListResultExportJobHistoryRequestOrder
	SetPageNumber(v string) *ListResultExportJobHistoryRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListResultExportJobHistoryRequest
	GetPageSize() *string
	SetRegionId(v string) *ListResultExportJobHistoryRequest
	GetRegionId() *string
	SetResourceGroup(v string) *ListResultExportJobHistoryRequest
	GetResourceGroup() *string
	SetStartTime(v string) *ListResultExportJobHistoryRequest
	GetStartTime() *string
	SetStatusList(v []*string) *ListResultExportJobHistoryRequest
	GetStatusList() []*string
}

type ListResultExportJobHistoryRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-7xv5ty5m9o4v****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// test1
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2023-05-25T06:54:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The order in which to sort the SQL statements by field, which contains the `Field` and `Type` fields. Specify the order in the JSON format. Example: `[{"Field":"CreateTimee", "Type": "desc" }]`.
	//
	// 	- `Field` specifies the field that is used to sort the SQL statements. Valid values:
	//
	//     	- `CreateTime`: the time when the result set export job was created.
	//
	//     	- `Status`: the execution status.
	//
	//     	- `DatabaseUser`: the name of the database account.
	//
	//     	- `TimeCost`: the execution duration.
	//
	//     	- `ResourceGroup`: the name of the resource group.
	//
	//     	- `ExportRows`: the number of exported rows.
	//
	//     	- `Progress`: the export progress.
	//
	// 	- `Type` specifies the sorting order. Valid values (case-insensitive):
	//
	//     	- `Desc`: descending order.
	//
	//     	- `Asc`: ascending order.
	Order *ListResultExportJobHistoryRequestOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the resource group that runs the result set export jobs. You can use this parameter to query the execution records of export jobs that are run in a specific resource group.
	//
	// example:
	//
	// user_default
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-01T12:01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of result set export jobs. You can use this parameter to query the execution records of export jobs that are in a specific state.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListResultExportJobHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResultExportJobHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListResultExportJobHistoryRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListResultExportJobHistoryRequest) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *ListResultExportJobHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListResultExportJobHistoryRequest) GetOrder() *ListResultExportJobHistoryRequestOrder {
	return s.Order
}

func (s *ListResultExportJobHistoryRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListResultExportJobHistoryRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListResultExportJobHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResultExportJobHistoryRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ListResultExportJobHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListResultExportJobHistoryRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListResultExportJobHistoryRequest) SetDBClusterId(v string) *ListResultExportJobHistoryRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetDatabaseUser(v string) *ListResultExportJobHistoryRequest {
	s.DatabaseUser = &v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetEndTime(v string) *ListResultExportJobHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetOrder(v *ListResultExportJobHistoryRequestOrder) *ListResultExportJobHistoryRequest {
	s.Order = v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetPageNumber(v string) *ListResultExportJobHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetPageSize(v string) *ListResultExportJobHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetRegionId(v string) *ListResultExportJobHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetResourceGroup(v string) *ListResultExportJobHistoryRequest {
	s.ResourceGroup = &v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetStartTime(v string) *ListResultExportJobHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *ListResultExportJobHistoryRequest) SetStatusList(v []*string) *ListResultExportJobHistoryRequest {
	s.StatusList = v
	return s
}

func (s *ListResultExportJobHistoryRequest) Validate() error {
	return dara.Validate(s)
}

type ListResultExportJobHistoryRequestOrder struct {
	// The field that is used to sort the SQL statements. Valid values:
	//
	// 	- CreateTime
	//
	// 	- DatabaseUser
	//
	// 	- TimeCost
	//
	// 	- ResourceGroup
	//
	// 	- Status
	//
	// 	- Progress
	//
	// 	- ExportRows
	//
	// example:
	//
	// DatabaseUser
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The sorting order. Valid values (case-insensitive):
	//
	// 	- **Desc**: descending order.
	//
	// 	- **Asc**: ascending order.
	//
	// example:
	//
	// Desc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResultExportJobHistoryRequestOrder) String() string {
	return dara.Prettify(s)
}

func (s ListResultExportJobHistoryRequestOrder) GoString() string {
	return s.String()
}

func (s *ListResultExportJobHistoryRequestOrder) GetField() *string {
	return s.Field
}

func (s *ListResultExportJobHistoryRequestOrder) GetType() *string {
	return s.Type
}

func (s *ListResultExportJobHistoryRequestOrder) SetField(v string) *ListResultExportJobHistoryRequestOrder {
	s.Field = &v
	return s
}

func (s *ListResultExportJobHistoryRequestOrder) SetType(v string) *ListResultExportJobHistoryRequestOrder {
	s.Type = &v
	return s
}

func (s *ListResultExportJobHistoryRequestOrder) Validate() error {
	return dara.Validate(s)
}

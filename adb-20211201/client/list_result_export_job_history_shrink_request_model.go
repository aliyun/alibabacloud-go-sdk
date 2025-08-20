// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResultExportJobHistoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListResultExportJobHistoryShrinkRequest
	GetDBClusterId() *string
	SetDatabaseUser(v string) *ListResultExportJobHistoryShrinkRequest
	GetDatabaseUser() *string
	SetEndTime(v string) *ListResultExportJobHistoryShrinkRequest
	GetEndTime() *string
	SetOrderShrink(v string) *ListResultExportJobHistoryShrinkRequest
	GetOrderShrink() *string
	SetPageNumber(v string) *ListResultExportJobHistoryShrinkRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListResultExportJobHistoryShrinkRequest
	GetPageSize() *string
	SetRegionId(v string) *ListResultExportJobHistoryShrinkRequest
	GetRegionId() *string
	SetResourceGroup(v string) *ListResultExportJobHistoryShrinkRequest
	GetResourceGroup() *string
	SetStartTime(v string) *ListResultExportJobHistoryShrinkRequest
	GetStartTime() *string
	SetStatusListShrink(v string) *ListResultExportJobHistoryShrinkRequest
	GetStatusListShrink() *string
}

type ListResultExportJobHistoryShrinkRequest struct {
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
	OrderShrink *string `json:"Order,omitempty" xml:"Order,omitempty"`
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
	StatusListShrink *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
}

func (s ListResultExportJobHistoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResultExportJobHistoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResultExportJobHistoryShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListResultExportJobHistoryShrinkRequest) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *ListResultExportJobHistoryShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListResultExportJobHistoryShrinkRequest) GetOrderShrink() *string {
	return s.OrderShrink
}

func (s *ListResultExportJobHistoryShrinkRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListResultExportJobHistoryShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListResultExportJobHistoryShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResultExportJobHistoryShrinkRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ListResultExportJobHistoryShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListResultExportJobHistoryShrinkRequest) GetStatusListShrink() *string {
	return s.StatusListShrink
}

func (s *ListResultExportJobHistoryShrinkRequest) SetDBClusterId(v string) *ListResultExportJobHistoryShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetDatabaseUser(v string) *ListResultExportJobHistoryShrinkRequest {
	s.DatabaseUser = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetEndTime(v string) *ListResultExportJobHistoryShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetOrderShrink(v string) *ListResultExportJobHistoryShrinkRequest {
	s.OrderShrink = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetPageNumber(v string) *ListResultExportJobHistoryShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetPageSize(v string) *ListResultExportJobHistoryShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetRegionId(v string) *ListResultExportJobHistoryShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetResourceGroup(v string) *ListResultExportJobHistoryShrinkRequest {
	s.ResourceGroup = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetStartTime(v string) *ListResultExportJobHistoryShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) SetStatusListShrink(v string) *ListResultExportJobHistoryShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *ListResultExportJobHistoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}

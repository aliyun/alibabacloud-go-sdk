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
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (3.0) clusters in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-7xv5ty5m9o4v****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database user.
	//
	// example:
	//
	// test1
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// The end time of the query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2023-05-25T06:54:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sorting conditions for SQL statements based on specified fields. The value is in JSON format and is an ordered JSON array that specifies composite sorting based on the sequence of the input array. The array contains the `Field` and `Type` fields. Example: `[{"Field":"CreateTimee", "Type": "desc" }]`. Details:
	//
	// - `Field`: the name of the field by which to sort the statements. Valid values:
	//
	//     - `CreateTime`: the execution start time.
	//
	//     - `Status`: the execution status.
	//
	//     - `DatabaseUser`: the username.
	//
	//     - `TimeCost`: the execution duration.
	//
	//     - `ResourceGroup`: the resource group.
	//
	//     - `ExportRows`: the number of exported rows.
	//
	//     - `Progress`: the export progress.
	//
	//
	// - `Type`: the sorting type. Valid values (case-insensitive):
	//
	//     - `Desc`: descending order.
	//
	//     - `Asc`: ascending order.
	Order *ListResultExportJobHistoryRequestOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The page number. The value must be an integer greater than 0 and less than or equal to the maximum value of the Integer data type.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30*	- (default)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group in which the task is executed. You can filter tasks by resource group.
	//
	// example:
	//
	// user_default
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The start time of the query. Specify the time in the <i>yyyy-MM-ddTHH:mm:ssZ</i> format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-01T12:01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of task execution statuses. You can filter tasks by status.
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
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResultExportJobHistoryRequestOrder struct {
	// The name of the field by which to sort. Valid values:
	//
	// - CreateTime
	//
	// - DatabaseUser
	//
	// - TimeCost
	//
	// - ResourceGroup
	//
	// - Status
	//
	// - Progress
	//
	// - ExportRows
	//
	// example:
	//
	// DatabaseUser
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The sorting type. Valid values:
	//
	// - **Desc**: descending order.
	//
	// - **Asc**: ascending order.
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

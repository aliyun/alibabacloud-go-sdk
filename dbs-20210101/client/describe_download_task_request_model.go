// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *DescribeDownloadTaskRequest
	GetBackupSetId() *string
	SetClusterName(v string) *DescribeDownloadTaskRequest
	GetClusterName() *string
	SetCurrentPage(v string) *DescribeDownloadTaskRequest
	GetCurrentPage() *string
	SetDatasourceId(v string) *DescribeDownloadTaskRequest
	GetDatasourceId() *string
	SetEndTime(v string) *DescribeDownloadTaskRequest
	GetEndTime() *string
	SetInstanceName(v string) *DescribeDownloadTaskRequest
	GetInstanceName() *string
	SetOrderColumn(v string) *DescribeDownloadTaskRequest
	GetOrderColumn() *string
	SetOrderDirect(v string) *DescribeDownloadTaskRequest
	GetOrderDirect() *string
	SetPageSize(v string) *DescribeDownloadTaskRequest
	GetPageSize() *string
	SetRegionCode(v string) *DescribeDownloadTaskRequest
	GetRegionCode() *string
	SetStartTime(v string) *DescribeDownloadTaskRequest
	GetStartTime() *string
	SetState(v string) *DescribeDownloadTaskRequest
	GetState() *string
	SetTaskType(v string) *DescribeDownloadTaskRequest
	GetTaskType() *string
}

type DescribeDownloadTaskRequest struct {
	// The ID of the backup set generated when you create a download task. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/26273.html) operation to query the ID.
	//
	// example:
	//
	// 216****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// example:
	//
	// dds-example
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the Database Backup (DBS) data source. Specify the parameter in the format of *ds-${Instance ID}_${regionId}*.
	//
	// example:
	//
	// ds-rm-2ze8g2am97624****_cn-hangzhou
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The end of the time range to query. Specify this parameter as a timestamp of the LONG type. Unit: milliseconds.
	//
	// example:
	//
	// 1661941556000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// > This parameter is required.
	//
	// example:
	//
	// rm-bp1imnmcjxdz7****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The column based on which the entries are sorted. By default, the entries are sorted by the time when the download task was created. Set the value to **gmt_create**.
	//
	// example:
	//
	// gmt_create
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The order in which you want to sort the entries. Valid values:
	//
	// 	- **asc**: the ascending order.
	//
	// 	- **desc**: the descending order. This is the default value.
	//
	// example:
	//
	// desc
	OrderDirect *string `json:"OrderDirect,omitempty" xml:"OrderDirect,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/26231.html) operation to query the region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The beginning of the time range to query. Specify this parameter as a timestamp of the LONG type. Unit: milliseconds.
	//
	// example:
	//
	// 1661941554000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the download task. Valid values:
	//
	// 	- **Initializing**: The download task is being initialized.
	//
	// 	- **queueing**: The download task is queuing.
	//
	// 	- **running**: The download task is running.
	//
	// 	- **failed**: The download task fails.
	//
	// 	- **finished**: The download task is complete.
	//
	// 	- **expired**: The download task expires.
	//
	// example:
	//
	// queueing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The type of the download task. Valid values:
	//
	// 	- **full**: downloads a full backup set.
	//
	// 	- **pitr**: downloads a backup set at a specific point in time.
	//
	// example:
	//
	// full
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeDownloadTaskRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeDownloadTaskRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeDownloadTaskRequest) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *DescribeDownloadTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDownloadTaskRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDownloadTaskRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *DescribeDownloadTaskRequest) GetOrderDirect() *string {
	return s.OrderDirect
}

func (s *DescribeDownloadTaskRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDownloadTaskRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeDownloadTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDownloadTaskRequest) GetState() *string {
	return s.State
}

func (s *DescribeDownloadTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDownloadTaskRequest) SetBackupSetId(v string) *DescribeDownloadTaskRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetClusterName(v string) *DescribeDownloadTaskRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetCurrentPage(v string) *DescribeDownloadTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetDatasourceId(v string) *DescribeDownloadTaskRequest {
	s.DatasourceId = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetEndTime(v string) *DescribeDownloadTaskRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetInstanceName(v string) *DescribeDownloadTaskRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetOrderColumn(v string) *DescribeDownloadTaskRequest {
	s.OrderColumn = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetOrderDirect(v string) *DescribeDownloadTaskRequest {
	s.OrderDirect = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetPageSize(v string) *DescribeDownloadTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetRegionCode(v string) *DescribeDownloadTaskRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetStartTime(v string) *DescribeDownloadTaskRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetState(v string) *DescribeDownloadTaskRequest {
	s.State = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetTaskType(v string) *DescribeDownloadTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeDownloadTaskRequest) Validate() error {
	return dara.Validate(s)
}

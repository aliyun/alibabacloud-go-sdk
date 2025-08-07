// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduleTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterDescription(v string) *DescribeScheduleTasksRequest
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *DescribeScheduleTasksRequest
	GetDBClusterId() *string
	SetOrderId(v string) *DescribeScheduleTasksRequest
	GetOrderId() *string
	SetOwnerAccount(v string) *DescribeScheduleTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeScheduleTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeScheduleTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScheduleTasksRequest
	GetPageSize() *int32
	SetPlannedEndTime(v string) *DescribeScheduleTasksRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *DescribeScheduleTasksRequest
	GetPlannedStartTime() *string
	SetRegionId(v string) *DescribeScheduleTasksRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeScheduleTasksRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeScheduleTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScheduleTasksRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeScheduleTasksRequest
	GetStatus() *string
	SetTaskAction(v string) *DescribeScheduleTasksRequest
	GetTaskAction() *string
}

type DescribeScheduleTasksRequest struct {
	// The description of the cluster.
	//
	// example:
	//
	// testdb
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// >
	//
	// 	- You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the information of all PolarDB clusters that are deployed in a specific region, such as the cluster IDs.
	//
	// 	- If you do not specify this parameter, all scheduled tasks on your clusters are queried.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the order.
	//
	// >  The order ID can contain only digits.
	//
	// example:
	//
	// 20951253014****
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return. Set this parameter to an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The latest start time of the task that you specified when you created the scheduled task. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-28T12:30Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest start time of the task that you specified when you created the scheduled task. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-28T12:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The ID of the region.
	//
	// >
	//
	// 	- You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the region information of all clusters in a specific account.
	//
	// 	- If you do not specify this parameter, scheduled tasks on your clusters that are deployed in all regions are queried.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The state of the tasks that you want to query. Valid values:
	//
	// 	- **pending**: The tasks are pending execution.
	//
	// 	- **executing**: The tasks are being executed.
	//
	// 	- **failure**: The tasks failed and need to be run again.
	//
	// 	- **finish**: The tasks are complete.
	//
	// 	- **cancel**: The tasks are canceled.
	//
	// 	- **expired**: The tasks are expired. The tasks are not started within the time periods that are specified to start the tasks.
	//
	// 	- **rollback**: The tasks are being rolled back.
	//
	// >  If you do not specify this parameter, all scheduled tasks in all states are queried.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of scheduled tasks that you want to query. Valid values:
	//
	// 	- **CreateDBNodes**
	//
	// 	- **ModifyDBNodeClass**
	//
	// 	- **UpgradeDBClusterVersion**
	//
	// 	- **ModifyDBClusterPrimaryZone**
	//
	// >
	//
	// 	- If you specify the `PlannedStartTime` parameter when you call the four preceding operations, the details of each task are returned. Otherwise, an empty string is returned for the `TimerInfos` parameter.
	//
	// 	- If you do not specify this parameter, all types of scheduled tasks on you clusters are queried.
	//
	// example:
	//
	// CreateDBNodes
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeScheduleTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduleTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeScheduleTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeScheduleTasksRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribeScheduleTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeScheduleTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScheduleTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScheduleTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScheduleTasksRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *DescribeScheduleTasksRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *DescribeScheduleTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScheduleTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeScheduleTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScheduleTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScheduleTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeScheduleTasksRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeScheduleTasksRequest) SetDBClusterDescription(v string) *DescribeScheduleTasksRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetDBClusterId(v string) *DescribeScheduleTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetOrderId(v string) *DescribeScheduleTasksRequest {
	s.OrderId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetOwnerAccount(v string) *DescribeScheduleTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetOwnerId(v int64) *DescribeScheduleTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetPageNumber(v int32) *DescribeScheduleTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetPageSize(v int32) *DescribeScheduleTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetPlannedEndTime(v string) *DescribeScheduleTasksRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetPlannedStartTime(v string) *DescribeScheduleTasksRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetRegionId(v string) *DescribeScheduleTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetResourceGroupId(v string) *DescribeScheduleTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetResourceOwnerAccount(v string) *DescribeScheduleTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetResourceOwnerId(v int64) *DescribeScheduleTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetStatus(v string) *DescribeScheduleTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetTaskAction(v string) *DescribeScheduleTasksRequest {
	s.TaskAction = &v
	return s
}

func (s *DescribeScheduleTasksRequest) Validate() error {
	return dara.Validate(s)
}

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
	// The cluster description.
	//
	// example:
	//
	// testdb
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// > - You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters in a specific region, including cluster IDs.
	//
	// >
	//
	// > - If this parameter is omitted, scheduled tasks for all clusters in your account are queried.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	//
	// > The order ID can contain only digits.
	//
	// example:
	//
	// 20951253014****
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30*	- (default), **50**, and **100**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The latest start time of the task. The time is in UTC. If the task does not start by this time, it expires.
	//
	// example:
	//
	// 2021-01-28T12:30Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest start time of the task. The time is in UTC.
	//
	// example:
	//
	// 2021-01-28T12:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The region ID.
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the available regions.
	//
	// >
	//
	// > - If this parameter is omitted, scheduled tasks in all regions in your account are queried.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task status. Valid values:
	//
	// - **pending**: The task is waiting to be executed.
	//
	// - **executing**: The task is being executed.
	//
	// - **failure**: The task failed and is waiting for a retry.
	//
	// - **finish**: The task is complete.
	//
	// - **cancel**: The task is canceled.
	//
	// - **expired**: The task has expired because it did not start within the scheduled time window.
	//
	// - **rollback**: The task is being rolled back.
	//
	// > If this parameter is omitted, scheduled tasks in all states are queried.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The action of the scheduled task. Valid values:
	//
	// - **CreateDBNodes**
	//
	// - **ModifyDBNodeClass**
	//
	// - **UpgradeDBClusterVersion**
	//
	// - **ModifyDBClusterPrimaryZone**
	//
	// > 	- Task details are returned only if you specify the `PlannedStartTime` parameter when you call one of the preceding API operations. Otherwise, the `TimerInfos` field in the response is empty.
	//
	// >
	//
	// > 	- If this parameter is omitted, scheduled tasks of all types are queried.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeTasksRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTasksRequest
	GetRegionId() *string
	SetResourceIds(v []*string) *DescribeTasksRequest
	GetResourceIds() []*string
	SetResourceOwnerAccount(v string) *DescribeTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTasksRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeTasksRequest
	GetStartTime() *string
	SetTaskAction(v string) *DescribeTasksRequest
	GetTaskAction() *string
	SetTaskGroupId(v string) *DescribeTasksRequest
	GetTaskGroupId() *string
	SetTaskIds(v string) *DescribeTasksRequest
	GetTaskIds() *string
	SetTaskStatus(v string) *DescribeTasksRequest
	GetTaskStatus() *string
}

type DescribeTasksRequest struct {
	// The end of the creation time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-11-23T15:16:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the results.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. Valid values of N: 1 to 100.
	ResourceIds          []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the creation time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-11-23T15:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the API operation associated with the task. Valid values:
	//
	// - ImportImage: import an image.
	//
	// - ExportImage: export an image.
	//
	// - RedeployInstance: redeploy an ECS instance.
	//
	// - ModifyDiskSpec: change the cloud disk type.
	//
	// - ArchiveSnapshot: archive a snapshot.
	//
	// example:
	//
	// ImportImage
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The task group ID.
	//
	// > This parameter is in invitational preview. When this parameter is specified, other query conditions do not take effect.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// g-2ze2op2grqpclwu7****
	TaskGroupId *string `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
	// The task IDs. You can specify up to 100 task IDs at a time. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// t-bp1hvgwromzv32iq****,t-bp179lofu2pv768w****
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	// The task status. Valid values:
	//
	// - Finished: The task is complete.
	//
	// - Processing: The task is running.
	//
	// - Failed: The task has failed.
	//
	// Default value: null.
	//
	// > Only tasks in the Finished, Processing, or Failed state can be queried. Other values do not take effect.
	//
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTasksRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTasksRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeTasksRequest) GetTaskGroupId() *string {
	return s.TaskGroupId
}

func (s *DescribeTasksRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *DescribeTasksRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTasksRequest) SetEndTime(v string) *DescribeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerId(v int64) *DescribeTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetRegionId(v string) *DescribeTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceIds(v []*string) *DescribeTasksRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerAccount(v string) *DescribeTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerId(v int64) *DescribeTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetStartTime(v string) *DescribeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskAction(v string) *DescribeTasksRequest {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskGroupId(v string) *DescribeTasksRequest {
	s.TaskGroupId = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskIds(v string) *DescribeTasksRequest {
	s.TaskIds = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskStatus(v string) *DescribeTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksRequest) Validate() error {
	return dara.Validate(s)
}

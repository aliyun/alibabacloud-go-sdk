// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CancelScheduleTasksRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CancelScheduleTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelScheduleTasksRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CancelScheduleTasksRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CancelScheduleTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelScheduleTasksRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *CancelScheduleTasksRequest
	GetTaskId() *string
}

type CancelScheduleTasksRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the information of all clusters that are deployed in a specified region, such as the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scheduled task that you want to cancel.
	//
	// > 	- You can call the [DescribeScheduleTasks](https://help.aliyun.com/document_detail/199648.html) operation to query the details of all scheduled tasks that belong to the current account, such as the task IDs.
	//
	// >	- You can cancel only the tasks whose status is `pending`.``
	//
	// This parameter is required.
	//
	// example:
	//
	// ec8c4723-eac5-4f12-becb-01ac08******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelScheduleTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelScheduleTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CancelScheduleTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelScheduleTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelScheduleTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CancelScheduleTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelScheduleTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelScheduleTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelScheduleTasksRequest) SetDBClusterId(v string) *CancelScheduleTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetOwnerAccount(v string) *CancelScheduleTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetOwnerId(v int64) *CancelScheduleTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetResourceGroupId(v string) *CancelScheduleTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetResourceOwnerAccount(v string) *CancelScheduleTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetResourceOwnerId(v int64) *CancelScheduleTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetTaskId(v string) *CancelScheduleTasksRequest {
	s.TaskId = &v
	return s
}

func (s *CancelScheduleTasksRequest) Validate() error {
	return dara.Validate(s)
}

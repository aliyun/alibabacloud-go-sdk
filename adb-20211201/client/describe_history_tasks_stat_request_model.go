// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromExecTime(v int32) *DescribeHistoryTasksStatRequest
	GetFromExecTime() *int32
	SetFromStartTime(v string) *DescribeHistoryTasksStatRequest
	GetFromStartTime() *string
	SetInstanceId(v string) *DescribeHistoryTasksStatRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DescribeHistoryTasksStatRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeHistoryTasksStatRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeHistoryTasksStatRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v int64) *DescribeHistoryTasksStatRequest
	GetResourceOwnerAccount() *int64
	SetResourceOwnerId(v int64) *DescribeHistoryTasksStatRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeHistoryTasksStatRequest
	GetSecurityToken() *string
	SetStatus(v string) *DescribeHistoryTasksStatRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeHistoryTasksStatRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeHistoryTasksStatRequest
	GetTaskType() *string
	SetToExecTime(v int32) *DescribeHistoryTasksStatRequest
	GetToExecTime() *int32
	SetToStartTime(v string) *DescribeHistoryTasksStatRequest
	GetToStartTime() *string
}

type DescribeHistoryTasksStatRequest struct {
	// example:
	//
	// 1
	FromExecTime *int32 `json:"FromExecTime,omitempty" xml:"FromExecTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-08T07:04:17Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// example:
	//
	// amv-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *int64  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// Scheduled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1564657730
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// all
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 10
	ToExecTime *int32 `json:"ToExecTime,omitempty" xml:"ToExecTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-24T10:01:37Z
	ToStartTime *string `json:"ToStartTime,omitempty" xml:"ToStartTime,omitempty"`
}

func (s DescribeHistoryTasksStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksStatRequest) GetFromExecTime() *int32 {
	return s.FromExecTime
}

func (s *DescribeHistoryTasksStatRequest) GetFromStartTime() *string {
	return s.FromStartTime
}

func (s *DescribeHistoryTasksStatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHistoryTasksStatRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHistoryTasksStatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryTasksStatRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHistoryTasksStatRequest) GetResourceOwnerAccount() *int64 {
	return s.ResourceOwnerAccount
}

func (s *DescribeHistoryTasksStatRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHistoryTasksStatRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeHistoryTasksStatRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeHistoryTasksStatRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHistoryTasksStatRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeHistoryTasksStatRequest) GetToExecTime() *int32 {
	return s.ToExecTime
}

func (s *DescribeHistoryTasksStatRequest) GetToStartTime() *string {
	return s.ToStartTime
}

func (s *DescribeHistoryTasksStatRequest) SetFromExecTime(v int32) *DescribeHistoryTasksStatRequest {
	s.FromExecTime = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetFromStartTime(v string) *DescribeHistoryTasksStatRequest {
	s.FromStartTime = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetInstanceId(v string) *DescribeHistoryTasksStatRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetOwnerId(v int64) *DescribeHistoryTasksStatRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetRegionId(v string) *DescribeHistoryTasksStatRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetResourceGroupId(v string) *DescribeHistoryTasksStatRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetResourceOwnerAccount(v int64) *DescribeHistoryTasksStatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetResourceOwnerId(v int64) *DescribeHistoryTasksStatRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetSecurityToken(v string) *DescribeHistoryTasksStatRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetStatus(v string) *DescribeHistoryTasksStatRequest {
	s.Status = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetTaskId(v string) *DescribeHistoryTasksStatRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetTaskType(v string) *DescribeHistoryTasksStatRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetToExecTime(v int32) *DescribeHistoryTasksStatRequest {
	s.ToExecTime = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetToStartTime(v string) *DescribeHistoryTasksStatRequest {
	s.ToStartTime = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) Validate() error {
	return dara.Validate(s)
}

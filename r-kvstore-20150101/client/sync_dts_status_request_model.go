// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDtsStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SyncDtsStatusRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *SyncDtsStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SyncDtsStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SyncDtsStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SyncDtsStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SyncDtsStatusRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *SyncDtsStatusRequest
	GetSecurityToken() *string
	SetStatus(v string) *SyncDtsStatusRequest
	GetStatus() *string
	SetTaskId(v string) *SyncDtsStatusRequest
	GetTaskId() *string
}

type SyncDtsStatusRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether to restrict specification changes on the instance. Valid values:
	//
	// 	- **0**: Unrestricted.
	//
	// 	- **1**: Restricted. Users are prompted when they perform a specification change.
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Instance ID of the DTS instance. You can log on to the [DTS console](https://dts.console.aliyun.com/) to view instance ID.
	//
	// > An instance may be involved in multiple DTS migration or synchronization tasks simultaneously. To remove the specification change restrictions on the instance, pass this parameter to avoid calling this operation multiple times.
	//
	// example:
	//
	// dtss0611o8vv90****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SyncDtsStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDtsStatusRequest) GoString() string {
	return s.String()
}

func (s *SyncDtsStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SyncDtsStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SyncDtsStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SyncDtsStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SyncDtsStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SyncDtsStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SyncDtsStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SyncDtsStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SyncDtsStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SyncDtsStatusRequest) SetInstanceId(v string) *SyncDtsStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncDtsStatusRequest) SetOwnerAccount(v string) *SyncDtsStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SyncDtsStatusRequest) SetOwnerId(v int64) *SyncDtsStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SyncDtsStatusRequest) SetRegionId(v string) *SyncDtsStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SyncDtsStatusRequest) SetResourceOwnerAccount(v string) *SyncDtsStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SyncDtsStatusRequest) SetResourceOwnerId(v int64) *SyncDtsStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SyncDtsStatusRequest) SetSecurityToken(v string) *SyncDtsStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *SyncDtsStatusRequest) SetStatus(v string) *SyncDtsStatusRequest {
	s.Status = &v
	return s
}

func (s *SyncDtsStatusRequest) SetTaskId(v string) *SyncDtsStatusRequest {
	s.TaskId = &v
	return s
}

func (s *SyncDtsStatusRequest) Validate() error {
	return dara.Validate(s)
}

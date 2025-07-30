// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceHARequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SwitchInstanceHARequest
	GetInstanceId() *string
	SetNodeId(v string) *SwitchInstanceHARequest
	GetNodeId() *string
	SetOwnerAccount(v string) *SwitchInstanceHARequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchInstanceHARequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SwitchInstanceHARequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchInstanceHARequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *SwitchInstanceHARequest
	GetSecurityToken() *string
	SetSwitchMode(v int32) *SwitchInstanceHARequest
	GetSwitchMode() *int32
	SetSwitchType(v string) *SwitchInstanceHARequest
	GetSwitchType() *string
}

type SwitchInstanceHARequest struct {
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the data shard. You can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/473782.html) operation to obtain the value of the CustinsId parameter. Separate multiple data shard IDs with commas (,). `all` indicates that all data shards are specified.
	//
	// > This parameter is available and required only for read/write splitting and cluster instances.
	//
	// example:
	//
	// 56****19,56****20
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The time when to perform the switchover. Default value: 0. Valid values:
	//
	// 	- **0**: immediately performs the switchover.
	//
	// 	- **1**: performs the switchover during the maintenance window.
	//
	// > You can call the [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) operation to modify the maintenance window of a Tair (Redis OSS-compatible) instance.
	//
	// example:
	//
	// 0
	SwitchMode *int32 `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
	// The switching mode. Valid values:
	//
	// 	- **AvailablePriority**: immediately performs a switchover by prioritizing availability. No latency of data synchronization between the master and replica nodes is considered. This may cause data loss.
	//
	// 	- **ReliabilityPriority**: performs a switchover by prioritizing reliability. Make sure that no latency of data synchronization between the master and replica nodes exists. This ensures data integrity. This mode may cause switchover failures in scenarios where a large volume of data is written and data synchronization latency consistently exists.
	//
	// >  You must evaluate the requirements for data and services based on your business scenarios and then select a switching mode.
	//
	// example:
	//
	// AvailablePriority
	SwitchType *string `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s SwitchInstanceHARequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceHARequest) GoString() string {
	return s.String()
}

func (s *SwitchInstanceHARequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchInstanceHARequest) GetNodeId() *string {
	return s.NodeId
}

func (s *SwitchInstanceHARequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchInstanceHARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchInstanceHARequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchInstanceHARequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchInstanceHARequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SwitchInstanceHARequest) GetSwitchMode() *int32 {
	return s.SwitchMode
}

func (s *SwitchInstanceHARequest) GetSwitchType() *string {
	return s.SwitchType
}

func (s *SwitchInstanceHARequest) SetInstanceId(v string) *SwitchInstanceHARequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetNodeId(v string) *SwitchInstanceHARequest {
	s.NodeId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetOwnerAccount(v string) *SwitchInstanceHARequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchInstanceHARequest) SetOwnerId(v int64) *SwitchInstanceHARequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetResourceOwnerAccount(v string) *SwitchInstanceHARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchInstanceHARequest) SetResourceOwnerId(v int64) *SwitchInstanceHARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSecurityToken(v string) *SwitchInstanceHARequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSwitchMode(v int32) *SwitchInstanceHARequest {
	s.SwitchMode = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSwitchType(v string) *SwitchInstanceHARequest {
	s.SwitchType = &v
	return s
}

func (s *SwitchInstanceHARequest) Validate() error {
	return dara.Validate(s)
}

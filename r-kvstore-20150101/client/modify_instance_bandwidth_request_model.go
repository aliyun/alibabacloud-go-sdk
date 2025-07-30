// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceBandwidthRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceBandwidthRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceBandwidthRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceBandwidthRequest
	GetSecurityToken() *string
	SetTargetIntranetBandwidth(v string) *ModifyInstanceBandwidthRequest
	GetTargetIntranetBandwidth() *string
}

type ModifyInstanceBandwidthRequest struct {
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The total expected bandwidth of the instance.
	//
	// > If the instance is a cluster instance, the TargetIntranetBandwidth must be divisible by the number of shards in the instance. This operation is not supported for read/write splitting instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 260
	TargetIntranetBandwidth *string `json:"TargetIntranetBandwidth,omitempty" xml:"TargetIntranetBandwidth,omitempty"`
}

func (s ModifyInstanceBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceBandwidthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceBandwidthRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceBandwidthRequest) GetTargetIntranetBandwidth() *string {
	return s.TargetIntranetBandwidth
}

func (s *ModifyInstanceBandwidthRequest) SetInstanceId(v string) *ModifyInstanceBandwidthRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceBandwidthRequest) SetOwnerAccount(v string) *ModifyInstanceBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceBandwidthRequest) SetOwnerId(v int64) *ModifyInstanceBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceBandwidthRequest) SetResourceOwnerAccount(v string) *ModifyInstanceBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceBandwidthRequest) SetResourceOwnerId(v int64) *ModifyInstanceBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceBandwidthRequest) SetSecurityToken(v string) *ModifyInstanceBandwidthRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceBandwidthRequest) SetTargetIntranetBandwidth(v string) *ModifyInstanceBandwidthRequest {
	s.TargetIntranetBandwidth = &v
	return s
}

func (s *ModifyInstanceBandwidthRequest) Validate() error {
	return dara.Validate(s)
}

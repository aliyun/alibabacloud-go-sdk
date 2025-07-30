// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIntranetAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandWidth(v int64) *ModifyIntranetAttributeRequest
	GetBandWidth() *int64
	SetInstanceId(v string) *ModifyIntranetAttributeRequest
	GetInstanceId() *string
	SetNodeId(v string) *ModifyIntranetAttributeRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *ModifyIntranetAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyIntranetAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyIntranetAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyIntranetAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyIntranetAttributeRequest
	GetSecurityToken() *string
}

type ModifyIntranetAttributeRequest struct {
	// The amount of bandwidth that you want to add. Unit: Mbit/s. The value must be an integer greater than or equal to 0. In most cases, the maximum bandwidth that can be added can be two times the default maximum bandwidth of the current instance type. For more information about the bandwidth specifications supported by different instance types, see [Overview](https://help.aliyun.com/document_detail/26350.html). The bandwidth is also subject to the following limits:
	//
	// 	- The bandwidth of an individual instance cannot exceed 75% of the bandwidth of the host. For more information about the host specifications and bandwidth, see [Instance types of hosts](https://help.aliyun.com/document_detail/206343.html).
	//
	// 	- The total bandwidth of all of the instances on the host cannot exceed 150% of the bandwidth of the host. You can configure resource overcommitment to handle traffic spikes. For more information, see [Configure resource overcommitment to reduce costs](https://help.aliyun.com/document_detail/183798.html).
	//
	// > If you do not specify this parameter for a standard instance, the bandwidth of the instance is set to two times that of the current bandwidth.
	//
	// example:
	//
	// 10
	BandWidth *int64 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the data node. You can call the [DescribeClusterMemberInfo](https://help.aliyun.com/document_detail/473783.html) operation to query the node ID. Separate multiple IDs with commas (,).
	//
	// >  This parameter is required if the instance is a [cluster](https://help.aliyun.com/document_detail/52228.html) instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyIntranetAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIntranetAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyIntranetAttributeRequest) GetBandWidth() *int64 {
	return s.BandWidth
}

func (s *ModifyIntranetAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyIntranetAttributeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyIntranetAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyIntranetAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyIntranetAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyIntranetAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyIntranetAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyIntranetAttributeRequest) SetBandWidth(v int64) *ModifyIntranetAttributeRequest {
	s.BandWidth = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetInstanceId(v string) *ModifyIntranetAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetNodeId(v string) *ModifyIntranetAttributeRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetOwnerAccount(v string) *ModifyIntranetAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetOwnerId(v int64) *ModifyIntranetAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetResourceOwnerAccount(v string) *ModifyIntranetAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetResourceOwnerId(v int64) *ModifyIntranetAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetSecurityToken(v string) *ModifyIntranetAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) Validate() error {
	return dara.Validate(s)
}

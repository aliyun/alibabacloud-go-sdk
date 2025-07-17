// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachVServerGroupsRequest
	GetClientToken() *string
	SetForceAttach(v bool) *AttachVServerGroupsRequest
	GetForceAttach() *bool
	SetOwnerId(v int64) *AttachVServerGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachVServerGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachVServerGroupsRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *AttachVServerGroupsRequest
	GetScalingGroupId() *string
	SetVServerGroups(v []*AttachVServerGroupsRequestVServerGroups) *AttachVServerGroupsRequest
	GetVServerGroups() []*AttachVServerGroupsRequestVServerGroups
}

type AttachVServerGroupsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to add the existing Elastic Compute Service (ECS) instances or elastic container instances in the scaling group to the new vServer group. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceAttach *bool  `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group. Examples: cn-hangzhou and cn-shanghai. For information about regions and zones, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The information about the vServer groups.
	//
	// This parameter is required.
	VServerGroups []*AttachVServerGroupsRequestVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
}

func (s AttachVServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachVServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachVServerGroupsRequest) GetForceAttach() *bool {
	return s.ForceAttach
}

func (s *AttachVServerGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachVServerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachVServerGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachVServerGroupsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *AttachVServerGroupsRequest) GetVServerGroups() []*AttachVServerGroupsRequestVServerGroups {
	return s.VServerGroups
}

func (s *AttachVServerGroupsRequest) SetClientToken(v string) *AttachVServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetForceAttach(v bool) *AttachVServerGroupsRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetOwnerId(v int64) *AttachVServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetRegionId(v string) *AttachVServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetResourceOwnerAccount(v string) *AttachVServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetScalingGroupId(v string) *AttachVServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetVServerGroups(v []*AttachVServerGroupsRequestVServerGroups) *AttachVServerGroupsRequest {
	s.VServerGroups = v
	return s
}

func (s *AttachVServerGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type AttachVServerGroupsRequestVServerGroups struct {
	// The ID of the CLB instance to which the new vServer group belongs.
	//
	// example:
	//
	// rsp-bp1jp1rge****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The attributes of the vServer group.
	VServerGroupAttributes []*AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s AttachVServerGroupsRequestVServerGroups) String() string {
	return dara.Prettify(s)
}

func (s AttachVServerGroupsRequestVServerGroups) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequestVServerGroups) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AttachVServerGroupsRequestVServerGroups) GetVServerGroupAttributes() []*AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	return s.VServerGroupAttributes
}

func (s *AttachVServerGroupsRequestVServerGroups) SetLoadBalancerId(v string) *AttachVServerGroupsRequestVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroups) SetVServerGroupAttributes(v []*AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) *AttachVServerGroupsRequestVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroups) Validate() error {
	return dara.Validate(s)
}

type AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes struct {
	// The port number over which Auto Scaling adds ECS instances or elastic container instances to the new vServer group. Valid values: 1 to 65535.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the vServer group.
	//
	// example:
	//
	// lb-bp1u7etiogg38yvwz****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The weight of an ECS instance or elastic container instance as a backend server. Valid values: 0 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) String() string {
	return dara.Prettify(s)
}

func (s AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GetPort() *int32 {
	return s.Port
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GetWeight() *int32 {
	return s.Weight
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetPort(v int32) *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetWeight(v int32) *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.Weight = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) Validate() error {
	return dara.Validate(s)
}

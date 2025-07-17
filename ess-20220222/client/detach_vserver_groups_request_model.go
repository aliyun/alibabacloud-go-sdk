// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachVServerGroupsRequest
	GetClientToken() *string
	SetForceDetach(v bool) *DetachVServerGroupsRequest
	GetForceDetach() *bool
	SetOwnerId(v int64) *DetachVServerGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachVServerGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachVServerGroupsRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DetachVServerGroupsRequest
	GetScalingGroupId() *string
	SetVServerGroups(v []*DetachVServerGroupsRequestVServerGroups) *DetachVServerGroupsRequest
	GetVServerGroups() []*DetachVServerGroupsRequestVServerGroups
}

type DetachVServerGroupsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to remove the existing instances in the scaling group from the vServer group marked for detachment.
	//
	// 	- true: If you set this parameter to `true`, the detachment of the load balancer from the scaling group causes automatic removal of the existing instances in the scaling group from the corresponding vServer group.
	//
	// 	- false: If you set this parameter to `false`, the detachment of the load balancer from the scaling group does not cause automatic removal of the existing instances in the scaling group from the corresponding vServer group.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceDetach *bool  `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group. Examples: cn-hangzhou and cn-shanghai.
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
	// asg-bp1fo0dbtsbmqa9h****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The collection of information about the vServer groups marked for detachment.
	//
	// This parameter is required.
	VServerGroups []*DetachVServerGroupsRequestVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
}

func (s DetachVServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachVServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachVServerGroupsRequest) GetForceDetach() *bool {
	return s.ForceDetach
}

func (s *DetachVServerGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachVServerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachVServerGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachVServerGroupsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DetachVServerGroupsRequest) GetVServerGroups() []*DetachVServerGroupsRequestVServerGroups {
	return s.VServerGroups
}

func (s *DetachVServerGroupsRequest) SetClientToken(v string) *DetachVServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetForceDetach(v bool) *DetachVServerGroupsRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetOwnerId(v int64) *DetachVServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetRegionId(v string) *DetachVServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetResourceOwnerAccount(v string) *DetachVServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetScalingGroupId(v string) *DetachVServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetVServerGroups(v []*DetachVServerGroupsRequestVServerGroups) *DetachVServerGroupsRequest {
	s.VServerGroups = v
	return s
}

func (s *DetachVServerGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type DetachVServerGroupsRequestVServerGroups struct {
	// The ID of the load balancer to which the vServer group belongs.
	//
	// >  You can detach vServer groups of up to five load balancers from a scaling group in one call.
	//
	// example:
	//
	// lb-bp1p90y3ya9h8s62d****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The attributes of the backend vServer group.
	VServerGroupAttributes []*DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s DetachVServerGroupsRequestVServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DetachVServerGroupsRequestVServerGroups) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequestVServerGroups) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DetachVServerGroupsRequestVServerGroups) GetVServerGroupAttributes() []*DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	return s.VServerGroupAttributes
}

func (s *DetachVServerGroupsRequestVServerGroups) SetLoadBalancerId(v string) *DetachVServerGroupsRequestVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *DetachVServerGroupsRequestVServerGroups) SetVServerGroupAttributes(v []*DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) *DetachVServerGroupsRequestVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

func (s *DetachVServerGroupsRequestVServerGroups) Validate() error {
	return dara.Validate(s)
}

type DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes struct {
	// The port number that Auto Scaling employs to incorporate instances into the vServer group. Valid values: 1 to 65535.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend vServer group.
	//
	// example:
	//
	// rsp-bp1jp1rge****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) String() string {
	return dara.Prettify(s)
}

func (s DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GetPort() *int32 {
	return s.Port
}

func (s *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetPort(v int32) *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

func (s *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) Validate() error {
	return dara.Validate(s)
}

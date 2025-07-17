// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAlbServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlbServerGroups(v []*AttachAlbServerGroupsRequestAlbServerGroups) *AttachAlbServerGroupsRequest
	GetAlbServerGroups() []*AttachAlbServerGroupsRequestAlbServerGroups
	SetClientToken(v string) *AttachAlbServerGroupsRequest
	GetClientToken() *string
	SetForceAttach(v bool) *AttachAlbServerGroupsRequest
	GetForceAttach() *bool
	SetOwnerId(v int64) *AttachAlbServerGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachAlbServerGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachAlbServerGroupsRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *AttachAlbServerGroupsRequest
	GetScalingGroupId() *string
}

type AttachAlbServerGroupsRequest struct {
	// The information about the ALB server groups.
	//
	// This parameter is required.
	AlbServerGroups []*AttachAlbServerGroupsRequestAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to add the existing Elastic Compute Service (ECS) instances or elastic container instances in the scaling group to the new ALB server group. Valid values:
	//
	// 	- true: adds the existing ECS instances or elastic container instances in the scaling group to the new ALB server group and returns the value of `ScalingActivityId`. You can query the value of ScalingActivityId to check whether the existing ECS instances are added to the ALB server group.
	//
	// 	- false: does not add the existing ECS instances or elastic container instances in the scaling group to the new ALB server group.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceAttach *bool  `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// Examples: `cn-hangzhou` and `cn-shanghai`. For more information about regions and zones, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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
}

func (s AttachAlbServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachAlbServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsRequest) GetAlbServerGroups() []*AttachAlbServerGroupsRequestAlbServerGroups {
	return s.AlbServerGroups
}

func (s *AttachAlbServerGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachAlbServerGroupsRequest) GetForceAttach() *bool {
	return s.ForceAttach
}

func (s *AttachAlbServerGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachAlbServerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachAlbServerGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachAlbServerGroupsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *AttachAlbServerGroupsRequest) SetAlbServerGroups(v []*AttachAlbServerGroupsRequestAlbServerGroups) *AttachAlbServerGroupsRequest {
	s.AlbServerGroups = v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetClientToken(v string) *AttachAlbServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetForceAttach(v bool) *AttachAlbServerGroupsRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetOwnerId(v int64) *AttachAlbServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetRegionId(v string) *AttachAlbServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetResourceOwnerAccount(v string) *AttachAlbServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetScalingGroupId(v string) *AttachAlbServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type AttachAlbServerGroupsRequestAlbServerGroups struct {
	// The ID of the ALB server group.
	//
	// You can attach only a limited number of ALB server groups to a scaling group. To view the predefined quota limit or manually request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-ddwb0y0g6y9bjm****
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	// The port used by ECS instances or elastic container instances after being added as backend servers to the ALB server group.
	//
	// Valid values: 1 to 65535.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The weight of an ECS instance or elastic container instance after being added as a backend server to the ALB server group. Valid values: 0 to 100.
	//
	// If you assign a higher weight to an instance, the instance is allocated a larger proportion of access requests. If you assign zero weight to an instance, the instance is allocated no access requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AttachAlbServerGroupsRequestAlbServerGroups) String() string {
	return dara.Prettify(s)
}

func (s AttachAlbServerGroupsRequestAlbServerGroups) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) GetAlbServerGroupId() *string {
	return s.AlbServerGroupId
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) GetWeight() *int32 {
	return s.Weight
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) SetAlbServerGroupId(v string) *AttachAlbServerGroupsRequestAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) SetPort(v int32) *AttachAlbServerGroupsRequestAlbServerGroups {
	s.Port = &v
	return s
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) SetWeight(v int32) *AttachAlbServerGroupsRequestAlbServerGroups {
	s.Weight = &v
	return s
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) Validate() error {
	return dara.Validate(s)
}

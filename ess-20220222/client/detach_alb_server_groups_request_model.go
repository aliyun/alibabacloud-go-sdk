// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAlbServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlbServerGroups(v []*DetachAlbServerGroupsRequestAlbServerGroups) *DetachAlbServerGroupsRequest
	GetAlbServerGroups() []*DetachAlbServerGroupsRequestAlbServerGroups
	SetClientToken(v string) *DetachAlbServerGroupsRequest
	GetClientToken() *string
	SetForceDetach(v bool) *DetachAlbServerGroupsRequest
	GetForceDetach() *bool
	SetOwnerId(v int64) *DetachAlbServerGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachAlbServerGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachAlbServerGroupsRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DetachAlbServerGroupsRequest
	GetScalingGroupId() *string
}

type DetachAlbServerGroupsRequest struct {
	// Details of the ALB server groups.
	//
	// This parameter is required.
	AlbServerGroups []*DetachAlbServerGroupsRequestAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to remove the existing Elastic Compute Service (ECS) instances from the Application Load Balancer (ALB) server group marked for detachment. Valid values:
	//
	// 	- true: removes the existing ECS instances from the ALB server group and returns the value of `ScalingActivityId`. You can query the value of ScalingActivityId to check whether the existing ECS instances are removed from the ALB server group.
	//
	// 	- false: does not remove the existing ECS instances from the ALB server group.
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
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachAlbServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachAlbServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsRequest) GetAlbServerGroups() []*DetachAlbServerGroupsRequestAlbServerGroups {
	return s.AlbServerGroups
}

func (s *DetachAlbServerGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachAlbServerGroupsRequest) GetForceDetach() *bool {
	return s.ForceDetach
}

func (s *DetachAlbServerGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachAlbServerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachAlbServerGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachAlbServerGroupsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DetachAlbServerGroupsRequest) SetAlbServerGroups(v []*DetachAlbServerGroupsRequestAlbServerGroups) *DetachAlbServerGroupsRequest {
	s.AlbServerGroups = v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetClientToken(v string) *DetachAlbServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetForceDetach(v bool) *DetachAlbServerGroupsRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetOwnerId(v int64) *DetachAlbServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetRegionId(v string) *DetachAlbServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetResourceOwnerAccount(v string) *DetachAlbServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetScalingGroupId(v string) *DetachAlbServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type DetachAlbServerGroupsRequestAlbServerGroups struct {
	// The ID of the ALB server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-ddwb0y0g6y9bjm****
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	// The port number used by the ECS instances in the ALB server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DetachAlbServerGroupsRequestAlbServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DetachAlbServerGroupsRequestAlbServerGroups) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsRequestAlbServerGroups) GetAlbServerGroupId() *string {
	return s.AlbServerGroupId
}

func (s *DetachAlbServerGroupsRequestAlbServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *DetachAlbServerGroupsRequestAlbServerGroups) SetAlbServerGroupId(v string) *DetachAlbServerGroupsRequestAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *DetachAlbServerGroupsRequestAlbServerGroups) SetPort(v int32) *DetachAlbServerGroupsRequestAlbServerGroups {
	s.Port = &v
	return s
}

func (s *DetachAlbServerGroupsRequestAlbServerGroups) Validate() error {
	return dara.Validate(s)
}

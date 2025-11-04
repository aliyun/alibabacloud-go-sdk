// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachServerGroupsRequest
	GetClientToken() *string
	SetForceDetach(v bool) *DetachServerGroupsRequest
	GetForceDetach() *bool
	SetOwnerId(v int64) *DetachServerGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachServerGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachServerGroupsRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DetachServerGroupsRequest
	GetScalingGroupId() *string
	SetServerGroups(v []*DetachServerGroupsRequestServerGroups) *DetachServerGroupsRequest
	GetServerGroups() []*DetachServerGroupsRequestServerGroups
}

type DetachServerGroupsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to remove the existing Elastic Compute Service (ECS) instances or elastic container instances in the scaling group from the server group marked for detachment.
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
	ForceDetach *bool  `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
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
	// The information about the server groups.
	//
	// This parameter is required.
	ServerGroups []*DetachServerGroupsRequestServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
}

func (s DetachServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DetachServerGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachServerGroupsRequest) GetForceDetach() *bool {
	return s.ForceDetach
}

func (s *DetachServerGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachServerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachServerGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachServerGroupsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DetachServerGroupsRequest) GetServerGroups() []*DetachServerGroupsRequestServerGroups {
	return s.ServerGroups
}

func (s *DetachServerGroupsRequest) SetClientToken(v string) *DetachServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachServerGroupsRequest) SetForceDetach(v bool) *DetachServerGroupsRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachServerGroupsRequest) SetOwnerId(v int64) *DetachServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachServerGroupsRequest) SetRegionId(v string) *DetachServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DetachServerGroupsRequest) SetResourceOwnerAccount(v string) *DetachServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachServerGroupsRequest) SetScalingGroupId(v string) *DetachServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DetachServerGroupsRequest) SetServerGroups(v []*DetachServerGroupsRequestServerGroups) *DetachServerGroupsRequest {
	s.ServerGroups = v
	return s
}

func (s *DetachServerGroupsRequest) Validate() error {
	if s.ServerGroups != nil {
		for _, item := range s.ServerGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachServerGroupsRequestServerGroups struct {
	// The port used by ECS instances or elastic container instances as backend servers of the server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-1gv2uidn2msy****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The type of the server group. Valid values:
	//
	// 	- ALB
	//
	// 	- NLB
	//
	// This parameter is required.
	//
	// example:
	//
	// ALB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetachServerGroupsRequestServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DetachServerGroupsRequestServerGroups) GoString() string {
	return s.String()
}

func (s *DetachServerGroupsRequestServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *DetachServerGroupsRequestServerGroups) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *DetachServerGroupsRequestServerGroups) GetType() *string {
	return s.Type
}

func (s *DetachServerGroupsRequestServerGroups) SetPort(v int32) *DetachServerGroupsRequestServerGroups {
	s.Port = &v
	return s
}

func (s *DetachServerGroupsRequestServerGroups) SetServerGroupId(v string) *DetachServerGroupsRequestServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *DetachServerGroupsRequestServerGroups) SetType(v string) *DetachServerGroupsRequestServerGroups {
	s.Type = &v
	return s
}

func (s *DetachServerGroupsRequestServerGroups) Validate() error {
	return dara.Validate(s)
}

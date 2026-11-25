// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachServerGroupsRequest
	GetClientToken() *string
	SetForceAttach(v bool) *AttachServerGroupsRequest
	GetForceAttach() *bool
	SetOwnerId(v int64) *AttachServerGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachServerGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachServerGroupsRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *AttachServerGroupsRequest
	GetScalingGroupId() *string
	SetServerGroups(v []*AttachServerGroupsRequestServerGroups) *AttachServerGroupsRequest
	GetServerGroups() []*AttachServerGroupsRequestServerGroups
}

type AttachServerGroupsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can generate the token value from your client and ensure that it is unique across different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to add the existing ECS instances or ECI instances in the scaling group to the newly added server groups. Valid values:
	//
	// - true: adds the instances.
	//
	// - false: does not add the instances.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceAttach *bool  `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the scaling group belongs.
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
	// The information about the load balancing server groups.
	//
	// This parameter is required.
	ServerGroups []*AttachServerGroupsRequestServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
}

func (s AttachServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *AttachServerGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachServerGroupsRequest) GetForceAttach() *bool {
	return s.ForceAttach
}

func (s *AttachServerGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachServerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachServerGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachServerGroupsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *AttachServerGroupsRequest) GetServerGroups() []*AttachServerGroupsRequestServerGroups {
	return s.ServerGroups
}

func (s *AttachServerGroupsRequest) SetClientToken(v string) *AttachServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachServerGroupsRequest) SetForceAttach(v bool) *AttachServerGroupsRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachServerGroupsRequest) SetOwnerId(v int64) *AttachServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachServerGroupsRequest) SetRegionId(v string) *AttachServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachServerGroupsRequest) SetResourceOwnerAccount(v string) *AttachServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachServerGroupsRequest) SetScalingGroupId(v string) *AttachServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *AttachServerGroupsRequest) SetServerGroups(v []*AttachServerGroupsRequestServerGroups) *AttachServerGroupsRequest {
	s.ServerGroups = v
	return s
}

func (s *AttachServerGroupsRequest) Validate() error {
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

type AttachServerGroupsRequestServerGroups struct {
	// The port number used by the ECS instance or ECI instance after Auto Scaling adds it to the server group.
	//
	// Valid values: 1 to 65535.
	//
	// > This parameter is required for ALB and NLB load balancing. You cannot modify the Settings of this parameter for GWLB load balancing, and the default value is 6081.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the load balancing server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-5yc3bd9lfyh*****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The type of the load balancing server group. Valid values:
	//
	// - ALB: Application Load Balancer (ALB).
	//
	// - NLB: Network Load Balancer (NLB).
	//
	// - GWLB: Gateway Load Balancer (GWLB).
	//
	// This parameter is required.
	//
	// example:
	//
	// ALB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the ECS instance or ECI instance as a backend server after Auto Scaling adds it to the server group. Valid values: 0 to 100.
	//
	// A higher weight indicates that Auto Scaling will allocate more access requests to the ECS instance or ECI instance. If the weight is set to 0, the instance does not receive access requests.
	//
	// > This parameter is required for ALB and NLB load balancing. You cannot modify the Settings of this parameter for GWLB load balancing.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AttachServerGroupsRequestServerGroups) String() string {
	return dara.Prettify(s)
}

func (s AttachServerGroupsRequestServerGroups) GoString() string {
	return s.String()
}

func (s *AttachServerGroupsRequestServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *AttachServerGroupsRequestServerGroups) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *AttachServerGroupsRequestServerGroups) GetType() *string {
	return s.Type
}

func (s *AttachServerGroupsRequestServerGroups) GetWeight() *int32 {
	return s.Weight
}

func (s *AttachServerGroupsRequestServerGroups) SetPort(v int32) *AttachServerGroupsRequestServerGroups {
	s.Port = &v
	return s
}

func (s *AttachServerGroupsRequestServerGroups) SetServerGroupId(v string) *AttachServerGroupsRequestServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *AttachServerGroupsRequestServerGroups) SetType(v string) *AttachServerGroupsRequestServerGroups {
	s.Type = &v
	return s
}

func (s *AttachServerGroupsRequestServerGroups) SetWeight(v int32) *AttachServerGroupsRequestServerGroups {
	s.Weight = &v
	return s
}

func (s *AttachServerGroupsRequestServerGroups) Validate() error {
	return dara.Validate(s)
}

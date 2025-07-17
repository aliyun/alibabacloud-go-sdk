// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDBInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachDBInstancesRequest
	GetClientToken() *string
	SetDBInstances(v []*string) *DetachDBInstancesRequest
	GetDBInstances() []*string
	SetForceDetach(v bool) *DetachDBInstancesRequest
	GetForceDetach() *bool
	SetOwnerId(v int64) *DetachDBInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachDBInstancesRequest
	GetRegionId() *string
	SetRemoveSecurityGroup(v bool) *DetachDBInstancesRequest
	GetRemoveSecurityGroup() *bool
	SetResourceOwnerAccount(v string) *DetachDBInstancesRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DetachDBInstancesRequest
	GetScalingGroupId() *string
}

type DetachDBInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the ApsaraDB RDS instances. You can specify up to five ApsaraDB RDS instances.
	//
	// This parameter is required.
	DBInstances []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// Specifies whether to remove the private IP addresses of the existing instances in the scaling group from the IP address whitelist of the ApsaraDB RDS instance. Valid values:
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
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to remove the security group. This parameter takes effect only if you set `AttachMode` to `SecurityGroup`. Valid values:
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
	RemoveSecurityGroup  *bool   `json:"RemoveSecurityGroup,omitempty" xml:"RemoveSecurityGroup,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachDBInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DetachDBInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachDBInstancesRequest) GetDBInstances() []*string {
	return s.DBInstances
}

func (s *DetachDBInstancesRequest) GetForceDetach() *bool {
	return s.ForceDetach
}

func (s *DetachDBInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachDBInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachDBInstancesRequest) GetRemoveSecurityGroup() *bool {
	return s.RemoveSecurityGroup
}

func (s *DetachDBInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachDBInstancesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DetachDBInstancesRequest) SetClientToken(v string) *DetachDBInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachDBInstancesRequest) SetDBInstances(v []*string) *DetachDBInstancesRequest {
	s.DBInstances = v
	return s
}

func (s *DetachDBInstancesRequest) SetForceDetach(v bool) *DetachDBInstancesRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachDBInstancesRequest) SetOwnerId(v int64) *DetachDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachDBInstancesRequest) SetRegionId(v string) *DetachDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DetachDBInstancesRequest) SetRemoveSecurityGroup(v bool) *DetachDBInstancesRequest {
	s.RemoveSecurityGroup = &v
	return s
}

func (s *DetachDBInstancesRequest) SetResourceOwnerAccount(v string) *DetachDBInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachDBInstancesRequest) SetScalingGroupId(v string) *DetachDBInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DetachDBInstancesRequest) Validate() error {
	return dara.Validate(s)
}

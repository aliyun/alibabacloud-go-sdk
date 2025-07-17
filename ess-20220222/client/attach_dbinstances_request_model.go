// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDBInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachMode(v string) *AttachDBInstancesRequest
	GetAttachMode() *string
	SetClientToken(v string) *AttachDBInstancesRequest
	GetClientToken() *string
	SetDBInstances(v []*string) *AttachDBInstancesRequest
	GetDBInstances() []*string
	SetForceAttach(v bool) *AttachDBInstancesRequest
	GetForceAttach() *bool
	SetOwnerId(v int64) *AttachDBInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachDBInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachDBInstancesRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *AttachDBInstancesRequest
	GetScalingGroupId() *string
	SetType(v string) *AttachDBInstancesRequest
	GetType() *string
}

type AttachDBInstancesRequest struct {
	// The mode in which you want to attach the database to the scaling group. Valid values:
	//
	// 	- SecurityIp: adds the private IP addresses of scaled out ECS instances to the IP address whitelist of the database. Take note that you can choose this mode only when the database that you want to attach is an ApsaraDB RDS instance.
	//
	// 	- SecurityGroup: adds the security group of the scaling configuration based on which ECS instances are created in the scaling group to the security group whitelist of the database for registration.
	//
	// Default value: SecurityIp.
	//
	// example:
	//
	// SecurityIp
	AttachMode *string `json:"AttachMode,omitempty" xml:"AttachMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests.
	//
	// The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the ApsaraDB RDS instances that you want to attach to the scaling group.
	//
	// This parameter is required.
	DBInstances []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// Specifies whether to add the private IP addresses of all ECS instances in the scaling group to the IP address whitelist of the ApsaraDB RDS instance that you want to attach to the scaling group. Valid values:
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
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1avr6ensitts3w****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The type of the database that you want to attach to the scaling group. Valid values:
	//
	// 	- RDS
	//
	// 	- Redis
	//
	// 	- MongoDB
	//
	// Default value: RDS.
	//
	// example:
	//
	// RDS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AttachDBInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachDBInstancesRequest) GetAttachMode() *string {
	return s.AttachMode
}

func (s *AttachDBInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachDBInstancesRequest) GetDBInstances() []*string {
	return s.DBInstances
}

func (s *AttachDBInstancesRequest) GetForceAttach() *bool {
	return s.ForceAttach
}

func (s *AttachDBInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachDBInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachDBInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachDBInstancesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *AttachDBInstancesRequest) GetType() *string {
	return s.Type
}

func (s *AttachDBInstancesRequest) SetAttachMode(v string) *AttachDBInstancesRequest {
	s.AttachMode = &v
	return s
}

func (s *AttachDBInstancesRequest) SetClientToken(v string) *AttachDBInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachDBInstancesRequest) SetDBInstances(v []*string) *AttachDBInstancesRequest {
	s.DBInstances = v
	return s
}

func (s *AttachDBInstancesRequest) SetForceAttach(v bool) *AttachDBInstancesRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachDBInstancesRequest) SetOwnerId(v int64) *AttachDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachDBInstancesRequest) SetRegionId(v string) *AttachDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *AttachDBInstancesRequest) SetResourceOwnerAccount(v string) *AttachDBInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachDBInstancesRequest) SetScalingGroupId(v string) *AttachDBInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *AttachDBInstancesRequest) SetType(v string) *AttachDBInstancesRequest {
	s.Type = &v
	return s
}

func (s *AttachDBInstancesRequest) Validate() error {
	return dara.Validate(s)
}

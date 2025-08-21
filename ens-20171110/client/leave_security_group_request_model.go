// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *LeaveSecurityGroupRequest
	GetInstanceId() *string
	SetNetworkInterfaceId(v string) *LeaveSecurityGroupRequest
	GetNetworkInterfaceId() *string
	SetSecurityGroupId(v string) *LeaveSecurityGroupRequest
	GetSecurityGroupId() *string
}

type LeaveSecurityGroupRequest struct {
	// The instance ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-58z57orgmt6d1****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s LeaveSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s LeaveSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LeaveSecurityGroupRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *LeaveSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *LeaveSecurityGroupRequest) SetInstanceId(v string) *LeaveSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetNetworkInterfaceId(v string) *LeaveSecurityGroupRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetSecurityGroupId(v string) *LeaveSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}

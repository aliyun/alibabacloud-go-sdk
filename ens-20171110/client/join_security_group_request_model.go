// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *JoinSecurityGroupRequest
	GetInstanceId() *string
	SetNetworkInterfaceId(v string) *JoinSecurityGroupRequest
	GetNetworkInterfaceId() *string
	SetSecurityGroupId(v string) *JoinSecurityGroupRequest
	GetSecurityGroupId() *string
}

type JoinSecurityGroupRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp67acfmxazb4ph***
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
	// sg-bp67acfmxazb4ph***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s JoinSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *JoinSecurityGroupRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *JoinSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *JoinSecurityGroupRequest) SetInstanceId(v string) *JoinSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetNetworkInterfaceId(v string) *JoinSecurityGroupRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetSecurityGroupId(v string) *JoinSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *JoinSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}

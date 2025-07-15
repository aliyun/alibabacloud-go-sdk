// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVpcAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveVpcAccessRequest
	GetInstanceId() *string
	SetNeedBatchWork(v bool) *RemoveVpcAccessRequest
	GetNeedBatchWork() *bool
	SetPort(v int32) *RemoveVpcAccessRequest
	GetPort() *int32
	SetSecurityToken(v string) *RemoveVpcAccessRequest
	GetSecurityToken() *string
	SetVpcId(v string) *RemoveVpcAccessRequest
	GetVpcId() *string
}

type RemoveVpcAccessRequest struct {
	// The ID of an ECS or SLB instance in the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf6bzcg1pr4oh5jjmxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether batch work is required.
	//
	// example:
	//
	// true
	NeedBatchWork *bool `json:"NeedBatchWork,omitempty" xml:"NeedBatchWork,omitempty"`
	// The port number that corresponds to the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf657qec7lx42paw3qxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s RemoveVpcAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveVpcAccessRequest) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveVpcAccessRequest) GetNeedBatchWork() *bool {
	return s.NeedBatchWork
}

func (s *RemoveVpcAccessRequest) GetPort() *int32 {
	return s.Port
}

func (s *RemoveVpcAccessRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveVpcAccessRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *RemoveVpcAccessRequest) SetInstanceId(v string) *RemoveVpcAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveVpcAccessRequest) SetNeedBatchWork(v bool) *RemoveVpcAccessRequest {
	s.NeedBatchWork = &v
	return s
}

func (s *RemoveVpcAccessRequest) SetPort(v int32) *RemoveVpcAccessRequest {
	s.Port = &v
	return s
}

func (s *RemoveVpcAccessRequest) SetSecurityToken(v string) *RemoveVpcAccessRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveVpcAccessRequest) SetVpcId(v string) *RemoveVpcAccessRequest {
	s.VpcId = &v
	return s
}

func (s *RemoveVpcAccessRequest) Validate() error {
	return dara.Validate(s)
}

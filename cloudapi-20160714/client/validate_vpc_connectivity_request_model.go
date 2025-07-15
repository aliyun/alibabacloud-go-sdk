// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateVpcConnectivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ValidateVpcConnectivityRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *ValidateVpcConnectivityRequest
	GetSecurityToken() *string
	SetVpcAccessId(v string) *ValidateVpcConnectivityRequest
	GetVpcAccessId() *string
}

type ValidateVpcConnectivityRequest struct {
	// The ID of the API Gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-hz-72bc18******
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the VPC access authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5f1b3216f9********e2c1297b6741dc
	VpcAccessId *string `json:"VpcAccessId,omitempty" xml:"VpcAccessId,omitempty"`
}

func (s ValidateVpcConnectivityRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateVpcConnectivityRequest) GoString() string {
	return s.String()
}

func (s *ValidateVpcConnectivityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ValidateVpcConnectivityRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ValidateVpcConnectivityRequest) GetVpcAccessId() *string {
	return s.VpcAccessId
}

func (s *ValidateVpcConnectivityRequest) SetInstanceId(v string) *ValidateVpcConnectivityRequest {
	s.InstanceId = &v
	return s
}

func (s *ValidateVpcConnectivityRequest) SetSecurityToken(v string) *ValidateVpcConnectivityRequest {
	s.SecurityToken = &v
	return s
}

func (s *ValidateVpcConnectivityRequest) SetVpcAccessId(v string) *ValidateVpcConnectivityRequest {
	s.VpcAccessId = &v
	return s
}

func (s *ValidateVpcConnectivityRequest) Validate() error {
	return dara.Validate(s)
}

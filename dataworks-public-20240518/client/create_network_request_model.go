// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateNetworkRequest
	GetClientToken() *string
	SetResourceGroupId(v string) *CreateNetworkRequest
	GetResourceGroupId() *string
	SetVpcId(v string) *CreateNetworkRequest
	GetVpcId() *string
	SetVswitchId(v string) *CreateNetworkRequest
	GetVswitchId() *string
}

type CreateNetworkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// eb870033-74c8-4b1b-9664-04c4e7cc3465
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the serverless resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNetworkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateNetworkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNetworkRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateNetworkRequest) SetClientToken(v string) *CreateNetworkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkRequest) SetResourceGroupId(v string) *CreateNetworkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateNetworkRequest) SetVpcId(v string) *CreateNetworkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkRequest) SetVswitchId(v string) *CreateNetworkRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateNetworkRequest) Validate() error {
	return dara.Validate(s)
}

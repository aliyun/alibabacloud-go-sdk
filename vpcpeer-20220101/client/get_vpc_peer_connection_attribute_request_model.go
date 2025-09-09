// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPeerConnectionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetVpcPeerConnectionAttributeRequest
	GetInstanceId() *string
	SetResourceOwnerAccount(v string) *GetVpcPeerConnectionAttributeRequest
	GetResourceOwnerAccount() *string
}

type GetVpcPeerConnectionAttributeRequest struct {
	// The ID of the VPC peering connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcc-lnk0m24khwvtkm****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s GetVpcPeerConnectionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVpcPeerConnectionAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVpcPeerConnectionAttributeRequest) SetInstanceId(v string) *GetVpcPeerConnectionAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeRequest) SetResourceOwnerAccount(v string) *GetVpcPeerConnectionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeRequest) Validate() error {
	return dara.Validate(s)
}

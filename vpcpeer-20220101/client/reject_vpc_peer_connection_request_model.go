// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectVpcPeerConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RejectVpcPeerConnectionRequest
	GetClientToken() *string
	SetDryRun(v bool) *RejectVpcPeerConnectionRequest
	GetDryRun() *bool
	SetInstanceId(v string) *RejectVpcPeerConnectionRequest
	GetInstanceId() *string
	SetResourceOwnerAccount(v string) *RejectVpcPeerConnectionRequest
	GetResourceOwnerAccount() *string
}

type RejectVpcPeerConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client. Make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs a dry run and does not reject the connection request. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the actual request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the VPC peering connection instance denied by the acceptor.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcc-lnk0m24khwvtkm0****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s RejectVpcPeerConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *RejectVpcPeerConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RejectVpcPeerConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RejectVpcPeerConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RejectVpcPeerConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RejectVpcPeerConnectionRequest) SetClientToken(v string) *RejectVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *RejectVpcPeerConnectionRequest) SetDryRun(v bool) *RejectVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *RejectVpcPeerConnectionRequest) SetInstanceId(v string) *RejectVpcPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectVpcPeerConnectionRequest) SetResourceOwnerAccount(v string) *RejectVpcPeerConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RejectVpcPeerConnectionRequest) Validate() error {
	return dara.Validate(s)
}

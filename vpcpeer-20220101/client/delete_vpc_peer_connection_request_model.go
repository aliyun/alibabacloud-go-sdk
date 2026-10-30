// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcPeerConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpcPeerConnectionRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteVpcPeerConnectionRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteVpcPeerConnectionRequest
	GetForce() *bool
	SetInstanceId(v string) *DeleteVpcPeerConnectionRequest
	GetInstanceId() *string
}

type DeleteVpcPeerConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs only a dry run and does not delete the VPC peering connection. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to force delete the VPC peering connection. Valid values:
	//
	// - **false*	- (default): Does not force delete the VPC peering connection. Before you delete the connection, you must delete the route entry that points to the VPC peering connection from the routing table.
	//
	// - **true**: The system will force delete the VPC peering connection. The system deletes the route entry that points to the VPC peering connection from the VPC routing table.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The instance ID of the VPC peering connection that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcc-lnk0m24khwvtkm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteVpcPeerConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcPeerConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpcPeerConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVpcPeerConnectionRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteVpcPeerConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVpcPeerConnectionRequest) SetClientToken(v string) *DeleteVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcPeerConnectionRequest) SetDryRun(v bool) *DeleteVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVpcPeerConnectionRequest) SetForce(v bool) *DeleteVpcPeerConnectionRequest {
	s.Force = &v
	return s
}

func (s *DeleteVpcPeerConnectionRequest) SetInstanceId(v string) *DeleteVpcPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVpcPeerConnectionRequest) Validate() error {
	return dara.Validate(s)
}

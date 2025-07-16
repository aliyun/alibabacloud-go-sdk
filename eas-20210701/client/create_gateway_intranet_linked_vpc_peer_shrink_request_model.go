// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIntranetLinkedVpcPeerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPeerVpcsShrink(v string) *CreateGatewayIntranetLinkedVpcPeerShrinkRequest
	GetPeerVpcsShrink() *string
	SetVpcId(v string) *CreateGatewayIntranetLinkedVpcPeerShrinkRequest
	GetVpcId() *string
}

type CreateGatewayIntranetLinkedVpcPeerShrinkRequest struct {
	// The list of VPC peers.
	PeerVpcsShrink *string `json:"PeerVpcs,omitempty" xml:"PeerVpcs,omitempty"`
	// The VPC ID. To obtain the VPC ID, see [ListGatewayIntranetLinkedVpc](https://help.aliyun.com/document_detail/2621223.html).
	//
	// example:
	//
	// vpc-2zetuli9ws0qgjd******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcPeerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcPeerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcPeerShrinkRequest) GetPeerVpcsShrink() *string {
	return s.PeerVpcsShrink
}

func (s *CreateGatewayIntranetLinkedVpcPeerShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGatewayIntranetLinkedVpcPeerShrinkRequest) SetPeerVpcsShrink(v string) *CreateGatewayIntranetLinkedVpcPeerShrinkRequest {
	s.PeerVpcsShrink = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerShrinkRequest) SetVpcId(v string) *CreateGatewayIntranetLinkedVpcPeerShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerShrinkRequest) Validate() error {
	return dara.Validate(s)
}

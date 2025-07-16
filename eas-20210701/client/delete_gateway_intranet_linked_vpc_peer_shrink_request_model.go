// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIntranetLinkedVpcPeerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPeerVpcsShrink(v string) *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest
	GetPeerVpcsShrink() *string
	SetVpcId(v string) *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest
	GetVpcId() *string
}

type DeleteGatewayIntranetLinkedVpcPeerShrinkRequest struct {
	// The VPC peer.
	PeerVpcsShrink *string `json:"PeerVpcs,omitempty" xml:"PeerVpcs,omitempty"`
	// The ID of the associated VPC. To obtain the VPC ID, see [ListGatewayIntranetLinkedVpc](https://help.aliyun.com/document_detail/2621223.html).
	//
	// example:
	//
	// vpc-2zetuli9ws0qgjd******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcPeerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcPeerShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest) GetPeerVpcsShrink() *string {
	return s.PeerVpcsShrink
}

func (s *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest) SetPeerVpcsShrink(v string) *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest {
	s.PeerVpcsShrink = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest) SetVpcId(v string) *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerShrinkRequest) Validate() error {
	return dara.Validate(s)
}

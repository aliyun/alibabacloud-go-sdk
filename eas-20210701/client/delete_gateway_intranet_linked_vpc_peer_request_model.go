// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIntranetLinkedVpcPeerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPeerVpcs(v []*DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) *DeleteGatewayIntranetLinkedVpcPeerRequest
	GetPeerVpcs() []*DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs
	SetVpcId(v string) *DeleteGatewayIntranetLinkedVpcPeerRequest
	GetVpcId() *string
}

type DeleteGatewayIntranetLinkedVpcPeerRequest struct {
	// The VPC peer.
	PeerVpcs []*DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs `json:"PeerVpcs,omitempty" xml:"PeerVpcs,omitempty" type:"Repeated"`
	// The ID of the associated VPC. To obtain the VPC ID, see [ListGatewayIntranetLinkedVpc](https://help.aliyun.com/document_detail/2621223.html).
	//
	// example:
	//
	// vpc-2zetuli9ws0qgjd******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcPeerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcPeerRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequest) GetPeerVpcs() []*DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs {
	return s.PeerVpcs
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequest) SetPeerVpcs(v []*DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) *DeleteGatewayIntranetLinkedVpcPeerRequest {
	s.PeerVpcs = v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequest) SetVpcId(v string) *DeleteGatewayIntranetLinkedVpcPeerRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequest) Validate() error {
	if s.PeerVpcs != nil {
		for _, item := range s.PeerVpcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs struct {
	// The region where the VPC peer resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the VPC peer.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) GetRegion() *string {
	return s.Region
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) SetRegion(v string) *DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs {
	s.Region = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) SetVpcId(v string) *DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs {
	s.VpcId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerRequestPeerVpcs) Validate() error {
	return dara.Validate(s)
}

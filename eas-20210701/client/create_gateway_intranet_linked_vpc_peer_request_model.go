// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIntranetLinkedVpcPeerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPeerVpcs(v []*CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) *CreateGatewayIntranetLinkedVpcPeerRequest
	GetPeerVpcs() []*CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs
	SetVpcId(v string) *CreateGatewayIntranetLinkedVpcPeerRequest
	GetVpcId() *string
}

type CreateGatewayIntranetLinkedVpcPeerRequest struct {
	// The list of VPC peers.
	PeerVpcs []*CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs `json:"PeerVpcs,omitempty" xml:"PeerVpcs,omitempty" type:"Repeated"`
	// The VPC ID. To obtain the VPC ID, see [ListGatewayIntranetLinkedVpc](https://help.aliyun.com/document_detail/2621223.html).
	//
	// example:
	//
	// vpc-2zetuli9ws0qgjd******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcPeerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcPeerRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequest) GetPeerVpcs() []*CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs {
	return s.PeerVpcs
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequest) SetPeerVpcs(v []*CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) *CreateGatewayIntranetLinkedVpcPeerRequest {
	s.PeerVpcs = v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequest) SetVpcId(v string) *CreateGatewayIntranetLinkedVpcPeerRequest {
	s.VpcId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequest) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs struct {
	// The region where the VPC peer resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the VPC peer. To obtain the VPC ID, see [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html).
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) GetRegion() *string {
	return s.Region
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) SetRegion(v string) *CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs {
	s.Region = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) SetVpcId(v string) *CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs {
	s.VpcId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerRequestPeerVpcs) Validate() error {
	return dara.Validate(s)
}

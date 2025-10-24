// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetLinkedVpcPeerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListGatewayIntranetLinkedVpcPeerResponseBody
	GetGatewayId() *string
	SetPeerVpcList(v []*ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) *ListGatewayIntranetLinkedVpcPeerResponseBody
	GetPeerVpcList() []*ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList
	SetRequestId(v string) *ListGatewayIntranetLinkedVpcPeerResponseBody
	GetRequestId() *string
}

type ListGatewayIntranetLinkedVpcPeerResponseBody struct {
	// The ID of the private gateway.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The VPC peers.
	PeerVpcList []*ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList `json:"PeerVpcList,omitempty" xml:"PeerVpcList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcPeerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcPeerResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBody) GetPeerVpcList() []*ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList {
	return s.PeerVpcList
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBody) SetGatewayId(v string) *ListGatewayIntranetLinkedVpcPeerResponseBody {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBody) SetPeerVpcList(v []*ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) *ListGatewayIntranetLinkedVpcPeerResponseBody {
	s.PeerVpcList = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBody) SetRequestId(v string) *ListGatewayIntranetLinkedVpcPeerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBody) Validate() error {
	if s.PeerVpcList != nil {
		for _, item := range s.PeerVpcList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList struct {
	// The IDs of the VPC peers.
	PeerVpcs []*ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs `json:"PeerVpcs,omitempty" xml:"PeerVpcs,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zetuli9ws0qgjd******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) GetPeerVpcs() []*ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs {
	return s.PeerVpcs
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) SetPeerVpcs(v []*ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList {
	s.PeerVpcs = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) SetVpcId(v string) *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList {
	s.VpcId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcList) Validate() error {
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

type ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs struct {
	// The region where the VPC peer resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the VPC peer.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) GetRegion() *string {
	return s.Region
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) SetRegion(v string) *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs {
	s.Region = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) SetStatus(v string) *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs {
	s.Status = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) SetVpcId(v string) *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs {
	s.VpcId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponseBodyPeerVpcListPeerVpcs) Validate() error {
	return dara.Validate(s)
}

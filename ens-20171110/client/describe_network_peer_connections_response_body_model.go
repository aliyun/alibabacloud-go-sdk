// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkPeerConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPeerConnects(v []*DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) *DescribeNetworkPeerConnectionsResponseBody
	GetNetworkPeerConnects() []*DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects
	SetPageNumber(v string) *DescribeNetworkPeerConnectionsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeNetworkPeerConnectionsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeNetworkPeerConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeNetworkPeerConnectionsResponseBody
	GetTotalCount() *string
}

type DescribeNetworkPeerConnectionsResponseBody struct {
	NetworkPeerConnects []*DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects `json:"NetworkPeerConnects,omitempty" xml:"NetworkPeerConnects,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkPeerConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPeerConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPeerConnectionsResponseBody) GetNetworkPeerConnects() []*DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	return s.NetworkPeerConnects
}

func (s *DescribeNetworkPeerConnectionsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeNetworkPeerConnectionsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeNetworkPeerConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkPeerConnectionsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeNetworkPeerConnectionsResponseBody) SetNetworkPeerConnects(v []*DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) *DescribeNetworkPeerConnectionsResponseBody {
	s.NetworkPeerConnects = v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBody) SetPageNumber(v string) *DescribeNetworkPeerConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBody) SetPageSize(v string) *DescribeNetworkPeerConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBody) SetRequestId(v string) *DescribeNetworkPeerConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBody) SetTotalCount(v string) *DescribeNetworkPeerConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBody) Validate() error {
	if s.NetworkPeerConnects != nil {
		for _, item := range s.NetworkPeerConnects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects struct {
	AcceptingNetwork *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork `json:"AcceptingNetwork,omitempty" xml:"AcceptingNetwork,omitempty" type:"Struct"`
	// example:
	//
	// n-5q28114****
	AcceptingNetworkId *string `json:"AcceptingNetworkId,omitempty" xml:"AcceptingNetworkId,omitempty"`
	// example:
	//
	// 2024-05-31T06:43:29Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// cn-dongguan-9
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// pcc-521pzxdg****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	Name    *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Network *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// example:
	//
	// n-5q73bm****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetAcceptingNetwork() *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork {
	return s.AcceptingNetwork
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetAcceptingNetworkId() *string {
	return s.AcceptingNetworkId
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetName() *string {
	return s.Name
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetNetwork() *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork {
	return s.Network
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetAcceptingNetwork(v *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.AcceptingNetwork = v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetAcceptingNetworkId(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.AcceptingNetworkId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetCreationTime(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetDescription(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.Description = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetEnsRegionId(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetInstanceId(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetName(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.Name = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetNetwork(v *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.Network = v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetNetworkId(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) SetStatus(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects {
	s.Status = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnects) Validate() error {
	if s.AcceptingNetwork != nil {
		if err := s.AcceptingNetwork.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork struct {
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// example:
	//
	// n-5q28114****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork) SetIpv4Cidrs(v []*string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork {
	s.Ipv4Cidrs = v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork) SetNetworkId(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsAcceptingNetwork) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork struct {
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// example:
	//
	// n-5q73bm****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork) SetIpv4Cidrs(v []*string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork {
	s.Ipv4Cidrs = v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork) SetNetworkId(v string) *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponseBodyNetworkPeerConnectsNetwork) Validate() error {
	return dara.Validate(s)
}

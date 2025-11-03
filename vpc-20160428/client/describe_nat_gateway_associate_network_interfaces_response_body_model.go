// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewayAssociateNetworkInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociateNetworkInterfaces(v *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody
	GetAssociateNetworkInterfaces() *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces
	SetCount(v int32) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody
	GetCount() *int32
	SetMaxResults(v int32) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody
	GetMaxResults() *int32
	SetNatGatewayId(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody
	GetNatGatewayId() *string
	SetNextToken(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody
	GetTotalCount() *int32
}

type DescribeNatGatewayAssociateNetworkInterfacesResponseBody struct {
	// The ENIs associated with the VPC NAT gateway.
	AssociateNetworkInterfaces *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces `json:"AssociateNetworkInterfaces,omitempty" xml:"AssociateNetworkInterfaces,omitempty" type:"Struct"`
	// Number of associated ENIs.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The number of entries to return per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the VPC NAT gateway.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// Indicates whether the token for the next query exists. Valid value:
	//
	// 	- If **NextToken*	- is empty, there is no next page.
	//
	// 	- If the value returned of **NextToken*	- is not empty, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) GetAssociateNetworkInterfaces() *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces {
	return s.AssociateNetworkInterfaces
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) SetAssociateNetworkInterfaces(v *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody {
	s.AssociateNetworkInterfaces = v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) SetCount(v int32) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) SetMaxResults(v int32) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) SetNatGatewayId(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) SetNextToken(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) SetRequestId(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) SetTotalCount(v int32) *DescribeNatGatewayAssociateNetworkInterfacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) Validate() error {
	if s.AssociateNetworkInterfaces != nil {
		if err := s.AssociateNetworkInterfaces.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces struct {
	AssociateNetworkInterface []*DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface `json:"AssociateNetworkInterface,omitempty" xml:"AssociateNetworkInterface,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces) GetAssociateNetworkInterface() []*DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface {
	return s.AssociateNetworkInterface
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces) SetAssociateNetworkInterface(v []*DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces {
	s.AssociateNetworkInterface = v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces) Validate() error {
	if s.AssociateNetworkInterface != nil {
		for _, item := range s.AssociateNetworkInterface {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface struct {
	// The IPv4 addresses of the ENIs.
	IPv4Sets *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets `json:"IPv4Sets,omitempty" xml:"IPv4Sets,omitempty" type:"Struct"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-gw8g131ef2dnbu3k****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the service resource.
	//
	// example:
	//
	// ep-8psre8c8936596cd****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The UID of the account to which the service resource belongs.
	//
	// example:
	//
	// 138859086900****
	ResourceOwnerId *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the service resource.
	//
	// example:
	//
	// PrivateLink
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceVpcId *string `json:"ResourceVpcId,omitempty" xml:"ResourceVpcId,omitempty"`
	// The ID of the tunnel index.
	//
	// example:
	//
	// 41a5489ea2a0****
	TunnelIndex *string `json:"TunnelIndex,omitempty" xml:"TunnelIndex,omitempty"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) GetIPv4Sets() *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets {
	return s.IPv4Sets
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) GetResourceVpcId() *string {
	return s.ResourceVpcId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) GetTunnelIndex() *string {
	return s.TunnelIndex
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) SetIPv4Sets(v *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface {
	s.IPv4Sets = v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) SetNetworkInterfaceId(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) SetResourceId(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface {
	s.ResourceId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) SetResourceOwnerId(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) SetResourceType(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface {
	s.ResourceType = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) SetResourceVpcId(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface {
	s.ResourceVpcId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) SetTunnelIndex(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface {
	s.TunnelIndex = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface) Validate() error {
	if s.IPv4Sets != nil {
		if err := s.IPv4Sets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets struct {
	IPv4Set []*DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set `json:"IPv4Set,omitempty" xml:"IPv4Set,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets) GetIPv4Set() []*DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set {
	return s.IPv4Set
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets) SetIPv4Set(v []*DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets {
	s.IPv4Set = v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets) Validate() error {
	if s.IPv4Set != nil {
		for _, item := range s.IPv4Set {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set struct {
	// The primary private IP address of the ENI.
	//
	// example:
	//
	// ``172.17.**.**``
	IPv4Address *string `json:"IPv4Address,omitempty" xml:"IPv4Address,omitempty"`
	// Indicates whether the IP address is the primary private IP address. Valid values:
	//
	// 	- true: Primary private IP address
	//
	// 	- false: Secondary private IP addresses
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set) GetIPv4Address() *string {
	return s.IPv4Address
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set) SetIPv4Address(v string) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set {
	s.IPv4Address = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set) SetPrimary(v bool) *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set {
	s.Primary = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set) Validate() error {
	return dara.Validate(s)
}

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
	AssociateNetworkInterfaces *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfaces `json:"AssociateNetworkInterfaces,omitempty" xml:"AssociateNetworkInterfaces,omitempty" type:"Struct"`
	Count                      *int32                                                                              `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterface struct {
	IPv4Sets *DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4Sets `json:"IPv4Sets,omitempty" xml:"IPv4Sets,omitempty" type:"Struct"`
	// example:
	//
	// eni-gw8g131ef2dnbu3k****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// example:
	//
	// ep-8psre8c8936596cd****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 138859086900****
	ResourceOwnerId *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// PrivateLink
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceVpcId *string `json:"ResourceVpcId,omitempty" xml:"ResourceVpcId,omitempty"`
	TunnelIndex   *string `json:"TunnelIndex,omitempty" xml:"TunnelIndex,omitempty"`
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeNatGatewayAssociateNetworkInterfacesResponseBodyAssociateNetworkInterfacesAssociateNetworkInterfaceIPv4SetsIPv4Set struct {
	// example:
	//
	// ``172.17.**.**``
	IPv4Address *string `json:"IPv4Address,omitempty" xml:"IPv4Address,omitempty"`
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

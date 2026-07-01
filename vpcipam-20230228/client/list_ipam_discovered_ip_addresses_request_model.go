// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamDiscoveredIpAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *ListIpamDiscoveredIpAddressesRequest
	GetCidr() *string
	SetIpVersion(v string) *ListIpamDiscoveredIpAddressesRequest
	GetIpVersion() *string
	SetIpamResourceDiscoveryId(v string) *ListIpamDiscoveredIpAddressesRequest
	GetIpamResourceDiscoveryId() *string
	SetMaxResults(v int32) *ListIpamDiscoveredIpAddressesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamDiscoveredIpAddressesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListIpamDiscoveredIpAddressesRequest
	GetRegionId() *string
	SetVSwitchId(v string) *ListIpamDiscoveredIpAddressesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ListIpamDiscoveredIpAddressesRequest
	GetVpcId() *string
}

type ListIpamDiscoveredIpAddressesRequest struct {
	// The CIDR block to search for used IP addresses in a VPC or vSwitch. To query a specific IP address, use a /32 prefix length.
	//
	// > Only IPv4 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The IP protocol version. Valid value:
	//
	// - **IPv4**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ID of the resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The maximum number of entries to return per page. Valid values: 1 to 200. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. Valid values:
	//
	// - Do not specify this parameter for your first request.
	//
	// - If a next page exists, set this parameter to the value of **NextToken*	- returned in the previous response.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource discovery instance is hosted.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of a discovered VSwitch.
	//
	// > You must specify at least one of VpcId and VSwitchId.
	//
	// example:
	//
	// vsw-bp1bmwg5u07e1l3q0is4w
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of a discovered VPC.
	//
	// > You must specify at least one of VpcId and VSwitchId.
	//
	// example:
	//
	// vpc-bp1fjfnrg3av6zb9e****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpamDiscoveredIpAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredIpAddressesRequest) GetCidr() *string {
	return s.Cidr
}

func (s *ListIpamDiscoveredIpAddressesRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListIpamDiscoveredIpAddressesRequest) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *ListIpamDiscoveredIpAddressesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamDiscoveredIpAddressesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamDiscoveredIpAddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamDiscoveredIpAddressesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListIpamDiscoveredIpAddressesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListIpamDiscoveredIpAddressesRequest) SetCidr(v string) *ListIpamDiscoveredIpAddressesRequest {
	s.Cidr = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesRequest) SetIpVersion(v string) *ListIpamDiscoveredIpAddressesRequest {
	s.IpVersion = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesRequest) SetIpamResourceDiscoveryId(v string) *ListIpamDiscoveredIpAddressesRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesRequest) SetMaxResults(v int32) *ListIpamDiscoveredIpAddressesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesRequest) SetNextToken(v string) *ListIpamDiscoveredIpAddressesRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesRequest) SetRegionId(v string) *ListIpamDiscoveredIpAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesRequest) SetVSwitchId(v string) *ListIpamDiscoveredIpAddressesRequest {
	s.VSwitchId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesRequest) SetVpcId(v string) *ListIpamDiscoveredIpAddressesRequest {
	s.VpcId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrantVSwitchEnisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListGrantVSwitchEnisRequest
	GetCenId() *string
	SetMaxResults(v int64) *ListGrantVSwitchEnisRequest
	GetMaxResults() *int64
	SetNetworkInterfaceId(v []*string) *ListGrantVSwitchEnisRequest
	GetNetworkInterfaceId() []*string
	SetNetworkInterfaceName(v string) *ListGrantVSwitchEnisRequest
	GetNetworkInterfaceName() *string
	SetNextToken(v string) *ListGrantVSwitchEnisRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListGrantVSwitchEnisRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListGrantVSwitchEnisRequest
	GetOwnerId() *int64
	SetPrimaryIpAddress(v string) *ListGrantVSwitchEnisRequest
	GetPrimaryIpAddress() *string
	SetResourceOwnerAccount(v string) *ListGrantVSwitchEnisRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListGrantVSwitchEnisRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *ListGrantVSwitchEnisRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ListGrantVSwitchEnisRequest
	GetVpcId() *string
}

type ListGrantVSwitchEnisRequest struct {
	// The ID of the CEN instance to which the VPC-connected instance is connected.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-a7syd349kne38g****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The maximum number of entries per page for a paged query. Valid values: 10 to 500.
	//
	// Default value:
	//
	// - If you do not set this parameter, the default value is 20.
	//
	// - If the value you set is greater than 500, the default value is 500.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The IDs of network interface controllers (NICs).
	NetworkInterfaceId []*string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" type:"Repeated"`
	// The name of the elastic network interfaces (ENIs). You can use this parameter to filter network interface controllers (NICs) by name.
	//
	// example:
	//
	// test-eni-name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If this is the first query or no next query exists, leave this parameter empty.
	//
	// - If a next query exists, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The primary private IP IPv4 address of the elastic network interfaces (ENIs). You can use this parameter to filter network interface controllers (NICs) by primary private IP address.
	//
	// example:
	//
	// ``192.168.**.**``
	PrimaryIpAddress     *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of a vSwitch in the VPC-connected instance.
	//
	// You can query network interface controller (NIC) information for only one vSwitch at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-p0w9s2ig1jnwgrbzl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC-connected instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-p0w9alkte4w2htrqe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListGrantVSwitchEnisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGrantVSwitchEnisRequest) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchEnisRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListGrantVSwitchEnisRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListGrantVSwitchEnisRequest) GetNetworkInterfaceId() []*string {
	return s.NetworkInterfaceId
}

func (s *ListGrantVSwitchEnisRequest) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *ListGrantVSwitchEnisRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGrantVSwitchEnisRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListGrantVSwitchEnisRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListGrantVSwitchEnisRequest) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *ListGrantVSwitchEnisRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListGrantVSwitchEnisRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListGrantVSwitchEnisRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListGrantVSwitchEnisRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGrantVSwitchEnisRequest) SetCenId(v string) *ListGrantVSwitchEnisRequest {
	s.CenId = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetMaxResults(v int64) *ListGrantVSwitchEnisRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetNetworkInterfaceId(v []*string) *ListGrantVSwitchEnisRequest {
	s.NetworkInterfaceId = v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetNetworkInterfaceName(v string) *ListGrantVSwitchEnisRequest {
	s.NetworkInterfaceName = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetNextToken(v string) *ListGrantVSwitchEnisRequest {
	s.NextToken = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetOwnerAccount(v string) *ListGrantVSwitchEnisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetOwnerId(v int64) *ListGrantVSwitchEnisRequest {
	s.OwnerId = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetPrimaryIpAddress(v string) *ListGrantVSwitchEnisRequest {
	s.PrimaryIpAddress = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetResourceOwnerAccount(v string) *ListGrantVSwitchEnisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetResourceOwnerId(v int64) *ListGrantVSwitchEnisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetVSwitchId(v string) *ListGrantVSwitchEnisRequest {
	s.VSwitchId = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) SetVpcId(v string) *ListGrantVSwitchEnisRequest {
	s.VpcId = &v
	return s
}

func (s *ListGrantVSwitchEnisRequest) Validate() error {
	return dara.Validate(s)
}

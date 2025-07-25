// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpamPoolCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *AddIpamPoolCidrRequest
	GetCidr() *string
	SetClientToken(v string) *AddIpamPoolCidrRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddIpamPoolCidrRequest
	GetDryRun() *bool
	SetIpamPoolId(v string) *AddIpamPoolCidrRequest
	GetIpamPoolId() *string
	SetNetmaskLength(v int32) *AddIpamPoolCidrRequest
	GetNetmaskLength() *int32
	SetRegionId(v string) *AddIpamPoolCidrRequest
	GetRegionId() *string
}

type AddIpamPoolCidrRequest struct {
	// The CIDR block that you want to provision.
	//
	// >  Only IPv4 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId    *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	NetmaskLength *int32  `json:"NetmaskLength,omitempty" xml:"NetmaskLength,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddIpamPoolCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIpamPoolCidrRequest) GoString() string {
	return s.String()
}

func (s *AddIpamPoolCidrRequest) GetCidr() *string {
	return s.Cidr
}

func (s *AddIpamPoolCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddIpamPoolCidrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddIpamPoolCidrRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *AddIpamPoolCidrRequest) GetNetmaskLength() *int32 {
	return s.NetmaskLength
}

func (s *AddIpamPoolCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddIpamPoolCidrRequest) SetCidr(v string) *AddIpamPoolCidrRequest {
	s.Cidr = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetClientToken(v string) *AddIpamPoolCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetDryRun(v bool) *AddIpamPoolCidrRequest {
	s.DryRun = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetIpamPoolId(v string) *AddIpamPoolCidrRequest {
	s.IpamPoolId = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetNetmaskLength(v int32) *AddIpamPoolCidrRequest {
	s.NetmaskLength = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetRegionId(v string) *AddIpamPoolCidrRequest {
	s.RegionId = &v
	return s
}

func (s *AddIpamPoolCidrRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *DeleteIpamPoolCidrRequest
	GetCidr() *string
	SetClientToken(v string) *DeleteIpamPoolCidrRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpamPoolCidrRequest
	GetDryRun() *bool
	SetIpamPoolId(v string) *DeleteIpamPoolCidrRequest
	GetIpamPoolId() *string
	SetRegionId(v string) *DeleteIpamPoolCidrRequest
	GetRegionId() *string
}

type DeleteIpamPoolCidrRequest struct {
	// The provisioned CIDR block to be deleted.
	//
	// >  Only IPv4 CIDR blocks are supported.
	//
	// This parameter is required.
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
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
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

func (s DeleteIpamPoolCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolCidrRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolCidrRequest) GetCidr() *string {
	return s.Cidr
}

func (s *DeleteIpamPoolCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpamPoolCidrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpamPoolCidrRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *DeleteIpamPoolCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpamPoolCidrRequest) SetCidr(v string) *DeleteIpamPoolCidrRequest {
	s.Cidr = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) SetClientToken(v string) *DeleteIpamPoolCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) SetDryRun(v bool) *DeleteIpamPoolCidrRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) SetIpamPoolId(v string) *DeleteIpamPoolCidrRequest {
	s.IpamPoolId = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) SetRegionId(v string) *DeleteIpamPoolCidrRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) Validate() error {
	return dara.Validate(s)
}

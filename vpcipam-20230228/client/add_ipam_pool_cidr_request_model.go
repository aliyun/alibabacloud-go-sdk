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
	// The CIDR block to provision.
	//
	// > Private top-level pools support provisioning only by specifying a CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The client token that is used to ensure the idempotence of the request. A client-generated value that is unique across different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the RequestId of the API request as the ClientToken. The RequestId is different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without provisioning a CIDR block for the IPAM pool. The system checks the required parameters, request format, and business limits. If the check fails, the corresponding error is returned. If the check succeeds, the error code DryRunOperation is returned.
	//
	// - **false*	- (default): sends the request. After the check succeeds, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// Provisions a CIDR block by specifying a netmask length.
	//
	// > Public IPv6 top-level pools support provisioning only by specifying a netmask length.
	//
	// example:
	//
	// 24
	NetmaskLength *int32 `json:"NetmaskLength,omitempty" xml:"NetmaskLength,omitempty"`
	// The ID of the IPAM hosted region.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to obtain the region ID.
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

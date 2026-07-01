// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamPoolAllocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *CreateIpamPoolAllocationRequest
	GetCidr() *string
	SetCidrMask(v int32) *CreateIpamPoolAllocationRequest
	GetCidrMask() *int32
	SetClientToken(v string) *CreateIpamPoolAllocationRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateIpamPoolAllocationRequest
	GetDryRun() *bool
	SetIpamPoolAllocationDescription(v string) *CreateIpamPoolAllocationRequest
	GetIpamPoolAllocationDescription() *string
	SetIpamPoolAllocationName(v string) *CreateIpamPoolAllocationRequest
	GetIpamPoolAllocationName() *string
	SetIpamPoolId(v string) *CreateIpamPoolAllocationRequest
	GetIpamPoolId() *string
	SetRegionId(v string) *CreateIpamPoolAllocationRequest
	GetRegionId() *string
}

type CreateIpamPoolAllocationRequest struct {
	// The CIDR block to allocate from the IPAM pool.
	//
	// > You must specify either the **Cidr*	- or **CidrMask*	- parameter.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The mask of the CIDR block to allocate from the IPAM pool.
	//
	// > You must specify either the **Cidr*	- or **CidrMask*	- parameter.
	//
	// example:
	//
	// 24
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a value on your client to make sure that the value is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the request as the ClientToken. The RequestId may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Sends a check request. The custom reserved CIDR block is not created. The system checks for required parameters, request format, and service limits. If the check fails, an error is returned. If the check passes, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): Sends a normal request. After the request passes the check, a 2xx HTTP status code is returned and the custom reserved CIDR block is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the custom reserved CIDR block.
	//
	// The description must be 1 to 256 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. The default value is an empty string.
	//
	// example:
	//
	// test description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The name of the custom reserved CIDR block.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The ID of the region where you want to create the custom reserved CIDR block.
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

func (s CreateIpamPoolAllocationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamPoolAllocationRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolAllocationRequest) GetCidr() *string {
	return s.Cidr
}

func (s *CreateIpamPoolAllocationRequest) GetCidrMask() *int32 {
	return s.CidrMask
}

func (s *CreateIpamPoolAllocationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpamPoolAllocationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateIpamPoolAllocationRequest) GetIpamPoolAllocationDescription() *string {
	return s.IpamPoolAllocationDescription
}

func (s *CreateIpamPoolAllocationRequest) GetIpamPoolAllocationName() *string {
	return s.IpamPoolAllocationName
}

func (s *CreateIpamPoolAllocationRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *CreateIpamPoolAllocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpamPoolAllocationRequest) SetCidr(v string) *CreateIpamPoolAllocationRequest {
	s.Cidr = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetCidrMask(v int32) *CreateIpamPoolAllocationRequest {
	s.CidrMask = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetClientToken(v string) *CreateIpamPoolAllocationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetDryRun(v bool) *CreateIpamPoolAllocationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetIpamPoolAllocationDescription(v string) *CreateIpamPoolAllocationRequest {
	s.IpamPoolAllocationDescription = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetIpamPoolAllocationName(v string) *CreateIpamPoolAllocationRequest {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetIpamPoolId(v string) *CreateIpamPoolAllocationRequest {
	s.IpamPoolId = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetRegionId(v string) *CreateIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) Validate() error {
	return dara.Validate(s)
}

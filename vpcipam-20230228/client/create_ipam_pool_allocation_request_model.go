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
	// Enter a CIDR block to reserve a custom CIDR block.
	//
	// **Usage notes*	- Specify at least one of **Cidr*	- and **CidrMask*	- .
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// Enter a mask to reserve a custom CIDR block.
	//
	// **Usage notes*	- Specify at least one of **Cidr*	- and **CidrMask*	- .
	//
	// example:
	//
	// 24
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// **Usage notes*	- If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
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
	// The description of the allocation.
	//
	// example:
	//
	// test description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The name of the allocation.
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
	// The region ID of the custom CIDR block that you want to reserve.
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

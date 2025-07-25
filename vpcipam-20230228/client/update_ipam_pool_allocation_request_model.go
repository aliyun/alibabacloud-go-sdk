// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamPoolAllocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateIpamPoolAllocationRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateIpamPoolAllocationRequest
	GetDryRun() *bool
	SetIpamPoolAllocationDescription(v string) *UpdateIpamPoolAllocationRequest
	GetIpamPoolAllocationDescription() *string
	SetIpamPoolAllocationId(v string) *UpdateIpamPoolAllocationRequest
	GetIpamPoolAllocationId() *string
	SetIpamPoolAllocationName(v string) *UpdateIpamPoolAllocationRequest
	GetIpamPoolAllocationName() *string
	SetRegionId(v string) *UpdateIpamPoolAllocationRequest
	GetRegionId() *string
}

type UpdateIpamPoolAllocationRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to precheck this request. Valid values:
	//
	// 	- **true**: performs a dry run without modifying the CIDR blocks of IPAM pools. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and the actual request. If the request passes the check, an HTTP 2xx status code is returned and the CIDR allocation information of the IPAM address pool is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the CIDR allocation of the IPAM pool.
	//
	// The description must be 1 to 256 characters in length and start with a letter, but cannot start with a `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// test description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The name of the CIDR allocation of the IPAM pool.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the region where you want to perform the operation. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateIpamPoolAllocationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamPoolAllocationRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolAllocationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIpamPoolAllocationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateIpamPoolAllocationRequest) GetIpamPoolAllocationDescription() *string {
	return s.IpamPoolAllocationDescription
}

func (s *UpdateIpamPoolAllocationRequest) GetIpamPoolAllocationId() *string {
	return s.IpamPoolAllocationId
}

func (s *UpdateIpamPoolAllocationRequest) GetIpamPoolAllocationName() *string {
	return s.IpamPoolAllocationName
}

func (s *UpdateIpamPoolAllocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpamPoolAllocationRequest) SetClientToken(v string) *UpdateIpamPoolAllocationRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetDryRun(v bool) *UpdateIpamPoolAllocationRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetIpamPoolAllocationDescription(v string) *UpdateIpamPoolAllocationRequest {
	s.IpamPoolAllocationDescription = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetIpamPoolAllocationId(v string) *UpdateIpamPoolAllocationRequest {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetIpamPoolAllocationName(v string) *UpdateIpamPoolAllocationRequest {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetRegionId(v string) *UpdateIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) Validate() error {
	return dara.Validate(s)
}

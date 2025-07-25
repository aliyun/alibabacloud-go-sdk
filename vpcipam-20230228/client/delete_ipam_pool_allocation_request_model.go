// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolAllocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpamPoolAllocationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpamPoolAllocationRequest
	GetDryRun() *bool
	SetIpamPoolAllocationId(v string) *DeleteIpamPoolAllocationRequest
	GetIpamPoolAllocationId() *string
	SetRegionId(v string) *DeleteIpamPoolAllocationRequest
	GetRegionId() *string
}

type DeleteIpamPoolAllocationRequest struct {
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
	// 	- **false*	- (default): sends the API request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the custom reserved CIDR block to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-alloc-c4vhvr3b22mmc6****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The region ID of the custom reserved CIDR block.
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

func (s DeleteIpamPoolAllocationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolAllocationRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolAllocationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpamPoolAllocationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpamPoolAllocationRequest) GetIpamPoolAllocationId() *string {
	return s.IpamPoolAllocationId
}

func (s *DeleteIpamPoolAllocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpamPoolAllocationRequest) SetClientToken(v string) *DeleteIpamPoolAllocationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamPoolAllocationRequest) SetDryRun(v bool) *DeleteIpamPoolAllocationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamPoolAllocationRequest) SetIpamPoolAllocationId(v string) *DeleteIpamPoolAllocationRequest {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *DeleteIpamPoolAllocationRequest) SetRegionId(v string) *DeleteIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamPoolAllocationRequest) Validate() error {
	return dara.Validate(s)
}

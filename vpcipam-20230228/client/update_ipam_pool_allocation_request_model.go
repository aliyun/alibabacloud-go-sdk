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
	// A client token to ensure the idempotence of the request. Generate a unique value from your client for each request. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the request ID as the client token. The request ID is different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Sends a check request without modifying the CIDR allocation. The system checks for required parameters, request format, and service limits. If the check fails, an error is returned. If the check passes, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): Sends a normal request. After the check passes, an HTTP 2xx status code is returned and the CIDR allocation is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the CIDR allocation.
	//
	// The description must be 1 to 256 characters long and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`. If you do not specify this parameter, the description is empty.
	//
	// example:
	//
	// test description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The ID of the CIDR allocation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The name of the CIDR allocation.
	//
	// The name must be 1 to 128 characters long. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the region where the CIDR allocation is located. To obtain a region ID, call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation.
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

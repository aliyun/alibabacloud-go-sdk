// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpamPoolAllocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpamPoolAllocationId(v string) *GetIpamPoolAllocationRequest
	GetIpamPoolAllocationId() *string
	SetRegionId(v string) *GetIpamPoolAllocationRequest
	GetRegionId() *string
}

type GetIpamPoolAllocationRequest struct {
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The region of the IPAM pool.
	//
	// >  If the IPAM pool to which CIDR allocation belongs has the region attribute, this parameter is the region of the IPAM pool. If not, this parameter is the IPAM hosted region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIpamPoolAllocationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIpamPoolAllocationRequest) GoString() string {
	return s.String()
}

func (s *GetIpamPoolAllocationRequest) GetIpamPoolAllocationId() *string {
	return s.IpamPoolAllocationId
}

func (s *GetIpamPoolAllocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIpamPoolAllocationRequest) SetIpamPoolAllocationId(v string) *GetIpamPoolAllocationRequest {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *GetIpamPoolAllocationRequest) SetRegionId(v string) *GetIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

func (s *GetIpamPoolAllocationRequest) Validate() error {
	return dara.Validate(s)
}

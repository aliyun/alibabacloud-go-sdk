// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpamPoolAllocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *GetIpamPoolAllocationResponseBody
	GetCidr() *string
	SetCreationTime(v string) *GetIpamPoolAllocationResponseBody
	GetCreationTime() *string
	SetIpamPoolAllocationDescription(v string) *GetIpamPoolAllocationResponseBody
	GetIpamPoolAllocationDescription() *string
	SetIpamPoolAllocationId(v string) *GetIpamPoolAllocationResponseBody
	GetIpamPoolAllocationId() *string
	SetIpamPoolAllocationName(v string) *GetIpamPoolAllocationResponseBody
	GetIpamPoolAllocationName() *string
	SetIpamPoolId(v string) *GetIpamPoolAllocationResponseBody
	GetIpamPoolId() *string
	SetRegionId(v string) *GetIpamPoolAllocationResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetIpamPoolAllocationResponseBody
	GetRequestId() *string
	SetResourceId(v string) *GetIpamPoolAllocationResponseBody
	GetResourceId() *string
	SetResourceOwnerId(v int64) *GetIpamPoolAllocationResponseBody
	GetResourceOwnerId() *int64
	SetResourceRegionId(v string) *GetIpamPoolAllocationResponseBody
	GetResourceRegionId() *string
	SetResourceType(v string) *GetIpamPoolAllocationResponseBody
	GetResourceType() *string
	SetSourceCidr(v string) *GetIpamPoolAllocationResponseBody
	GetSourceCidr() *string
	SetStatus(v string) *GetIpamPoolAllocationResponseBody
	GetStatus() *string
}

type GetIpamPoolAllocationResponseBody struct {
	// The allocated CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2024-10-15T10:24:19+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the CIDR allocation of the IPAM pool.
	//
	// The description must be 1 to 256 characters in length and start with a letter, but cannot start with a `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// ipam pool allocation description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
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
	// ipam pool allocation name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The region of the IPAM pool.
	//
	// >  If the IPAM pool to which the CIDR block allocation belongs has the region attribute, this parameter is the region of the IPAM pool. If not, this parameter is the IPAM hosted region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3748DEFF-68BE-5EED-9937-7C1D0C21BAB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource to which the CIDR block is allocated.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 1616080591216318
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The effective region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource to which the CIDR block is allocated. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **IpamPool**
	//
	// 	- **Custom**
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.0.0/16
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The instance state. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIpamPoolAllocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIpamPoolAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpamPoolAllocationResponseBody) GetCidr() *string {
	return s.Cidr
}

func (s *GetIpamPoolAllocationResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetIpamPoolAllocationResponseBody) GetIpamPoolAllocationDescription() *string {
	return s.IpamPoolAllocationDescription
}

func (s *GetIpamPoolAllocationResponseBody) GetIpamPoolAllocationId() *string {
	return s.IpamPoolAllocationId
}

func (s *GetIpamPoolAllocationResponseBody) GetIpamPoolAllocationName() *string {
	return s.IpamPoolAllocationName
}

func (s *GetIpamPoolAllocationResponseBody) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *GetIpamPoolAllocationResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIpamPoolAllocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIpamPoolAllocationResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetIpamPoolAllocationResponseBody) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetIpamPoolAllocationResponseBody) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *GetIpamPoolAllocationResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetIpamPoolAllocationResponseBody) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *GetIpamPoolAllocationResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetIpamPoolAllocationResponseBody) SetCidr(v string) *GetIpamPoolAllocationResponseBody {
	s.Cidr = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetCreationTime(v string) *GetIpamPoolAllocationResponseBody {
	s.CreationTime = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetIpamPoolAllocationDescription(v string) *GetIpamPoolAllocationResponseBody {
	s.IpamPoolAllocationDescription = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetIpamPoolAllocationId(v string) *GetIpamPoolAllocationResponseBody {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetIpamPoolAllocationName(v string) *GetIpamPoolAllocationResponseBody {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetIpamPoolId(v string) *GetIpamPoolAllocationResponseBody {
	s.IpamPoolId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetRegionId(v string) *GetIpamPoolAllocationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetRequestId(v string) *GetIpamPoolAllocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetResourceId(v string) *GetIpamPoolAllocationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetResourceOwnerId(v int64) *GetIpamPoolAllocationResponseBody {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetResourceRegionId(v string) *GetIpamPoolAllocationResponseBody {
	s.ResourceRegionId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetResourceType(v string) *GetIpamPoolAllocationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetSourceCidr(v string) *GetIpamPoolAllocationResponseBody {
	s.SourceCidr = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetStatus(v string) *GetIpamPoolAllocationResponseBody {
	s.Status = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) Validate() error {
	return dara.Validate(s)
}

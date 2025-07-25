// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamPoolAllocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *CreateIpamPoolAllocationResponseBody
	GetCidr() *string
	SetIpamPoolAllocationId(v string) *CreateIpamPoolAllocationResponseBody
	GetIpamPoolAllocationId() *string
	SetRequestId(v string) *CreateIpamPoolAllocationResponseBody
	GetRequestId() *string
	SetSourceCidr(v string) *CreateIpamPoolAllocationResponseBody
	GetSourceCidr() *string
}

type CreateIpamPoolAllocationResponseBody struct {
	// The custom reserved CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ID of the custom reserved CIDR block.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CE9CDAE5-341E-5D0B-AC8A-2BAC707D3EB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.0.0/16
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
}

func (s CreateIpamPoolAllocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamPoolAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolAllocationResponseBody) GetCidr() *string {
	return s.Cidr
}

func (s *CreateIpamPoolAllocationResponseBody) GetIpamPoolAllocationId() *string {
	return s.IpamPoolAllocationId
}

func (s *CreateIpamPoolAllocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpamPoolAllocationResponseBody) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *CreateIpamPoolAllocationResponseBody) SetCidr(v string) *CreateIpamPoolAllocationResponseBody {
	s.Cidr = &v
	return s
}

func (s *CreateIpamPoolAllocationResponseBody) SetIpamPoolAllocationId(v string) *CreateIpamPoolAllocationResponseBody {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *CreateIpamPoolAllocationResponseBody) SetRequestId(v string) *CreateIpamPoolAllocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpamPoolAllocationResponseBody) SetSourceCidr(v string) *CreateIpamPoolAllocationResponseBody {
	s.SourceCidr = &v
	return s
}

func (s *CreateIpamPoolAllocationResponseBody) Validate() error {
	return dara.Validate(s)
}

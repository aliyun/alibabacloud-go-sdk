// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v []*string) *DescribeDcdnIpaDomainCidrResponseBody
	GetCidr() []*string
	SetRequestId(v string) *DescribeDcdnIpaDomainCidrResponseBody
	GetRequestId() *string
}

type DescribeDcdnIpaDomainCidrResponseBody struct {
	// The back-to-origin IPv4 and IPv6 CIDR blocks.
	//
	// example:
	//
	// ["1.1.1.0/24","2.2.2.0/24","1111:2222:3333:4444:5555:0:0:0/80"]
	Cidr []*string `json:"Cidr,omitempty" xml:"Cidr,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpaDomainCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainCidrResponseBody) GetCidr() []*string {
	return s.Cidr
}

func (s *DescribeDcdnIpaDomainCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnIpaDomainCidrResponseBody) SetCidr(v []*string) *DescribeDcdnIpaDomainCidrResponseBody {
	s.Cidr = v
	return s
}

func (s *DescribeDcdnIpaDomainCidrResponseBody) SetRequestId(v string) *DescribeDcdnIpaDomainCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnIpaDomainCidrResponseBody) Validate() error {
	return dara.Validate(s)
}

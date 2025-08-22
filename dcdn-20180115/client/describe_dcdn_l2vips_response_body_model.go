// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnL2VipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnL2VipsResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDcdnL2VipsResponseBody
	GetRequestId() *string
	SetVips(v []*string) *DescribeDcdnL2VipsResponseBody
	GetVips() []*string
}

type DescribeDcdnL2VipsResponseBody struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 820E7900-5CA9-4AEF-B0DD-20ED5F64BE55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The virtual IP addresses (VIPs).
	Vips []*string `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Repeated"`
}

func (s DescribeDcdnL2VipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnL2VipsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnL2VipsResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnL2VipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnL2VipsResponseBody) GetVips() []*string {
	return s.Vips
}

func (s *DescribeDcdnL2VipsResponseBody) SetDomainName(v string) *DescribeDcdnL2VipsResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnL2VipsResponseBody) SetRequestId(v string) *DescribeDcdnL2VipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnL2VipsResponseBody) SetVips(v []*string) *DescribeDcdnL2VipsResponseBody {
	s.Vips = v
	return s
}

func (s *DescribeDcdnL2VipsResponseBody) Validate() error {
	return dara.Validate(s)
}

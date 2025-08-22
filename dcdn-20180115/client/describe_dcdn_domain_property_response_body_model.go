// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainPropertyResponseBody
	GetDomainName() *string
	SetProtocol(v string) *DescribeDcdnDomainPropertyResponseBody
	GetProtocol() *string
	SetRequestId(v string) *DescribeDcdnDomainPropertyResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainPropertyResponseBody struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **udp**
	//
	// 	- **tcp**
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPropertyResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainPropertyResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDcdnDomainPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainPropertyResponseBody) SetDomainName(v string) *DescribeDcdnDomainPropertyResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainPropertyResponseBody) SetProtocol(v string) *DescribeDcdnDomainPropertyResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeDcdnDomainPropertyResponseBody) SetRequestId(v string) *DescribeDcdnDomainPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}

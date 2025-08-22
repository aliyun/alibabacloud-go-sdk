// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnL2VipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnL2VipsRequest
	GetDomainName() *string
}

type DescribeDcdnL2VipsRequest struct {
	// The domain name. You can specify only one domain name in each request. If you do not specify this parameter, the origin CIDR blocks of all domain names in your account in the whitelist are returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnL2VipsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnL2VipsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnL2VipsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnL2VipsRequest) SetDomainName(v string) *DescribeDcdnL2VipsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnL2VipsRequest) Validate() error {
	return dara.Validate(s)
}

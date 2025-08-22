// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnIpaDomainCidrRequest
	GetDomainName() *string
}

type DescribeDcdnIpaDomainCidrRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnIpaDomainCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainCidrRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainCidrRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnIpaDomainCidrRequest) SetDomainName(v string) *DescribeDcdnIpaDomainCidrRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaDomainCidrRequest) Validate() error {
	return dara.Validate(s)
}

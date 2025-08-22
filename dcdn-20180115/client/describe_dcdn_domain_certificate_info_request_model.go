// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCertificateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainCertificateInfoRequest
	GetDomainName() *string
}

type DescribeDcdnDomainCertificateInfoRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnDomainCertificateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainCertificateInfoRequest) SetDomainName(v string) *DescribeDcdnDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCertificateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainCertificateInfoRequest
	GetDomainName() *string
}

type DescribeDomainCertificateInfoRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainCertificateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainCertificateInfoRequest) SetDomainName(v string) *DescribeDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCertificateInfoRequest) Validate() error {
	return dara.Validate(s)
}

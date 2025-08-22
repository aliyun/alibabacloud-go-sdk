// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnOriginSiteHealthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnOriginSiteHealthStatusRequest
	GetDomainName() *string
}

type DescribeDcdnOriginSiteHealthStatusRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnOriginSiteHealthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnOriginSiteHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnOriginSiteHealthStatusRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnOriginSiteHealthStatusRequest) SetDomainName(v string) *DescribeDcdnOriginSiteHealthStatusRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnOriginSiteHealthStatusRequest) Validate() error {
	return dara.Validate(s)
}

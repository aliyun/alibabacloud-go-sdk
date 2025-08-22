// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnWafDomainDetailRequest
	GetDomainName() *string
}

type DescribeDcdnWafDomainDetailRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request. Exact match is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnWafDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainDetailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafDomainDetailRequest) SetDomainName(v string) *DescribeDcdnWafDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainCnameRequest
	GetDomainName() *string
}

type DescribeDcdnDomainCnameRequest struct {
	// The accelerated domain name. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnDomainCnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainCnameRequest) SetDomainName(v string) *DescribeDcdnDomainCnameRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCnameRequest) Validate() error {
	return dara.Validate(s)
}

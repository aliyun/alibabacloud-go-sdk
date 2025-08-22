// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainPropertyRequest
	GetDomainName() *string
}

type DescribeDcdnDomainPropertyRequest struct {
	// The accelerated domain name that you want to query. You can specify only one domain name in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnDomainPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainPropertyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPropertyRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainPropertyRequest) SetDomainName(v string) *DescribeDcdnDomainPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainPropertyRequest) Validate() error {
	return dara.Validate(s)
}

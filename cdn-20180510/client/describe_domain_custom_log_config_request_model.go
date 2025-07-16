// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCustomLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainCustomLogConfigRequest
	GetDomainName() *string
}

type DescribeDomainCustomLogConfigRequest struct {
	// The domain name. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainCustomLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCustomLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCustomLogConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainCustomLogConfigRequest) SetDomainName(v string) *DescribeDomainCustomLogConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCustomLogConfigRequest) Validate() error {
	return dara.Validate(s)
}

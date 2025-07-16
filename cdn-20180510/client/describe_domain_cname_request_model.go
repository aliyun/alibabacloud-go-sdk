// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainCnameRequest
	GetDomainName() *string
}

type DescribeDomainCnameRequest struct {
	// The accelerated domain name that you want to query. Separate multiple domain names with commas (,). This parameter cannot be left empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// pay.slci6c.mbolsos.com,mch.b7r2v7.mbolsos.com,p.h99e.mbolsos.com,p.xmko.mbolsos.com,p.f2kd.mbolsos.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainCnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCnameRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainCnameRequest) SetDomainName(v string) *DescribeDomainCnameRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCnameRequest) Validate() error {
	return dara.Validate(s)
}

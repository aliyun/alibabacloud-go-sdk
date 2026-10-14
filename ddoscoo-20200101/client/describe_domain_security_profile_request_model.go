// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecurityProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainSecurityProfileRequest
	GetDomain() *string
}

type DescribeDomainSecurityProfileRequest struct {
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainSecurityProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecurityProfileRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecurityProfileRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainSecurityProfileRequest) SetDomain(v string) *DescribeDomainSecurityProfileRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainSecurityProfileRequest) Validate() error {
	return dara.Validate(s)
}

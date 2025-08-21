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
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
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

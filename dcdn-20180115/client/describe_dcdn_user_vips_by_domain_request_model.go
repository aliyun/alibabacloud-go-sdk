// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserVipsByDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailable(v string) *DescribeDcdnUserVipsByDomainRequest
	GetAvailable() *string
	SetDomainName(v string) *DescribeDcdnUserVipsByDomainRequest
	GetDomainName() *string
}

type DescribeDcdnUserVipsByDomainRequest struct {
	// Specifies whether to query the virtual IP addresses of only healthy POPs. Valid value:
	//
	// 	- **on**: queries healthy VIPs.
	//
	// 	- **off**: queries all VIPs.
	//
	// example:
	//
	// on
	Available *string `json:"Available,omitempty" xml:"Available,omitempty"`
	// The accelerated domain name. You can specify only one domain name.
	//
	// Enumeration values: example.com
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnUserVipsByDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserVipsByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserVipsByDomainRequest) GetAvailable() *string {
	return s.Available
}

func (s *DescribeDcdnUserVipsByDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnUserVipsByDomainRequest) SetAvailable(v string) *DescribeDcdnUserVipsByDomainRequest {
	s.Available = &v
	return s
}

func (s *DescribeDcdnUserVipsByDomainRequest) SetDomainName(v string) *DescribeDcdnUserVipsByDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserVipsByDomainRequest) Validate() error {
	return dara.Validate(s)
}

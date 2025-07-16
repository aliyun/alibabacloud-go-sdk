// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVipsByDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailable(v string) *DescribeUserVipsByDomainRequest
	GetAvailable() *string
	SetDomainName(v string) *DescribeUserVipsByDomainRequest
	GetDomainName() *string
}

type DescribeUserVipsByDomainRequest struct {
	// Specifies whether to query the virtual IP addresses of only healthy CDN POPs. Valid values:
	//
	// 	- **on**: healthy CDN edge nodes.
	//
	// 	- **off**: all CDN edge nodes.
	//
	// example:
	//
	// on
	Available *string `json:"Available,omitempty" xml:"Available,omitempty"`
	// The accelerated domain name. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeUserVipsByDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVipsByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserVipsByDomainRequest) GetAvailable() *string {
	return s.Available
}

func (s *DescribeUserVipsByDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUserVipsByDomainRequest) SetAvailable(v string) *DescribeUserVipsByDomainRequest {
	s.Available = &v
	return s
}

func (s *DescribeUserVipsByDomainRequest) SetDomainName(v string) *DescribeUserVipsByDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUserVipsByDomainRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnFullDomainsBlockIPConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIPList(v string) *DescribeCdnFullDomainsBlockIPConfigRequest
	GetIPList() *string
}

type DescribeCdnFullDomainsBlockIPConfigRequest struct {
	// The IP address or CIDR block to query. Separate multiple values with commas (,). You can specify up to 50 IP addresses or CIDR blocks.
	//
	// example:
	//
	// 1.XXX.XXX.1,2.XXX.XXX.2
	IPList *string `json:"IPList,omitempty" xml:"IPList,omitempty"`
}

func (s DescribeCdnFullDomainsBlockIPConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnFullDomainsBlockIPConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnFullDomainsBlockIPConfigRequest) GetIPList() *string {
	return s.IPList
}

func (s *DescribeCdnFullDomainsBlockIPConfigRequest) SetIPList(v string) *DescribeCdnFullDomainsBlockIPConfigRequest {
	s.IPList = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPConfigRequest) Validate() error {
	return dara.Validate(s)
}

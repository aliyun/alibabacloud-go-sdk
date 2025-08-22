// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnFullDomainsBlockIPConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIPList(v string) *DescribeDcdnFullDomainsBlockIPConfigRequest
	GetIPList() *string
}

type DescribeDcdnFullDomainsBlockIPConfigRequest struct {
	// The IP address or CIDR block to query. Separate multiple values with commas (,). You can specify up to 50 IP addresses or CIDR blocks.
	//
	// example:
	//
	// 10.XX.XX.10/24
	IPList *string `json:"IPList,omitempty" xml:"IPList,omitempty"`
}

func (s DescribeDcdnFullDomainsBlockIPConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnFullDomainsBlockIPConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnFullDomainsBlockIPConfigRequest) GetIPList() *string {
	return s.IPList
}

func (s *DescribeDcdnFullDomainsBlockIPConfigRequest) SetIPList(v string) *DescribeDcdnFullDomainsBlockIPConfigRequest {
	s.IPList = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPConfigRequest) Validate() error {
	return dara.Validate(s)
}

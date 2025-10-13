// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomDNS interface {
	dara.Model
	String() string
	GoString() string
	SetDnsOptions(v []*DNSOption) *CustomDNS
	GetDnsOptions() []*DNSOption
	SetNameServers(v []*string) *CustomDNS
	GetNameServers() []*string
	SetSearches(v []*string) *CustomDNS
	GetSearches() []*string
}

type CustomDNS struct {
	DnsOptions  []*DNSOption `json:"dnsOptions,omitempty" xml:"dnsOptions,omitempty" type:"Repeated"`
	NameServers []*string    `json:"nameServers,omitempty" xml:"nameServers,omitempty" type:"Repeated"`
	Searches    []*string    `json:"searches,omitempty" xml:"searches,omitempty" type:"Repeated"`
}

func (s CustomDNS) String() string {
	return dara.Prettify(s)
}

func (s CustomDNS) GoString() string {
	return s.String()
}

func (s *CustomDNS) GetDnsOptions() []*DNSOption {
	return s.DnsOptions
}

func (s *CustomDNS) GetNameServers() []*string {
	return s.NameServers
}

func (s *CustomDNS) GetSearches() []*string {
	return s.Searches
}

func (s *CustomDNS) SetDnsOptions(v []*DNSOption) *CustomDNS {
	s.DnsOptions = v
	return s
}

func (s *CustomDNS) SetNameServers(v []*string) *CustomDNS {
	s.NameServers = v
	return s
}

func (s *CustomDNS) SetSearches(v []*string) *CustomDNS {
	s.Searches = v
	return s
}

func (s *CustomDNS) Validate() error {
	if s.DnsOptions != nil {
		for _, item := range s.DnsOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

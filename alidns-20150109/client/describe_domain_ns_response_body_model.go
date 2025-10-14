// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainNsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllAliDns(v bool) *DescribeDomainNsResponseBody
	GetAllAliDns() *bool
	SetDetectFailedReasonCode(v string) *DescribeDomainNsResponseBody
	GetDetectFailedReasonCode() *string
	SetDnsServers(v *DescribeDomainNsResponseBodyDnsServers) *DescribeDomainNsResponseBody
	GetDnsServers() *DescribeDomainNsResponseBodyDnsServers
	SetExpectDnsServers(v *DescribeDomainNsResponseBodyExpectDnsServers) *DescribeDomainNsResponseBody
	GetExpectDnsServers() *DescribeDomainNsResponseBodyExpectDnsServers
	SetIncludeAliDns(v bool) *DescribeDomainNsResponseBody
	GetIncludeAliDns() *bool
	SetRequestId(v string) *DescribeDomainNsResponseBody
	GetRequestId() *string
}

type DescribeDomainNsResponseBody struct {
	// Indicates whether all the name servers are Alibaba Cloud DNS servers.
	//
	// example:
	//
	// true
	AllAliDns *bool `json:"AllAliDns,omitempty" xml:"AllAliDns,omitempty"`
	// The cause code of the detection failure.
	//
	// example:
	//
	// DnsCheck.Failed
	DetectFailedReasonCode *string `json:"DetectFailedReasonCode,omitempty" xml:"DetectFailedReasonCode,omitempty"`
	// The DNS server names configured for the domain name.
	DnsServers *DescribeDomainNsResponseBodyDnsServers `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Struct"`
	// The Domain Name System (DNS) server names assigned by Alibaba Cloud DNS.
	ExpectDnsServers *DescribeDomainNsResponseBodyExpectDnsServers `json:"ExpectDnsServers,omitempty" xml:"ExpectDnsServers,omitempty" type:"Struct"`
	// Indicates whether the name servers include Alibaba Cloud DNS servers.
	//
	// example:
	//
	// true
	IncludeAliDns *bool `json:"IncludeAliDns,omitempty" xml:"IncludeAliDns,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16C7DC7A-2FA7-4D14-8B12-88A2BB6373DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainNsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainNsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainNsResponseBody) GetAllAliDns() *bool {
	return s.AllAliDns
}

func (s *DescribeDomainNsResponseBody) GetDetectFailedReasonCode() *string {
	return s.DetectFailedReasonCode
}

func (s *DescribeDomainNsResponseBody) GetDnsServers() *DescribeDomainNsResponseBodyDnsServers {
	return s.DnsServers
}

func (s *DescribeDomainNsResponseBody) GetExpectDnsServers() *DescribeDomainNsResponseBodyExpectDnsServers {
	return s.ExpectDnsServers
}

func (s *DescribeDomainNsResponseBody) GetIncludeAliDns() *bool {
	return s.IncludeAliDns
}

func (s *DescribeDomainNsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainNsResponseBody) SetAllAliDns(v bool) *DescribeDomainNsResponseBody {
	s.AllAliDns = &v
	return s
}

func (s *DescribeDomainNsResponseBody) SetDetectFailedReasonCode(v string) *DescribeDomainNsResponseBody {
	s.DetectFailedReasonCode = &v
	return s
}

func (s *DescribeDomainNsResponseBody) SetDnsServers(v *DescribeDomainNsResponseBodyDnsServers) *DescribeDomainNsResponseBody {
	s.DnsServers = v
	return s
}

func (s *DescribeDomainNsResponseBody) SetExpectDnsServers(v *DescribeDomainNsResponseBodyExpectDnsServers) *DescribeDomainNsResponseBody {
	s.ExpectDnsServers = v
	return s
}

func (s *DescribeDomainNsResponseBody) SetIncludeAliDns(v bool) *DescribeDomainNsResponseBody {
	s.IncludeAliDns = &v
	return s
}

func (s *DescribeDomainNsResponseBody) SetRequestId(v string) *DescribeDomainNsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainNsResponseBody) Validate() error {
	if s.DnsServers != nil {
		if err := s.DnsServers.Validate(); err != nil {
			return err
		}
	}
	if s.ExpectDnsServers != nil {
		if err := s.ExpectDnsServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainNsResponseBodyDnsServers struct {
	DnsServer []*string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty" type:"Repeated"`
}

func (s DescribeDomainNsResponseBodyDnsServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainNsResponseBodyDnsServers) GoString() string {
	return s.String()
}

func (s *DescribeDomainNsResponseBodyDnsServers) GetDnsServer() []*string {
	return s.DnsServer
}

func (s *DescribeDomainNsResponseBodyDnsServers) SetDnsServer(v []*string) *DescribeDomainNsResponseBodyDnsServers {
	s.DnsServer = v
	return s
}

func (s *DescribeDomainNsResponseBodyDnsServers) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainNsResponseBodyExpectDnsServers struct {
	ExpectDnsServer []*string `json:"ExpectDnsServer,omitempty" xml:"ExpectDnsServer,omitempty" type:"Repeated"`
}

func (s DescribeDomainNsResponseBodyExpectDnsServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainNsResponseBodyExpectDnsServers) GoString() string {
	return s.String()
}

func (s *DescribeDomainNsResponseBodyExpectDnsServers) GetExpectDnsServer() []*string {
	return s.ExpectDnsServer
}

func (s *DescribeDomainNsResponseBodyExpectDnsServers) SetExpectDnsServer(v []*string) *DescribeDomainNsResponseBodyExpectDnsServers {
	s.ExpectDnsServer = v
	return s
}

func (s *DescribeDomainNsResponseBodyExpectDnsServers) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHichinaDomainDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNewDnsServers(v *ModifyHichinaDomainDNSResponseBodyNewDnsServers) *ModifyHichinaDomainDNSResponseBody
	GetNewDnsServers() *ModifyHichinaDomainDNSResponseBodyNewDnsServers
	SetOriginalDnsServers(v *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers) *ModifyHichinaDomainDNSResponseBody
	GetOriginalDnsServers() *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers
	SetRequestId(v string) *ModifyHichinaDomainDNSResponseBody
	GetRequestId() *string
}

type ModifyHichinaDomainDNSResponseBody struct {
	// The DNS server names after modification.
	NewDnsServers *ModifyHichinaDomainDNSResponseBodyNewDnsServers `json:"NewDnsServers,omitempty" xml:"NewDnsServers,omitempty" type:"Struct"`
	// The DNS server names before modification.
	OriginalDnsServers *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers `json:"OriginalDnsServers,omitempty" xml:"OriginalDnsServers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHichinaDomainDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHichinaDomainDNSResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHichinaDomainDNSResponseBody) GetNewDnsServers() *ModifyHichinaDomainDNSResponseBodyNewDnsServers {
	return s.NewDnsServers
}

func (s *ModifyHichinaDomainDNSResponseBody) GetOriginalDnsServers() *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers {
	return s.OriginalDnsServers
}

func (s *ModifyHichinaDomainDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHichinaDomainDNSResponseBody) SetNewDnsServers(v *ModifyHichinaDomainDNSResponseBodyNewDnsServers) *ModifyHichinaDomainDNSResponseBody {
	s.NewDnsServers = v
	return s
}

func (s *ModifyHichinaDomainDNSResponseBody) SetOriginalDnsServers(v *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers) *ModifyHichinaDomainDNSResponseBody {
	s.OriginalDnsServers = v
	return s
}

func (s *ModifyHichinaDomainDNSResponseBody) SetRequestId(v string) *ModifyHichinaDomainDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHichinaDomainDNSResponseBody) Validate() error {
	if s.NewDnsServers != nil {
		if err := s.NewDnsServers.Validate(); err != nil {
			return err
		}
	}
	if s.OriginalDnsServers != nil {
		if err := s.OriginalDnsServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyHichinaDomainDNSResponseBodyNewDnsServers struct {
	DnsServer []*string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty" type:"Repeated"`
}

func (s ModifyHichinaDomainDNSResponseBodyNewDnsServers) String() string {
	return dara.Prettify(s)
}

func (s ModifyHichinaDomainDNSResponseBodyNewDnsServers) GoString() string {
	return s.String()
}

func (s *ModifyHichinaDomainDNSResponseBodyNewDnsServers) GetDnsServer() []*string {
	return s.DnsServer
}

func (s *ModifyHichinaDomainDNSResponseBodyNewDnsServers) SetDnsServer(v []*string) *ModifyHichinaDomainDNSResponseBodyNewDnsServers {
	s.DnsServer = v
	return s
}

func (s *ModifyHichinaDomainDNSResponseBodyNewDnsServers) Validate() error {
	return dara.Validate(s)
}

type ModifyHichinaDomainDNSResponseBodyOriginalDnsServers struct {
	DnsServer []*string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty" type:"Repeated"`
}

func (s ModifyHichinaDomainDNSResponseBodyOriginalDnsServers) String() string {
	return dara.Prettify(s)
}

func (s ModifyHichinaDomainDNSResponseBodyOriginalDnsServers) GoString() string {
	return s.String()
}

func (s *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers) GetDnsServer() []*string {
	return s.DnsServer
}

func (s *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers) SetDnsServer(v []*string) *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers {
	s.DnsServer = v
	return s
}

func (s *ModifyHichinaDomainDNSResponseBodyOriginalDnsServers) Validate() error {
	return dara.Validate(s)
}

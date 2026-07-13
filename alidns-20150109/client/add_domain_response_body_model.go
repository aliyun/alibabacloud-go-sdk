// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDnsServers(v *AddDomainResponseBodyDnsServers) *AddDomainResponseBody
	GetDnsServers() *AddDomainResponseBodyDnsServers
	SetDomainId(v string) *AddDomainResponseBody
	GetDomainId() *string
	SetDomainName(v string) *AddDomainResponseBody
	GetDomainName() *string
	SetGroupId(v string) *AddDomainResponseBody
	GetGroupId() *string
	SetGroupName(v string) *AddDomainResponseBody
	GetGroupName() *string
	SetPunyCode(v string) *AddDomainResponseBody
	GetPunyCode() *string
	SetRequestId(v string) *AddDomainResponseBody
	GetRequestId() *string
}

type AddDomainResponseBody struct {
	DnsServers *AddDomainResponseBodyDnsServers `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Struct"`
	// The ID of the domain name.
	//
	// example:
	//
	// xxxxx6615cf240c697f9f7e207xxxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the domain name group.
	//
	// example:
	//
	// defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the domain name group.
	//
	// example:
	//
	// MyGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The Punycode for the Chinese domain name. This parameter is returned only for Chinese domain names.
	//
	// example:
	//
	// xn--fsq270a.com
	PunyCode *string `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// xxx508EF-00FD-xxx9-95A4-1E10BACxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainResponseBody) GetDnsServers() *AddDomainResponseBodyDnsServers {
	return s.DnsServers
}

func (s *AddDomainResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *AddDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDomainResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *AddDomainResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *AddDomainResponseBody) GetPunyCode() *string {
	return s.PunyCode
}

func (s *AddDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDomainResponseBody) SetDnsServers(v *AddDomainResponseBodyDnsServers) *AddDomainResponseBody {
	s.DnsServers = v
	return s
}

func (s *AddDomainResponseBody) SetDomainId(v string) *AddDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *AddDomainResponseBody) SetDomainName(v string) *AddDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *AddDomainResponseBody) SetGroupId(v string) *AddDomainResponseBody {
	s.GroupId = &v
	return s
}

func (s *AddDomainResponseBody) SetGroupName(v string) *AddDomainResponseBody {
	s.GroupName = &v
	return s
}

func (s *AddDomainResponseBody) SetPunyCode(v string) *AddDomainResponseBody {
	s.PunyCode = &v
	return s
}

func (s *AddDomainResponseBody) SetRequestId(v string) *AddDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainResponseBody) Validate() error {
	if s.DnsServers != nil {
		if err := s.DnsServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDomainResponseBodyDnsServers struct {
	DnsServer []*string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty" type:"Repeated"`
}

func (s AddDomainResponseBodyDnsServers) String() string {
	return dara.Prettify(s)
}

func (s AddDomainResponseBodyDnsServers) GoString() string {
	return s.String()
}

func (s *AddDomainResponseBodyDnsServers) GetDnsServer() []*string {
	return s.DnsServer
}

func (s *AddDomainResponseBodyDnsServers) SetDnsServer(v []*string) *AddDomainResponseBodyDnsServers {
	s.DnsServer = v
	return s
}

func (s *AddDomainResponseBodyDnsServers) Validate() error {
	return dara.Validate(s)
}

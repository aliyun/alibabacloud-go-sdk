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
	// The Domain Name System (DNS) servers configured for the domain name.
	DnsServers *AddDomainResponseBodyDnsServers `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Struct"`
	// The ID of the domain name.
	//
	// example:
	//
	// 00efd71a-770e-4255-b54e-6fe5659baffe
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// dns-example.top
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the domain name group.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the domain name group.
	//
	// example:
	//
	// MyGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The Punycode for the domain name. This parameter is returned only for Chinese domain names.
	//
	// example:
	//
	// xn--fsq270a.com
	PunyCode *string `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
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
	return dara.Validate(s)
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

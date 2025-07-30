// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainDnsChallengeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainDnsChallenge(v *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) *GetDomainDnsChallengeResponseBody
	GetDomainDnsChallenge() *GetDomainDnsChallengeResponseBodyDomainDnsChallenge
	SetRequestId(v string) *GetDomainDnsChallengeResponseBody
	GetRequestId() *string
}

type GetDomainDnsChallengeResponseBody struct {
	// The DNS challenge records.
	DomainDnsChallenge *GetDomainDnsChallengeResponseBodyDomainDnsChallenge `json:"DomainDnsChallenge,omitempty" xml:"DomainDnsChallenge,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainDnsChallengeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDomainDnsChallengeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainDnsChallengeResponseBody) GetDomainDnsChallenge() *GetDomainDnsChallengeResponseBodyDomainDnsChallenge {
	return s.DomainDnsChallenge
}

func (s *GetDomainDnsChallengeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDomainDnsChallengeResponseBody) SetDomainDnsChallenge(v *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) *GetDomainDnsChallengeResponseBody {
	s.DomainDnsChallenge = v
	return s
}

func (s *GetDomainDnsChallengeResponseBody) SetRequestId(v string) *GetDomainDnsChallengeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainDnsChallengeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDomainDnsChallengeResponseBodyDomainDnsChallenge struct {
	// The name of the DNS challenge record.
	//
	// example:
	//
	// _idaas-challenge.${domain}
	DnsChallengeName *string `json:"DnsChallengeName,omitempty" xml:"DnsChallengeName,omitempty"`
	// The value of the DNS challenge record.
	//
	// example:
	//
	// exmple123xxx
	DnsChallengeValue *string `json:"DnsChallengeValue,omitempty" xml:"DnsChallengeValue,omitempty"`
	// The type of the DNS challenge record.
	//
	// example:
	//
	// txt
	DnsType *string `json:"DnsType,omitempty" xml:"DnsType,omitempty"`
}

func (s GetDomainDnsChallengeResponseBodyDomainDnsChallenge) String() string {
	return dara.Prettify(s)
}

func (s GetDomainDnsChallengeResponseBodyDomainDnsChallenge) GoString() string {
	return s.String()
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) GetDnsChallengeName() *string {
	return s.DnsChallengeName
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) GetDnsChallengeValue() *string {
	return s.DnsChallengeValue
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) GetDnsType() *string {
	return s.DnsType
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) SetDnsChallengeName(v string) *GetDomainDnsChallengeResponseBodyDomainDnsChallenge {
	s.DnsChallengeName = &v
	return s
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) SetDnsChallengeValue(v string) *GetDomainDnsChallengeResponseBodyDomainDnsChallenge {
	s.DnsChallengeValue = &v
	return s
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) SetDnsType(v string) *GetDomainDnsChallengeResponseBodyDomainDnsChallenge {
	s.DnsType = &v
	return s
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) Validate() error {
	return dara.Validate(s)
}

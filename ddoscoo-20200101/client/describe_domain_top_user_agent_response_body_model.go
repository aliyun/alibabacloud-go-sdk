// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopUserAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainTopUa(v []*DescribeDomainTopUserAgentResponseBodyDomainTopUa) *DescribeDomainTopUserAgentResponseBody
	GetDomainTopUa() []*DescribeDomainTopUserAgentResponseBodyDomainTopUa
	SetRequestId(v string) *DescribeDomainTopUserAgentResponseBody
	GetRequestId() *string
}

type DescribeDomainTopUserAgentResponseBody struct {
	// The information about the user agents.
	DomainTopUa []*DescribeDomainTopUserAgentResponseBodyDomainTopUa `json:"DomainTopUa,omitempty" xml:"DomainTopUa,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainTopUserAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUserAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUserAgentResponseBody) GetDomainTopUa() []*DescribeDomainTopUserAgentResponseBodyDomainTopUa {
	return s.DomainTopUa
}

func (s *DescribeDomainTopUserAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTopUserAgentResponseBody) SetDomainTopUa(v []*DescribeDomainTopUserAgentResponseBodyDomainTopUa) *DescribeDomainTopUserAgentResponseBody {
	s.DomainTopUa = v
	return s
}

func (s *DescribeDomainTopUserAgentResponseBody) SetRequestId(v string) *DescribeDomainTopUserAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopUserAgentResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainTopUserAgentResponseBodyDomainTopUa struct {
	// The domain name of the website.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The page views.
	//
	// example:
	//
	// 22121
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// The Base64-encoded user agent.
	//
	// example:
	//
	// TW96aWxsYS81LjAgKFgxMTsgTGludXggeDg2XzY0KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvOTYuMC40NjY0LjExMCACYWZhcmkvNTM3LjM2
	UserAgent *string `json:"UserAgent,omitempty" xml:"UserAgent,omitempty"`
}

func (s DescribeDomainTopUserAgentResponseBodyDomainTopUa) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUserAgentResponseBodyDomainTopUa) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUserAgentResponseBodyDomainTopUa) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopUserAgentResponseBodyDomainTopUa) GetPv() *int64 {
	return s.Pv
}

func (s *DescribeDomainTopUserAgentResponseBodyDomainTopUa) GetUserAgent() *string {
	return s.UserAgent
}

func (s *DescribeDomainTopUserAgentResponseBodyDomainTopUa) SetDomain(v string) *DescribeDomainTopUserAgentResponseBodyDomainTopUa {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopUserAgentResponseBodyDomainTopUa) SetPv(v int64) *DescribeDomainTopUserAgentResponseBodyDomainTopUa {
	s.Pv = &v
	return s
}

func (s *DescribeDomainTopUserAgentResponseBodyDomainTopUa) SetUserAgent(v string) *DescribeDomainTopUserAgentResponseBodyDomainTopUa {
	s.UserAgent = &v
	return s
}

func (s *DescribeDomainTopUserAgentResponseBodyDomainTopUa) Validate() error {
	return dara.Validate(s)
}

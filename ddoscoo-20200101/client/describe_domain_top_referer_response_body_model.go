// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopRefererResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainTopReferer(v []*DescribeDomainTopRefererResponseBodyDomainTopReferer) *DescribeDomainTopRefererResponseBody
	GetDomainTopReferer() []*DescribeDomainTopRefererResponseBodyDomainTopReferer
	SetRequestId(v string) *DescribeDomainTopRefererResponseBody
	GetRequestId() *string
}

type DescribeDomainTopRefererResponseBody struct {
	// The information about top referers.
	DomainTopReferer []*DescribeDomainTopRefererResponseBodyDomainTopReferer `json:"DomainTopReferer,omitempty" xml:"DomainTopReferer,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainTopRefererResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopRefererResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopRefererResponseBody) GetDomainTopReferer() []*DescribeDomainTopRefererResponseBodyDomainTopReferer {
	return s.DomainTopReferer
}

func (s *DescribeDomainTopRefererResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTopRefererResponseBody) SetDomainTopReferer(v []*DescribeDomainTopRefererResponseBodyDomainTopReferer) *DescribeDomainTopRefererResponseBody {
	s.DomainTopReferer = v
	return s
}

func (s *DescribeDomainTopRefererResponseBody) SetRequestId(v string) *DescribeDomainTopRefererResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopRefererResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainTopRefererResponseBodyDomainTopReferer struct {
	// The domain name of the website.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The page views.
	//
	// example:
	//
	// 257031
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// The Base64-encoded referer.
	//
	// example:
	//
	// aHR0cHM6Ly9zZXJ2aWNld2VjaGF0LmNvbS93eGY3ZDc5YWY0YzU4ZDH3NTEvNC9wYWdlLWZyYW1lLmh0bWw=
	Referer *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
}

func (s DescribeDomainTopRefererResponseBodyDomainTopReferer) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopRefererResponseBodyDomainTopReferer) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopRefererResponseBodyDomainTopReferer) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopRefererResponseBodyDomainTopReferer) GetPv() *int64 {
	return s.Pv
}

func (s *DescribeDomainTopRefererResponseBodyDomainTopReferer) GetReferer() *string {
	return s.Referer
}

func (s *DescribeDomainTopRefererResponseBodyDomainTopReferer) SetDomain(v string) *DescribeDomainTopRefererResponseBodyDomainTopReferer {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopRefererResponseBodyDomainTopReferer) SetPv(v int64) *DescribeDomainTopRefererResponseBodyDomainTopReferer {
	s.Pv = &v
	return s
}

func (s *DescribeDomainTopRefererResponseBodyDomainTopReferer) SetReferer(v string) *DescribeDomainTopRefererResponseBodyDomainTopReferer {
	s.Referer = &v
	return s
}

func (s *DescribeDomainTopRefererResponseBodyDomainTopReferer) Validate() error {
	return dara.Validate(s)
}

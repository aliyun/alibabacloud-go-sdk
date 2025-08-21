// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopHttpMethodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainTopMethod(v []*DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) *DescribeDomainTopHttpMethodResponseBody
	GetDomainTopMethod() []*DescribeDomainTopHttpMethodResponseBodyDomainTopMethod
	SetRequestId(v string) *DescribeDomainTopHttpMethodResponseBody
	GetRequestId() *string
}

type DescribeDomainTopHttpMethodResponseBody struct {
	// The information about top HTTP methods.
	DomainTopMethod []*DescribeDomainTopHttpMethodResponseBodyDomainTopMethod `json:"DomainTopMethod,omitempty" xml:"DomainTopMethod,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainTopHttpMethodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopHttpMethodResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopHttpMethodResponseBody) GetDomainTopMethod() []*DescribeDomainTopHttpMethodResponseBodyDomainTopMethod {
	return s.DomainTopMethod
}

func (s *DescribeDomainTopHttpMethodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTopHttpMethodResponseBody) SetDomainTopMethod(v []*DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) *DescribeDomainTopHttpMethodResponseBody {
	s.DomainTopMethod = v
	return s
}

func (s *DescribeDomainTopHttpMethodResponseBody) SetRequestId(v string) *DescribeDomainTopHttpMethodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopHttpMethodResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainTopHttpMethodResponseBodyDomainTopMethod struct {
	// The domain name of the website.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The HTTP method.
	//
	// example:
	//
	// GET
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The page views.
	//
	// example:
	//
	// 22121
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
}

func (s DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) GetPv() *int64 {
	return s.Pv
}

func (s *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) SetDomain(v string) *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) SetHttpMethod(v string) *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod {
	s.HttpMethod = &v
	return s
}

func (s *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) SetPv(v int64) *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod {
	s.Pv = &v
	return s
}

func (s *DescribeDomainTopHttpMethodResponseBodyDomainTopMethod) Validate() error {
	return dara.Validate(s)
}

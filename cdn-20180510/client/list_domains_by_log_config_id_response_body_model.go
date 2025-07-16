// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsByLogConfigIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *ListDomainsByLogConfigIdResponseBodyDomains) *ListDomainsByLogConfigIdResponseBody
	GetDomains() *ListDomainsByLogConfigIdResponseBodyDomains
	SetRequestId(v string) *ListDomainsByLogConfigIdResponseBody
	GetRequestId() *string
}

type ListDomainsByLogConfigIdResponseBody struct {
	// The domain names.
	Domains *ListDomainsByLogConfigIdResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDomainsByLogConfigIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsByLogConfigIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsByLogConfigIdResponseBody) GetDomains() *ListDomainsByLogConfigIdResponseBodyDomains {
	return s.Domains
}

func (s *ListDomainsByLogConfigIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainsByLogConfigIdResponseBody) SetDomains(v *ListDomainsByLogConfigIdResponseBodyDomains) *ListDomainsByLogConfigIdResponseBody {
	s.Domains = v
	return s
}

func (s *ListDomainsByLogConfigIdResponseBody) SetRequestId(v string) *ListDomainsByLogConfigIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainsByLogConfigIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDomainsByLogConfigIdResponseBodyDomains struct {
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s ListDomainsByLogConfigIdResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsByLogConfigIdResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *ListDomainsByLogConfigIdResponseBodyDomains) GetDomain() []*string {
	return s.Domain
}

func (s *ListDomainsByLogConfigIdResponseBodyDomains) SetDomain(v []*string) *ListDomainsByLogConfigIdResponseBodyDomains {
	s.Domain = v
	return s
}

func (s *ListDomainsByLogConfigIdResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

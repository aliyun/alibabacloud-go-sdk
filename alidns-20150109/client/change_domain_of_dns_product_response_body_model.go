// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainOfDnsProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOriginalDomain(v string) *ChangeDomainOfDnsProductResponseBody
	GetOriginalDomain() *string
	SetRequestId(v string) *ChangeDomainOfDnsProductResponseBody
	GetRequestId() *string
}

type ChangeDomainOfDnsProductResponseBody struct {
	// The domain name that was originally attached. If this parameter is empty, it indicates that this is the first time a domain name is attached to the product.
	//
	// example:
	//
	// www.example.com
	OriginalDomain *string `json:"OriginalDomain,omitempty" xml:"OriginalDomain,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeDomainOfDnsProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainOfDnsProductResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDomainOfDnsProductResponseBody) GetOriginalDomain() *string {
	return s.OriginalDomain
}

func (s *ChangeDomainOfDnsProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeDomainOfDnsProductResponseBody) SetOriginalDomain(v string) *ChangeDomainOfDnsProductResponseBody {
	s.OriginalDomain = &v
	return s
}

func (s *ChangeDomainOfDnsProductResponseBody) SetRequestId(v string) *ChangeDomainOfDnsProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDomainOfDnsProductResponseBody) Validate() error {
	return dara.Validate(s)
}

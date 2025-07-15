// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6TranslatorEntryId(v string) *CreateIPv6TranslatorEntryResponseBody
	GetIpv6TranslatorEntryId() *string
	SetRequestId(v string) *CreateIPv6TranslatorEntryResponseBody
	GetRequestId() *string
}

type CreateIPv6TranslatorEntryResponseBody struct {
	// The ID of the IPv6 Translation Service instance.
	//
	// example:
	//
	// ipv6transentry-xxxxxxxx
	Ipv6TranslatorEntryId *string `json:"Ipv6TranslatorEntryId,omitempty" xml:"Ipv6TranslatorEntryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DCE5D25-FFC9-492A-8371-12A4E0EE2E05
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIPv6TranslatorEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorEntryResponseBody) GetIpv6TranslatorEntryId() *string {
	return s.Ipv6TranslatorEntryId
}

func (s *CreateIPv6TranslatorEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIPv6TranslatorEntryResponseBody) SetIpv6TranslatorEntryId(v string) *CreateIPv6TranslatorEntryResponseBody {
	s.Ipv6TranslatorEntryId = &v
	return s
}

func (s *CreateIPv6TranslatorEntryResponseBody) SetRequestId(v string) *CreateIPv6TranslatorEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIPv6TranslatorEntryResponseBody) Validate() error {
	return dara.Validate(s)
}

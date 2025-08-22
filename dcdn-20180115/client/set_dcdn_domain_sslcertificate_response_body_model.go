// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainSSLCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDcdnDomainSSLCertificateResponseBody
	GetRequestId() *string
}

type SetDcdnDomainSSLCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A7C69682-7F88-40DD-A198-10D0309E439D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnDomainSSLCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainSSLCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSSLCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDcdnDomainSSLCertificateResponseBody) SetRequestId(v string) *SetDcdnDomainSSLCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

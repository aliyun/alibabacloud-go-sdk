// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainSSLCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCdnDomainSSLCertificateResponseBody
	GetRequestId() *string
}

type SetCdnDomainSSLCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A7C69682-7F88-40DD-A198-10D0309E439D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCdnDomainSSLCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainSSLCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetCdnDomainSSLCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCdnDomainSSLCertificateResponseBody) SetRequestId(v string) *SetCdnDomainSSLCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCdnDomainSSLCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

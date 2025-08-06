// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainSSLCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetVodDomainSSLCertificateResponseBody
	GetRequestId() *string
}

type SetVodDomainSSLCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F4C6D5BE-BF13-45*****6C-516EA8906DCD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVodDomainSSLCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainSSLCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetVodDomainSSLCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetVodDomainSSLCertificateResponseBody) SetRequestId(v string) *SetVodDomainSSLCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVodDomainSSLCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

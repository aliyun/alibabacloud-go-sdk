// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainCSRCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCdnDomainCSRCertificateResponseBody
	GetRequestId() *string
}

type SetCdnDomainCSRCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCdnDomainCSRCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainCSRCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetCdnDomainCSRCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCdnDomainCSRCertificateResponseBody) SetRequestId(v string) *SetCdnDomainCSRCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCdnDomainCSRCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainCSRCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDcdnDomainCSRCertificateResponseBody
	GetRequestId() *string
}

type SetDcdnDomainCSRCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnDomainCSRCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainCSRCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCSRCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDcdnDomainCSRCertificateResponseBody) SetRequestId(v string) *SetDcdnDomainCSRCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDcdnDomainCSRCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainSMCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDcdnDomainSMCertificateResponseBody
	GetRequestId() *string
}

type SetDcdnDomainSMCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnDomainSMCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainSMCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSMCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDcdnDomainSMCertificateResponseBody) SetRequestId(v string) *SetDcdnDomainSMCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDcdnDomainSMCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

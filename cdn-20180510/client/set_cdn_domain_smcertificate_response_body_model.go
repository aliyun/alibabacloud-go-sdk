// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainSMCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCdnDomainSMCertificateResponseBody
	GetRequestId() *string
}

type SetCdnDomainSMCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCdnDomainSMCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainSMCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetCdnDomainSMCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCdnDomainSMCertificateResponseBody) SetRequestId(v string) *SetCdnDomainSMCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCdnDomainSMCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

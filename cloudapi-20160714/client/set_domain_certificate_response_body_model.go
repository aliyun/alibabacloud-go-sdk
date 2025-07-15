// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDomainCertificateResponseBody
	GetRequestId() *string
}

type SetDomainCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDomainCertificateResponseBody) SetRequestId(v string) *SetDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDomainCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

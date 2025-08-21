// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVsDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetVsDomainCertificateResponseBody
	GetRequestId() *string
}

type SetVsDomainCertificateResponseBody struct {
	// example:
	//
	// 119F7639-4646-51A4-B6C1-300D391C0104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVsDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetVsDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetVsDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetVsDomainCertificateResponseBody) SetRequestId(v string) *SetVsDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVsDomainCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

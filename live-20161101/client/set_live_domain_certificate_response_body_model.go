// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveDomainCertificateResponseBody
	GetRequestId() *string
}

type SetLiveDomainCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveDomainCertificateResponseBody) SetRequestId(v string) *SetLiveDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveDomainCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

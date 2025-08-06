// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetVodDomainCertificateResponseBody
	GetRequestId() *string
}

type SetVodDomainCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-****-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVodDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetVodDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetVodDomainCertificateResponseBody) SetRequestId(v string) *SetVodDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVodDomainCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

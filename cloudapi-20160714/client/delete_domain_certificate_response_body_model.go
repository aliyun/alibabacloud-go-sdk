// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDomainCertificateResponseBody
	GetRequestId() *string
}

type DeleteDomainCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE5722A6-AE78-4741-A9B0-6C817D360510
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainCertificateResponseBody) SetRequestId(v string) *DeleteDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

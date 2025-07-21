// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCertificateResponseBody
	GetRequestId() *string
}

type DeleteCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// d97f6c33-ca26-4de2-a580-0e2fd1c5bfb0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCertificateResponseBody) SetRequestId(v string) *DeleteCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

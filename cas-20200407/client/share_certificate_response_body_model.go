// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ShareCertificateResponseBody
	GetRequestId() *string
}

type ShareCertificateResponseBody struct {
	// example:
	//
	// A79D2C05-4B0B-57CA-873D-4FA985F2D26E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ShareCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ShareCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ShareCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShareCertificateResponseBody) SetRequestId(v string) *ShareCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShareCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

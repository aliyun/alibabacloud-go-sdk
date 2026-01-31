// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPendingCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelPendingCertificateResponseBody
	GetRequestId() *string
}

type CancelPendingCertificateResponseBody struct {
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelPendingCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelPendingCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPendingCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelPendingCertificateResponseBody) SetRequestId(v string) *CancelPendingCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPendingCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

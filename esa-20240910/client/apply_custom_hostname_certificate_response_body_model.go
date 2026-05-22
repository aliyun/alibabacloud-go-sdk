// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCustomHostnameCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyCustomHostnameCertificateResponseBody
	GetRequestId() *string
}

type ApplyCustomHostnameCertificateResponseBody struct {
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyCustomHostnameCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyCustomHostnameCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCustomHostnameCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyCustomHostnameCertificateResponseBody) SetRequestId(v string) *ApplyCustomHostnameCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCustomHostnameCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

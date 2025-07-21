// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCertificateStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCertificateStatusResponseBody
	GetRequestId() *string
}

type UpdateCertificateStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// e3f57fe0-9ded-40b0-9caf-a3815f2148c1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCertificateStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCertificateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCertificateStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCertificateStatusResponseBody) SetRequestId(v string) *UpdateCertificateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCertificateStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

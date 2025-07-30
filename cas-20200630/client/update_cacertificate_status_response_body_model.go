// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCACertificateStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCACertificateStatusResponseBody
	GetRequestId() *string
}

type UpdateCACertificateStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCACertificateStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCACertificateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCACertificateStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCACertificateStatusResponseBody) SetRequestId(v string) *UpdateCACertificateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCACertificateStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

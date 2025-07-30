// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRevokeClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRevokeClientCertificateResponseBody
	GetRequestId() *string
}

type CreateRevokeClientCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRevokeClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRevokeClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRevokeClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRevokeClientCertificateResponseBody) SetRequestId(v string) *CreateRevokeClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRevokeClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

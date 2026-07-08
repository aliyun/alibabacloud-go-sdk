// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeWHClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeWHClientCertificateResponseBody
	GetRequestId() *string
}

type RevokeWHClientCertificateResponseBody struct {
	// The unique identifier for the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeWHClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeWHClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeWHClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeWHClientCertificateResponseBody) SetRequestId(v string) *RevokeWHClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeWHClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

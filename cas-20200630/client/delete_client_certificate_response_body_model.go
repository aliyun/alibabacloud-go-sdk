// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClientCertificateResponseBody
	GetRequestId() *string
}

type DeleteClientCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientCertificateResponseBody) SetRequestId(v string) *DeleteClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

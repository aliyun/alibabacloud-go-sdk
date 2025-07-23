// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserCertificateResponseBody
	GetRequestId() *string
}

type DeleteUserCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3E50D480-9765-5CFD-9650-9BACCECA5135
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserCertificateResponseBody) SetRequestId(v string) *DeleteUserCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCertificateRequestResponseBody
	GetRequestId() *string
}

type DeleteCertificateRequestResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCertificateRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateRequestResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCertificateRequestResponseBody) SetRequestId(v string) *DeleteCertificateRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCertificateRequestResponseBody) Validate() error {
	return dara.Validate(s)
}

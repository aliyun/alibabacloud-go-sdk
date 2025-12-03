// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServerCertificateResponseBody
	GetRequestId() *string
}

type DeleteServerCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServerCertificateResponseBody) SetRequestId(v string) *DeleteServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServerCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

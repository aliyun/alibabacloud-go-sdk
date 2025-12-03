// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCACertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCACertificateResponseBody
	GetRequestId() *string
}

type DeleteCACertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCACertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCACertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCACertificateResponseBody) SetRequestId(v string) *DeleteCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCACertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

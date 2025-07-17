// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCertificateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCertificateResponseBody
	GetSuccess() *bool
}

type DeleteCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D9A61DC0-B922-421B-B706
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCertificateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCertificateResponseBody) SetRequestId(v string) *DeleteCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCertificateResponseBody) SetSuccess(v bool) *DeleteCertificateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeCertificateResponseBody
	GetRequestId() *string
}

type RevokeCertificateResponseBody struct {
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeCertificateResponseBody) SetRequestId(v string) *RevokeCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

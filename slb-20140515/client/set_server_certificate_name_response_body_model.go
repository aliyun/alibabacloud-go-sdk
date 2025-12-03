// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetServerCertificateNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetServerCertificateNameResponseBody
	GetRequestId() *string
}

type SetServerCertificateNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FE7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetServerCertificateNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetServerCertificateNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetServerCertificateNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetServerCertificateNameResponseBody) SetRequestId(v string) *SetServerCertificateNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetServerCertificateNameResponseBody) Validate() error {
	return dara.Validate(s)
}

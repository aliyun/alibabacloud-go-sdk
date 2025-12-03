// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCACertificateNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCACertificateNameResponseBody
	GetRequestId() *string
}

type SetCACertificateNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FE7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCACertificateNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCACertificateNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetCACertificateNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCACertificateNameResponseBody) SetRequestId(v string) *SetCACertificateNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCACertificateNameResponseBody) Validate() error {
	return dara.Validate(s)
}

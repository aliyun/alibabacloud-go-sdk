// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyInstanceLDAPAuthServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyInstanceLDAPAuthServerResponseBody
	GetRequestId() *string
}

type VerifyInstanceLDAPAuthServerResponseBody struct {
	// example:
	//
	// C9E97677-BD74-584B-AFCE-948C2A70BB83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyInstanceLDAPAuthServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyInstanceLDAPAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyInstanceLDAPAuthServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyInstanceLDAPAuthServerResponseBody) SetRequestId(v string) *VerifyInstanceLDAPAuthServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerResponseBody) Validate() error {
	return dara.Validate(s)
}

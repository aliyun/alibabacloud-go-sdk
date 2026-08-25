// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMFAAuthenticationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetMFAAuthenticationStatusResponseBody
	GetRequestId() *string
}

type SetMFAAuthenticationStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 14E2B1A9-7713-5E6F-8409-8DE12DF51AF4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetMFAAuthenticationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetMFAAuthenticationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetMFAAuthenticationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetMFAAuthenticationStatusResponseBody) SetRequestId(v string) *SetMFAAuthenticationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMFAAuthenticationStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAuthenticationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyAuthenticationResponseBody
	GetRequestId() *string
}

type VerifyAuthenticationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyAuthenticationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyAuthenticationResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyAuthenticationResponseBody) SetRequestId(v string) *VerifyAuthenticationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyAuthenticationResponseBody) Validate() error {
	return dara.Validate(s)
}

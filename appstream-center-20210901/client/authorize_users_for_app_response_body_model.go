// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeUsersForAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeUsersForAppResponseBody
	GetRequestId() *string
}

type AuthorizeUsersForAppResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeUsersForAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeUsersForAppResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeUsersForAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeUsersForAppResponseBody) SetRequestId(v string) *AuthorizeUsersForAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeUsersForAppResponseBody) Validate() error {
	return dara.Validate(s)
}

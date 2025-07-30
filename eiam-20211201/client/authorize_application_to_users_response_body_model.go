// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeApplicationToUsersResponseBody
	GetRequestId() *string
}

type AuthorizeApplicationToUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeApplicationToUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeApplicationToUsersResponseBody) SetRequestId(v string) *AuthorizeApplicationToUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeApplicationToUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

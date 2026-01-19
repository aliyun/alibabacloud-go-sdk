// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeResourceServerScopesToUserResponseBody
	GetRequestId() *string
}

type AuthorizeResourceServerScopesToUserResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeResourceServerScopesToUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeResourceServerScopesToUserResponseBody) SetRequestId(v string) *AuthorizeResourceServerScopesToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToUserResponseBody) Validate() error {
	return dara.Validate(s)
}

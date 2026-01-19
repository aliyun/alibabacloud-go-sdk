// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerToClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeResourceServerToClientResponseBody
	GetRequestId() *string
}

type AuthorizeResourceServerToClientResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeResourceServerToClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerToClientResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerToClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeResourceServerToClientResponseBody) SetRequestId(v string) *AuthorizeResourceServerToClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeResourceServerToClientResponseBody) Validate() error {
	return dara.Validate(s)
}

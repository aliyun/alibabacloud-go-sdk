// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeResourceServerScopesToGroupResponseBody
	GetRequestId() *string
}

type AuthorizeResourceServerScopesToGroupResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeResourceServerScopesToGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeResourceServerScopesToGroupResponseBody) SetRequestId(v string) *AuthorizeResourceServerScopesToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

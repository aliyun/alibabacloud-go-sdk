// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeApplicationToGroupsResponseBody
	GetRequestId() *string
}

type AuthorizeApplicationToGroupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeApplicationToGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeApplicationToGroupsResponseBody) SetRequestId(v string) *AuthorizeApplicationToGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeApplicationToGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

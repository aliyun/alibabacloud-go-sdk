// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuthorizeRoleResponseBody
	GetCode() *int32
	SetMessage(v string) *AuthorizeRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthorizeRoleResponseBody
	GetRequestId() *string
}

type AuthorizeRoleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57609587-DFA2-*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeRoleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuthorizeRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthorizeRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeRoleResponseBody) SetCode(v int32) *AuthorizeRoleResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeRoleResponseBody) SetMessage(v string) *AuthorizeRoleResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeRoleResponseBody) SetRequestId(v string) *AuthorizeRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeRoleResponseBody) Validate() error {
	return dara.Validate(s)
}

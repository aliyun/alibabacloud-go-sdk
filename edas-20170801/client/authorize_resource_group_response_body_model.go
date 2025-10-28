// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuthorizeResourceGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *AuthorizeResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthorizeResourceGroupResponseBody
	GetRequestId() *string
}

type AuthorizeResourceGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57609587-DFA2-**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuthorizeResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthorizeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeResourceGroupResponseBody) SetCode(v int32) *AuthorizeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeResourceGroupResponseBody) SetMessage(v string) *AuthorizeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeResourceGroupResponseBody) SetRequestId(v string) *AuthorizeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

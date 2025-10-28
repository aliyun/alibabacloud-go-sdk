// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuthorizeApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *AuthorizeApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthorizeApplicationResponseBody
	GetRequestId() *string
}

type AuthorizeApplicationResponseBody struct {
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
	// E3DA95D3-**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuthorizeApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthorizeApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeApplicationResponseBody) SetCode(v int32) *AuthorizeApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeApplicationResponseBody) SetMessage(v string) *AuthorizeApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeApplicationResponseBody) SetRequestId(v string) *AuthorizeApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

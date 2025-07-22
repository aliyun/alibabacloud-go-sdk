// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenResponseBody
	GetCode() *string
	SetMessage(v string) *GetTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTokenResponseBody
	GetSuccess() *bool
	SetToken(v string) *GetTokenResponseBody
	GetToken() *string
}

type GetTokenResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: All errors, except for parameter validation errors, are internal errors.
	//
	// 	- ValidationError: A parameter validation error.
	//
	// example:
	//
	// ValidationError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The temporary authentication information of the DSW instance.
	//
	// example:
	//
	// *******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetTokenResponseBody) SetCode(v string) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v bool) *GetTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenResponseBody) SetToken(v string) *GetTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GetTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

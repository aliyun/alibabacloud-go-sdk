// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterAccountResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RegisterAccountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RegisterAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *RegisterAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RegisterAccountResponseBody
	GetSuccess() *bool
}

type RegisterAccountResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterAccountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RegisterAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RegisterAccountResponseBody) SetCode(v string) *RegisterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterAccountResponseBody) SetHttpStatusCode(v int32) *RegisterAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RegisterAccountResponseBody) SetMessage(v string) *RegisterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterAccountResponseBody) SetRequestId(v string) *RegisterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterAccountResponseBody) SetSuccess(v bool) *RegisterAccountResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

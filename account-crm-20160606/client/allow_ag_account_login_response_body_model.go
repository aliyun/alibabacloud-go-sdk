// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllowAgAccountLoginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AllowAgAccountLoginResponseBody
	GetCode() *string
	SetData(v bool) *AllowAgAccountLoginResponseBody
	GetData() *bool
	SetMessage(v string) *AllowAgAccountLoginResponseBody
	GetMessage() *string
	SetRequestId(v string) *AllowAgAccountLoginResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AllowAgAccountLoginResponseBody
	GetSuccess() *bool
}

type AllowAgAccountLoginResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AllowAgAccountLoginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllowAgAccountLoginResponseBody) GoString() string {
	return s.String()
}

func (s *AllowAgAccountLoginResponseBody) GetCode() *string {
	return s.Code
}

func (s *AllowAgAccountLoginResponseBody) GetData() *bool {
	return s.Data
}

func (s *AllowAgAccountLoginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AllowAgAccountLoginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllowAgAccountLoginResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AllowAgAccountLoginResponseBody) SetCode(v string) *AllowAgAccountLoginResponseBody {
	s.Code = &v
	return s
}

func (s *AllowAgAccountLoginResponseBody) SetData(v bool) *AllowAgAccountLoginResponseBody {
	s.Data = &v
	return s
}

func (s *AllowAgAccountLoginResponseBody) SetMessage(v string) *AllowAgAccountLoginResponseBody {
	s.Message = &v
	return s
}

func (s *AllowAgAccountLoginResponseBody) SetRequestId(v string) *AllowAgAccountLoginResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllowAgAccountLoginResponseBody) SetSuccess(v bool) *AllowAgAccountLoginResponseBody {
	s.Success = &v
	return s
}

func (s *AllowAgAccountLoginResponseBody) Validate() error {
	return dara.Validate(s)
}

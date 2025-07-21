// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTerminalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddTerminalResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddTerminalResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddTerminalResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddTerminalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTerminalResponseBody
	GetSuccess() *bool
}

type AddTerminalResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTerminalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *AddTerminalResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddTerminalResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddTerminalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddTerminalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTerminalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTerminalResponseBody) SetCode(v string) *AddTerminalResponseBody {
	s.Code = &v
	return s
}

func (s *AddTerminalResponseBody) SetHttpStatusCode(v int32) *AddTerminalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddTerminalResponseBody) SetMessage(v string) *AddTerminalResponseBody {
	s.Message = &v
	return s
}

func (s *AddTerminalResponseBody) SetRequestId(v string) *AddTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTerminalResponseBody) SetSuccess(v bool) *AddTerminalResponseBody {
	s.Success = &v
	return s
}

func (s *AddTerminalResponseBody) Validate() error {
	return dara.Validate(s)
}

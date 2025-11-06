// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UnbindResponseBody
	GetCode() *int32
	SetData(v bool) *UnbindResponseBody
	GetData() *bool
	SetMessage(v string) *UnbindResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindResponseBody
	GetSuccess() *bool
}

type UnbindResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnbindResponseBody) GetData() *bool {
	return s.Data
}

func (s *UnbindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindResponseBody) SetCode(v int32) *UnbindResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindResponseBody) SetData(v bool) *UnbindResponseBody {
	s.Data = &v
	return s
}

func (s *UnbindResponseBody) SetMessage(v string) *UnbindResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindResponseBody) SetRequestId(v string) *UnbindResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindResponseBody) SetSuccess(v bool) *UnbindResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindResponseBody) Validate() error {
	return dara.Validate(s)
}

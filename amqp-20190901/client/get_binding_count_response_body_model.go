// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetBindingCountResponseBody
	GetCode() *int32
	SetData(v int32) *GetBindingCountResponseBody
	GetData() *int32
	SetMessage(v string) *GetBindingCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBindingCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBindingCountResponseBody
	GetSuccess() *bool
}

type GetBindingCountResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBindingCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBindingCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetBindingCountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetBindingCountResponseBody) GetData() *int32 {
	return s.Data
}

func (s *GetBindingCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBindingCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBindingCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBindingCountResponseBody) SetCode(v int32) *GetBindingCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetBindingCountResponseBody) SetData(v int32) *GetBindingCountResponseBody {
	s.Data = &v
	return s
}

func (s *GetBindingCountResponseBody) SetMessage(v string) *GetBindingCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetBindingCountResponseBody) SetRequestId(v string) *GetBindingCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBindingCountResponseBody) SetSuccess(v bool) *GetBindingCountResponseBody {
	s.Success = &v
	return s
}

func (s *GetBindingCountResponseBody) Validate() error {
	return dara.Validate(s)
}

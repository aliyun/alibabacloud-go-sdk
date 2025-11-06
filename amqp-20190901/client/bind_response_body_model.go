// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BindResponseBody
	GetCode() *int32
	SetData(v bool) *BindResponseBody
	GetData() *bool
	SetMessage(v string) *BindResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindResponseBody
	GetSuccess() *bool
}

type BindResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindResponseBody) GoString() string {
	return s.String()
}

func (s *BindResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BindResponseBody) GetData() *bool {
	return s.Data
}

func (s *BindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindResponseBody) SetCode(v int32) *BindResponseBody {
	s.Code = &v
	return s
}

func (s *BindResponseBody) SetData(v bool) *BindResponseBody {
	s.Data = &v
	return s
}

func (s *BindResponseBody) SetMessage(v string) *BindResponseBody {
	s.Message = &v
	return s
}

func (s *BindResponseBody) SetRequestId(v string) *BindResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindResponseBody) SetSuccess(v bool) *BindResponseBody {
	s.Success = &v
	return s
}

func (s *BindResponseBody) Validate() error {
	return dara.Validate(s)
}

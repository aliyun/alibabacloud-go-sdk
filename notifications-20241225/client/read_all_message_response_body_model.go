// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadAllMessageResponseBody
	GetCode() *string
	SetData(v bool) *ReadAllMessageResponseBody
	GetData() *bool
	SetMessage(v string) *ReadAllMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadAllMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadAllMessageResponseBody
	GetSuccess() *bool
}

type ReadAllMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadAllMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadAllMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadAllMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadAllMessageResponseBody) GetData() *bool {
	return s.Data
}

func (s *ReadAllMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadAllMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadAllMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadAllMessageResponseBody) SetCode(v string) *ReadAllMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ReadAllMessageResponseBody) SetData(v bool) *ReadAllMessageResponseBody {
	s.Data = &v
	return s
}

func (s *ReadAllMessageResponseBody) SetMessage(v string) *ReadAllMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ReadAllMessageResponseBody) SetRequestId(v string) *ReadAllMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadAllMessageResponseBody) SetSuccess(v bool) *ReadAllMessageResponseBody {
	s.Success = &v
	return s
}

func (s *ReadAllMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

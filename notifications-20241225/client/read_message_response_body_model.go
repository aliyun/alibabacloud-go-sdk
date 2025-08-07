// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMessageResponseBody
	GetCode() *string
	SetData(v bool) *ReadMessageResponseBody
	GetData() *bool
	SetMessage(v string) *ReadMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMessageResponseBody
	GetSuccess() *bool
}

type ReadMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMessageResponseBody) GetData() *bool {
	return s.Data
}

func (s *ReadMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMessageResponseBody) SetCode(v string) *ReadMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageResponseBody) SetData(v bool) *ReadMessageResponseBody {
	s.Data = &v
	return s
}

func (s *ReadMessageResponseBody) SetMessage(v string) *ReadMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageResponseBody) SetRequestId(v string) *ReadMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageResponseBody) SetSuccess(v bool) *ReadMessageResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

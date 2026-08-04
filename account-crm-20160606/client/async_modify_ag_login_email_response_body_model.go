// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncModifyAgLoginEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AsyncModifyAgLoginEmailResponseBody
	GetCode() *string
	SetMessage(v string) *AsyncModifyAgLoginEmailResponseBody
	GetMessage() *string
	SetRequestId(v string) *AsyncModifyAgLoginEmailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AsyncModifyAgLoginEmailResponseBody
	GetSuccess() *bool
	SetTraceNo(v string) *AsyncModifyAgLoginEmailResponseBody
	GetTraceNo() *string
}

type AsyncModifyAgLoginEmailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceNo   *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s AsyncModifyAgLoginEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsyncModifyAgLoginEmailResponseBody) GoString() string {
	return s.String()
}

func (s *AsyncModifyAgLoginEmailResponseBody) GetCode() *string {
	return s.Code
}

func (s *AsyncModifyAgLoginEmailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AsyncModifyAgLoginEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncModifyAgLoginEmailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncModifyAgLoginEmailResponseBody) GetTraceNo() *string {
	return s.TraceNo
}

func (s *AsyncModifyAgLoginEmailResponseBody) SetCode(v string) *AsyncModifyAgLoginEmailResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncModifyAgLoginEmailResponseBody) SetMessage(v string) *AsyncModifyAgLoginEmailResponseBody {
	s.Message = &v
	return s
}

func (s *AsyncModifyAgLoginEmailResponseBody) SetRequestId(v string) *AsyncModifyAgLoginEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncModifyAgLoginEmailResponseBody) SetSuccess(v bool) *AsyncModifyAgLoginEmailResponseBody {
	s.Success = &v
	return s
}

func (s *AsyncModifyAgLoginEmailResponseBody) SetTraceNo(v string) *AsyncModifyAgLoginEmailResponseBody {
	s.TraceNo = &v
	return s
}

func (s *AsyncModifyAgLoginEmailResponseBody) Validate() error {
	return dara.Validate(s)
}

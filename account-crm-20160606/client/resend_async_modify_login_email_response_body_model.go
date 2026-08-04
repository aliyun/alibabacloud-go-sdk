// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendAsyncModifyLoginEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResendAsyncModifyLoginEmailResponseBody
	GetCode() *string
	SetMessage(v string) *ResendAsyncModifyLoginEmailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResendAsyncModifyLoginEmailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResendAsyncModifyLoginEmailResponseBody
	GetSuccess() *bool
	SetTraceNo(v string) *ResendAsyncModifyLoginEmailResponseBody
	GetTraceNo() *string
}

type ResendAsyncModifyLoginEmailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceNo   *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s ResendAsyncModifyLoginEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResendAsyncModifyLoginEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendAsyncModifyLoginEmailResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResendAsyncModifyLoginEmailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResendAsyncModifyLoginEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResendAsyncModifyLoginEmailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResendAsyncModifyLoginEmailResponseBody) GetTraceNo() *string {
	return s.TraceNo
}

func (s *ResendAsyncModifyLoginEmailResponseBody) SetCode(v string) *ResendAsyncModifyLoginEmailResponseBody {
	s.Code = &v
	return s
}

func (s *ResendAsyncModifyLoginEmailResponseBody) SetMessage(v string) *ResendAsyncModifyLoginEmailResponseBody {
	s.Message = &v
	return s
}

func (s *ResendAsyncModifyLoginEmailResponseBody) SetRequestId(v string) *ResendAsyncModifyLoginEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendAsyncModifyLoginEmailResponseBody) SetSuccess(v bool) *ResendAsyncModifyLoginEmailResponseBody {
	s.Success = &v
	return s
}

func (s *ResendAsyncModifyLoginEmailResponseBody) SetTraceNo(v string) *ResendAsyncModifyLoginEmailResponseBody {
	s.TraceNo = &v
	return s
}

func (s *ResendAsyncModifyLoginEmailResponseBody) Validate() error {
	return dara.Validate(s)
}

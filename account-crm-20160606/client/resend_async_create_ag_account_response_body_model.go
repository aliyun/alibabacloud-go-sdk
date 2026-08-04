// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendAsyncCreateAgAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResendAsyncCreateAgAccountResponseBody
	GetCode() *string
	SetMessage(v string) *ResendAsyncCreateAgAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResendAsyncCreateAgAccountResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ResendAsyncCreateAgAccountResponseBody
	GetSuccess() *string
	SetTraceNo(v string) *ResendAsyncCreateAgAccountResponseBody
	GetTraceNo() *string
}

type ResendAsyncCreateAgAccountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceNo   *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s ResendAsyncCreateAgAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResendAsyncCreateAgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ResendAsyncCreateAgAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResendAsyncCreateAgAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResendAsyncCreateAgAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResendAsyncCreateAgAccountResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ResendAsyncCreateAgAccountResponseBody) GetTraceNo() *string {
	return s.TraceNo
}

func (s *ResendAsyncCreateAgAccountResponseBody) SetCode(v string) *ResendAsyncCreateAgAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ResendAsyncCreateAgAccountResponseBody) SetMessage(v string) *ResendAsyncCreateAgAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ResendAsyncCreateAgAccountResponseBody) SetRequestId(v string) *ResendAsyncCreateAgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendAsyncCreateAgAccountResponseBody) SetSuccess(v string) *ResendAsyncCreateAgAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ResendAsyncCreateAgAccountResponseBody) SetTraceNo(v string) *ResendAsyncCreateAgAccountResponseBody {
	s.TraceNo = &v
	return s
}

func (s *ResendAsyncCreateAgAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

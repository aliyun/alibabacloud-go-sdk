// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateAgAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AsyncCreateAgAccountResponseBody
	GetCode() *string
	SetMessage(v string) *AsyncCreateAgAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *AsyncCreateAgAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AsyncCreateAgAccountResponseBody
	GetSuccess() *bool
	SetTraceNo(v string) *AsyncCreateAgAccountResponseBody
	GetTraceNo() *string
}

type AsyncCreateAgAccountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceNo   *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s AsyncCreateAgAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateAgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AsyncCreateAgAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *AsyncCreateAgAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AsyncCreateAgAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncCreateAgAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncCreateAgAccountResponseBody) GetTraceNo() *string {
	return s.TraceNo
}

func (s *AsyncCreateAgAccountResponseBody) SetCode(v string) *AsyncCreateAgAccountResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncCreateAgAccountResponseBody) SetMessage(v string) *AsyncCreateAgAccountResponseBody {
	s.Message = &v
	return s
}

func (s *AsyncCreateAgAccountResponseBody) SetRequestId(v string) *AsyncCreateAgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncCreateAgAccountResponseBody) SetSuccess(v bool) *AsyncCreateAgAccountResponseBody {
	s.Success = &v
	return s
}

func (s *AsyncCreateAgAccountResponseBody) SetTraceNo(v string) *AsyncCreateAgAccountResponseBody {
	s.TraceNo = &v
	return s
}

func (s *AsyncCreateAgAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAgOneKeyDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyAgOneKeyDeleteTaskResponseBody
	GetCode() *string
	SetData(v string) *ApplyAgOneKeyDeleteTaskResponseBody
	GetData() *string
	SetMessage(v string) *ApplyAgOneKeyDeleteTaskResponseBody
	GetMessage() *string
	SetNeedAbandonSpAfterPay(v bool) *ApplyAgOneKeyDeleteTaskResponseBody
	GetNeedAbandonSpAfterPay() *bool
	SetRequestId(v string) *ApplyAgOneKeyDeleteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyAgOneKeyDeleteTaskResponseBody
	GetSuccess() *bool
}

type ApplyAgOneKeyDeleteTaskResponseBody struct {
	Code                  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data                  *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message               *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NeedAbandonSpAfterPay *bool   `json:"NeedAbandonSpAfterPay,omitempty" xml:"NeedAbandonSpAfterPay,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success               *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyAgOneKeyDeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyAgOneKeyDeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) GetNeedAbandonSpAfterPay() *bool {
	return s.NeedAbandonSpAfterPay
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) SetCode(v string) *ApplyAgOneKeyDeleteTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) SetData(v string) *ApplyAgOneKeyDeleteTaskResponseBody {
	s.Data = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) SetMessage(v string) *ApplyAgOneKeyDeleteTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) SetNeedAbandonSpAfterPay(v bool) *ApplyAgOneKeyDeleteTaskResponseBody {
	s.NeedAbandonSpAfterPay = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) SetRequestId(v string) *ApplyAgOneKeyDeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) SetSuccess(v bool) *ApplyAgOneKeyDeleteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillConfirmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MonthBillConfirmResponseBody
	GetCode() *string
	SetMessage(v string) *MonthBillConfirmResponseBody
	GetMessage() *string
	SetRequestId(v string) *MonthBillConfirmResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MonthBillConfirmResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MonthBillConfirmResponseBody
	GetTraceId() *string
}

type MonthBillConfirmResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MonthBillConfirmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MonthBillConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *MonthBillConfirmResponseBody) GetCode() *string {
	return s.Code
}

func (s *MonthBillConfirmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MonthBillConfirmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MonthBillConfirmResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MonthBillConfirmResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MonthBillConfirmResponseBody) SetCode(v string) *MonthBillConfirmResponseBody {
	s.Code = &v
	return s
}

func (s *MonthBillConfirmResponseBody) SetMessage(v string) *MonthBillConfirmResponseBody {
	s.Message = &v
	return s
}

func (s *MonthBillConfirmResponseBody) SetRequestId(v string) *MonthBillConfirmResponseBody {
	s.RequestId = &v
	return s
}

func (s *MonthBillConfirmResponseBody) SetSuccess(v bool) *MonthBillConfirmResponseBody {
	s.Success = &v
	return s
}

func (s *MonthBillConfirmResponseBody) SetTraceId(v string) *MonthBillConfirmResponseBody {
	s.TraceId = &v
	return s
}

func (s *MonthBillConfirmResponseBody) Validate() error {
	return dara.Validate(s)
}

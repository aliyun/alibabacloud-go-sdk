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
	// example:
	//
	// success
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
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

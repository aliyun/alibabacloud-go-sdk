// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderPayCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightOrderPayCheckResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightOrderPayCheckResponseBody
	GetMessage() *string
	SetRequestId(v string) *IntlFlightOrderPayCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightOrderPayCheckResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightOrderPayCheckResponseBody
	GetTraceId() *string
}

type IntlFlightOrderPayCheckResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightOrderPayCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayCheckResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayCheckResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightOrderPayCheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightOrderPayCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightOrderPayCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightOrderPayCheckResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightOrderPayCheckResponseBody) SetCode(v string) *IntlFlightOrderPayCheckResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightOrderPayCheckResponseBody) SetMessage(v string) *IntlFlightOrderPayCheckResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightOrderPayCheckResponseBody) SetRequestId(v string) *IntlFlightOrderPayCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightOrderPayCheckResponseBody) SetSuccess(v bool) *IntlFlightOrderPayCheckResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightOrderPayCheckResponseBody) SetTraceId(v string) *IntlFlightOrderPayCheckResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightOrderPayCheckResponseBody) Validate() error {
	return dara.Validate(s)
}

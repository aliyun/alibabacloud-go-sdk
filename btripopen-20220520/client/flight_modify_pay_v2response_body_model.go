// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyPayV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightModifyPayV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightModifyPayV2ResponseBody
	GetMessage() *string
	SetModule(v map[string]interface{}) *FlightModifyPayV2ResponseBody
	GetModule() map[string]interface{}
	SetRequestId(v string) *FlightModifyPayV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightModifyPayV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightModifyPayV2ResponseBody
	GetTraceId() *string
}

type FlightModifyPayV2ResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	Module    map[string]interface{} `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                  `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightModifyPayV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyPayV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightModifyPayV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightModifyPayV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightModifyPayV2ResponseBody) GetModule() map[string]interface{} {
	return s.Module
}

func (s *FlightModifyPayV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightModifyPayV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightModifyPayV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightModifyPayV2ResponseBody) SetCode(v string) *FlightModifyPayV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightModifyPayV2ResponseBody) SetMessage(v string) *FlightModifyPayV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightModifyPayV2ResponseBody) SetModule(v map[string]interface{}) *FlightModifyPayV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightModifyPayV2ResponseBody) SetRequestId(v string) *FlightModifyPayV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightModifyPayV2ResponseBody) SetSuccess(v bool) *FlightModifyPayV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightModifyPayV2ResponseBody) SetTraceId(v string) *FlightModifyPayV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightModifyPayV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

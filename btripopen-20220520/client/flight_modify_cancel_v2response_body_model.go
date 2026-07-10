// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyCancelV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightModifyCancelV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightModifyCancelV2ResponseBody
	GetMessage() *string
	SetModule(v map[string]interface{}) *FlightModifyCancelV2ResponseBody
	GetModule() map[string]interface{}
	SetRequestId(v string) *FlightModifyCancelV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightModifyCancelV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightModifyCancelV2ResponseBody
	GetTraceId() *string
}

type FlightModifyCancelV2ResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	Module    map[string]interface{} `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                  `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightModifyCancelV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyCancelV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightModifyCancelV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightModifyCancelV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightModifyCancelV2ResponseBody) GetModule() map[string]interface{} {
	return s.Module
}

func (s *FlightModifyCancelV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightModifyCancelV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightModifyCancelV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightModifyCancelV2ResponseBody) SetCode(v string) *FlightModifyCancelV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightModifyCancelV2ResponseBody) SetMessage(v string) *FlightModifyCancelV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightModifyCancelV2ResponseBody) SetModule(v map[string]interface{}) *FlightModifyCancelV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightModifyCancelV2ResponseBody) SetRequestId(v string) *FlightModifyCancelV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightModifyCancelV2ResponseBody) SetSuccess(v bool) *FlightModifyCancelV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightModifyCancelV2ResponseBody) SetTraceId(v string) *FlightModifyCancelV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightModifyCancelV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

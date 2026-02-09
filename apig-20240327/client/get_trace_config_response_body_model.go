// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTraceConfigResponseBody
	GetCode() *int32
	SetData(v *GetTraceConfigResponseBodyData) *GetTraceConfigResponseBody
	GetData() *GetTraceConfigResponseBodyData
	SetMessage(v string) *GetTraceConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTraceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTraceConfigResponseBody
	GetSuccess() *bool
}

type GetTraceConfigResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetTraceConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32AF2C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTraceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTraceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTraceConfigResponseBody) GetData() *GetTraceConfigResponseBodyData {
	return s.Data
}

func (s *GetTraceConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTraceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTraceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTraceConfigResponseBody) SetCode(v int32) *GetTraceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetTraceConfigResponseBody) SetData(v *GetTraceConfigResponseBodyData) *GetTraceConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetTraceConfigResponseBody) SetMessage(v string) *GetTraceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetTraceConfigResponseBody) SetRequestId(v string) *GetTraceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceConfigResponseBody) SetSuccess(v bool) *GetTraceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetTraceConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTraceConfigResponseBodyData struct {
	// Indicates whether tracing analysis is enabled. Valid values: true and false
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The sampling rate.
	//
	// example:
	//
	// 50
	SampleRatio *int32 `json:"sampleRatio,omitempty" xml:"sampleRatio,omitempty"`
	// The service ID. This parameter exists when the traceType value is SKYWALKING.
	//
	// example:
	//
	// ss-co370icmjeu****
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service port. This parameter exists when the traceType value is SKYWALKING.
	//
	// example:
	//
	// 8090
	ServicePort *string `json:"servicePort,omitempty" xml:"servicePort,omitempty"`
	// The type of tracing analysis. Valid values:
	//
	// 	- XTRACE
	//
	// 	- SKYWALKING
	//
	// 	- OPENTELEMETRY
	//
	// 	- OTSKYWALKING
	//
	// example:
	//
	// SKYWALKING
	TraceType *string `json:"traceType,omitempty" xml:"traceType,omitempty"`
}

func (s GetTraceConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTraceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTraceConfigResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetTraceConfigResponseBodyData) GetSampleRatio() *int32 {
	return s.SampleRatio
}

func (s *GetTraceConfigResponseBodyData) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetTraceConfigResponseBodyData) GetServicePort() *string {
	return s.ServicePort
}

func (s *GetTraceConfigResponseBodyData) GetTraceType() *string {
	return s.TraceType
}

func (s *GetTraceConfigResponseBodyData) SetEnable(v bool) *GetTraceConfigResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) SetSampleRatio(v int32) *GetTraceConfigResponseBodyData {
	s.SampleRatio = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) SetServiceId(v string) *GetTraceConfigResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) SetServicePort(v string) *GetTraceConfigResponseBodyData {
	s.ServicePort = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) SetTraceType(v string) *GetTraceConfigResponseBodyData {
	s.TraceType = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

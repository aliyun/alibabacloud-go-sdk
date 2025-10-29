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
	// Response Code
	//
	// example:
	//
	// 200
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// Response Data
	Data *GetTraceConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Error Message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32AF2C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Boolean	Request Result, with the following values:
	//
	// true: Request succeeded.
	//
	// false: Request failed.
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
	// Whether to Enable Tracing:
	//
	// true: Enabled
	//
	// false: Disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Sampling Rate
	//
	// example:
	//
	// 50
	SampleRatio *int32 `json:"sampleRatio,omitempty" xml:"sampleRatio,omitempty"`
	// Service ID, present when the tracing type is SKYWALKING
	//
	// example:
	//
	// ss-co370icmjeu****
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// 服务端口，链路追踪类型为SKYWALKING时存在该参数
	//
	// example:
	//
	// 8090
	ServicePort *string `json:"servicePort,omitempty" xml:"servicePort,omitempty"`
	// Tracing Type:
	//
	// - XTRACE
	//
	// - SKYWALKING
	//
	// - OPENTELEMETRY
	//
	// - OTSKYWALKING
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

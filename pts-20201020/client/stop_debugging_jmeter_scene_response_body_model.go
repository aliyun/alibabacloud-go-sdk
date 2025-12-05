// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDebuggingJMeterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopDebuggingJMeterSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StopDebuggingJMeterSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopDebuggingJMeterSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopDebuggingJMeterSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopDebuggingJMeterSceneResponseBody
	GetSuccess() *bool
}

type StopDebuggingJMeterSceneResponseBody struct {
	// The system status code. If the request was returned, this parameter is left empty.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the request was returned, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was returned, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDebuggingJMeterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDebuggingJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopDebuggingJMeterSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopDebuggingJMeterSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopDebuggingJMeterSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopDebuggingJMeterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDebuggingJMeterSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDebuggingJMeterSceneResponseBody) SetCode(v string) *StopDebuggingJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) SetHttpStatusCode(v int32) *StopDebuggingJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) SetMessage(v string) *StopDebuggingJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) SetRequestId(v string) *StopDebuggingJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) SetSuccess(v bool) *StopDebuggingJMeterSceneResponseBody {
	s.Success = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) Validate() error {
	return dara.Validate(s)
}

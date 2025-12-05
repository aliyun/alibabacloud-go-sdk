// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTestingJMeterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopTestingJMeterSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StopTestingJMeterSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopTestingJMeterSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopTestingJMeterSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopTestingJMeterSceneResponseBody
	GetSuccess() *bool
}

type StopTestingJMeterSceneResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
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

func (s StopTestingJMeterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTestingJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopTestingJMeterSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopTestingJMeterSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopTestingJMeterSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopTestingJMeterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTestingJMeterSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopTestingJMeterSceneResponseBody) SetCode(v string) *StopTestingJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) SetHttpStatusCode(v int32) *StopTestingJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) SetMessage(v string) *StopTestingJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) SetRequestId(v string) *StopTestingJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) SetSuccess(v bool) *StopTestingJMeterSceneResponseBody {
	s.Success = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) Validate() error {
	return dara.Validate(s)
}

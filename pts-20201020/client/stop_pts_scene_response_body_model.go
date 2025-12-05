// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopPtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StopPtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopPtsSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopPtsSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopPtsSceneResponseBody
	GetSuccess() *bool
}

type StopPtsSceneResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
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
	// DD6F2ED8-E31B-497F-85AB-C4E358A5F6F9
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

func (s StopPtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopPtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopPtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopPtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopPtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopPtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopPtsSceneResponseBody) SetCode(v string) *StopPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopPtsSceneResponseBody) SetHttpStatusCode(v int32) *StopPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopPtsSceneResponseBody) SetMessage(v string) *StopPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopPtsSceneResponseBody) SetRequestId(v string) *StopPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPtsSceneResponseBody) SetSuccess(v bool) *StopPtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *StopPtsSceneResponseBody) Validate() error {
	return dara.Validate(s)
}

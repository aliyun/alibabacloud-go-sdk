// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDebugPtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopDebugPtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StopDebugPtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopDebugPtsSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopDebugPtsSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopDebugPtsSceneResponseBody
	GetSuccess() *bool
}

type StopDebugPtsSceneResponseBody struct {
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
	// The returned message. If the request was successful, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AE6505C-55CE-444A-B73B-810D0ED27C66
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

func (s StopDebugPtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDebugPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopDebugPtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopDebugPtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopDebugPtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopDebugPtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDebugPtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDebugPtsSceneResponseBody) SetCode(v string) *StopDebugPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) SetHttpStatusCode(v int32) *StopDebugPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) SetMessage(v string) *StopDebugPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) SetRequestId(v string) *StopDebugPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) SetSuccess(v bool) *StopDebugPtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) Validate() error {
	return dara.Validate(s)
}

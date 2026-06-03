// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteScriptRecordingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteScriptRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteScriptRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteScriptRecordingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScriptRecordingResponseBody
	GetSuccess() *bool
}

type DeleteScriptRecordingResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 56666881-887A-530A-B679-C3B6B6B142C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScriptRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScriptRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteScriptRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteScriptRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteScriptRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScriptRecordingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScriptRecordingResponseBody) SetCode(v string) *DeleteScriptRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScriptRecordingResponseBody) SetHttpStatusCode(v int32) *DeleteScriptRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteScriptRecordingResponseBody) SetMessage(v string) *DeleteScriptRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScriptRecordingResponseBody) SetRequestId(v string) *DeleteScriptRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScriptRecordingResponseBody) SetSuccess(v bool) *DeleteScriptRecordingResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScriptRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}

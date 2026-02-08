// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptWaveformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteScriptWaveformResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteScriptWaveformResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteScriptWaveformResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteScriptWaveformResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScriptWaveformResponseBody
	GetSuccess() *bool
}

type DeleteScriptWaveformResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 0FD69531-F82C-542D-A5E7-CB54544041C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScriptWaveformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptWaveformResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScriptWaveformResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteScriptWaveformResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteScriptWaveformResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteScriptWaveformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScriptWaveformResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScriptWaveformResponseBody) SetCode(v string) *DeleteScriptWaveformResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScriptWaveformResponseBody) SetHttpStatusCode(v int32) *DeleteScriptWaveformResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteScriptWaveformResponseBody) SetMessage(v string) *DeleteScriptWaveformResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScriptWaveformResponseBody) SetRequestId(v string) *DeleteScriptWaveformResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScriptWaveformResponseBody) SetSuccess(v bool) *DeleteScriptWaveformResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScriptWaveformResponseBody) Validate() error {
	return dara.Validate(s)
}

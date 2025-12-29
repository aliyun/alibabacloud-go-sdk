// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushVoiceBoxCommandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PushVoiceBoxCommandsResponseBody
	GetCode() *int32
	SetMessage(v string) *PushVoiceBoxCommandsResponseBody
	GetMessage() *string
	SetRequestId(v string) *PushVoiceBoxCommandsResponseBody
	GetRequestId() *string
	SetResult(v bool) *PushVoiceBoxCommandsResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *PushVoiceBoxCommandsResponseBody
	GetStatusCode() *int32
}

type PushVoiceBoxCommandsResponseBody struct {
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	StatusCode *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PushVoiceBoxCommandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushVoiceBoxCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PushVoiceBoxCommandsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PushVoiceBoxCommandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushVoiceBoxCommandsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PushVoiceBoxCommandsResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushVoiceBoxCommandsResponseBody) SetCode(v int32) *PushVoiceBoxCommandsResponseBody {
	s.Code = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) SetMessage(v string) *PushVoiceBoxCommandsResponseBody {
	s.Message = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) SetRequestId(v string) *PushVoiceBoxCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) SetResult(v bool) *PushVoiceBoxCommandsResponseBody {
	s.Result = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) SetStatusCode(v int32) *PushVoiceBoxCommandsResponseBody {
	s.StatusCode = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) Validate() error {
	return dara.Validate(s)
}

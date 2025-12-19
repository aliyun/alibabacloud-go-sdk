// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPptOutlineGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunPptOutlineGenerationResponseBody
	GetCode() *string
	SetHeader(v *RunPptOutlineGenerationResponseBodyHeader) *RunPptOutlineGenerationResponseBody
	GetHeader() *RunPptOutlineGenerationResponseBodyHeader
	SetHttpStatusCode(v string) *RunPptOutlineGenerationResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *RunPptOutlineGenerationResponseBody
	GetMessage() *string
	SetPayload(v *RunPptOutlineGenerationResponseBodyPayload) *RunPptOutlineGenerationResponseBody
	GetPayload() *RunPptOutlineGenerationResponseBodyPayload
	SetRequestId(v string) *RunPptOutlineGenerationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunPptOutlineGenerationResponseBody
	GetSuccess() *bool
}

type RunPptOutlineGenerationResponseBody struct {
	// example:
	//
	// success
	Code   *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Header *RunPptOutlineGenerationResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Payload *RunPptOutlineGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunPptOutlineGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunPptOutlineGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunPptOutlineGenerationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunPptOutlineGenerationResponseBody) GetHeader() *RunPptOutlineGenerationResponseBodyHeader {
	return s.Header
}

func (s *RunPptOutlineGenerationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RunPptOutlineGenerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunPptOutlineGenerationResponseBody) GetPayload() *RunPptOutlineGenerationResponseBodyPayload {
	return s.Payload
}

func (s *RunPptOutlineGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunPptOutlineGenerationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunPptOutlineGenerationResponseBody) SetCode(v string) *RunPptOutlineGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBody) SetHeader(v *RunPptOutlineGenerationResponseBodyHeader) *RunPptOutlineGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunPptOutlineGenerationResponseBody) SetHttpStatusCode(v string) *RunPptOutlineGenerationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBody) SetMessage(v string) *RunPptOutlineGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBody) SetPayload(v *RunPptOutlineGenerationResponseBodyPayload) *RunPptOutlineGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunPptOutlineGenerationResponseBody) SetRequestId(v string) *RunPptOutlineGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBody) SetSuccess(v bool) *RunPptOutlineGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBody) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunPptOutlineGenerationResponseBodyHeader struct {
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-started
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 1a3d7c9f-3a6d-4e49-b176-2d8721a27397
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// 8996314ce5514867943c71935e6a45af
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0bc1ec3a17435601877224179ecc8a
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunPptOutlineGenerationResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunPptOutlineGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunPptOutlineGenerationResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunPptOutlineGenerationResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunPptOutlineGenerationResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunPptOutlineGenerationResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunPptOutlineGenerationResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunPptOutlineGenerationResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunPptOutlineGenerationResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunPptOutlineGenerationResponseBodyHeader) SetErrorCode(v string) *RunPptOutlineGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyHeader) SetErrorMessage(v string) *RunPptOutlineGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyHeader) SetEvent(v string) *RunPptOutlineGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyHeader) SetSessionId(v string) *RunPptOutlineGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyHeader) SetStatusCode(v int32) *RunPptOutlineGenerationResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyHeader) SetTaskId(v string) *RunPptOutlineGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyHeader) SetTraceId(v string) *RunPptOutlineGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunPptOutlineGenerationResponseBodyPayload struct {
	Output *RunPptOutlineGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
}

func (s RunPptOutlineGenerationResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunPptOutlineGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunPptOutlineGenerationResponseBodyPayload) GetOutput() *RunPptOutlineGenerationResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunPptOutlineGenerationResponseBodyPayload) SetOutput(v *RunPptOutlineGenerationResponseBodyPayloadOutput) *RunPptOutlineGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyPayload) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunPptOutlineGenerationResponseBodyPayloadOutput struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunPptOutlineGenerationResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunPptOutlineGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunPptOutlineGenerationResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunPptOutlineGenerationResponseBodyPayloadOutput) SetText(v string) *RunPptOutlineGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunPptOutlineGenerationResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

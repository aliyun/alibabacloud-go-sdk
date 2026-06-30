// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentTaskResponseBody
	GetCode() *string
	SetData(v *CreateAgentTaskResponseBodyData) *CreateAgentTaskResponseBody
	GetData() *CreateAgentTaskResponseBodyData
	SetMessage(v string) *CreateAgentTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentTaskResponseBody
	GetSuccess() *bool
}

type CreateAgentTaskResponseBody struct {
	// The result code. A value of **200*	- indicates success. Any other value indicates failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *CreateAgentTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error details when the request fails. The value is **successful*	- when the request succeeds.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. You can use this field to determine whether the request succeeded: true indicates success. false/null indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentTaskResponseBody) GetData() *CreateAgentTaskResponseBodyData {
	return s.Data
}

func (s *CreateAgentTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentTaskResponseBody) SetCode(v string) *CreateAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentTaskResponseBody) SetData(v *CreateAgentTaskResponseBodyData) *CreateAgentTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateAgentTaskResponseBody) SetMessage(v string) *CreateAgentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentTaskResponseBody) SetRequestId(v string) *CreateAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentTaskResponseBody) SetSuccess(v bool) *CreateAgentTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentTaskResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 7C1DEF5F-2C18-4D36-99C6-8***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 6fa76916-3ce6-45d8-ac64-01b7f31***
	Vid *string `json:"Vid,omitempty" xml:"Vid,omitempty"`
}

func (s CreateAgentTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgentTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAgentTaskResponseBodyData) GetVid() *string {
	return s.Vid
}

func (s *CreateAgentTaskResponseBodyData) SetTaskId(v string) *CreateAgentTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateAgentTaskResponseBodyData) SetVid(v string) *CreateAgentTaskResponseBodyData {
	s.Vid = &v
	return s
}

func (s *CreateAgentTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

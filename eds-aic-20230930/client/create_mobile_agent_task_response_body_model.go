// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMobileAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMobileAgentTaskResponseBody
	GetCode() *string
	SetData(v *CreateMobileAgentTaskResponseBodyData) *CreateMobileAgentTaskResponseBody
	GetData() *CreateMobileAgentTaskResponseBodyData
	SetMessage(v string) *CreateMobileAgentTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMobileAgentTaskResponseBody
	GetRequestId() *string
}

type CreateMobileAgentTaskResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result object.
	Data *CreateMobileAgentTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMobileAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMobileAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMobileAgentTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMobileAgentTaskResponseBody) GetData() *CreateMobileAgentTaskResponseBodyData {
	return s.Data
}

func (s *CreateMobileAgentTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMobileAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMobileAgentTaskResponseBody) SetCode(v string) *CreateMobileAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMobileAgentTaskResponseBody) SetData(v *CreateMobileAgentTaskResponseBodyData) *CreateMobileAgentTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateMobileAgentTaskResponseBody) SetMessage(v string) *CreateMobileAgentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMobileAgentTaskResponseBody) SetRequestId(v string) *CreateMobileAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMobileAgentTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMobileAgentTaskResponseBodyData struct {
	// The session ID. Tasks with the same session ID share the same context, such as ses-260702-21b****.
	//
	// example:
	//
	// ses-261005-6ekkb1m9ls00
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-22ex666a653gq****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateMobileAgentTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMobileAgentTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMobileAgentTaskResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateMobileAgentTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMobileAgentTaskResponseBodyData) SetSessionId(v string) *CreateMobileAgentTaskResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateMobileAgentTaskResponseBodyData) SetTaskId(v string) *CreateMobileAgentTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateMobileAgentTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAgentConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstanceAgentConfigResponseBody
	GetCode() *string
	SetData(v *UpdateInstanceAgentConfigResponseBodyData) *UpdateInstanceAgentConfigResponseBody
	GetData() *UpdateInstanceAgentConfigResponseBodyData
	SetMessage(v string) *UpdateInstanceAgentConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceAgentConfigResponseBody
	GetRequestId() *string
}

type UpdateInstanceAgentConfigResponseBody struct {
	// The status code.
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned results.
	Data *UpdateInstanceAgentConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message. If `code == Success`, this field is empty. Otherwise, this field contains the error message of the request.
	//
	// example:
	//
	// Analysis ID: 56dae746-ef55-4f77-8373-cb3594c41457
	//
	// Analysis failed for instance i-bp164ns76pzirbwv0snt. Cause: Not get GPU trace data for \\"56dae746-ef55-4f77-8373-cb3594c41457\\" \\"[\\"93811\\"]\\"!
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateInstanceAgentConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAgentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAgentConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceAgentConfigResponseBody) GetData() *UpdateInstanceAgentConfigResponseBodyData {
	return s.Data
}

func (s *UpdateInstanceAgentConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceAgentConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceAgentConfigResponseBody) SetCode(v string) *UpdateInstanceAgentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceAgentConfigResponseBody) SetData(v *UpdateInstanceAgentConfigResponseBodyData) *UpdateInstanceAgentConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateInstanceAgentConfigResponseBody) SetMessage(v string) *UpdateInstanceAgentConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceAgentConfigResponseBody) SetRequestId(v string) *UpdateInstanceAgentConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceAgentConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceAgentConfigResponseBodyData struct {
	// The task ID. You can use this task ID to call the GetAgentTask operation to check the task execution status.
	//
	// example:
	//
	// bbe94a98-4192-4172-b856-95777e0a55d7
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateInstanceAgentConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAgentConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAgentConfigResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateInstanceAgentConfigResponseBodyData) SetTaskId(v string) *UpdateInstanceAgentConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateInstanceAgentConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

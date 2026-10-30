// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAgentConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateClusterAgentConfigResponseBody
	GetCode() *string
	SetData(v *UpdateClusterAgentConfigResponseBodyData) *UpdateClusterAgentConfigResponseBody
	GetData() *UpdateClusterAgentConfigResponseBodyData
	SetMessage(v string) *UpdateClusterAgentConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateClusterAgentConfigResponseBody
	GetRequestId() *string
}

type UpdateClusterAgentConfigResponseBody struct {
	// The status code.
	//
	// - If the value is `Success`, the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. If the authorization fails, check the `message` field for the detailed fault message.
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *UpdateClusterAgentConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If the value of `code` is `Success`, this field is empty.
	//
	// - Otherwise, this field contains the error message of the request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateClusterAgentConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAgentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterAgentConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateClusterAgentConfigResponseBody) GetData() *UpdateClusterAgentConfigResponseBodyData {
	return s.Data
}

func (s *UpdateClusterAgentConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateClusterAgentConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterAgentConfigResponseBody) SetCode(v string) *UpdateClusterAgentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateClusterAgentConfigResponseBody) SetData(v *UpdateClusterAgentConfigResponseBodyData) *UpdateClusterAgentConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateClusterAgentConfigResponseBody) SetMessage(v string) *UpdateClusterAgentConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClusterAgentConfigResponseBody) SetRequestId(v string) *UpdateClusterAgentConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterAgentConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateClusterAgentConfigResponseBodyData struct {
	// The task ID.
	//
	// You can call the GetAgentTask operation by using this task ID to check the execution status of the task.
	//
	// example:
	//
	// bbe94a98-4192-4172-b856-95777e0a55d7
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateClusterAgentConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAgentConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateClusterAgentConfigResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateClusterAgentConfigResponseBodyData) SetTaskId(v string) *UpdateClusterAgentConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateClusterAgentConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

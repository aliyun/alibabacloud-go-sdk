// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAgentConfigResponseBody
	GetCode() *string
	SetData(v *DeleteAgentConfigResponseBodyData) *DeleteAgentConfigResponseBody
	GetData() *DeleteAgentConfigResponseBodyData
	SetMessage(v string) *DeleteAgentConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAgentConfigResponseBody
	GetRequestId() *string
}

type DeleteAgentConfigResponseBody struct {
	// The status code.
	//
	// - `code == Success`: Authorization succeeded.
	//
	// - Other status codes indicate that authorization failed. If authorization failed, check the `message` field for detailed error messages.
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *DeleteAgentConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
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

func (s DeleteAgentConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAgentConfigResponseBody) GetData() *DeleteAgentConfigResponseBodyData {
	return s.Data
}

func (s *DeleteAgentConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentConfigResponseBody) SetCode(v string) *DeleteAgentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentConfigResponseBody) SetData(v *DeleteAgentConfigResponseBodyData) *DeleteAgentConfigResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAgentConfigResponseBody) SetMessage(v string) *DeleteAgentConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentConfigResponseBody) SetRequestId(v string) *DeleteAgentConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAgentConfigResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// The ID of the currently running task
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteAgentConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAgentConfigResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteAgentConfigResponseBodyData) SetTaskId(v string) *DeleteAgentConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteAgentConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

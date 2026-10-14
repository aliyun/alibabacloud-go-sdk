// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgentConfigResponseBody
	GetCode() *string
	SetData(v *UpdateAgentConfigResponseBodyData) *UpdateAgentConfigResponseBody
	GetData() *UpdateAgentConfigResponseBodyData
	SetMessage(v string) *UpdateAgentConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgentConfigResponseBody
	GetRequestId() *string
}

type UpdateAgentConfigResponseBody struct {
	// The status code.
	//
	// - If `code` is `Success`, the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. If the authorization failed, check the `message` field for the detailed fault information.
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *UpdateAgentConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code` is `Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error information.
	//
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAgentConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgentConfigResponseBody) GetData() *UpdateAgentConfigResponseBodyData {
	return s.Data
}

func (s *UpdateAgentConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentConfigResponseBody) SetCode(v string) *UpdateAgentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentConfigResponseBody) SetData(v *UpdateAgentConfigResponseBodyData) *UpdateAgentConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAgentConfigResponseBody) SetMessage(v string) *UpdateAgentConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentConfigResponseBody) SetRequestId(v string) *UpdateAgentConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentConfigResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 1234567890
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateAgentConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAgentConfigResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateAgentConfigResponseBodyData) SetTaskId(v string) *UpdateAgentConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateAgentConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

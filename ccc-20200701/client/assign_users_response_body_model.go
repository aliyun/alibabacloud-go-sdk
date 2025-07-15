// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AssignUsersResponseBody
	GetCode() *string
	SetData(v string) *AssignUsersResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *AssignUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AssignUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssignUsersResponseBody
	GetRequestId() *string
	SetWorkflowId(v string) *AssignUsersResponseBody
	GetWorkflowId() *string
}

type AssignUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1ca2b084-6f0a-454b-9851-29768a9a5832
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1ca2b084-6f0a-454b-9851-29768a9a5832
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s AssignUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AssignUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *AssignUsersResponseBody) GetData() *string {
	return s.Data
}

func (s *AssignUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AssignUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssignUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignUsersResponseBody) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *AssignUsersResponseBody) SetCode(v string) *AssignUsersResponseBody {
	s.Code = &v
	return s
}

func (s *AssignUsersResponseBody) SetData(v string) *AssignUsersResponseBody {
	s.Data = &v
	return s
}

func (s *AssignUsersResponseBody) SetHttpStatusCode(v int32) *AssignUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AssignUsersResponseBody) SetMessage(v string) *AssignUsersResponseBody {
	s.Message = &v
	return s
}

func (s *AssignUsersResponseBody) SetRequestId(v string) *AssignUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignUsersResponseBody) SetWorkflowId(v string) *AssignUsersResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *AssignUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

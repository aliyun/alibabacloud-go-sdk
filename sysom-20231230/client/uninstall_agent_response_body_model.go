// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallAgentResponseBody
	GetRequestId() *string
	SetCode(v string) *UninstallAgentResponseBody
	GetCode() *string
	SetData(v *UninstallAgentResponseBodyData) *UninstallAgentResponseBody
	GetData() *UninstallAgentResponseBodyData
	SetMessage(v string) *UninstallAgentResponseBody
	GetMessage() *string
}

type UninstallAgentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *UninstallAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UninstallAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallAgentResponseBody) GetData() *UninstallAgentResponseBodyData {
	return s.Data
}

func (s *UninstallAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallAgentResponseBody) SetRequestId(v string) *UninstallAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAgentResponseBody) SetCode(v string) *UninstallAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallAgentResponseBody) SetData(v *UninstallAgentResponseBodyData) *UninstallAgentResponseBody {
	s.Data = v
	return s
}

func (s *UninstallAgentResponseBody) SetMessage(v string) *UninstallAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallAgentResponseBody) Validate() error {
	return dara.Validate(s)
}

type UninstallAgentResponseBodyData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UninstallAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UninstallAgentResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallAgentResponseBodyData) SetTaskId(v string) *UninstallAgentResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UninstallAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

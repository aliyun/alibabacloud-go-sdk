// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentForClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallAgentForClusterResponseBody
	GetRequestId() *string
	SetCode(v string) *UninstallAgentForClusterResponseBody
	GetCode() *string
	SetData(v *UninstallAgentForClusterResponseBodyData) *UninstallAgentForClusterResponseBody
	GetData() *UninstallAgentForClusterResponseBodyData
	SetMessage(v string) *UninstallAgentForClusterResponseBody
	GetMessage() *string
}

type UninstallAgentForClusterResponseBody struct {
	// example:
	//
	// 44841312-7227-55C9-AE03-D59729BFAE38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *UninstallAgentForClusterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.NotAuthorizedInstance Instance 21 is not authorized
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UninstallAgentForClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentForClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallAgentForClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallAgentForClusterResponseBody) GetData() *UninstallAgentForClusterResponseBodyData {
	return s.Data
}

func (s *UninstallAgentForClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallAgentForClusterResponseBody) SetRequestId(v string) *UninstallAgentForClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAgentForClusterResponseBody) SetCode(v string) *UninstallAgentForClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallAgentForClusterResponseBody) SetData(v *UninstallAgentForClusterResponseBodyData) *UninstallAgentForClusterResponseBody {
	s.Data = v
	return s
}

func (s *UninstallAgentForClusterResponseBody) SetMessage(v string) *UninstallAgentForClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallAgentForClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type UninstallAgentForClusterResponseBodyData struct {
	// example:
	//
	// 049ea0609515414b9e19c3389d7ba638
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UninstallAgentForClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentForClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallAgentForClusterResponseBodyData) SetTaskId(v string) *UninstallAgentForClusterResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UninstallAgentForClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentForClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallAgentForClusterResponseBody
	GetRequestId() *string
	SetCode(v string) *InstallAgentForClusterResponseBody
	GetCode() *string
	SetData(v *InstallAgentForClusterResponseBodyData) *InstallAgentForClusterResponseBody
	GetData() *InstallAgentForClusterResponseBodyData
	SetMessage(v string) *InstallAgentForClusterResponseBody
	GetMessage() *string
}

type InstallAgentForClusterResponseBody struct {
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *InstallAgentForClusterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s InstallAgentForClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentForClusterResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAgentForClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAgentForClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallAgentForClusterResponseBody) GetData() *InstallAgentForClusterResponseBodyData {
	return s.Data
}

func (s *InstallAgentForClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallAgentForClusterResponseBody) SetRequestId(v string) *InstallAgentForClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAgentForClusterResponseBody) SetCode(v string) *InstallAgentForClusterResponseBody {
	s.Code = &v
	return s
}

func (s *InstallAgentForClusterResponseBody) SetData(v *InstallAgentForClusterResponseBodyData) *InstallAgentForClusterResponseBody {
	s.Data = v
	return s
}

func (s *InstallAgentForClusterResponseBody) SetMessage(v string) *InstallAgentForClusterResponseBody {
	s.Message = &v
	return s
}

func (s *InstallAgentForClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type InstallAgentForClusterResponseBodyData struct {
	// example:
	//
	// 049ea0609515414b9e19c3389d7ba638
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InstallAgentForClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentForClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallAgentForClusterResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallAgentForClusterResponseBodyData) SetTaskId(v string) *InstallAgentForClusterResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *InstallAgentForClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}

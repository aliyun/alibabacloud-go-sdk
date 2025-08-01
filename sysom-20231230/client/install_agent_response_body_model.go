// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallAgentResponseBody
	GetRequestId() *string
	SetCode(v string) *InstallAgentResponseBody
	GetCode() *string
	SetData(v *InstallAgentResponseBodyData) *InstallAgentResponseBody
	GetData() *InstallAgentResponseBodyData
	SetMessage(v string) *InstallAgentResponseBody
	GetMessage() *string
}

type InstallAgentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *InstallAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s InstallAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallAgentResponseBody) GetData() *InstallAgentResponseBodyData {
	return s.Data
}

func (s *InstallAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallAgentResponseBody) SetRequestId(v string) *InstallAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAgentResponseBody) SetCode(v string) *InstallAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InstallAgentResponseBody) SetData(v *InstallAgentResponseBodyData) *InstallAgentResponseBody {
	s.Data = v
	return s
}

func (s *InstallAgentResponseBody) SetMessage(v string) *InstallAgentResponseBody {
	s.Message = &v
	return s
}

func (s *InstallAgentResponseBody) Validate() error {
	return dara.Validate(s)
}

type InstallAgentResponseBodyData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InstallAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallAgentResponseBodyData) SetTaskId(v string) *InstallAgentResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *InstallAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

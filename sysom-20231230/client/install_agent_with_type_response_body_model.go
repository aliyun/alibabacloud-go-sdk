// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentWithTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallAgentWithTypeResponseBody
	GetCode() *string
	SetData(v *InstallAgentWithTypeResponseBodyData) *InstallAgentWithTypeResponseBody
	GetData() *InstallAgentWithTypeResponseBodyData
	SetMessage(v string) *InstallAgentWithTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallAgentWithTypeResponseBody
	GetRequestId() *string
}

type InstallAgentWithTypeResponseBody struct {
	// The status code.
	//
	// - `code == Success` indicates that the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. Check the `message` field for the detailed fault information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *InstallAgentWithTypeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which can be used for end-to-end diagnostics.
	//
	// example:
	//
	// 1D8887FC-4BDB-5A1C-AB19-135C29A9E481
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InstallAgentWithTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentWithTypeResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAgentWithTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallAgentWithTypeResponseBody) GetData() *InstallAgentWithTypeResponseBodyData {
	return s.Data
}

func (s *InstallAgentWithTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallAgentWithTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAgentWithTypeResponseBody) SetCode(v string) *InstallAgentWithTypeResponseBody {
	s.Code = &v
	return s
}

func (s *InstallAgentWithTypeResponseBody) SetData(v *InstallAgentWithTypeResponseBodyData) *InstallAgentWithTypeResponseBody {
	s.Data = v
	return s
}

func (s *InstallAgentWithTypeResponseBody) SetMessage(v string) *InstallAgentWithTypeResponseBody {
	s.Message = &v
	return s
}

func (s *InstallAgentWithTypeResponseBody) SetRequestId(v string) *InstallAgentWithTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAgentWithTypeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallAgentWithTypeResponseBodyData struct {
	// The task ID.
	//
	// You can use this task ID to call the GetAgentTask operation to check the task execution status.
	//
	// example:
	//
	// 26b3cd97389c43dcad6bc4901c36fcec
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s InstallAgentWithTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentWithTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallAgentWithTypeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallAgentWithTypeResponseBodyData) SetTaskId(v string) *InstallAgentWithTypeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *InstallAgentWithTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

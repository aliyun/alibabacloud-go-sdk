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
	// Request ID, which can be used for end-to-end diagnosis
	//
	// example:
	//
	// 1D8887FC-4BDB-5A1C-AB19-135C29A9E481
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status code
	//
	// - `code == Success` indicates that authorization succeeded.
	//
	// - Any other status code indicates that authorization failed. If authorization fails, view the `message` field to obtain detailed error information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data
	Data *InstallAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Error message
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallAgentResponseBodyData struct {
	// Job ID.
	//
	// You can use this job ID to invoke the GetAgentTask API to view the job execution status.
	//
	// example:
	//
	// 26b3cd97389c43dcad6bc4901c36fcec
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

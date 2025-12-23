// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeIdResponseBodyData) *DescribeIdResponseBody
	GetData() *DescribeIdResponseBodyData
	SetRequestId(v string) *DescribeIdResponseBody
	GetRequestId() *string
}

type DescribeIdResponseBody struct {
	// example:
	//
	// p-123****
	Data *DescribeIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIdResponseBody) GetData() *DescribeIdResponseBodyData {
	return s.Data
}

func (s *DescribeIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIdResponseBody) SetData(v *DescribeIdResponseBodyData) *DescribeIdResponseBody {
	s.Data = v
	return s
}

func (s *DescribeIdResponseBody) SetRequestId(v string) *DescribeIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIdResponseBodyData struct {
	// example:
	//
	// p-123****
	InputId *string `json:"inputId,omitempty" xml:"inputId,omitempty"`
	// example:
	//
	// PROJECT
	InputIdType *string `json:"inputIdType,omitempty" xml:"inputIdType,omitempty"`
	// example:
	//
	// mt-123****
	ManualTaskId *string `json:"manualTaskId,omitempty" xml:"manualTaskId,omitempty"`
	// example:
	//
	// mti-123****
	ManualTaskInstanceId *string `json:"manualTaskInstanceId,omitempty" xml:"manualTaskInstanceId,omitempty"`
	// example:
	//
	// p-123****
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// t-123****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// ti-123****
	TaskInstanceId *string `json:"taskInstanceId,omitempty" xml:"taskInstanceId,omitempty"`
	// example:
	//
	// w-123****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// example:
	//
	// wi-123****
	WorkflowInstanceId *string `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	// example:
	//
	// workspace-123****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeIdResponseBodyData) GetInputId() *string {
	return s.InputId
}

func (s *DescribeIdResponseBodyData) GetInputIdType() *string {
	return s.InputIdType
}

func (s *DescribeIdResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *DescribeIdResponseBodyData) GetManualTaskInstanceId() *string {
	return s.ManualTaskInstanceId
}

func (s *DescribeIdResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeIdResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeIdResponseBodyData) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *DescribeIdResponseBodyData) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DescribeIdResponseBodyData) GetWorkflowInstanceId() *string {
	return s.WorkflowInstanceId
}

func (s *DescribeIdResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeIdResponseBodyData) SetInputId(v string) *DescribeIdResponseBodyData {
	s.InputId = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetInputIdType(v string) *DescribeIdResponseBodyData {
	s.InputIdType = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetManualTaskId(v string) *DescribeIdResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetManualTaskInstanceId(v string) *DescribeIdResponseBodyData {
	s.ManualTaskInstanceId = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetProjectId(v string) *DescribeIdResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetTaskId(v string) *DescribeIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetTaskInstanceId(v string) *DescribeIdResponseBodyData {
	s.TaskInstanceId = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetWorkflowId(v string) *DescribeIdResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetWorkflowInstanceId(v string) *DescribeIdResponseBodyData {
	s.WorkflowInstanceId = &v
	return s
}

func (s *DescribeIdResponseBodyData) SetWorkspaceId(v string) *DescribeIdResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}

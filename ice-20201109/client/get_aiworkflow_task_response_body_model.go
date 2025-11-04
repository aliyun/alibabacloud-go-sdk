// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIWorkflowTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAIWorkflowTaskResponseBody
	GetRequestId() *string
	SetWorkflowTask(v *GetAIWorkflowTaskResponseBodyWorkflowTask) *GetAIWorkflowTaskResponseBody
	GetWorkflowTask() *GetAIWorkflowTaskResponseBodyWorkflowTask
}

type GetAIWorkflowTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the workflow task.
	WorkflowTask *GetAIWorkflowTaskResponseBodyWorkflowTask `json:"WorkflowTask,omitempty" xml:"WorkflowTask,omitempty" type:"Struct"`
}

func (s GetAIWorkflowTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIWorkflowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIWorkflowTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIWorkflowTaskResponseBody) GetWorkflowTask() *GetAIWorkflowTaskResponseBodyWorkflowTask {
	return s.WorkflowTask
}

func (s *GetAIWorkflowTaskResponseBody) SetRequestId(v string) *GetAIWorkflowTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBody) SetWorkflowTask(v *GetAIWorkflowTaskResponseBodyWorkflowTask) *GetAIWorkflowTaskResponseBody {
	s.WorkflowTask = v
	return s
}

func (s *GetAIWorkflowTaskResponseBody) Validate() error {
	if s.WorkflowTask != nil {
		if err := s.WorkflowTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIWorkflowTaskResponseBodyWorkflowTask struct {
	// The time when the task was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-07-28T02:17:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-08-19T02:28:22Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input parameters of the workflow task.
	//
	// example:
	//
	// {\\"source_language_id\\":\\"en\\",\\"live_id\\":123,\\"live_url\\":{\\"url\\":\\"rtmp://test.com.cn/video/638d9088fe4f15ce\\"}}
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// The results of the workflow nodes. The structure of this JSON varies based on the workflow\\"s configuration.
	//
	// example:
	//
	// {...}
	NodeResults *string `json:"NodeResults,omitempty" xml:"NodeResults,omitempty"`
	// The node output.
	//
	// example:
	//
	// {
	//
	// "result":"test"
	//
	// }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The task state.
	//
	// Valid values:
	//
	// 	- running
	//
	// 	- stopped
	//
	// 	- failed
	//
	// 	- partial-succeeded
	//
	// 	- succeeded
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the workflow task.
	//
	// example:
	//
	// ********-67fd-43aa-9cc1-3e7f********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The version of the workflow template that was executed.
	//
	// example:
	//
	// ****ec0a-e3b9-40b1-abf2-6549d00e****
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The workflow template information.
	Workflow *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow `json:"Workflow,omitempty" xml:"Workflow,omitempty" type:"Struct"`
}

func (s GetAIWorkflowTaskResponseBodyWorkflowTask) String() string {
	return dara.Prettify(s)
}

func (s GetAIWorkflowTaskResponseBodyWorkflowTask) GoString() string {
	return s.String()
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetInputs() *string {
	return s.Inputs
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetNodeResults() *string {
	return s.NodeResults
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetOutputs() *string {
	return s.Outputs
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetStatus() *string {
	return s.Status
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetUserData() *string {
	return s.UserData
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetVersion() *string {
	return s.Version
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) GetWorkflow() *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	return s.Workflow
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetCreateTime(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.CreateTime = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetFinishTime(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.FinishTime = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetInputs(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.Inputs = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetNodeResults(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.NodeResults = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetOutputs(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.Outputs = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetStatus(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.Status = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetTaskId(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.TaskId = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetUserData(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.UserData = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetVersion(v string) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.Version = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) SetWorkflow(v *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) *GetAIWorkflowTaskResponseBodyWorkflowTask {
	s.Workflow = v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTask) Validate() error {
	if s.Workflow != nil {
		if err := s.Workflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow struct {
	// The time when the template was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-08-20T01:35:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The workflow\\"s topological structure.
	//
	// example:
	//
	// {
	//
	// "nodes":[...],
	//
	// "edges":[...]
	//
	// }
	Graph *string `json:"Graph,omitempty" xml:"Graph,omitempty"`
	// The time when the template was last modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-08-20T01:35:04Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the workflow template.
	//
	// example:
	//
	// RealtimeTranslation
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Workflow template status. Valid values:
	//
	// 	- Draft
	//
	// 	- Published
	//
	// 	- Editing
	//
	// example:
	//
	// Draft
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scenario type of the template.
	//
	// example:
	//
	// Live
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The template version.
	//
	// example:
	//
	// ****ec0a-e3b9-40b1-abf2-6549d00e****
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the workflow template.
	//
	// example:
	//
	// ****3f44-f1f6-477e-9364-c5e6c49e****
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) String() string {
	return dara.Prettify(s)
}

func (s GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GoString() string {
	return s.String()
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetGraph() *string {
	return s.Graph
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetName() *string {
	return s.Name
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetStatus() *string {
	return s.Status
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetType() *string {
	return s.Type
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetVersion() *string {
	return s.Version
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetCreateTime(v string) *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.CreateTime = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetGraph(v string) *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.Graph = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetModifiedTime(v string) *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.ModifiedTime = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetName(v string) *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.Name = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetStatus(v string) *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.Status = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetType(v string) *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.Type = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetVersion(v string) *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.Version = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetWorkflowId(v string) *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.WorkflowId = &v
	return s
}

func (s *GetAIWorkflowTaskResponseBodyWorkflowTaskWorkflow) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *CreateWorkflowDefinitionRequest
	GetProjectId() *int64
	SetSpec(v string) *CreateWorkflowDefinitionRequest
	GetSpec() *string
}

type CreateWorkflowDefinitionRequest struct {
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Configuration page to obtain the workspace ID. You can also call the ListProjects operation to query the workspace ID of the current tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec information that describes the workflow. For more information about the specifications, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow).
	//
	// > How do I quickly obtain a FlowSpec template?
	//
	// > - Open a workflow in Data Studio and click Show Spec in the upper-right corner to obtain the FlowSpec description for the current workflow. You can quickly build a compliant template based on this FlowSpec description.
	//
	// 	Notice: This operation only supports creating workflows. Internal nodes described in the FlowSpec are not created.</notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "kind": "CycleWorkflow",
	//
	//     "version": "1.1.0",
	//
	//     "spec": {
	//
	//         "name": "OpenAPITestWorkflowDemo",
	//
	//         "type": "CycleWorkflow",
	//
	//         "workflows": [
	//
	//             {
	//
	//                 "script": {
	//
	//                     "path": "XX/OpenAPITest/WorkflowTest/OpenAPITestWorkflowDemo",
	//
	//                     "runtime": {
	//
	//                         "command": "WORKFLOW"
	//
	//                     }
	//
	//                 },
	//
	//                 "trigger": {
	//
	//                     "type": "Scheduler",
	//
	//                     "cron": "00 02 00 	- 	- ?",
	//
	//                     "startTime": "1970-01-01 00:00:00",
	//
	//                     "endTime": "9999-01-01 00:00:00",
	//
	//                     "timezone": "Asia/Shanghai",
	//
	//                     "delaySeconds": 0
	//
	//                 },
	//
	//                 "strategy": {
	//
	//                     "timeout": 0,
	//
	//                     "instanceMode": "T+1",
	//
	//                     "rerunMode": "Allowed",
	//
	//                     "rerunTimes": 3,
	//
	//                     "rerunInterval": 180000,
	//
	//                     "failureStrategy": "Break"
	//
	//                 },
	//
	//                 "name": "OpenAPITestWorkflowDemo",
	//
	//                 "outputs": {
	//
	//                     "nodeOutputs": [
	//
	//                         {
	//
	//                             "data": "workflow_output",
	//
	//                             "artifactType": "NodeOutput",
	//
	//                             "refTableName": "OpenAPITestWorkflowDemo"
	//
	//                         }
	//
	//                     ]
	//
	//                 },
	//
	//                 "dependencies": []
	//
	//             }
	//
	//         ]
	//
	//     }
	//
	// }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateWorkflowDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateWorkflowDefinitionRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateWorkflowDefinitionRequest) SetProjectId(v int64) *CreateWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateWorkflowDefinitionRequest) SetSpec(v string) *CreateWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

func (s *CreateWorkflowDefinitionRequest) Validate() error {
	return dara.Validate(s)
}

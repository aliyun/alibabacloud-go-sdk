// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateWorkflowDefinitionRequest
	GetId() *int64
	SetProjectId(v int64) *UpdateWorkflowDefinitionRequest
	GetProjectId() *int64
	SetSpec(v string) *UpdateWorkflowDefinitionRequest
	GetSpec() *string
}

type UpdateWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the workflow. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
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
	//         "name": "OpenAPI Test Workflow Demo",
	//
	//         "type": "CycleWorkflow",
	//
	//         "id": "652567824470354XXXX",
	//
	//         "workflows": [
	//
	//             {
	//
	//                 "id": "652567824470354XXXX",
	//
	//                 "script": {
	//
	//                     "path": "XX/OpenAPI_Test/Workflow_Test/OpenAPI_Test_Workflow_Demo",
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
	//                 "name": "OpenAPI Test Workflow Demo",
	//
	//                 "inputs": {},
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
	//                             "refTableName": "OpenAPI_Test_Workflow_Demo"
	//
	//                         }
	//
	//                     ]
	//
	//                 },
	//
	//                 "nodes": [],
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

func (s UpdateWorkflowDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateWorkflowDefinitionRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateWorkflowDefinitionRequest) SetId(v int64) *UpdateWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) SetProjectId(v int64) *UpdateWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) SetSpec(v string) *UpdateWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) Validate() error {
	return dara.Validate(s)
}

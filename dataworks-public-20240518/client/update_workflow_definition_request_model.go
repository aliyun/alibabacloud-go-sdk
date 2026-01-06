// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateWorkflowDefinitionRequest
	GetId() *string
	SetProjectId(v int64) *UpdateWorkflowDefinitionRequest
	GetProjectId() *int64
	SetSpec(v string) *UpdateWorkflowDefinitionRequest
	GetSpec() *string
}

type UpdateWorkflowDefinitionRequest struct {
	// The unique identifier of the Data Studio workflow.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the Data Studio workflow.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
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

func (s *UpdateWorkflowDefinitionRequest) GetId() *string {
	return s.Id
}

func (s *UpdateWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateWorkflowDefinitionRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateWorkflowDefinitionRequest) SetId(v string) *UpdateWorkflowDefinitionRequest {
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

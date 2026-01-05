// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *ImportWorkflowDefinitionRequest
	GetProjectId() *int64
	SetSpec(v string) *ImportWorkflowDefinitionRequest
	GetSpec() *string
}

type ImportWorkflowDefinitionRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec information for this workflow. For more information, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow/).
	//
	// > How to quickly obtain a FlowSpec template?
	//
	// 	- Open a workflow in Data Studio, then click "Show Spec" in the top-right corner to retrieve the FlowSpec description for the current workflow. You can use this FlowSpec description to quickly build a template that meets your requirements.
	//
	// > This interface supports creating both the workflow and its internal nodes simultaneously. Therefore, please pay close attention to the ID specified in the FlowSpec. If the provided ID already exists, the operation will be treated as an update. A create operation is performed only if the ID is omitted or does not exist.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "version": "1.1.0",
	//
	//     "kind": "CycleWorkflow",
	//
	//     "spec": {
	//
	//         "name": "Asynchronous_Workflow_Creation_Test",
	//
	//         "id": "632647691239009XXXX",
	//
	//         "type": "CycleWorkflow",
	//
	//         "workflows": [
	//
	//             {
	//
	//                 "script": {
	//
	//                     "path": "XX/OpenAPI_Test/Workflow_Test/Asynchronous_Workflow_Creation_Test",
	//
	//                     "runtime": {
	//
	//                         "command": "WORKFLOW"
	//
	//                     }
	//
	//                 },
	//
	//                 "id": "632647691239009XXXX",
	//
	//                 "trigger": {
	//
	//                     "type": "Scheduler",
	//
	//                     "cron": "00 03 00 	- 	- ?",
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
	//                 "name": "Asynchronous_Workflow_Creation_Test",
	//
	//                 "inputs": {},
	//
	//                 "outputs": {
	//
	//                     "nodeOutputs": [
	//
	//                         {
	//
	//                             "data": "632647691239009XXXX",
	//
	//                             "artifactType": "NodeOutput",
	//
	//                             "refTableName": "Asynchronous_Workflow_Creation_Test"
	//
	//                         }
	//
	//                     ]
	//
	//                 },
	//
	//                 "nodes": [
	//
	//                     {
	//
	//                         "recurrence": "Normal",
	//
	//                         "id": "742981001612325XXXX",
	//
	//                         "timeout": 0,
	//
	//                         "instanceMode": "T+1",
	//
	//                         "rerunMode": "Allowed",
	//
	//                         "rerunTimes": 3,
	//
	//                         "rerunInterval": 180000,
	//
	//                         "script": {
	//
	//                             "path": "XX/OpenAPI_Test/Workflow_Test/Asynchronous_Workflow_Creation_Test/111",
	//
	//                             "runtime": {
	//
	//                                 "command": "ODPS_SQL"
	//
	//                             },
	//
	//                             "content": "select now();\\n"
	//
	//                         },
	//
	//                         "trigger": {
	//
	//                             "type": "Scheduler",
	//
	//                             "cron": "00 24 00 	- 	- ?",
	//
	//                             "startTime": "1970-01-01 00:00:00",
	//
	//                             "endTime": "9999-01-01 00:00:00",
	//
	//                             "timezone": "Asia/Shanghai",
	//
	//                             "delaySeconds": 0
	//
	//                         },
	//
	//                         "name": "111",
	//
	//                         "inputs": {},
	//
	//                         "outputs": {
	//
	//                             "nodeOutputs": [
	//
	//                                 {
	//
	//                                     "data": "742981001612325XXXX",
	//
	//                                     "artifactType": "NodeOutput",
	//
	//                                     "refTableName": "111"
	//
	//                                 }
	//
	//                             ]
	//
	//                         }
	//
	//                     },
	//
	//                     {
	//
	//                         "recurrence": "Normal",
	//
	//                         "id": "595182137303408XXXX",
	//
	//                         "timeout": 0,
	//
	//                         "instanceMode": "T+1",
	//
	//                         "rerunMode": "Allowed",
	//
	//                         "rerunTimes": 3,
	//
	//                         "rerunInterval": 180000,
	//
	//                         "script": {
	//
	//                             "path": "XX/OpenAPI_Test/Workflow_Test/Asynchronous_Workflow_Creation_Test/222",
	//
	//                             "runtime": {
	//
	//                                 "command": "ODPS_SQL"
	//
	//                             },
	//
	//                             "content": "select now();\\n select 1;"
	//
	//                         },
	//
	//                         "trigger": {
	//
	//                             "type": "Scheduler",
	//
	//                             "cron": "00 00 00 	- 	- ?",
	//
	//                             "startTime": "1970-01-01 00:00:00",
	//
	//                             "endTime": "9999-01-01 00:00:00",
	//
	//                             "timezone": "Asia/Shanghai",
	//
	//                             "delaySeconds": 0
	//
	//                         },
	//
	//                         "name": "222",
	//
	//                         "inputs": {},
	//
	//                         "outputs": {
	//
	//                             "nodeOutputs": [
	//
	//                                 {
	//
	//                                     "data": "595182137303408XXXX",
	//
	//                                     "artifactType": "NodeOutput",
	//
	//                                     "refTableName": "222"
	//
	//                                 }
	//
	//                             ]
	//
	//                         }
	//
	//                     }
	//
	//                 ],
	//
	//                 "dependencies": [
	//
	//                     {
	//
	//                         "nodeId": "595182137303408XXXX",
	//
	//                         "depends": [
	//
	//                             {
	//
	//                                 "type": "Normal",
	//
	//                                 "output": "742981001612325XXXX",
	//
	//                                 "refTableName": "111"
	//
	//                             }
	//
	//                         ]
	//
	//                     }
	//
	//                 ]
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

func (s ImportWorkflowDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *ImportWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ImportWorkflowDefinitionRequest) GetSpec() *string {
	return s.Spec
}

func (s *ImportWorkflowDefinitionRequest) SetProjectId(v int64) *ImportWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *ImportWorkflowDefinitionRequest) SetSpec(v string) *ImportWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

func (s *ImportWorkflowDefinitionRequest) Validate() error {
	return dara.Validate(s)
}

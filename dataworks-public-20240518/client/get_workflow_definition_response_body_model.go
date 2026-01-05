// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkflowDefinitionResponseBody
	GetRequestId() *string
	SetWorkflowDefinition(v *GetWorkflowDefinitionResponseBodyWorkflowDefinition) *GetWorkflowDefinitionResponseBody
	GetWorkflowDefinition() *GetWorkflowDefinitionResponseBodyWorkflowDefinition
}

type GetWorkflowDefinitionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F2BDD628-8A21-5BD1-B930-1A2D5989XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the workflow.
	WorkflowDefinition *GetWorkflowDefinitionResponseBodyWorkflowDefinition `json:"WorkflowDefinition,omitempty" xml:"WorkflowDefinition,omitempty" type:"Struct"`
}

func (s GetWorkflowDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowDefinitionResponseBody) GetWorkflowDefinition() *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	return s.WorkflowDefinition
}

func (s *GetWorkflowDefinitionResponseBody) SetRequestId(v string) *GetWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBody) SetWorkflowDefinition(v *GetWorkflowDefinitionResponseBodyWorkflowDefinition) *GetWorkflowDefinitionResponseBody {
	s.WorkflowDefinition = v
	return s
}

func (s *GetWorkflowDefinitionResponseBody) Validate() error {
	if s.WorkflowDefinition != nil {
		if err := s.WorkflowDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowDefinitionResponseBodyWorkflowDefinition struct {
	// The time when the workflow was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1708481905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the workflow was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1708481905000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// OpenAPI test workflow Demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the workflow.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the workflow belongs.
	//
	// example:
	//
	// 307XXX
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the workflow. For more information, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow/).
	//
	// example:
	//
	// {
	//
	//     "metadata": {
	//
	//         "tenantId": "52425742456XXXX",
	//
	//         "projectId": "307XXXX",
	//
	//         "uuid": "463497880880954XXXX"
	//
	//     },
	//
	//     "kind": "CycleWorkflow",
	//
	//     "version": "1.1.0",
	//
	//     "spec": {
	//
	//         "name": "OpenAPI_Test_Workflow_Demo",
	//
	//         "id": "463497880880954XXXX",
	//
	//         "type": "CycleWorkflow",
	//
	//         "owner": "110755000425XXXX",
	//
	//         "workflows": [
	//
	//             {
	//
	//                 "script": {
	//
	//                     "path": "XX/OpenAPI_Test/Workflow_Test/OpenAPI_Test_Workflow_Demo",
	//
	//                     "runtime": {
	//
	//                         "command": "WORKFLOW"
	//
	//                     },
	//
	//                     "id": "698002781368644XXXX"
	//
	//                 },
	//
	//                 "id": "463497880880954XXXX",
	//
	//                 "trigger": {
	//
	//                     "type": "Scheduler",
	//
	//                     "id": "652567824470354XXXX",
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
	//                 "name": "OpenAPI_Test_Workflow_Demo",
	//
	//                 "owner": "110755000425XXXX",
	//
	//                 "metadata": {
	//
	//                     "owner": "110755000425XXXX",
	//
	//                     "ownerName": "XXXX@test.XXXX.com",
	//
	//                     "tenantId": "52425742456XXXX",
	//
	//                     "project": {
	//
	//                         "mode": "STANDARD",
	//
	//                         "projectId": "307303",
	//
	//                         "projectIdentifier": "lwttest_standard",
	//
	//                         "projectName": "XXXX",
	//
	//                         "projectOwnerId": "110755000425XXXX",
	//
	//                         "simple": false,
	//
	//                         "tenantId": "52425742456XXXX"
	//
	//                     },
	//
	//                     "projectId": "307XXXX"
	//
	//                 },
	//
	//                 "inputs": {},
	//
	//                 "outputs": {
	//
	//                     "nodeOutputs": [
	//
	//                         {
	//
	//                             "data": "463497880880954XXXX",
	//
	//                             "artifactType": "NodeOutput",
	//
	//                             "refTableName": "OpenAPI_Test_Workflow_Demo",
	//
	//                             "isDefault": true
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
	// The ID of the workflow on the scheduling side after publishing.
	//
	// example:
	//
	// 700006657495
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowDefinitionResponseBodyWorkflowDefinition) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDefinitionResponseBodyWorkflowDefinition) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) GetId() *string {
	return s.Id
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) GetName() *string {
	return s.Name
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) GetOwner() *string {
	return s.Owner
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) GetSpec() *string {
	return s.Spec
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetCreateTime(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetId(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Id = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetModifyTime(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.ModifyTime = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetName(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Name = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetOwner(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Owner = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetProjectId(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.ProjectId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetSpec(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Spec = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetWorkflowId(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) Validate() error {
	return dara.Validate(s)
}

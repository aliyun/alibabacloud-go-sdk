// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerId(v string) *CreateNodeRequest
	GetContainerId() *string
	SetProjectId(v int64) *CreateNodeRequest
	GetProjectId() *int64
	SetScene(v string) *CreateNodeRequest
	GetScene() *string
	SetSpec(v string) *CreateNodeRequest
	GetSpec() *string
}

type CreateNodeRequest struct {
	// The unique identifier of a container in which you want to create the node. The container can be a workflow or a container node. Specify this parameter when you need to create the node inside a container.
	//
	// 	Notice: If this parameter is specified, the path field defined in FlowSpec becomes invalid.
	//
	// 	Notice: This field was of the Long type in SDK versions earlier than 8.0.0 and is of the String type in SDK 8.0.0 and later. **This change does not affect normal SDK usage, and the parameter is still returned in the type defined in the SDK**. Only when you upgrade across SDK version 8.0.0, the type change may cause project compilation failures, and you need to manually correct the data type.
	//
	// example:
	//
	// 23451286945488XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace for this API call operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scenario in which the node is created. This parameter determines whether the node is created in the manual node area or the data development area. DATAWORKS_MANUAL_WORKFLOW can be used only when ContainerId is specified and the container is a manual workflow.
	//
	// Valid values:
	//
	// - DATAWORKS_PROJECT: project directory.
	//
	// - DATAWORKS_MANUAL_WORKFLOW: manual workflow.
	//
	// - DATAWORKS_MANUAL_TASK: manual task.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATAWORKS_PROJECT
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The FlowSpec information that describes the node. For more information about the specification, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow).
	//
	// > How to quickly obtain a FlowSpec template?
	//
	// > - In DataStudio, open a node, click Versions on the right side, view the latest version, and then view the scheduling configuration. This provides the FlowSpec description for the current node. You can use the FlowSpec description in the version to quickly build a template that meets your requirements.
	//
	// > How to specify the node content?
	//
	// > - Specify the node content in the $.spec.nodes[*].script.content field.
	//
	// > How to configure the content of a batch synchronization node?
	//
	// > - Write a script by following Step 4 in [Configure a batch synchronization node by using the code editor](https://www.alibabacloud.com/help/en/dataworks/user-guide/configure-a-batch-synchronization-node-by-using-the-code-editor), and specify the content in the $.spec.nodes[*].script.content field. Alternatively, create a batch synchronization node on the page and obtain the script content by viewing the version.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "version": "1.1.0",
	//
	//     "kind": "Node",
	//
	//     "spec": {
	//
	//         "nodes": [
	//
	//             {
	//
	//                 "recurrence": "Normal",
	//
	//                 "timeout": 0,
	//
	//                 "instanceMode": "T+1",
	//
	//                 "rerunMode": "Allowed",
	//
	//                 "rerunTimes": 3,
	//
	//                 "rerunInterval": 180000,
	//
	//                 "datasource": {
	//
	//                     "name": "odps_test"
	//
	//                 },
	//
	//                 "script": {
	//
	//                     "path": "XX/OpenAPI_Test/odpsSQL_Test",
	//
	//                     "runtime": {
	//
	//                         "command": "ODPS_SQL"
	//
	//                     },
	//
	//                     "content": "select now();"
	//
	//                 },
	//
	//                 "trigger": {
	//
	//                     "type": "Scheduler",
	//
	//                     "cron": "00 00 00 	- 	- ?",
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
	//                 "runtimeResource": {
	//
	//                     "resourceGroup": "S_res_group_XXXX_XXXX"
	//
	//                 },
	//
	//                 "name": "odpsSQL_Test",
	//
	//                 "inputs": {
	//
	//                     "nodeOutputs": [
	//
	//                         {
	//
	//                             "data": "project_root",
	//
	//                             "artifactType": "NodeOutput"
	//
	//                         }
	//
	//                     ]
	//
	//                 },
	//
	//                 "outputs": {
	//
	//                     "nodeOutputs": [
	//
	//                         {
	//
	//                             "data": "output_data",
	//
	//                             "artifactType": "NodeOutput",
	//
	//                             "refTableName": "odpsSQL_Test"
	//
	//                         }
	//
	//                     ]
	//
	//                 }
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

func (s CreateNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *CreateNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateNodeRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateNodeRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateNodeRequest) SetContainerId(v string) *CreateNodeRequest {
	s.ContainerId = &v
	return s
}

func (s *CreateNodeRequest) SetProjectId(v int64) *CreateNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateNodeRequest) SetScene(v string) *CreateNodeRequest {
	s.Scene = &v
	return s
}

func (s *CreateNodeRequest) SetSpec(v string) *CreateNodeRequest {
	s.Spec = &v
	return s
}

func (s *CreateNodeRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateNodeRequest
	GetId() *string
	SetProjectId(v int64) *UpdateNodeRequest
	GetProjectId() *int64
	SetSpec(v string) *UpdateNodeRequest
	GetSpec() *string
}

type UpdateNodeRequest struct {
	// The unique identifier of the Data Studio node.
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
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the Data Studio node.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
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
	//                 "id": "860438872620113XXXX",
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
	//                     "name": "odps_test",
	//
	//                     "type": "odps"
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
	//                             "data": "lwttest_standard_root",
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
	//         ],
	//
	//         "flow": [
	//
	//             {
	//
	//                 "nodeId": "860438872620113XXXX",
	//
	//                 "depends": [
	//
	//                     {
	//
	//                         "type": "Normal",
	//
	//                         "output": "project_root"
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

func (s UpdateNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeRequest) GetId() *string {
	return s.Id
}

func (s *UpdateNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateNodeRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateNodeRequest) SetId(v string) *UpdateNodeRequest {
	s.Id = &v
	return s
}

func (s *UpdateNodeRequest) SetProjectId(v int64) *UpdateNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateNodeRequest) SetSpec(v string) *UpdateNodeRequest {
	s.Spec = &v
	return s
}

func (s *UpdateNodeRequest) Validate() error {
	return dara.Validate(s)
}

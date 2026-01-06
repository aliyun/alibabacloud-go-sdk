// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNode(v *GetNodeResponseBodyNode) *GetNodeResponseBody
	GetNode() *GetNodeResponseBodyNode
	SetRequestId(v string) *GetNodeResponseBody
	GetRequestId() *string
}

type GetNodeResponseBody struct {
	// Data Studio node details.
	Node *GetNodeResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBody) GetNode() *GetNodeResponseBodyNode {
	return s.Node
}

func (s *GetNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeResponseBody) SetNode(v *GetNodeResponseBodyNode) *GetNodeResponseBody {
	s.Node = v
	return s
}

func (s *GetNodeResponseBody) SetRequestId(v string) *GetNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeResponseBody) Validate() error {
	if s.Node != nil {
		if err := s.Node.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodeResponseBodyNode struct {
	// The time when the node was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1700539206000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The unique identifier of the Data Studio node.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the node was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1700539206000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the node.
	//
	// example:
	//
	// 196596664824XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about this node. For more information, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow).
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
	//                 "id": "860438872620113XXXX",
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
	//                     "language": "odps-sql",
	//
	//                     "path": "XX/OpenAPI_Test/ODPS_SQL_Test",
	//
	//                     "runtime": {
	//
	//                         "command": "ODPS_SQL",
	//
	//                         "commandTypeId": 10
	//
	//                     },
	//
	//                     "content": "select now();",
	//
	//                     "id": "853573334108680XXXX"
	//
	//                 },
	//
	//                 "trigger": {
	//
	//                     "type": "Scheduler",
	//
	//                     "id": "543680677872062XXXX",
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
	//                     "resourceGroup": "S_res_group_XXXX_XXXX",
	//
	//                     "id": "623731286945488XXXX",
	//
	//                     "resourceGroupId": "7201XXXX"
	//
	//                 },
	//
	//                 "name": "ODPS_SQL_Test",
	//
	//                 "owner": "110755000425XXXX",
	//
	//                 "metadata": {
	//
	//                     "owner": "110755000425XXXX",
	//
	//                     "ownerName": "XXXXX@test.XXX.com",
	//
	//                     "projectId": "307XXX"
	//
	//                 },
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
	//                             "data": "860438872620113XXXX",
	//
	//                             "artifactType": "NodeOutput",
	//
	//                             "refTableName": "ODPS_SQL_Test",
	//
	//                             "isDefault": true
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
	//                         "output": "lwttest_standard_root"
	//
	//                     }
	//
	//                 ]
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "metadata": {
	//
	//         "uuid": "860438872620113XXXX"
	//
	//     }
	//
	// }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The ID of the corresponding scheduling task after the node is published.
	//
	// example:
	//
	// 700006680527
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetNodeResponseBodyNode) String() string {
	return dara.Prettify(s)
}

func (s GetNodeResponseBodyNode) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyNode) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNodeResponseBodyNode) GetId() *string {
	return s.Id
}

func (s *GetNodeResponseBodyNode) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetNodeResponseBodyNode) GetName() *string {
	return s.Name
}

func (s *GetNodeResponseBodyNode) GetOwner() *string {
	return s.Owner
}

func (s *GetNodeResponseBodyNode) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeResponseBodyNode) GetSpec() *string {
	return s.Spec
}

func (s *GetNodeResponseBodyNode) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetNodeResponseBodyNode) SetCreateTime(v int64) *GetNodeResponseBodyNode {
	s.CreateTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetId(v string) *GetNodeResponseBodyNode {
	s.Id = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetModifyTime(v int64) *GetNodeResponseBodyNode {
	s.ModifyTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetName(v string) *GetNodeResponseBodyNode {
	s.Name = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetOwner(v string) *GetNodeResponseBodyNode {
	s.Owner = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetProjectId(v int64) *GetNodeResponseBodyNode {
	s.ProjectId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetSpec(v string) *GetNodeResponseBodyNode {
	s.Spec = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetTaskId(v int64) *GetNodeResponseBodyNode {
	s.TaskId = &v
	return s
}

func (s *GetNodeResponseBodyNode) Validate() error {
	return dara.Validate(s)
}

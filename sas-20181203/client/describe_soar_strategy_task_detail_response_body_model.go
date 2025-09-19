// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSoarStrategyTaskDetailResponseBody
	GetRequestId() *string
	SetTaskDetail(v *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) *DescribeSoarStrategyTaskDetailResponseBody
	GetTaskDetail() *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail
}

type DescribeSoarStrategyTaskDetailResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBBEB173-1F43-505F-A876-C03ECDF6CE4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the task.
	TaskDetail *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Struct"`
}

func (s DescribeSoarStrategyTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarStrategyTaskDetailResponseBody) GetTaskDetail() *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail {
	return s.TaskDetail
}

func (s *DescribeSoarStrategyTaskDetailResponseBody) SetRequestId(v string) *DescribeSoarStrategyTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponseBody) SetTaskDetail(v *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) *DescribeSoarStrategyTaskDetailResponseBody {
	s.TaskDetail = v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSoarStrategyTaskDetailResponseBodyTaskDetail struct {
	// The operational log information of the task.
	//
	// example:
	//
	// {
	//
	// "resultContent": "{\\"failedNum\\":0,\\"totalNum\\":1,\\"successNum\\":1}",
	//
	// "resultStatus": 0,
	//
	// "status": 2
	//
	// }
	LogInfo *string `json:"LogInfo,omitempty" xml:"LogInfo,omitempty"`
	// The parameters of the task.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "name": "vulList",
	//
	//         "associationProperty": "sasAllVul",
	//
	//         "value": [
	//
	//             {
	//
	//                 "regionId": "cn-hangzhou",
	//
	//                 "instanceId": "i-bp10i3dtyq7x4****5e",
	//
	//                 "instanceName": "xxx",
	//
	//                 "vulId": 1222,
	//
	//                 "vulName": "centos:7:cesa-2024:1249",
	//
	//                 "vulAliasName": "CESA-2024:1249",
	//
	//                 "vulTag": "oval",
	//
	//                 "vulUuid": "3c5eb76a-df89-****-85ef-67562cdc2344",
	//
	//                 "vulType": "cve",
	//
	//                 "vulModifyTs": 1721324258000
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     {
	//
	//         "name": "snapshotConfig",
	//
	//         "associationProperty": "snapshotConfig",
	//
	//         "value": {
	//
	//             "ttl": 1
	//
	//         }
	//
	//     },
	//
	//     {
	//
	//         "name": "notifyConfig",
	//
	//         "associationProperty": "notifyConfig",
	//
	//         "value": {
	//
	//             "ding": [
	//
	//                 {
	//
	//                     "value": 2195,
	//
	//                     "label": "test"
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     }
	//
	// ]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The process information of the task.
	//
	// example:
	//
	// {
	//
	//     "edges": [
	//
	//         {
	//
	//             "level": 0,
	//
	//             "removeFlag": 0,
	//
	//             "source": 1,
	//
	//             "target": 8
	//
	//         }
	//
	//     ],
	//
	//     "nodes": [
	//
	//         {
	//
	//             "actionId": "Action_014s73k",
	//
	//             "iconUrl": "https://img.alicdn.com/tfs/TB1T*****jSZLeXXb9kVXa-12-14.svg",
	//
	//             "id": 1,
	//
	//             "label": "describeDisks",
	//
	//             "nodeName": "DescribeDisks",
	//
	//             "status": 0,
	//
	//             "type": "openAPI"
	//
	//         }
	//
	//     ]
	//
	// }
	ProcessInfo *string `json:"ProcessInfo,omitempty" xml:"ProcessInfo,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// task-1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) GetLogInfo() *string {
	return s.LogInfo
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) GetParams() *string {
	return s.Params
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) GetProcessInfo() *string {
	return s.ProcessInfo
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) SetLogInfo(v string) *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail {
	s.LogInfo = &v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) SetParams(v string) *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail {
	s.Params = &v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) SetProcessInfo(v string) *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail {
	s.ProcessInfo = &v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) SetTaskName(v string) *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponseBodyTaskDetail) Validate() error {
	return dara.Validate(s)
}

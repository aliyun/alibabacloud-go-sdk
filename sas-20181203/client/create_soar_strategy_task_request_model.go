// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoarStrategyTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v int64) *CreateSoarStrategyTaskRequest
	GetStrategyId() *int64
	SetStrategyName(v string) *CreateSoarStrategyTaskRequest
	GetStrategyName() *string
	SetStrategyTaskName(v string) *CreateSoarStrategyTaskRequest
	GetStrategyTaskName() *string
	SetStrategyTaskParams(v string) *CreateSoarStrategyTaskRequest
	GetStrategyTaskParams() *string
	SetStrategyTaskPlanExeTime(v int64) *CreateSoarStrategyTaskRequest
	GetStrategyTaskPlanExeTime() *int64
}

type CreateSoarStrategyTaskRequest struct {
	// The ID of the policy.
	//
	// >  You can call the [DescribeSoarSubscribedStrategy](~~DescribeSoarSubscribedStrategy~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13840
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the policy. Set the value to Automated Batch Vulnerability Fixing Policy for Multiple Servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// Automated Batch Vulnerability Fixing Policy for Multiple Servers
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The name of.the policy task.
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	StrategyTaskName *string `json:"StrategyTaskName,omitempty" xml:"StrategyTaskName,omitempty"`
	// The parameters of the policy task. The value is a JSON array.
	//
	// Vulnerability-related parameters:
	//
	// 	- name: vluList
	//
	// 	- associationProperty: sasAllVul
	//
	// 	- value: basic vulnerability information
	//
	// Snapshot-related parameters:
	//
	// 	- name: snapshotConfig
	//
	// 	- associationProperty: snapshotConfig
	//
	// 	- value: retention period
	//
	// Notification-related parameters:
	//
	// 	- name: notifyConfig
	//
	// 	- associationProperty: notifyConfig
	//
	// 	- value: email or DingTalk configuration information
	//
	// This parameter is required.
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
	//                 "instanceId": "i-bp10i******68lo5e",
	//
	//                 "instanceName": "instance****",
	//
	//                 "vulId": 3974347681,
	//
	//                 "vulName": "centos:7:cesa-2024:1249",
	//
	//                 "vulAliasName": "CESA-2024:1249",
	//
	//                 "vulTag": "oval",
	//
	//                 "vulUuid": "3c5eb76a-******-85ef-67562cdc2344",
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
	StrategyTaskParams *string `json:"StrategyTaskParams,omitempty" xml:"StrategyTaskParams,omitempty"`
	// The timestamp when the task is scheduled to start. Unit: milliseconds.
	//
	// example:
	//
	// 1586739841000
	StrategyTaskPlanExeTime *int64 `json:"StrategyTaskPlanExeTime,omitempty" xml:"StrategyTaskPlanExeTime,omitempty"`
}

func (s CreateSoarStrategyTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSoarStrategyTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSoarStrategyTaskRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *CreateSoarStrategyTaskRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CreateSoarStrategyTaskRequest) GetStrategyTaskName() *string {
	return s.StrategyTaskName
}

func (s *CreateSoarStrategyTaskRequest) GetStrategyTaskParams() *string {
	return s.StrategyTaskParams
}

func (s *CreateSoarStrategyTaskRequest) GetStrategyTaskPlanExeTime() *int64 {
	return s.StrategyTaskPlanExeTime
}

func (s *CreateSoarStrategyTaskRequest) SetStrategyId(v int64) *CreateSoarStrategyTaskRequest {
	s.StrategyId = &v
	return s
}

func (s *CreateSoarStrategyTaskRequest) SetStrategyName(v string) *CreateSoarStrategyTaskRequest {
	s.StrategyName = &v
	return s
}

func (s *CreateSoarStrategyTaskRequest) SetStrategyTaskName(v string) *CreateSoarStrategyTaskRequest {
	s.StrategyTaskName = &v
	return s
}

func (s *CreateSoarStrategyTaskRequest) SetStrategyTaskParams(v string) *CreateSoarStrategyTaskRequest {
	s.StrategyTaskParams = &v
	return s
}

func (s *CreateSoarStrategyTaskRequest) SetStrategyTaskPlanExeTime(v int64) *CreateSoarStrategyTaskRequest {
	s.StrategyTaskPlanExeTime = &v
	return s
}

func (s *CreateSoarStrategyTaskRequest) Validate() error {
	return dara.Validate(s)
}

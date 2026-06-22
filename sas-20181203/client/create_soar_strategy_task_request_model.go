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
	// The policy ID.
	//
	// >Call the [DescribeSoarSubscribedStrategy](~~DescribeSoarSubscribedStrategy~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13840
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The policy name. Fixed value: Automated Batch Vulnerability Fix Policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Automated Batch Vulnerability Fixing Policy for Multiple Servers
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The name of the policy task.
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	StrategyTaskName *string `json:"StrategyTaskName,omitempty" xml:"StrategyTaskName,omitempty"`
	// The parameter information of the policy. A string in JSONArray format with the following items:
	//
	// Vulnerability configuration item. Valid values:
	//
	// - name: vluList
	//
	// - associationProperty: sasAllVul
	//
	// - value: basic information about the vulnerability
	//
	// Snapshot configuration item. Valid values:
	//
	// - name: snapshotConfig
	//
	// - associationProperty: snapshotConfig
	//
	// - value: storage time information
	//
	// Notification configuration. Valid values:
	//
	// - name: notifyConfig
	//
	// - associationProperty: notifyConfig
	//
	// - value: email or DingTalk configuration information.
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
	// The planned execution timestamp of the policy task. Unit: milliseconds.
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

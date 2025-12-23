// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManualTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEmrClusterId(v string) *DescribeManualTaskInstanceResponseBody
	GetEmrClusterId() *string
	SetEndTime(v string) *DescribeManualTaskInstanceResponseBody
	GetEndTime() *string
	SetExternalAppId(v string) *DescribeManualTaskInstanceResponseBody
	GetExternalAppId() *string
	SetManualTaskInstanceId(v string) *DescribeManualTaskInstanceResponseBody
	GetManualTaskInstanceId() *string
	SetManualTaskInstanceName(v string) *DescribeManualTaskInstanceResponseBody
	GetManualTaskInstanceName() *string
	SetResourceGroupId(v string) *DescribeManualTaskInstanceResponseBody
	GetResourceGroupId() *string
	SetStartTime(v string) *DescribeManualTaskInstanceResponseBody
	GetStartTime() *string
	SetStatus(v string) *DescribeManualTaskInstanceResponseBody
	GetStatus() *string
	SetSubmitTime(v string) *DescribeManualTaskInstanceResponseBody
	GetSubmitTime() *string
	SetTaskParams(v string) *DescribeManualTaskInstanceResponseBody
	GetTaskParams() *string
	SetTaskType(v string) *DescribeManualTaskInstanceResponseBody
	GetTaskType() *string
	SetRequestId(v string) *DescribeManualTaskInstanceResponseBody
	GetRequestId() *string
}

type DescribeManualTaskInstanceResponseBody struct {
	// example:
	//
	// c-b933c5aac7f7***
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// application_123_***
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// example:
	//
	// mti-0k5vype05xm****
	ManualTaskInstanceId *string `json:"ManualTaskInstanceId,omitempty" xml:"ManualTaskInstanceId,omitempty"`
	// example:
	//
	// test
	ManualTaskInstanceName *string `json:"ManualTaskInstanceName,omitempty" xml:"ManualTaskInstanceName,omitempty"`
	// example:
	//
	// wg-123abc***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeManualTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeManualTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskInstanceResponseBody) GetEmrClusterId() *string {
	return s.EmrClusterId
}

func (s *DescribeManualTaskInstanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeManualTaskInstanceResponseBody) GetExternalAppId() *string {
	return s.ExternalAppId
}

func (s *DescribeManualTaskInstanceResponseBody) GetManualTaskInstanceId() *string {
	return s.ManualTaskInstanceId
}

func (s *DescribeManualTaskInstanceResponseBody) GetManualTaskInstanceName() *string {
	return s.ManualTaskInstanceName
}

func (s *DescribeManualTaskInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeManualTaskInstanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeManualTaskInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeManualTaskInstanceResponseBody) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *DescribeManualTaskInstanceResponseBody) GetTaskParams() *string {
	return s.TaskParams
}

func (s *DescribeManualTaskInstanceResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeManualTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeManualTaskInstanceResponseBody) SetEmrClusterId(v string) *DescribeManualTaskInstanceResponseBody {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetEndTime(v string) *DescribeManualTaskInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetExternalAppId(v string) *DescribeManualTaskInstanceResponseBody {
	s.ExternalAppId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetManualTaskInstanceId(v string) *DescribeManualTaskInstanceResponseBody {
	s.ManualTaskInstanceId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetManualTaskInstanceName(v string) *DescribeManualTaskInstanceResponseBody {
	s.ManualTaskInstanceName = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetResourceGroupId(v string) *DescribeManualTaskInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetStartTime(v string) *DescribeManualTaskInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetStatus(v string) *DescribeManualTaskInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetSubmitTime(v string) *DescribeManualTaskInstanceResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetTaskParams(v string) *DescribeManualTaskInstanceResponseBody {
	s.TaskParams = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetTaskType(v string) *DescribeManualTaskInstanceResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetRequestId(v string) *DescribeManualTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManualTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeManualTaskResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeManualTaskResponseBody
	GetDescription() *string
	SetManualTaskId(v string) *DescribeManualTaskResponseBody
	GetManualTaskId() *string
	SetManualTaskName(v string) *DescribeManualTaskResponseBody
	GetManualTaskName() *string
	SetParentDirectoryId(v string) *DescribeManualTaskResponseBody
	GetParentDirectoryId() *string
	SetProjectId(v string) *DescribeManualTaskResponseBody
	GetProjectId() *string
	SetResourceIds(v string) *DescribeManualTaskResponseBody
	GetResourceIds() *string
	SetTaskParams(v string) *DescribeManualTaskResponseBody
	GetTaskParams() *string
	SetTaskType(v string) *DescribeManualTaskResponseBody
	GetTaskType() *string
	SetUpdateTime(v string) *DescribeManualTaskResponseBody
	GetUpdateTime() *string
	SetRequestId(v string) *DescribeManualTaskResponseBody
	GetRequestId() *string
}

type DescribeManualTaskResponseBody struct {
	// example:
	//
	// 2024-03-27 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mt-3q9jo749ne5****
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// example:
	//
	// test
	ManualTaskName *string `json:"ManualTaskName,omitempty" xml:"ManualTaskName,omitempty"`
	// example:
	//
	// mtd-oy98v7n43el****
	ParentDirectoryId *string `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// r-oy98v7n43el****
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
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
	// example:
	//
	// 2024-03-27 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeManualTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeManualTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeManualTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeManualTaskResponseBody) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *DescribeManualTaskResponseBody) GetManualTaskName() *string {
	return s.ManualTaskName
}

func (s *DescribeManualTaskResponseBody) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *DescribeManualTaskResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeManualTaskResponseBody) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *DescribeManualTaskResponseBody) GetTaskParams() *string {
	return s.TaskParams
}

func (s *DescribeManualTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeManualTaskResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeManualTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeManualTaskResponseBody) SetCreateTime(v string) *DescribeManualTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetDescription(v string) *DescribeManualTaskResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetManualTaskId(v string) *DescribeManualTaskResponseBody {
	s.ManualTaskId = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetManualTaskName(v string) *DescribeManualTaskResponseBody {
	s.ManualTaskName = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetParentDirectoryId(v string) *DescribeManualTaskResponseBody {
	s.ParentDirectoryId = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetProjectId(v string) *DescribeManualTaskResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetResourceIds(v string) *DescribeManualTaskResponseBody {
	s.ResourceIds = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetTaskParams(v string) *DescribeManualTaskResponseBody {
	s.TaskParams = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetTaskType(v string) *DescribeManualTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetUpdateTime(v string) *DescribeManualTaskResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetRequestId(v string) *DescribeManualTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeManualTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListManualTasksResponseBodyData) *ListManualTasksResponseBody
	GetData() []*ListManualTasksResponseBodyData
	SetNextToken(v string) *ListManualTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListManualTasksResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *ListManualTasksResponseBody
	GetTotalSize() *int32
}

type ListManualTasksResponseBody struct {
	Data []*ListManualTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListManualTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListManualTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListManualTasksResponseBody) GetData() []*ListManualTasksResponseBodyData {
	return s.Data
}

func (s *ListManualTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListManualTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListManualTasksResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListManualTasksResponseBody) SetData(v []*ListManualTasksResponseBodyData) *ListManualTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListManualTasksResponseBody) SetNextToken(v string) *ListManualTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListManualTasksResponseBody) SetRequestId(v string) *ListManualTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManualTasksResponseBody) SetTotalSize(v int32) *ListManualTasksResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListManualTasksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListManualTasksResponseBodyData struct {
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
}

func (s ListManualTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListManualTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListManualTasksResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListManualTasksResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListManualTasksResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *ListManualTasksResponseBodyData) GetManualTaskName() *string {
	return s.ManualTaskName
}

func (s *ListManualTasksResponseBodyData) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *ListManualTasksResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListManualTasksResponseBodyData) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *ListManualTasksResponseBodyData) GetTaskParams() *string {
	return s.TaskParams
}

func (s *ListManualTasksResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *ListManualTasksResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListManualTasksResponseBodyData) SetCreateTime(v string) *ListManualTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetDescription(v string) *ListManualTasksResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetManualTaskId(v string) *ListManualTasksResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetManualTaskName(v string) *ListManualTasksResponseBodyData {
	s.ManualTaskName = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetParentDirectoryId(v string) *ListManualTasksResponseBodyData {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetProjectId(v string) *ListManualTasksResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetResourceIds(v string) *ListManualTasksResponseBodyData {
	s.ResourceIds = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetTaskParams(v string) *ListManualTasksResponseBodyData {
	s.TaskParams = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetTaskType(v string) *ListManualTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetUpdateTime(v string) *ListManualTasksResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListManualTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}

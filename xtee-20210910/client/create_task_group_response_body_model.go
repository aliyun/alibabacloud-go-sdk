// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTaskGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreateTaskGroupResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateTaskGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTaskGroupResponseBody
	GetRequestId() *string
	SetResultObject(v *CreateTaskGroupResponseBodyResultObject) *CreateTaskGroupResponseBody
	GetResultObject() *CreateTaskGroupResponseBodyResultObject
}

type CreateTaskGroupResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	ResultObject *CreateTaskGroupResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CreateTaskGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTaskGroupResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateTaskGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTaskGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskGroupResponseBody) GetResultObject() *CreateTaskGroupResponseBodyResultObject {
	return s.ResultObject
}

func (s *CreateTaskGroupResponseBody) SetCode(v string) *CreateTaskGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetHttpStatusCode(v string) *CreateTaskGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetMessage(v string) *CreateTaskGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetRequestId(v string) *CreateTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetResultObject(v *CreateTaskGroupResponseBodyResultObject) *CreateTaskGroupResponseBody {
	s.ResultObject = v
	return s
}

func (s *CreateTaskGroupResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTaskGroupResponseBodyResultObject struct {
	// The creation time.
	//
	// example:
	//
	// 1750645267000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator user ID.
	//
	// example:
	//
	// 345298
	CreatorUserId *int32 `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// The group status.
	//
	// example:
	//
	// RUNNING
	GroupStatus *string `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	// The task group name.
	SampleNames []*string `json:"SampleNames,omitempty" xml:"SampleNames,omitempty" type:"Repeated"`
	// The number of subtasks generated from task parsing and splitting.
	//
	// example:
	//
	// 3
	SubTaskCount *int32 `json:"SubTaskCount,omitempty" xml:"SubTaskCount,omitempty"`
	// The subtask list.
	SubTaskList []*CreateTaskGroupResponseBodyResultObjectSubTaskList `json:"SubTaskList,omitempty" xml:"SubTaskList,omitempty" type:"Repeated"`
	// The scenario.
	//
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// The task group ID.
	//
	// > This parameter is in invitational preview. When this parameter is used, other query conditions become invalid.
	//
	// example:
	//
	// g-0jlcreertd0p471l6f72
	TaskGroupId *int32 `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
	// The task group name.
	//
	// example:
	//
	// GroupTest
	TaskGroupName *string `json:"TaskGroupName,omitempty" xml:"TaskGroupName,omitempty"`
}

func (s CreateTaskGroupResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskGroupResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBodyResultObject) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateTaskGroupResponseBodyResultObject) GetCreatorUserId() *int32 {
	return s.CreatorUserId
}

func (s *CreateTaskGroupResponseBodyResultObject) GetGroupStatus() *string {
	return s.GroupStatus
}

func (s *CreateTaskGroupResponseBodyResultObject) GetSampleNames() []*string {
	return s.SampleNames
}

func (s *CreateTaskGroupResponseBodyResultObject) GetSubTaskCount() *int32 {
	return s.SubTaskCount
}

func (s *CreateTaskGroupResponseBodyResultObject) GetSubTaskList() []*CreateTaskGroupResponseBodyResultObjectSubTaskList {
	return s.SubTaskList
}

func (s *CreateTaskGroupResponseBodyResultObject) GetTab() *string {
	return s.Tab
}

func (s *CreateTaskGroupResponseBodyResultObject) GetTaskGroupId() *int32 {
	return s.TaskGroupId
}

func (s *CreateTaskGroupResponseBodyResultObject) GetTaskGroupName() *string {
	return s.TaskGroupName
}

func (s *CreateTaskGroupResponseBodyResultObject) SetCreateTime(v int64) *CreateTaskGroupResponseBodyResultObject {
	s.CreateTime = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetCreatorUserId(v int32) *CreateTaskGroupResponseBodyResultObject {
	s.CreatorUserId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetGroupStatus(v string) *CreateTaskGroupResponseBodyResultObject {
	s.GroupStatus = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetSampleNames(v []*string) *CreateTaskGroupResponseBodyResultObject {
	s.SampleNames = v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetSubTaskCount(v int32) *CreateTaskGroupResponseBodyResultObject {
	s.SubTaskCount = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetSubTaskList(v []*CreateTaskGroupResponseBodyResultObjectSubTaskList) *CreateTaskGroupResponseBodyResultObject {
	s.SubTaskList = v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetTab(v string) *CreateTaskGroupResponseBodyResultObject {
	s.Tab = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetTaskGroupId(v int32) *CreateTaskGroupResponseBodyResultObject {
	s.TaskGroupId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetTaskGroupName(v string) *CreateTaskGroupResponseBodyResultObject {
	s.TaskGroupName = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) Validate() error {
	if s.SubTaskList != nil {
		for _, item := range s.SubTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTaskGroupResponseBodyResultObjectSubTaskList struct {
	// The reviewer.
	//
	// example:
	//
	// 303872
	Checker *string `json:"Checker,omitempty" xml:"Checker,omitempty"`
	// The file name.
	//
	// example:
	//
	// model_test_enorl_202606040953+10.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The number of rows in the file.
	//
	// example:
	//
	// 12
	FileRows *string `json:"FileRows,omitempty" xml:"FileRows,omitempty"`
	// The task end time.
	//
	// example:
	//
	// 2026-01-12 15:47:23
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The user group name.
	//
	// example:
	//
	// pts-demo
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether the task is billed.
	//
	// example:
	//
	// true
	IsCharge *string `json:"IsCharge,omitempty" xml:"IsCharge,omitempty"`
	// The model scenario.
	//
	// example:
	//
	// rfs
	ModelScene *string `json:"ModelScene,omitempty" xml:"ModelScene,omitempty"`
	// The sample IDs.
	//
	// example:
	//
	// 4
	SampleId *string `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
	// The sample name.
	//
	// example:
	//
	// fs
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// The service code.
	//
	// example:
	//
	// anti_fraud_customed
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The service name.
	//
	// example:
	//
	// 多头风险前筛
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The subtask ID.
	//
	// example:
	//
	// 3
	SubTaskId *int64 `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
	// The scenario.
	//
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// The task group ID.
	//
	// example:
	//
	// 4
	TaskGroupId *string `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
	// The task name.
	//
	// example:
	//
	// 阿里多头_68629_20260115
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The execution status of the import task. Valid values:
	//
	// - DOING: Running.
	//
	// - FINISH: Completed.
	//
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s CreateTaskGroupResponseBodyResultObjectSubTaskList) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskGroupResponseBodyResultObjectSubTaskList) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetChecker() *string {
	return s.Checker
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetFileName() *string {
	return s.FileName
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetFileRows() *string {
	return s.FileRows
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetIsCharge() *string {
	return s.IsCharge
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetModelScene() *string {
	return s.ModelScene
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetSampleId() *string {
	return s.SampleId
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetSampleName() *string {
	return s.SampleName
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetSubTaskId() *int64 {
	return s.SubTaskId
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetTab() *string {
	return s.Tab
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetTaskGroupId() *string {
	return s.TaskGroupId
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetChecker(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.Checker = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetFileName(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.FileName = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetFileRows(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.FileRows = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetFinishTime(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.FinishTime = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetGroupName(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.GroupName = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetIsCharge(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.IsCharge = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetModelScene(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.ModelScene = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetSampleId(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.SampleId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetSampleName(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.SampleName = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetServiceCode(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.ServiceCode = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetServiceName(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.ServiceName = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetSubTaskId(v int64) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.SubTaskId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetTab(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.Tab = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetTaskGroupId(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.TaskGroupId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetTaskName(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.TaskName = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) SetTaskStatus(v string) *CreateTaskGroupResponseBodyResultObjectSubTaskList {
	s.TaskStatus = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObjectSubTaskList) Validate() error {
	return dara.Validate(s)
}

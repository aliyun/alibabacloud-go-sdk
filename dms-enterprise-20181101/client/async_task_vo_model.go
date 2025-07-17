// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncTaskVO interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *AsyncTaskVO
	GetDatasetId() *string
	SetId(v int64) *AsyncTaskVO
	GetId() *int64
	SetRemark(v string) *AsyncTaskVO
	GetRemark() *string
	SetTaskName(v string) *AsyncTaskVO
	GetTaskName() *string
	SetTaskStatus(v int32) *AsyncTaskVO
	GetTaskStatus() *int32
	SetTaskType(v int32) *AsyncTaskVO
	GetTaskType() *int32
	SetUserId(v int64) *AsyncTaskVO
	GetUserId() *int64
}

type AsyncTaskVO struct {
	DatasetId  *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType   *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UserId     *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AsyncTaskVO) String() string {
	return dara.Prettify(s)
}

func (s AsyncTaskVO) GoString() string {
	return s.String()
}

func (s *AsyncTaskVO) GetDatasetId() *string {
	return s.DatasetId
}

func (s *AsyncTaskVO) GetId() *int64 {
	return s.Id
}

func (s *AsyncTaskVO) GetRemark() *string {
	return s.Remark
}

func (s *AsyncTaskVO) GetTaskName() *string {
	return s.TaskName
}

func (s *AsyncTaskVO) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *AsyncTaskVO) GetTaskType() *int32 {
	return s.TaskType
}

func (s *AsyncTaskVO) GetUserId() *int64 {
	return s.UserId
}

func (s *AsyncTaskVO) SetDatasetId(v string) *AsyncTaskVO {
	s.DatasetId = &v
	return s
}

func (s *AsyncTaskVO) SetId(v int64) *AsyncTaskVO {
	s.Id = &v
	return s
}

func (s *AsyncTaskVO) SetRemark(v string) *AsyncTaskVO {
	s.Remark = &v
	return s
}

func (s *AsyncTaskVO) SetTaskName(v string) *AsyncTaskVO {
	s.TaskName = &v
	return s
}

func (s *AsyncTaskVO) SetTaskStatus(v int32) *AsyncTaskVO {
	s.TaskStatus = &v
	return s
}

func (s *AsyncTaskVO) SetTaskType(v int32) *AsyncTaskVO {
	s.TaskType = &v
	return s
}

func (s *AsyncTaskVO) SetUserId(v int64) *AsyncTaskVO {
	s.UserId = &v
	return s
}

func (s *AsyncTaskVO) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTaskListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeTaskListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeTaskListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeTaskListResponseBodyResultObject) *DescribeTaskListResponseBody
	GetResultObject() []*DescribeTaskListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeTaskListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeTaskListResponseBody
	GetTotalPage() *int32
}

type DescribeTaskListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeTaskListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTaskListResponseBody) GetResultObject() []*DescribeTaskListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTaskListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeTaskListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeTaskListResponseBody) SetRequestId(v string) *DescribeTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskListResponseBody) SetCurrentPage(v int32) *DescribeTaskListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTaskListResponseBody) SetPageSize(v int32) *DescribeTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTaskListResponseBody) SetResultObject(v []*DescribeTaskListResponseBodyResultObject) *DescribeTaskListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTaskListResponseBody) SetTotalItem(v int32) *DescribeTaskListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeTaskListResponseBody) SetTotalPage(v int32) *DescribeTaskListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeTaskListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTaskListResponseBodyResultObject struct {
	// Completion time, in milliseconds.
	//
	// example:
	//
	// 1753804800000
	CompletionTime *int64 `json:"completionTime,omitempty" xml:"completionTime,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1753804800000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Total number of mark information.
	//
	// example:
	//
	// 100
	Mark *int32 `json:"mark,omitempty" xml:"mark,omitempty"`
	// Remark.
	//
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// Scene name
	//
	// example:
	//
	// 样本调度
	SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty"`
	// Data status.
	//
	// -1: Failed
	//
	// 0: Deleted
	//
	// 1: Pending
	//
	// 2: Success
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 1102
	TaskLogId *int64 `json:"taskLogId,omitempty" xml:"taskLogId,omitempty"`
	// Task type
	//
	// 1: Data upload
	//
	// 2: Supplemental upload
	//
	// 3: Labeling
	//
	// example:
	//
	// 1
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s DescribeTaskListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTaskListResponseBodyResultObject) GetCompletionTime() *int64 {
	return s.CompletionTime
}

func (s *DescribeTaskListResponseBodyResultObject) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeTaskListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeTaskListResponseBodyResultObject) GetMark() *int32 {
	return s.Mark
}

func (s *DescribeTaskListResponseBodyResultObject) GetRemark() *string {
	return s.Remark
}

func (s *DescribeTaskListResponseBodyResultObject) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeTaskListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeTaskListResponseBodyResultObject) GetTaskLogId() *int64 {
	return s.TaskLogId
}

func (s *DescribeTaskListResponseBodyResultObject) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskListResponseBodyResultObject) SetCompletionTime(v int64) *DescribeTaskListResponseBodyResultObject {
	s.CompletionTime = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) SetCreateTime(v int64) *DescribeTaskListResponseBodyResultObject {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) SetId(v int64) *DescribeTaskListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) SetMark(v int32) *DescribeTaskListResponseBodyResultObject {
	s.Mark = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) SetRemark(v string) *DescribeTaskListResponseBodyResultObject {
	s.Remark = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) SetSceneName(v string) *DescribeTaskListResponseBodyResultObject {
	s.SceneName = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) SetStatus(v string) *DescribeTaskListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) SetTaskLogId(v int64) *DescribeTaskListResponseBodyResultObject {
	s.TaskLogId = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) SetTaskType(v string) *DescribeTaskListResponseBodyResultObject {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

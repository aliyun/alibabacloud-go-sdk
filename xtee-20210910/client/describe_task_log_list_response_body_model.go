// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskLogListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTaskLogListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeTaskLogListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeTaskLogListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeTaskLogListResponseBodyResultObject) *DescribeTaskLogListResponseBody
	GetResultObject() []*DescribeTaskLogListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeTaskLogListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeTaskLogListResponseBody
	GetTotalPage() *int32
}

type DescribeTaskLogListResponseBody struct {
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
	// Number of items per page, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeTaskLogListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeTaskLogListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskLogListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskLogListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskLogListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTaskLogListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTaskLogListResponseBody) GetResultObject() []*DescribeTaskLogListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTaskLogListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeTaskLogListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeTaskLogListResponseBody) SetRequestId(v string) *DescribeTaskLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskLogListResponseBody) SetCurrentPage(v int32) *DescribeTaskLogListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTaskLogListResponseBody) SetPageSize(v int32) *DescribeTaskLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTaskLogListResponseBody) SetResultObject(v []*DescribeTaskLogListResponseBodyResultObject) *DescribeTaskLogListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTaskLogListResponseBody) SetTotalItem(v int32) *DescribeTaskLogListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeTaskLogListResponseBody) SetTotalPage(v int32) *DescribeTaskLogListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeTaskLogListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTaskLogListResponseBodyResultObject struct {
	// Completion time, in milliseconds.
	//
	// example:
	//
	// 1753804800000
	CompletionTime *int64 `json:"completionTime,omitempty" xml:"completionTime,omitempty"`
	// Task creation time.
	//
	// example:
	//
	// 1753804800000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Log ID.
	//
	// example:
	//
	// 2793
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
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
	// coupon_abuse_detection
	SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty"`
	// Status.
	//
	// 0: Deleted
	//
	// 1: Normal
	//
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s DescribeTaskLogListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskLogListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTaskLogListResponseBodyResultObject) GetCompletionTime() *int64 {
	return s.CompletionTime
}

func (s *DescribeTaskLogListResponseBodyResultObject) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeTaskLogListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeTaskLogListResponseBodyResultObject) GetRemark() *string {
	return s.Remark
}

func (s *DescribeTaskLogListResponseBodyResultObject) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeTaskLogListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeTaskLogListResponseBodyResultObject) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskLogListResponseBodyResultObject) SetCompletionTime(v int64) *DescribeTaskLogListResponseBodyResultObject {
	s.CompletionTime = &v
	return s
}

func (s *DescribeTaskLogListResponseBodyResultObject) SetCreateTime(v int64) *DescribeTaskLogListResponseBodyResultObject {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskLogListResponseBodyResultObject) SetId(v int64) *DescribeTaskLogListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeTaskLogListResponseBodyResultObject) SetRemark(v string) *DescribeTaskLogListResponseBodyResultObject {
	s.Remark = &v
	return s
}

func (s *DescribeTaskLogListResponseBodyResultObject) SetSceneName(v string) *DescribeTaskLogListResponseBodyResultObject {
	s.SceneName = &v
	return s
}

func (s *DescribeTaskLogListResponseBodyResultObject) SetStatus(v string) *DescribeTaskLogListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeTaskLogListResponseBodyResultObject) SetTaskType(v string) *DescribeTaskLogListResponseBodyResultObject {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskLogListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

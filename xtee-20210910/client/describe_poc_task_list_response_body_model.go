// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePocTaskListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribePocTaskListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribePocTaskListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribePocTaskListResponseBodyResultObject) *DescribePocTaskListResponseBody
	GetResultObject() []*DescribePocTaskListResponseBodyResultObject
	SetTotalItem(v int32) *DescribePocTaskListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribePocTaskListResponseBody
	GetTotalPage() *int32
}

type DescribePocTaskListResponseBody struct {
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
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribePocTaskListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
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

func (s DescribePocTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePocTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePocTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePocTaskListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePocTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePocTaskListResponseBody) GetResultObject() []*DescribePocTaskListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribePocTaskListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribePocTaskListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribePocTaskListResponseBody) SetRequestId(v string) *DescribePocTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePocTaskListResponseBody) SetCurrentPage(v int32) *DescribePocTaskListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribePocTaskListResponseBody) SetPageSize(v int32) *DescribePocTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePocTaskListResponseBody) SetResultObject(v []*DescribePocTaskListResponseBodyResultObject) *DescribePocTaskListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribePocTaskListResponseBody) SetTotalItem(v int32) *DescribePocTaskListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribePocTaskListResponseBody) SetTotalPage(v int32) *DescribePocTaskListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribePocTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePocTaskListResponseBodyResultObject struct {
	// Creation time.
	//
	// example:
	//
	// 1753804800000
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Download URL.
	//
	// example:
	//
	// https://xxxxx-oss-xxxxx.xxxxxx.aliyuncs.com/xx/xx/xxx/xxxxxx.csv?Expires=1753433384&OSSAccessKeyId=xxxxxxxxx&Signature=%2F%xxxxxxxxxxxx%3D
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// File type.
	//
	// example:
	//
	// EXCEL
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// Service name.
	//
	// example:
	//
	// 注册风险识别
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// Status.
	//
	// example:
	//
	// WAIT_CHECK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 7
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// Task name.
	//
	// example:
	//
	// 任务
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// Last modified time.
	//
	// example:
	//
	// 1753804800000
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s DescribePocTaskListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribePocTaskListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribePocTaskListResponseBodyResultObject) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePocTaskListResponseBodyResultObject) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribePocTaskListResponseBodyResultObject) GetFileType() *string {
	return s.FileType
}

func (s *DescribePocTaskListResponseBodyResultObject) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribePocTaskListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribePocTaskListResponseBodyResultObject) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePocTaskListResponseBodyResultObject) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribePocTaskListResponseBodyResultObject) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribePocTaskListResponseBodyResultObject) SetCreateTime(v string) *DescribePocTaskListResponseBodyResultObject {
	s.CreateTime = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) SetDownloadUrl(v string) *DescribePocTaskListResponseBodyResultObject {
	s.DownloadUrl = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) SetFileType(v string) *DescribePocTaskListResponseBodyResultObject {
	s.FileType = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) SetServiceName(v string) *DescribePocTaskListResponseBodyResultObject {
	s.ServiceName = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) SetStatus(v string) *DescribePocTaskListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) SetTaskId(v string) *DescribePocTaskListResponseBodyResultObject {
	s.TaskId = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) SetTaskName(v string) *DescribePocTaskListResponseBodyResultObject {
	s.TaskName = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) SetUpdateTime(v string) *DescribePocTaskListResponseBodyResultObject {
	s.UpdateTime = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

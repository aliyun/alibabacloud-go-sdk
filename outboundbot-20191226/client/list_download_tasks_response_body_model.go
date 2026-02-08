// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownloadTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDownloadTasksResponseBody
	GetCode() *string
	SetDownloadTasks(v *ListDownloadTasksResponseBodyDownloadTasks) *ListDownloadTasksResponseBody
	GetDownloadTasks() *ListDownloadTasksResponseBodyDownloadTasks
	SetHttpStatusCode(v int32) *ListDownloadTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDownloadTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDownloadTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDownloadTasksResponseBody
	GetSuccess() *bool
}

type ListDownloadTasksResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of download jobs
	DownloadTasks *ListDownloadTasksResponseBodyDownloadTasks `json:"DownloadTasks,omitempty" xml:"DownloadTasks,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 904CFA7B-8AD9-50FF-9B3E-404B20B9EE31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDownloadTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDownloadTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDownloadTasksResponseBody) GetDownloadTasks() *ListDownloadTasksResponseBodyDownloadTasks {
	return s.DownloadTasks
}

func (s *ListDownloadTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDownloadTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDownloadTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDownloadTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDownloadTasksResponseBody) SetCode(v string) *ListDownloadTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListDownloadTasksResponseBody) SetDownloadTasks(v *ListDownloadTasksResponseBodyDownloadTasks) *ListDownloadTasksResponseBody {
	s.DownloadTasks = v
	return s
}

func (s *ListDownloadTasksResponseBody) SetHttpStatusCode(v int32) *ListDownloadTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDownloadTasksResponseBody) SetMessage(v string) *ListDownloadTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListDownloadTasksResponseBody) SetRequestId(v string) *ListDownloadTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDownloadTasksResponseBody) SetSuccess(v bool) *ListDownloadTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListDownloadTasksResponseBody) Validate() error {
	if s.DownloadTasks != nil {
		if err := s.DownloadTasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDownloadTasksResponseBodyDownloadTasks struct {
	// Array of download jobs
	List []*ListDownloadTasksResponseBodyDownloadTasksList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDownloadTasksResponseBodyDownloadTasks) String() string {
	return dara.Prettify(s)
}

func (s ListDownloadTasksResponseBodyDownloadTasks) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) GetList() []*ListDownloadTasksResponseBodyDownloadTasksList {
	return s.List
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) SetList(v []*ListDownloadTasksResponseBodyDownloadTasksList) *ListDownloadTasksResponseBodyDownloadTasks {
	s.List = v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) SetPageNumber(v int32) *ListDownloadTasksResponseBodyDownloadTasks {
	s.PageNumber = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) SetPageSize(v int32) *ListDownloadTasksResponseBodyDownloadTasks {
	s.PageSize = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) SetTotalCount(v int32) *ListDownloadTasksResponseBodyDownloadTasks {
	s.TotalCount = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDownloadTasksResponseBodyDownloadTasksList struct {
	// List of download task files
	DownloadTaskFiles []*ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles `json:"DownloadTaskFiles,omitempty" xml:"DownloadTaskFiles,omitempty" type:"Repeated"`
	// Expiration time
	//
	// example:
	//
	// 1646792941
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Export status.
	//
	// - Pending: Pending
	//
	// - InProgress: In progress
	//
	// - Finished: Finished
	//
	// - Failed: Failed
	//
	// - Expired: Expired
	//
	// - Cancelled: Cancelled
	//
	// example:
	//
	// Empty
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Job ID
	//
	// example:
	//
	// 6b0e547e-501c-480a-812f-d27e28e74f9a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Title
	//
	// example:
	//
	// GUI-TEST-新建风险
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDownloadTasksResponseBodyDownloadTasksList) String() string {
	return dara.Prettify(s)
}

func (s ListDownloadTasksResponseBodyDownloadTasksList) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) GetDownloadTaskFiles() []*ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	return s.DownloadTaskFiles
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) GetStatus() *string {
	return s.Status
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) GetTitle() *string {
	return s.Title
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetDownloadTaskFiles(v []*ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.DownloadTaskFiles = v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetExpireTime(v int64) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.ExpireTime = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetStatus(v string) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.Status = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetTaskId(v string) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.TaskId = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetTitle(v string) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.Title = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) Validate() error {
	if s.DownloadTaskFiles != nil {
		for _, item := range s.DownloadTaskFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles struct {
	// File ID
	//
	// example:
	//
	// 01f8d11abdb9406996c0c46e7*******
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Completion progress [Deprecated]
	//
	// example:
	//
	// 10
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// File status
	//
	// - Pending: Pending
	//
	// - InProgress: In progress
	//
	// - Finished: Finished
	//
	// - Failed: Failed
	//
	// - Expired: Expired
	//
	// - Cancelled: Cancelled
	//
	// example:
	//
	// Empty
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Title
	//
	// example:
	//
	// GUI-TEST-新建风险
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) String() string {
	return dara.Prettify(s)
}

func (s ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) GetFileId() *string {
	return s.FileId
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) GetProgress() *int32 {
	return s.Progress
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) GetStatus() *string {
	return s.Status
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) GetTitle() *string {
	return s.Title
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) SetFileId(v string) *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	s.FileId = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) SetProgress(v int32) *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	s.Progress = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) SetStatus(v string) *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	s.Status = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) SetTitle(v string) *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	s.Title = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledPreloadJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *GetScheduledPreloadJobResponseBody
	GetAliUid() *string
	SetCreatedAt(v string) *GetScheduledPreloadJobResponseBody
	GetCreatedAt() *string
	SetDomains(v string) *GetScheduledPreloadJobResponseBody
	GetDomains() *string
	SetErrorInfo(v string) *GetScheduledPreloadJobResponseBody
	GetErrorInfo() *string
	SetFailedFileOss(v string) *GetScheduledPreloadJobResponseBody
	GetFailedFileOss() *string
	SetFileId(v string) *GetScheduledPreloadJobResponseBody
	GetFileId() *string
	SetId(v string) *GetScheduledPreloadJobResponseBody
	GetId() *string
	SetInsertWay(v string) *GetScheduledPreloadJobResponseBody
	GetInsertWay() *string
	SetName(v string) *GetScheduledPreloadJobResponseBody
	GetName() *string
	SetRequestId(v string) *GetScheduledPreloadJobResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetScheduledPreloadJobResponseBody
	GetSiteId() *int64
	SetTaskSubmitted(v int32) *GetScheduledPreloadJobResponseBody
	GetTaskSubmitted() *int32
	SetTaskType(v string) *GetScheduledPreloadJobResponseBody
	GetTaskType() *string
	SetUrlCount(v int32) *GetScheduledPreloadJobResponseBody
	GetUrlCount() *int32
	SetUrlSubmitted(v int32) *GetScheduledPreloadJobResponseBody
	GetUrlSubmitted() *int32
}

type GetScheduledPreloadJobResponseBody struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 15685865xxx14622
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2024-06-02T02:23:26Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The domain names to be prefetched.
	//
	// example:
	//
	// testurl.com
	Domains *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// invalid domain:test.com
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The URL of the OSS object that stores a list of URLs that failed the conditional check for prefetching.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	FailedFileOss *string `json:"FailedFileOss,omitempty" xml:"FailedFileOss,omitempty"`
	// The ID of the URL list file, which can be used during downloads.
	//
	// example:
	//
	// 665d3b48621bccf3fe29e1a7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the prefetch task.
	//
	// example:
	//
	// 665d3af3621bccf3fe29e1a4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The method to submit the URLs to be prefetched.
	//
	// example:
	//
	// oss
	InsertWay *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	// The task name.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 190007158391808
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The number of submitted prefetch tasks.
	//
	// example:
	//
	// 1
	TaskSubmitted *int32 `json:"TaskSubmitted,omitempty" xml:"TaskSubmitted,omitempty"`
	// The task type. Valid values: refresh and preload.
	//
	// example:
	//
	// preload
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of URLs.
	//
	// example:
	//
	// 2
	UrlCount *int32 `json:"UrlCount,omitempty" xml:"UrlCount,omitempty"`
	// The number of submitted URLs.
	//
	// example:
	//
	// 1
	UrlSubmitted *int32 `json:"UrlSubmitted,omitempty" xml:"UrlSubmitted,omitempty"`
}

func (s GetScheduledPreloadJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledPreloadJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledPreloadJobResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *GetScheduledPreloadJobResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetScheduledPreloadJobResponseBody) GetDomains() *string {
	return s.Domains
}

func (s *GetScheduledPreloadJobResponseBody) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *GetScheduledPreloadJobResponseBody) GetFailedFileOss() *string {
	return s.FailedFileOss
}

func (s *GetScheduledPreloadJobResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *GetScheduledPreloadJobResponseBody) GetId() *string {
	return s.Id
}

func (s *GetScheduledPreloadJobResponseBody) GetInsertWay() *string {
	return s.InsertWay
}

func (s *GetScheduledPreloadJobResponseBody) GetName() *string {
	return s.Name
}

func (s *GetScheduledPreloadJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledPreloadJobResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetScheduledPreloadJobResponseBody) GetTaskSubmitted() *int32 {
	return s.TaskSubmitted
}

func (s *GetScheduledPreloadJobResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetScheduledPreloadJobResponseBody) GetUrlCount() *int32 {
	return s.UrlCount
}

func (s *GetScheduledPreloadJobResponseBody) GetUrlSubmitted() *int32 {
	return s.UrlSubmitted
}

func (s *GetScheduledPreloadJobResponseBody) SetAliUid(v string) *GetScheduledPreloadJobResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetCreatedAt(v string) *GetScheduledPreloadJobResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetDomains(v string) *GetScheduledPreloadJobResponseBody {
	s.Domains = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetErrorInfo(v string) *GetScheduledPreloadJobResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetFailedFileOss(v string) *GetScheduledPreloadJobResponseBody {
	s.FailedFileOss = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetFileId(v string) *GetScheduledPreloadJobResponseBody {
	s.FileId = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetId(v string) *GetScheduledPreloadJobResponseBody {
	s.Id = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetInsertWay(v string) *GetScheduledPreloadJobResponseBody {
	s.InsertWay = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetName(v string) *GetScheduledPreloadJobResponseBody {
	s.Name = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetRequestId(v string) *GetScheduledPreloadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetSiteId(v int64) *GetScheduledPreloadJobResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetTaskSubmitted(v int32) *GetScheduledPreloadJobResponseBody {
	s.TaskSubmitted = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetTaskType(v string) *GetScheduledPreloadJobResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetUrlCount(v int32) *GetScheduledPreloadJobResponseBody {
	s.UrlCount = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetUrlSubmitted(v int32) *GetScheduledPreloadJobResponseBody {
	s.UrlSubmitted = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) Validate() error {
	return dara.Validate(s)
}

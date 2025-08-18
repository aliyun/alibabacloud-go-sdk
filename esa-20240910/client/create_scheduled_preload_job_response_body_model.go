// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPreloadJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *CreateScheduledPreloadJobResponseBody
	GetAliUid() *string
	SetCreatedAt(v string) *CreateScheduledPreloadJobResponseBody
	GetCreatedAt() *string
	SetDomains(v string) *CreateScheduledPreloadJobResponseBody
	GetDomains() *string
	SetErrorInfo(v string) *CreateScheduledPreloadJobResponseBody
	GetErrorInfo() *string
	SetFailedFileOss(v string) *CreateScheduledPreloadJobResponseBody
	GetFailedFileOss() *string
	SetFileId(v string) *CreateScheduledPreloadJobResponseBody
	GetFileId() *string
	SetId(v string) *CreateScheduledPreloadJobResponseBody
	GetId() *string
	SetInsertWay(v string) *CreateScheduledPreloadJobResponseBody
	GetInsertWay() *string
	SetName(v string) *CreateScheduledPreloadJobResponseBody
	GetName() *string
	SetRequestId(v string) *CreateScheduledPreloadJobResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *CreateScheduledPreloadJobResponseBody
	GetSiteId() *int64
	SetTaskSubmitted(v int32) *CreateScheduledPreloadJobResponseBody
	GetTaskSubmitted() *int32
	SetTaskType(v string) *CreateScheduledPreloadJobResponseBody
	GetTaskType() *string
	SetUrlCount(v int32) *CreateScheduledPreloadJobResponseBody
	GetUrlCount() *int32
	SetUrlSubmitted(v int32) *CreateScheduledPreloadJobResponseBody
	GetUrlSubmitted() *int32
}

type CreateScheduledPreloadJobResponseBody struct {
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
	// 2023-06-05T10:04:20+0800
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The domain names to be prefetched.
	//
	// example:
	//
	// testurl.com
	Domains *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// The error message. Multiple error messages are separated by commas (,). Valid values:
	//
	// 	- **InvalidUrl**: The URL format is invalid.
	//
	// 	- **InvalidDomain**: The domain name fails the domain ownership verification.
	//
	// 	- **QuotaExcess**: the quota limit has been reached.
	//
	// 	- **OtherErrors**: other errors.
	//
	// example:
	//
	// InvalidDomain
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
	// The ID of the scheduled prefetch task.
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
	// 15C66C7B-671A-4297-9187-2C4477247B78
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
	// The task type (refresh or preload).
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

func (s CreateScheduledPreloadJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadJobResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *CreateScheduledPreloadJobResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateScheduledPreloadJobResponseBody) GetDomains() *string {
	return s.Domains
}

func (s *CreateScheduledPreloadJobResponseBody) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *CreateScheduledPreloadJobResponseBody) GetFailedFileOss() *string {
	return s.FailedFileOss
}

func (s *CreateScheduledPreloadJobResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *CreateScheduledPreloadJobResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateScheduledPreloadJobResponseBody) GetInsertWay() *string {
	return s.InsertWay
}

func (s *CreateScheduledPreloadJobResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateScheduledPreloadJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledPreloadJobResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateScheduledPreloadJobResponseBody) GetTaskSubmitted() *int32 {
	return s.TaskSubmitted
}

func (s *CreateScheduledPreloadJobResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateScheduledPreloadJobResponseBody) GetUrlCount() *int32 {
	return s.UrlCount
}

func (s *CreateScheduledPreloadJobResponseBody) GetUrlSubmitted() *int32 {
	return s.UrlSubmitted
}

func (s *CreateScheduledPreloadJobResponseBody) SetAliUid(v string) *CreateScheduledPreloadJobResponseBody {
	s.AliUid = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetCreatedAt(v string) *CreateScheduledPreloadJobResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetDomains(v string) *CreateScheduledPreloadJobResponseBody {
	s.Domains = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetErrorInfo(v string) *CreateScheduledPreloadJobResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetFailedFileOss(v string) *CreateScheduledPreloadJobResponseBody {
	s.FailedFileOss = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetFileId(v string) *CreateScheduledPreloadJobResponseBody {
	s.FileId = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetId(v string) *CreateScheduledPreloadJobResponseBody {
	s.Id = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetInsertWay(v string) *CreateScheduledPreloadJobResponseBody {
	s.InsertWay = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetName(v string) *CreateScheduledPreloadJobResponseBody {
	s.Name = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetRequestId(v string) *CreateScheduledPreloadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetSiteId(v int64) *CreateScheduledPreloadJobResponseBody {
	s.SiteId = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetTaskSubmitted(v int32) *CreateScheduledPreloadJobResponseBody {
	s.TaskSubmitted = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetTaskType(v string) *CreateScheduledPreloadJobResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetUrlCount(v int32) *CreateScheduledPreloadJobResponseBody {
	s.UrlCount = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetUrlSubmitted(v int32) *CreateScheduledPreloadJobResponseBody {
	s.UrlSubmitted = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) Validate() error {
	return dara.Validate(s)
}

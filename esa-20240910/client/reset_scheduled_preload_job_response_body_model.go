// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetScheduledPreloadJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *ResetScheduledPreloadJobResponseBody
	GetAliUid() *string
	SetCreatedAt(v string) *ResetScheduledPreloadJobResponseBody
	GetCreatedAt() *string
	SetDomains(v string) *ResetScheduledPreloadJobResponseBody
	GetDomains() *string
	SetErrorInfo(v string) *ResetScheduledPreloadJobResponseBody
	GetErrorInfo() *string
	SetFailedFileOss(v string) *ResetScheduledPreloadJobResponseBody
	GetFailedFileOss() *string
	SetFileId(v string) *ResetScheduledPreloadJobResponseBody
	GetFileId() *string
	SetId(v string) *ResetScheduledPreloadJobResponseBody
	GetId() *string
	SetInsertWay(v string) *ResetScheduledPreloadJobResponseBody
	GetInsertWay() *string
	SetName(v string) *ResetScheduledPreloadJobResponseBody
	GetName() *string
	SetRequestId(v string) *ResetScheduledPreloadJobResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *ResetScheduledPreloadJobResponseBody
	GetSiteId() *int64
	SetTaskSubmitted(v int32) *ResetScheduledPreloadJobResponseBody
	GetTaskSubmitted() *int32
	SetTaskType(v string) *ResetScheduledPreloadJobResponseBody
	GetTaskType() *string
	SetUrlCount(v int32) *ResetScheduledPreloadJobResponseBody
	GetUrlCount() *int32
	SetUrlSubmitted(v int32) *ResetScheduledPreloadJobResponseBody
	GetUrlSubmitted() *int32
}

type ResetScheduledPreloadJobResponseBody struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 15685865xxx14622
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when the job was created, in ISO 8601 format (for example, 2024-01-01T00:00:00+Z).
	//
	// example:
	//
	// 2024-06-02T02:23:26Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The list of preload domains.
	//
	// example:
	//
	// testurl.com
	Domains *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// The error message.
	//
	// example:
	//
	// invalid domain:test.com
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The OSS URL of the failed file.
	//
	// example:
	//
	// https://xxxobject.oss-cn-region.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	FailedFileOss *string `json:"FailedFileOss,omitempty" xml:"FailedFileOss,omitempty"`
	// The file ID of the URL list, which is used for downloading.
	//
	// example:
	//
	// 665d3b48621bccf3fe29e1a7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 665d3af3621bccf3fe29e1a4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The URL insertion method.
	//
	// example:
	//
	// oss
	InsertWay *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	// The job name.
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
	// The site ID.
	//
	// example:
	//
	// 190007158391808
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The number of URLs that have been submitted to the system for preloading.
	//
	// example:
	//
	// 1
	TaskSubmitted *int32 `json:"TaskSubmitted,omitempty" xml:"TaskSubmitted,omitempty"`
	// The task type (refresh/preload).
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

func (s ResetScheduledPreloadJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetScheduledPreloadJobResponseBody) GoString() string {
	return s.String()
}

func (s *ResetScheduledPreloadJobResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *ResetScheduledPreloadJobResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ResetScheduledPreloadJobResponseBody) GetDomains() *string {
	return s.Domains
}

func (s *ResetScheduledPreloadJobResponseBody) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *ResetScheduledPreloadJobResponseBody) GetFailedFileOss() *string {
	return s.FailedFileOss
}

func (s *ResetScheduledPreloadJobResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *ResetScheduledPreloadJobResponseBody) GetId() *string {
	return s.Id
}

func (s *ResetScheduledPreloadJobResponseBody) GetInsertWay() *string {
	return s.InsertWay
}

func (s *ResetScheduledPreloadJobResponseBody) GetName() *string {
	return s.Name
}

func (s *ResetScheduledPreloadJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetScheduledPreloadJobResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ResetScheduledPreloadJobResponseBody) GetTaskSubmitted() *int32 {
	return s.TaskSubmitted
}

func (s *ResetScheduledPreloadJobResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *ResetScheduledPreloadJobResponseBody) GetUrlCount() *int32 {
	return s.UrlCount
}

func (s *ResetScheduledPreloadJobResponseBody) GetUrlSubmitted() *int32 {
	return s.UrlSubmitted
}

func (s *ResetScheduledPreloadJobResponseBody) SetAliUid(v string) *ResetScheduledPreloadJobResponseBody {
	s.AliUid = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetCreatedAt(v string) *ResetScheduledPreloadJobResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetDomains(v string) *ResetScheduledPreloadJobResponseBody {
	s.Domains = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetErrorInfo(v string) *ResetScheduledPreloadJobResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetFailedFileOss(v string) *ResetScheduledPreloadJobResponseBody {
	s.FailedFileOss = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetFileId(v string) *ResetScheduledPreloadJobResponseBody {
	s.FileId = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetId(v string) *ResetScheduledPreloadJobResponseBody {
	s.Id = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetInsertWay(v string) *ResetScheduledPreloadJobResponseBody {
	s.InsertWay = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetName(v string) *ResetScheduledPreloadJobResponseBody {
	s.Name = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetRequestId(v string) *ResetScheduledPreloadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetSiteId(v int64) *ResetScheduledPreloadJobResponseBody {
	s.SiteId = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetTaskSubmitted(v int32) *ResetScheduledPreloadJobResponseBody {
	s.TaskSubmitted = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetTaskType(v string) *ResetScheduledPreloadJobResponseBody {
	s.TaskType = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetUrlCount(v int32) *ResetScheduledPreloadJobResponseBody {
	s.UrlCount = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetUrlSubmitted(v int32) *ResetScheduledPreloadJobResponseBody {
	s.UrlSubmitted = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) Validate() error {
	return dara.Validate(s)
}

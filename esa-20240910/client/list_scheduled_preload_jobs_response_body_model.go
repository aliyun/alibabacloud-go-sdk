// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPreloadJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListScheduledPreloadJobsResponseBodyJobs) *ListScheduledPreloadJobsResponseBody
	GetJobs() []*ListScheduledPreloadJobsResponseBodyJobs
	SetRequestId(v string) *ListScheduledPreloadJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListScheduledPreloadJobsResponseBody
	GetTotalCount() *string
}

type ListScheduledPreloadJobsResponseBody struct {
	Jobs       []*ListScheduledPreloadJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScheduledPreloadJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPreloadJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadJobsResponseBody) GetJobs() []*ListScheduledPreloadJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListScheduledPreloadJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduledPreloadJobsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListScheduledPreloadJobsResponseBody) SetJobs(v []*ListScheduledPreloadJobsResponseBodyJobs) *ListScheduledPreloadJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListScheduledPreloadJobsResponseBody) SetRequestId(v string) *ListScheduledPreloadJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBody) SetTotalCount(v string) *ListScheduledPreloadJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScheduledPreloadJobsResponseBodyJobs struct {
	AliUid         *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreatedAt      *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Domains        *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	ErrorInfo      *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ExecutionCount *int32  `json:"ExecutionCount,omitempty" xml:"ExecutionCount,omitempty"`
	FailedFileOss  *string `json:"FailedFileOss,omitempty" xml:"FailedFileOss,omitempty"`
	FileId         *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InsertWay      *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SiteId         *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TaskSubmitted  *int32  `json:"TaskSubmitted,omitempty" xml:"TaskSubmitted,omitempty"`
	TaskType       *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UrlCount       *int32  `json:"UrlCount,omitempty" xml:"UrlCount,omitempty"`
	UrlSubmitted   *int32  `json:"UrlSubmitted,omitempty" xml:"UrlSubmitted,omitempty"`
}

func (s ListScheduledPreloadJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPreloadJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetAliUid() *string {
	return s.AliUid
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetDomains() *string {
	return s.Domains
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetExecutionCount() *int32 {
	return s.ExecutionCount
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetFailedFileOss() *string {
	return s.FailedFileOss
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetFileId() *string {
	return s.FileId
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetId() *string {
	return s.Id
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetInsertWay() *string {
	return s.InsertWay
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetName() *string {
	return s.Name
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetTaskSubmitted() *int32 {
	return s.TaskSubmitted
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetTaskType() *string {
	return s.TaskType
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetUrlCount() *int32 {
	return s.UrlCount
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) GetUrlSubmitted() *int32 {
	return s.UrlSubmitted
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetAliUid(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.AliUid = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetCreatedAt(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.CreatedAt = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetDomains(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.Domains = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetErrorInfo(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.ErrorInfo = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetExecutionCount(v int32) *ListScheduledPreloadJobsResponseBodyJobs {
	s.ExecutionCount = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetFailedFileOss(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.FailedFileOss = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetFileId(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.FileId = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetId(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetInsertWay(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.InsertWay = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetName(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetSiteId(v int64) *ListScheduledPreloadJobsResponseBodyJobs {
	s.SiteId = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetTaskSubmitted(v int32) *ListScheduledPreloadJobsResponseBodyJobs {
	s.TaskSubmitted = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetTaskType(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.TaskType = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetUrlCount(v int32) *ListScheduledPreloadJobsResponseBodyJobs {
	s.UrlCount = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetUrlSubmitted(v int32) *ListScheduledPreloadJobsResponseBodyJobs {
	s.UrlSubmitted = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) Validate() error {
	return dara.Validate(s)
}

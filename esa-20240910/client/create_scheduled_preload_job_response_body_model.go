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
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreatedAt     *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Domains       *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	ErrorInfo     *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	FailedFileOss *string `json:"FailedFileOss,omitempty" xml:"FailedFileOss,omitempty"`
	FileId        *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InsertWay     *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId        *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TaskSubmitted *int32  `json:"TaskSubmitted,omitempty" xml:"TaskSubmitted,omitempty"`
	TaskType      *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UrlCount      *int32  `json:"UrlCount,omitempty" xml:"UrlCount,omitempty"`
	UrlSubmitted  *int32  `json:"UrlSubmitted,omitempty" xml:"UrlSubmitted,omitempty"`
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

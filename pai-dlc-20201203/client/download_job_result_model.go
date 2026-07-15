// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadJobResult interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *DownloadJobResult
	GetDisplayName() *string
	SetDownloadJobId(v string) *DownloadJobResult
	GetDownloadJobId() *string
	SetDownloadUrl(v string) *DownloadJobResult
	GetDownloadUrl() *string
	SetEndTime(v string) *DownloadJobResult
	GetEndTime() *string
	SetFileType(v string) *DownloadJobResult
	GetFileType() *string
	SetGmtCreated(v string) *DownloadJobResult
	GetGmtCreated() *string
	SetGmtModified(v string) *DownloadJobResult
	GetGmtModified() *string
	SetLogCount(v int32) *DownloadJobResult
	GetLogCount() *int32
	SetPodIds(v []*string) *DownloadJobResult
	GetPodIds() []*string
	SetPodUids(v []*string) *DownloadJobResult
	GetPodUids() []*string
	SetSourceJobId(v string) *DownloadJobResult
	GetSourceJobId() *string
	SetStartTime(v string) *DownloadJobResult
	GetStartTime() *string
	SetStatus(v string) *DownloadJobResult
	GetStatus() *string
	SetTenantId(v string) *DownloadJobResult
	GetTenantId() *string
	SetType(v string) *DownloadJobResult
	GetType() *string
	SetUrlExpireTime(v string) *DownloadJobResult
	GetUrlExpireTime() *string
	SetUserId(v string) *DownloadJobResult
	GetUserId() *string
	SetWorkspaceId(v string) *DownloadJobResult
	GetWorkspaceId() *string
}

type DownloadJobResult struct {
	// The display name of the download job.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the download job.
	DownloadJobId *string `json:"DownloadJobId,omitempty" xml:"DownloadJobId,omitempty"`
	// A temporary, pre-signed URL for downloading the result file.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The completion time of the download job, in UTC format.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of file to download, such as `logs` or `output`.
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The creation time of the download job, in UTC format.
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The last modification time of the download job, in UTC format.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The number of log entries included in the download.
	LogCount *int32 `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	// The IDs of the pods for the source job.
	PodIds []*string `json:"PodIds,omitempty" xml:"PodIds,omitempty" type:"Repeated"`
	// The UIDs of the pods for the source job.
	PodUids []*string `json:"PodUids,omitempty" xml:"PodUids,omitempty" type:"Repeated"`
	// The ID of the source job whose results are downloaded.
	SourceJobId *string `json:"SourceJobId,omitempty" xml:"SourceJobId,omitempty"`
	// The start time of the download job, in UTC format.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the download job. Valid values: `Running`, `Succeeded`, and `Failed`.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant that the job belongs to.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The type of the download job.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the download URL expires, in UTC format.
	UrlExpireTime *string `json:"UrlExpireTime,omitempty" xml:"UrlExpireTime,omitempty"`
	// The ID of the user who created the job.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the workspace where the job was created.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DownloadJobResult) String() string {
	return dara.Prettify(s)
}

func (s DownloadJobResult) GoString() string {
	return s.String()
}

func (s *DownloadJobResult) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DownloadJobResult) GetDownloadJobId() *string {
	return s.DownloadJobId
}

func (s *DownloadJobResult) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DownloadJobResult) GetEndTime() *string {
	return s.EndTime
}

func (s *DownloadJobResult) GetFileType() *string {
	return s.FileType
}

func (s *DownloadJobResult) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DownloadJobResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DownloadJobResult) GetLogCount() *int32 {
	return s.LogCount
}

func (s *DownloadJobResult) GetPodIds() []*string {
	return s.PodIds
}

func (s *DownloadJobResult) GetPodUids() []*string {
	return s.PodUids
}

func (s *DownloadJobResult) GetSourceJobId() *string {
	return s.SourceJobId
}

func (s *DownloadJobResult) GetStartTime() *string {
	return s.StartTime
}

func (s *DownloadJobResult) GetStatus() *string {
	return s.Status
}

func (s *DownloadJobResult) GetTenantId() *string {
	return s.TenantId
}

func (s *DownloadJobResult) GetType() *string {
	return s.Type
}

func (s *DownloadJobResult) GetUrlExpireTime() *string {
	return s.UrlExpireTime
}

func (s *DownloadJobResult) GetUserId() *string {
	return s.UserId
}

func (s *DownloadJobResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DownloadJobResult) SetDisplayName(v string) *DownloadJobResult {
	s.DisplayName = &v
	return s
}

func (s *DownloadJobResult) SetDownloadJobId(v string) *DownloadJobResult {
	s.DownloadJobId = &v
	return s
}

func (s *DownloadJobResult) SetDownloadUrl(v string) *DownloadJobResult {
	s.DownloadUrl = &v
	return s
}

func (s *DownloadJobResult) SetEndTime(v string) *DownloadJobResult {
	s.EndTime = &v
	return s
}

func (s *DownloadJobResult) SetFileType(v string) *DownloadJobResult {
	s.FileType = &v
	return s
}

func (s *DownloadJobResult) SetGmtCreated(v string) *DownloadJobResult {
	s.GmtCreated = &v
	return s
}

func (s *DownloadJobResult) SetGmtModified(v string) *DownloadJobResult {
	s.GmtModified = &v
	return s
}

func (s *DownloadJobResult) SetLogCount(v int32) *DownloadJobResult {
	s.LogCount = &v
	return s
}

func (s *DownloadJobResult) SetPodIds(v []*string) *DownloadJobResult {
	s.PodIds = v
	return s
}

func (s *DownloadJobResult) SetPodUids(v []*string) *DownloadJobResult {
	s.PodUids = v
	return s
}

func (s *DownloadJobResult) SetSourceJobId(v string) *DownloadJobResult {
	s.SourceJobId = &v
	return s
}

func (s *DownloadJobResult) SetStartTime(v string) *DownloadJobResult {
	s.StartTime = &v
	return s
}

func (s *DownloadJobResult) SetStatus(v string) *DownloadJobResult {
	s.Status = &v
	return s
}

func (s *DownloadJobResult) SetTenantId(v string) *DownloadJobResult {
	s.TenantId = &v
	return s
}

func (s *DownloadJobResult) SetType(v string) *DownloadJobResult {
	s.Type = &v
	return s
}

func (s *DownloadJobResult) SetUrlExpireTime(v string) *DownloadJobResult {
	s.UrlExpireTime = &v
	return s
}

func (s *DownloadJobResult) SetUserId(v string) *DownloadJobResult {
	s.UserId = &v
	return s
}

func (s *DownloadJobResult) SetWorkspaceId(v string) *DownloadJobResult {
	s.WorkspaceId = &v
	return s
}

func (s *DownloadJobResult) Validate() error {
	return dara.Validate(s)
}

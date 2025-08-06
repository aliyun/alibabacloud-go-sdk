// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaExportJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedJobIds(v []*string) *GetMediaExportJobsResponseBody
	GetFailedJobIds() []*string
	SetMediaExportJobList(v []*GetMediaExportJobsResponseBodyMediaExportJobList) *GetMediaExportJobsResponseBody
	GetMediaExportJobList() []*GetMediaExportJobsResponseBodyMediaExportJobList
	SetNonExistJobIds(v []*string) *GetMediaExportJobsResponseBody
	GetNonExistJobIds() []*string
	SetRequestId(v string) *GetMediaExportJobsResponseBody
	GetRequestId() *string
}

type GetMediaExportJobsResponseBody struct {
	FailedJobIds       []*string                                           `json:"FailedJobIds,omitempty" xml:"FailedJobIds,omitempty" type:"Repeated"`
	MediaExportJobList []*GetMediaExportJobsResponseBodyMediaExportJobList `json:"MediaExportJobList,omitempty" xml:"MediaExportJobList,omitempty" type:"Repeated"`
	NonExistJobIds     []*string                                           `json:"NonExistJobIds,omitempty" xml:"NonExistJobIds,omitempty" type:"Repeated"`
	RequestId          *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaExportJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaExportJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaExportJobsResponseBody) GetFailedJobIds() []*string {
	return s.FailedJobIds
}

func (s *GetMediaExportJobsResponseBody) GetMediaExportJobList() []*GetMediaExportJobsResponseBodyMediaExportJobList {
	return s.MediaExportJobList
}

func (s *GetMediaExportJobsResponseBody) GetNonExistJobIds() []*string {
	return s.NonExistJobIds
}

func (s *GetMediaExportJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaExportJobsResponseBody) SetFailedJobIds(v []*string) *GetMediaExportJobsResponseBody {
	s.FailedJobIds = v
	return s
}

func (s *GetMediaExportJobsResponseBody) SetMediaExportJobList(v []*GetMediaExportJobsResponseBodyMediaExportJobList) *GetMediaExportJobsResponseBody {
	s.MediaExportJobList = v
	return s
}

func (s *GetMediaExportJobsResponseBody) SetNonExistJobIds(v []*string) *GetMediaExportJobsResponseBody {
	s.NonExistJobIds = v
	return s
}

func (s *GetMediaExportJobsResponseBody) SetRequestId(v string) *GetMediaExportJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaExportJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaExportJobsResponseBodyMediaExportJobList struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FileURL          *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName          *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Match            *string `json:"Match,omitempty" xml:"Match,omitempty"`
	MediaType        *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	SortBy           *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMediaExportJobsResponseBodyMediaExportJobList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaExportJobsResponseBodyMediaExportJobList) GoString() string {
	return s.String()
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetAppId() *string {
	return s.AppId
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetFileURL() *string {
	return s.FileURL
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetJobName() *string {
	return s.JobName
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetMatch() *string {
	return s.Match
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetMediaType() *string {
	return s.MediaType
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetSortBy() *string {
	return s.SortBy
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) GetStatus() *string {
	return s.Status
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetAppId(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.AppId = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetCreationTime(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.CreationTime = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetErrorCode(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.ErrorCode = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetErrorMessage(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.ErrorMessage = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetFileURL(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.FileURL = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetJobId(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.JobId = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetJobName(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.JobName = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetMatch(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.Match = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetMediaType(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.MediaType = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetModificationTime(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.ModificationTime = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetSortBy(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.SortBy = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) SetStatus(v string) *GetMediaExportJobsResponseBodyMediaExportJobList {
	s.Status = &v
	return s
}

func (s *GetMediaExportJobsResponseBodyMediaExportJobList) Validate() error {
	return dara.Validate(s)
}

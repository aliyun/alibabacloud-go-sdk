// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaExportJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaExportJobList(v []*ListMediaExportJobsResponseBodyMediaExportJobList) *ListMediaExportJobsResponseBody
	GetMediaExportJobList() []*ListMediaExportJobsResponseBodyMediaExportJobList
	SetRequestId(v string) *ListMediaExportJobsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListMediaExportJobsResponseBody
	GetTotal() *int64
}

type ListMediaExportJobsResponseBody struct {
	MediaExportJobList []*ListMediaExportJobsResponseBodyMediaExportJobList `json:"MediaExportJobList,omitempty" xml:"MediaExportJobList,omitempty" type:"Repeated"`
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total              *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMediaExportJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaExportJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaExportJobsResponseBody) GetMediaExportJobList() []*ListMediaExportJobsResponseBodyMediaExportJobList {
	return s.MediaExportJobList
}

func (s *ListMediaExportJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaExportJobsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListMediaExportJobsResponseBody) SetMediaExportJobList(v []*ListMediaExportJobsResponseBodyMediaExportJobList) *ListMediaExportJobsResponseBody {
	s.MediaExportJobList = v
	return s
}

func (s *ListMediaExportJobsResponseBody) SetRequestId(v string) *ListMediaExportJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaExportJobsResponseBody) SetTotal(v int64) *ListMediaExportJobsResponseBody {
	s.Total = &v
	return s
}

func (s *ListMediaExportJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMediaExportJobsResponseBodyMediaExportJobList struct {
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

func (s ListMediaExportJobsResponseBodyMediaExportJobList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaExportJobsResponseBodyMediaExportJobList) GoString() string {
	return s.String()
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetAppId() *string {
	return s.AppId
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetFileURL() *string {
	return s.FileURL
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetJobName() *string {
	return s.JobName
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetMatch() *string {
	return s.Match
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetMediaType() *string {
	return s.MediaType
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetSortBy() *string {
	return s.SortBy
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) GetStatus() *string {
	return s.Status
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetAppId(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.AppId = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetCreationTime(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.CreationTime = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetErrorCode(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.ErrorCode = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetErrorMessage(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.ErrorMessage = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetFileURL(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.FileURL = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetJobId(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.JobId = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetJobName(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.JobName = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetMatch(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.Match = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetMediaType(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.MediaType = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetModificationTime(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.ModificationTime = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetSortBy(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.SortBy = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) SetStatus(v string) *ListMediaExportJobsResponseBodyMediaExportJobList {
	s.Status = &v
	return s
}

func (s *ListMediaExportJobsResponseBodyMediaExportJobList) Validate() error {
	return dara.Validate(s)
}

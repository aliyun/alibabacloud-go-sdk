// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCategoryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCategoryJobList(v *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList) *ListAIVideoCategoryJobResponseBody
	GetAIVideoCategoryJobList() *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList
	SetNonExistAIVideoCategoryJobIds(v *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds) *ListAIVideoCategoryJobResponseBody
	GetNonExistAIVideoCategoryJobIds() *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds
	SetRequestId(v string) *ListAIVideoCategoryJobResponseBody
	GetRequestId() *string
}

type ListAIVideoCategoryJobResponseBody struct {
	AIVideoCategoryJobList        *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList        `json:"AIVideoCategoryJobList,omitempty" xml:"AIVideoCategoryJobList,omitempty" type:"Struct"`
	NonExistAIVideoCategoryJobIds *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds `json:"NonExistAIVideoCategoryJobIds,omitempty" xml:"NonExistAIVideoCategoryJobIds,omitempty" type:"Struct"`
	RequestId                     *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIVideoCategoryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCategoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIVideoCategoryJobResponseBody) GetAIVideoCategoryJobList() *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList {
	return s.AIVideoCategoryJobList
}

func (s *ListAIVideoCategoryJobResponseBody) GetNonExistAIVideoCategoryJobIds() *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds {
	return s.NonExistAIVideoCategoryJobIds
}

func (s *ListAIVideoCategoryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIVideoCategoryJobResponseBody) SetAIVideoCategoryJobList(v *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList) *ListAIVideoCategoryJobResponseBody {
	s.AIVideoCategoryJobList = v
	return s
}

func (s *ListAIVideoCategoryJobResponseBody) SetNonExistAIVideoCategoryJobIds(v *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds) *ListAIVideoCategoryJobResponseBody {
	s.NonExistAIVideoCategoryJobIds = v
	return s
}

func (s *ListAIVideoCategoryJobResponseBody) SetRequestId(v string) *ListAIVideoCategoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIVideoCategoryJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList struct {
	AIVideoCategoryJob []*ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob `json:"AIVideoCategoryJob,omitempty" xml:"AIVideoCategoryJob,omitempty" type:"Repeated"`
}

func (s ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList) GoString() string {
	return s.String()
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList) GetAIVideoCategoryJob() []*ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob {
	return s.AIVideoCategoryJob
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList) SetAIVideoCategoryJob(v []*ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList {
	s.AIVideoCategoryJob = v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) GoString() string {
	return s.String()
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) GetCode() *string {
	return s.Code
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) GetData() *string {
	return s.Data
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) SetCode(v string) *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob {
	s.Code = &v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) SetCreationTime(v string) *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) SetData(v string) *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob {
	s.Data = &v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) SetJobId(v string) *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob {
	s.JobId = &v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) SetMediaId(v string) *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob {
	s.MediaId = &v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) SetMessage(v string) *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob {
	s.Message = &v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) SetStatus(v string) *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob {
	s.Status = &v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyAIVideoCategoryJobListAIVideoCategoryJob) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds) GoString() string {
	return s.String()
}

func (s *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds) SetString_(v []*string) *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds {
	s.String_ = v
	return s
}

func (s *ListAIVideoCategoryJobResponseBodyNonExistAIVideoCategoryJobIds) Validate() error {
	return dara.Validate(s)
}

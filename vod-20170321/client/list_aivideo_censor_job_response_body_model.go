// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCensorJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCensorJobList(v *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList) *ListAIVideoCensorJobResponseBody
	GetAIVideoCensorJobList() *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList
	SetNonExistAIVideoCensorJobIds(v *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds) *ListAIVideoCensorJobResponseBody
	GetNonExistAIVideoCensorJobIds() *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds
	SetRequestId(v string) *ListAIVideoCensorJobResponseBody
	GetRequestId() *string
}

type ListAIVideoCensorJobResponseBody struct {
	AIVideoCensorJobList        *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList        `json:"AIVideoCensorJobList,omitempty" xml:"AIVideoCensorJobList,omitempty" type:"Struct"`
	NonExistAIVideoCensorJobIds *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds `json:"NonExistAIVideoCensorJobIds,omitempty" xml:"NonExistAIVideoCensorJobIds,omitempty" type:"Struct"`
	RequestId                   *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIVideoCensorJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCensorJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIVideoCensorJobResponseBody) GetAIVideoCensorJobList() *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList {
	return s.AIVideoCensorJobList
}

func (s *ListAIVideoCensorJobResponseBody) GetNonExistAIVideoCensorJobIds() *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds {
	return s.NonExistAIVideoCensorJobIds
}

func (s *ListAIVideoCensorJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIVideoCensorJobResponseBody) SetAIVideoCensorJobList(v *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList) *ListAIVideoCensorJobResponseBody {
	s.AIVideoCensorJobList = v
	return s
}

func (s *ListAIVideoCensorJobResponseBody) SetNonExistAIVideoCensorJobIds(v *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds) *ListAIVideoCensorJobResponseBody {
	s.NonExistAIVideoCensorJobIds = v
	return s
}

func (s *ListAIVideoCensorJobResponseBody) SetRequestId(v string) *ListAIVideoCensorJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIVideoCensorJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCensorJobResponseBodyAIVideoCensorJobList struct {
	AIVideoCensorJob []*ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob `json:"AIVideoCensorJob,omitempty" xml:"AIVideoCensorJob,omitempty" type:"Repeated"`
}

func (s ListAIVideoCensorJobResponseBodyAIVideoCensorJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCensorJobResponseBodyAIVideoCensorJobList) GoString() string {
	return s.String()
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList) GetAIVideoCensorJob() []*ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob {
	return s.AIVideoCensorJob
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList) SetAIVideoCensorJob(v []*ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList {
	s.AIVideoCensorJob = v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) GoString() string {
	return s.String()
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) GetCode() *string {
	return s.Code
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) GetData() *string {
	return s.Data
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) SetCode(v string) *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob {
	s.Code = &v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) SetCreationTime(v string) *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) SetData(v string) *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob {
	s.Data = &v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) SetJobId(v string) *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob {
	s.JobId = &v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) SetMediaId(v string) *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob {
	s.MediaId = &v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) SetMessage(v string) *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob {
	s.Message = &v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) SetStatus(v string) *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob {
	s.Status = &v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyAIVideoCensorJobListAIVideoCensorJob) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds) GoString() string {
	return s.String()
}

func (s *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds) SetString_(v []*string) *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds {
	s.String_ = v
	return s
}

func (s *ListAIVideoCensorJobResponseBodyNonExistAIVideoCensorJobIds) Validate() error {
	return dara.Validate(s)
}

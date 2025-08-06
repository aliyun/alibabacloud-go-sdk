// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIASRJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIASRJobList(v *ListAIASRJobResponseBodyAIASRJobList) *ListAIASRJobResponseBody
	GetAIASRJobList() *ListAIASRJobResponseBodyAIASRJobList
	SetNonExistAIASRJobIds(v *ListAIASRJobResponseBodyNonExistAIASRJobIds) *ListAIASRJobResponseBody
	GetNonExistAIASRJobIds() *ListAIASRJobResponseBodyNonExistAIASRJobIds
	SetRequestId(v string) *ListAIASRJobResponseBody
	GetRequestId() *string
}

type ListAIASRJobResponseBody struct {
	AIASRJobList        *ListAIASRJobResponseBodyAIASRJobList        `json:"AIASRJobList,omitempty" xml:"AIASRJobList,omitempty" type:"Struct"`
	NonExistAIASRJobIds *ListAIASRJobResponseBodyNonExistAIASRJobIds `json:"NonExistAIASRJobIds,omitempty" xml:"NonExistAIASRJobIds,omitempty" type:"Struct"`
	RequestId           *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIASRJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIASRJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIASRJobResponseBody) GetAIASRJobList() *ListAIASRJobResponseBodyAIASRJobList {
	return s.AIASRJobList
}

func (s *ListAIASRJobResponseBody) GetNonExistAIASRJobIds() *ListAIASRJobResponseBodyNonExistAIASRJobIds {
	return s.NonExistAIASRJobIds
}

func (s *ListAIASRJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIASRJobResponseBody) SetAIASRJobList(v *ListAIASRJobResponseBodyAIASRJobList) *ListAIASRJobResponseBody {
	s.AIASRJobList = v
	return s
}

func (s *ListAIASRJobResponseBody) SetNonExistAIASRJobIds(v *ListAIASRJobResponseBodyNonExistAIASRJobIds) *ListAIASRJobResponseBody {
	s.NonExistAIASRJobIds = v
	return s
}

func (s *ListAIASRJobResponseBody) SetRequestId(v string) *ListAIASRJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIASRJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIASRJobResponseBodyAIASRJobList struct {
	AIASRJob []*ListAIASRJobResponseBodyAIASRJobListAIASRJob `json:"AIASRJob,omitempty" xml:"AIASRJob,omitempty" type:"Repeated"`
}

func (s ListAIASRJobResponseBodyAIASRJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIASRJobResponseBodyAIASRJobList) GoString() string {
	return s.String()
}

func (s *ListAIASRJobResponseBodyAIASRJobList) GetAIASRJob() []*ListAIASRJobResponseBodyAIASRJobListAIASRJob {
	return s.AIASRJob
}

func (s *ListAIASRJobResponseBodyAIASRJobList) SetAIASRJob(v []*ListAIASRJobResponseBodyAIASRJobListAIASRJob) *ListAIASRJobResponseBodyAIASRJobList {
	s.AIASRJob = v
	return s
}

func (s *ListAIASRJobResponseBodyAIASRJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIASRJobResponseBodyAIASRJobListAIASRJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIASRJobResponseBodyAIASRJobListAIASRJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIASRJobResponseBodyAIASRJobListAIASRJob) GoString() string {
	return s.String()
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) GetCode() *string {
	return s.Code
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) GetData() *string {
	return s.Data
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) SetCode(v string) *ListAIASRJobResponseBodyAIASRJobListAIASRJob {
	s.Code = &v
	return s
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) SetCreationTime(v string) *ListAIASRJobResponseBodyAIASRJobListAIASRJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) SetData(v string) *ListAIASRJobResponseBodyAIASRJobListAIASRJob {
	s.Data = &v
	return s
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) SetJobId(v string) *ListAIASRJobResponseBodyAIASRJobListAIASRJob {
	s.JobId = &v
	return s
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) SetMediaId(v string) *ListAIASRJobResponseBodyAIASRJobListAIASRJob {
	s.MediaId = &v
	return s
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) SetMessage(v string) *ListAIASRJobResponseBodyAIASRJobListAIASRJob {
	s.Message = &v
	return s
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) SetStatus(v string) *ListAIASRJobResponseBodyAIASRJobListAIASRJob {
	s.Status = &v
	return s
}

func (s *ListAIASRJobResponseBodyAIASRJobListAIASRJob) Validate() error {
	return dara.Validate(s)
}

type ListAIASRJobResponseBodyNonExistAIASRJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIASRJobResponseBodyNonExistAIASRJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIASRJobResponseBodyNonExistAIASRJobIds) GoString() string {
	return s.String()
}

func (s *ListAIASRJobResponseBodyNonExistAIASRJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIASRJobResponseBodyNonExistAIASRJobIds) SetString_(v []*string) *ListAIASRJobResponseBodyNonExistAIASRJobIds {
	s.String_ = v
	return s
}

func (s *ListAIASRJobResponseBodyNonExistAIASRJobIds) Validate() error {
	return dara.Validate(s)
}

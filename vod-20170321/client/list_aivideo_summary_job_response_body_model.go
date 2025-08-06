// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoSummaryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoSummaryJobList(v *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList) *ListAIVideoSummaryJobResponseBody
	GetAIVideoSummaryJobList() *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList
	SetNonExistAIVideoSummaryJobIds(v *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds) *ListAIVideoSummaryJobResponseBody
	GetNonExistAIVideoSummaryJobIds() *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds
	SetRequestId(v string) *ListAIVideoSummaryJobResponseBody
	GetRequestId() *string
}

type ListAIVideoSummaryJobResponseBody struct {
	AIVideoSummaryJobList        *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList        `json:"AIVideoSummaryJobList,omitempty" xml:"AIVideoSummaryJobList,omitempty" type:"Struct"`
	NonExistAIVideoSummaryJobIds *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds `json:"NonExistAIVideoSummaryJobIds,omitempty" xml:"NonExistAIVideoSummaryJobIds,omitempty" type:"Struct"`
	RequestId                    *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIVideoSummaryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoSummaryJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIVideoSummaryJobResponseBody) GetAIVideoSummaryJobList() *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList {
	return s.AIVideoSummaryJobList
}

func (s *ListAIVideoSummaryJobResponseBody) GetNonExistAIVideoSummaryJobIds() *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds {
	return s.NonExistAIVideoSummaryJobIds
}

func (s *ListAIVideoSummaryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIVideoSummaryJobResponseBody) SetAIVideoSummaryJobList(v *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList) *ListAIVideoSummaryJobResponseBody {
	s.AIVideoSummaryJobList = v
	return s
}

func (s *ListAIVideoSummaryJobResponseBody) SetNonExistAIVideoSummaryJobIds(v *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds) *ListAIVideoSummaryJobResponseBody {
	s.NonExistAIVideoSummaryJobIds = v
	return s
}

func (s *ListAIVideoSummaryJobResponseBody) SetRequestId(v string) *ListAIVideoSummaryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIVideoSummaryJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList struct {
	AIVideoSummaryJob []*ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob `json:"AIVideoSummaryJob,omitempty" xml:"AIVideoSummaryJob,omitempty" type:"Repeated"`
}

func (s ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList) GoString() string {
	return s.String()
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList) GetAIVideoSummaryJob() []*ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob {
	return s.AIVideoSummaryJob
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList) SetAIVideoSummaryJob(v []*ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList {
	s.AIVideoSummaryJob = v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) GoString() string {
	return s.String()
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) GetCode() *string {
	return s.Code
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) GetData() *string {
	return s.Data
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) SetCode(v string) *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob {
	s.Code = &v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) SetCreationTime(v string) *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) SetData(v string) *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob {
	s.Data = &v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) SetJobId(v string) *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob {
	s.JobId = &v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) SetMediaId(v string) *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob {
	s.MediaId = &v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) SetMessage(v string) *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob {
	s.Message = &v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) SetStatus(v string) *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob {
	s.Status = &v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyAIVideoSummaryJobListAIVideoSummaryJob) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds) GoString() string {
	return s.String()
}

func (s *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds) SetString_(v []*string) *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds {
	s.String_ = v
	return s
}

func (s *ListAIVideoSummaryJobResponseBodyNonExistAIVideoSummaryJobIds) Validate() error {
	return dara.Validate(s)
}

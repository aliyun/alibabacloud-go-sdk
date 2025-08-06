// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoTagJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoTagJobList(v *ListAIVideoTagJobResponseBodyAIVideoTagJobList) *ListAIVideoTagJobResponseBody
	GetAIVideoTagJobList() *ListAIVideoTagJobResponseBodyAIVideoTagJobList
	SetNonExistAIVideoTagJobIds(v *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds) *ListAIVideoTagJobResponseBody
	GetNonExistAIVideoTagJobIds() *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds
	SetRequestId(v string) *ListAIVideoTagJobResponseBody
	GetRequestId() *string
}

type ListAIVideoTagJobResponseBody struct {
	AIVideoTagJobList        *ListAIVideoTagJobResponseBodyAIVideoTagJobList        `json:"AIVideoTagJobList,omitempty" xml:"AIVideoTagJobList,omitempty" type:"Struct"`
	NonExistAIVideoTagJobIds *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds `json:"NonExistAIVideoTagJobIds,omitempty" xml:"NonExistAIVideoTagJobIds,omitempty" type:"Struct"`
	RequestId                *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIVideoTagJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTagJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIVideoTagJobResponseBody) GetAIVideoTagJobList() *ListAIVideoTagJobResponseBodyAIVideoTagJobList {
	return s.AIVideoTagJobList
}

func (s *ListAIVideoTagJobResponseBody) GetNonExistAIVideoTagJobIds() *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds {
	return s.NonExistAIVideoTagJobIds
}

func (s *ListAIVideoTagJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIVideoTagJobResponseBody) SetAIVideoTagJobList(v *ListAIVideoTagJobResponseBodyAIVideoTagJobList) *ListAIVideoTagJobResponseBody {
	s.AIVideoTagJobList = v
	return s
}

func (s *ListAIVideoTagJobResponseBody) SetNonExistAIVideoTagJobIds(v *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds) *ListAIVideoTagJobResponseBody {
	s.NonExistAIVideoTagJobIds = v
	return s
}

func (s *ListAIVideoTagJobResponseBody) SetRequestId(v string) *ListAIVideoTagJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIVideoTagJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoTagJobResponseBodyAIVideoTagJobList struct {
	AIVideoTagJob []*ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob `json:"AIVideoTagJob,omitempty" xml:"AIVideoTagJob,omitempty" type:"Repeated"`
}

func (s ListAIVideoTagJobResponseBodyAIVideoTagJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTagJobResponseBodyAIVideoTagJobList) GoString() string {
	return s.String()
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobList) GetAIVideoTagJob() []*ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob {
	return s.AIVideoTagJob
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobList) SetAIVideoTagJob(v []*ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) *ListAIVideoTagJobResponseBodyAIVideoTagJobList {
	s.AIVideoTagJob = v
	return s
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) GoString() string {
	return s.String()
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) GetCode() *string {
	return s.Code
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) GetData() *string {
	return s.Data
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) SetCode(v string) *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob {
	s.Code = &v
	return s
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) SetCreationTime(v string) *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) SetData(v string) *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob {
	s.Data = &v
	return s
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) SetJobId(v string) *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob {
	s.JobId = &v
	return s
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) SetMediaId(v string) *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob {
	s.MediaId = &v
	return s
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) SetMessage(v string) *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob {
	s.Message = &v
	return s
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) SetStatus(v string) *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob {
	s.Status = &v
	return s
}

func (s *ListAIVideoTagJobResponseBodyAIVideoTagJobListAIVideoTagJob) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds) GoString() string {
	return s.String()
}

func (s *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds) SetString_(v []*string) *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds {
	s.String_ = v
	return s
}

func (s *ListAIVideoTagJobResponseBodyNonExistAIVideoTagJobIds) Validate() error {
	return dara.Validate(s)
}

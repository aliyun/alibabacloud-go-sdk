// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoFaceRecogJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoFaceRecogJobList(v *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList) *ListAIVideoFaceRecogJobResponseBody
	GetAIVideoFaceRecogJobList() *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList
	SetNonExistAIVideoFaceRecogJobIds(v *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds) *ListAIVideoFaceRecogJobResponseBody
	GetNonExistAIVideoFaceRecogJobIds() *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds
	SetRequestId(v string) *ListAIVideoFaceRecogJobResponseBody
	GetRequestId() *string
}

type ListAIVideoFaceRecogJobResponseBody struct {
	AIVideoFaceRecogJobList        *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList        `json:"AIVideoFaceRecogJobList,omitempty" xml:"AIVideoFaceRecogJobList,omitempty" type:"Struct"`
	NonExistAIVideoFaceRecogJobIds *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds `json:"NonExistAIVideoFaceRecogJobIds,omitempty" xml:"NonExistAIVideoFaceRecogJobIds,omitempty" type:"Struct"`
	RequestId                      *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIVideoFaceRecogJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoFaceRecogJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIVideoFaceRecogJobResponseBody) GetAIVideoFaceRecogJobList() *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList {
	return s.AIVideoFaceRecogJobList
}

func (s *ListAIVideoFaceRecogJobResponseBody) GetNonExistAIVideoFaceRecogJobIds() *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds {
	return s.NonExistAIVideoFaceRecogJobIds
}

func (s *ListAIVideoFaceRecogJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIVideoFaceRecogJobResponseBody) SetAIVideoFaceRecogJobList(v *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList) *ListAIVideoFaceRecogJobResponseBody {
	s.AIVideoFaceRecogJobList = v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBody) SetNonExistAIVideoFaceRecogJobIds(v *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds) *ListAIVideoFaceRecogJobResponseBody {
	s.NonExistAIVideoFaceRecogJobIds = v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBody) SetRequestId(v string) *ListAIVideoFaceRecogJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList struct {
	AIVideoFaceRecogJob []*ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob `json:"AIVideoFaceRecogJob,omitempty" xml:"AIVideoFaceRecogJob,omitempty" type:"Repeated"`
}

func (s ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList) GoString() string {
	return s.String()
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList) GetAIVideoFaceRecogJob() []*ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob {
	return s.AIVideoFaceRecogJob
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList) SetAIVideoFaceRecogJob(v []*ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList {
	s.AIVideoFaceRecogJob = v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) GoString() string {
	return s.String()
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) GetCode() *string {
	return s.Code
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) GetData() *string {
	return s.Data
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) SetCode(v string) *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob {
	s.Code = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) SetCreationTime(v string) *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) SetData(v string) *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob {
	s.Data = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) SetJobId(v string) *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob {
	s.JobId = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) SetMediaId(v string) *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob {
	s.MediaId = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) SetMessage(v string) *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob {
	s.Message = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) SetStatus(v string) *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob {
	s.Status = &v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJobListAIVideoFaceRecogJob) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds) GoString() string {
	return s.String()
}

func (s *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds) SetString_(v []*string) *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds {
	s.String_ = v
	return s
}

func (s *ListAIVideoFaceRecogJobResponseBodyNonExistAIVideoFaceRecogJobIds) Validate() error {
	return dara.Validate(s)
}

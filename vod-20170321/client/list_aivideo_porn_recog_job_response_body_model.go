// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoPornRecogJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoPornRecogJobList(v *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList) *ListAIVideoPornRecogJobResponseBody
	GetAIVideoPornRecogJobList() *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList
	SetNonExistPornRecogJobIds(v *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds) *ListAIVideoPornRecogJobResponseBody
	GetNonExistPornRecogJobIds() *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds
	SetRequestId(v string) *ListAIVideoPornRecogJobResponseBody
	GetRequestId() *string
}

type ListAIVideoPornRecogJobResponseBody struct {
	AIVideoPornRecogJobList *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList `json:"AIVideoPornRecogJobList,omitempty" xml:"AIVideoPornRecogJobList,omitempty" type:"Struct"`
	NonExistPornRecogJobIds *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds `json:"NonExistPornRecogJobIds,omitempty" xml:"NonExistPornRecogJobIds,omitempty" type:"Struct"`
	RequestId               *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIVideoPornRecogJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoPornRecogJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIVideoPornRecogJobResponseBody) GetAIVideoPornRecogJobList() *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList {
	return s.AIVideoPornRecogJobList
}

func (s *ListAIVideoPornRecogJobResponseBody) GetNonExistPornRecogJobIds() *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds {
	return s.NonExistPornRecogJobIds
}

func (s *ListAIVideoPornRecogJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIVideoPornRecogJobResponseBody) SetAIVideoPornRecogJobList(v *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList) *ListAIVideoPornRecogJobResponseBody {
	s.AIVideoPornRecogJobList = v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBody) SetNonExistPornRecogJobIds(v *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds) *ListAIVideoPornRecogJobResponseBody {
	s.NonExistPornRecogJobIds = v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBody) SetRequestId(v string) *ListAIVideoPornRecogJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList struct {
	AIVideoPornRecogJob []*ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob `json:"AIVideoPornRecogJob,omitempty" xml:"AIVideoPornRecogJob,omitempty" type:"Repeated"`
}

func (s ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList) GoString() string {
	return s.String()
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList) GetAIVideoPornRecogJob() []*ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob {
	return s.AIVideoPornRecogJob
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList) SetAIVideoPornRecogJob(v []*ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList {
	s.AIVideoPornRecogJob = v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) GoString() string {
	return s.String()
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) GetCode() *string {
	return s.Code
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) GetData() *string {
	return s.Data
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) SetCode(v string) *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob {
	s.Code = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) SetCreationTime(v string) *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) SetData(v string) *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob {
	s.Data = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) SetJobId(v string) *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob {
	s.JobId = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) SetMediaId(v string) *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob {
	s.MediaId = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) SetMessage(v string) *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob {
	s.Message = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) SetStatus(v string) *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob {
	s.Status = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyAIVideoPornRecogJobListAIVideoPornRecogJob) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds) GoString() string {
	return s.String()
}

func (s *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds) SetString_(v []*string) *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds {
	s.String_ = v
	return s
}

func (s *ListAIVideoPornRecogJobResponseBodyNonExistPornRecogJobIds) Validate() error {
	return dara.Validate(s)
}

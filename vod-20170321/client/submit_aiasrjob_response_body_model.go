// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIASRJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIASRJob(v *SubmitAIASRJobResponseBodyAIASRJob) *SubmitAIASRJobResponseBody
	GetAIASRJob() *SubmitAIASRJobResponseBodyAIASRJob
	SetRequestId(v string) *SubmitAIASRJobResponseBody
	GetRequestId() *string
}

type SubmitAIASRJobResponseBody struct {
	AIASRJob  *SubmitAIASRJobResponseBodyAIASRJob `json:"AIASRJob,omitempty" xml:"AIASRJob,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIASRJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIASRJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIASRJobResponseBody) GetAIASRJob() *SubmitAIASRJobResponseBodyAIASRJob {
	return s.AIASRJob
}

func (s *SubmitAIASRJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIASRJobResponseBody) SetAIASRJob(v *SubmitAIASRJobResponseBodyAIASRJob) *SubmitAIASRJobResponseBody {
	s.AIASRJob = v
	return s
}

func (s *SubmitAIASRJobResponseBody) SetRequestId(v string) *SubmitAIASRJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIASRJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIASRJobResponseBodyAIASRJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIASRJobResponseBodyAIASRJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIASRJobResponseBodyAIASRJob) GoString() string {
	return s.String()
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) SetCode(v string) *SubmitAIASRJobResponseBodyAIASRJob {
	s.Code = &v
	return s
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) SetCreationTime(v string) *SubmitAIASRJobResponseBodyAIASRJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) SetData(v string) *SubmitAIASRJobResponseBodyAIASRJob {
	s.Data = &v
	return s
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) SetJobId(v string) *SubmitAIASRJobResponseBodyAIASRJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) SetMediaId(v string) *SubmitAIASRJobResponseBodyAIASRJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) SetMessage(v string) *SubmitAIASRJobResponseBodyAIASRJob {
	s.Message = &v
	return s
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) SetStatus(v string) *SubmitAIASRJobResponseBodyAIASRJob {
	s.Status = &v
	return s
}

func (s *SubmitAIASRJobResponseBodyAIASRJob) Validate() error {
	return dara.Validate(s)
}

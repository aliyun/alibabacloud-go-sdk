// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoTagJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoTagJob(v *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) *SubmitAIVideoTagJobResponseBody
	GetAIVideoTagJob() *SubmitAIVideoTagJobResponseBodyAIVideoTagJob
	SetRequestId(v string) *SubmitAIVideoTagJobResponseBody
	GetRequestId() *string
}

type SubmitAIVideoTagJobResponseBody struct {
	AIVideoTagJob *SubmitAIVideoTagJobResponseBodyAIVideoTagJob `json:"AIVideoTagJob,omitempty" xml:"AIVideoTagJob,omitempty" type:"Struct"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIVideoTagJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoTagJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoTagJobResponseBody) GetAIVideoTagJob() *SubmitAIVideoTagJobResponseBodyAIVideoTagJob {
	return s.AIVideoTagJob
}

func (s *SubmitAIVideoTagJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIVideoTagJobResponseBody) SetAIVideoTagJob(v *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) *SubmitAIVideoTagJobResponseBody {
	s.AIVideoTagJob = v
	return s
}

func (s *SubmitAIVideoTagJobResponseBody) SetRequestId(v string) *SubmitAIVideoTagJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIVideoTagJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIVideoTagJobResponseBodyAIVideoTagJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIVideoTagJobResponseBodyAIVideoTagJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoTagJobResponseBodyAIVideoTagJob) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) SetCode(v string) *SubmitAIVideoTagJobResponseBodyAIVideoTagJob {
	s.Code = &v
	return s
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) SetCreationTime(v string) *SubmitAIVideoTagJobResponseBodyAIVideoTagJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) SetData(v string) *SubmitAIVideoTagJobResponseBodyAIVideoTagJob {
	s.Data = &v
	return s
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) SetJobId(v string) *SubmitAIVideoTagJobResponseBodyAIVideoTagJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) SetMediaId(v string) *SubmitAIVideoTagJobResponseBodyAIVideoTagJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) SetMessage(v string) *SubmitAIVideoTagJobResponseBodyAIVideoTagJob {
	s.Message = &v
	return s
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) SetStatus(v string) *SubmitAIVideoTagJobResponseBodyAIVideoTagJob {
	s.Status = &v
	return s
}

func (s *SubmitAIVideoTagJobResponseBodyAIVideoTagJob) Validate() error {
	return dara.Validate(s)
}

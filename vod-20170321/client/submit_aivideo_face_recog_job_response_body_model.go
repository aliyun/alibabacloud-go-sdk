// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoFaceRecogJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoFaceRecogJob(v *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) *SubmitAIVideoFaceRecogJobResponseBody
	GetAIVideoFaceRecogJob() *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob
	SetRequestId(v string) *SubmitAIVideoFaceRecogJobResponseBody
	GetRequestId() *string
}

type SubmitAIVideoFaceRecogJobResponseBody struct {
	AIVideoFaceRecogJob *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob `json:"AIVideoFaceRecogJob,omitempty" xml:"AIVideoFaceRecogJob,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIVideoFaceRecogJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoFaceRecogJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoFaceRecogJobResponseBody) GetAIVideoFaceRecogJob() *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob {
	return s.AIVideoFaceRecogJob
}

func (s *SubmitAIVideoFaceRecogJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIVideoFaceRecogJobResponseBody) SetAIVideoFaceRecogJob(v *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) *SubmitAIVideoFaceRecogJobResponseBody {
	s.AIVideoFaceRecogJob = v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBody) SetRequestId(v string) *SubmitAIVideoFaceRecogJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) SetCode(v string) *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob {
	s.Code = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) SetCreationTime(v string) *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) SetData(v string) *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob {
	s.Data = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) SetJobId(v string) *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) SetMediaId(v string) *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) SetMessage(v string) *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob {
	s.Message = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) SetStatus(v string) *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob {
	s.Status = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobResponseBodyAIVideoFaceRecogJob) Validate() error {
	return dara.Validate(s)
}

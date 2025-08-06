// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoPornRecogJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoPornRecogJob(v *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) *SubmitAIVideoPornRecogJobResponseBody
	GetAIVideoPornRecogJob() *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob
	SetRequestId(v string) *SubmitAIVideoPornRecogJobResponseBody
	GetRequestId() *string
}

type SubmitAIVideoPornRecogJobResponseBody struct {
	AIVideoPornRecogJob *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob `json:"AIVideoPornRecogJob,omitempty" xml:"AIVideoPornRecogJob,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIVideoPornRecogJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoPornRecogJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoPornRecogJobResponseBody) GetAIVideoPornRecogJob() *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob {
	return s.AIVideoPornRecogJob
}

func (s *SubmitAIVideoPornRecogJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIVideoPornRecogJobResponseBody) SetAIVideoPornRecogJob(v *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) *SubmitAIVideoPornRecogJobResponseBody {
	s.AIVideoPornRecogJob = v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBody) SetRequestId(v string) *SubmitAIVideoPornRecogJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) SetCode(v string) *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob {
	s.Code = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) SetCreationTime(v string) *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) SetData(v string) *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob {
	s.Data = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) SetJobId(v string) *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) SetMediaId(v string) *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) SetMessage(v string) *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob {
	s.Message = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) SetStatus(v string) *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob {
	s.Status = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobResponseBodyAIVideoPornRecogJob) Validate() error {
	return dara.Validate(s)
}

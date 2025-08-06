// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoTerrorismRecogJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoTerrorismRecogJob(v *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) *SubmitAIVideoTerrorismRecogJobResponseBody
	GetAIVideoTerrorismRecogJob() *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob
	SetRequestId(v string) *SubmitAIVideoTerrorismRecogJobResponseBody
	GetRequestId() *string
}

type SubmitAIVideoTerrorismRecogJobResponseBody struct {
	AIVideoTerrorismRecogJob *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob `json:"AIVideoTerrorismRecogJob,omitempty" xml:"AIVideoTerrorismRecogJob,omitempty" type:"Struct"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIVideoTerrorismRecogJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoTerrorismRecogJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBody) GetAIVideoTerrorismRecogJob() *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob {
	return s.AIVideoTerrorismRecogJob
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBody) SetAIVideoTerrorismRecogJob(v *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) *SubmitAIVideoTerrorismRecogJobResponseBody {
	s.AIVideoTerrorismRecogJob = v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBody) SetRequestId(v string) *SubmitAIVideoTerrorismRecogJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) SetCode(v string) *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob {
	s.Code = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) SetCreationTime(v string) *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) SetData(v string) *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob {
	s.Data = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) SetJobId(v string) *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) SetMediaId(v string) *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) SetMessage(v string) *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob {
	s.Message = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) SetStatus(v string) *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob {
	s.Status = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJob) Validate() error {
	return dara.Validate(s)
}

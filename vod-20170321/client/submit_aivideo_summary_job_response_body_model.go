// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoSummaryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoSummaryJob(v *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) *SubmitAIVideoSummaryJobResponseBody
	GetAIVideoSummaryJob() *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob
	SetRequestId(v string) *SubmitAIVideoSummaryJobResponseBody
	GetRequestId() *string
}

type SubmitAIVideoSummaryJobResponseBody struct {
	AIVideoSummaryJob *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob `json:"AIVideoSummaryJob,omitempty" xml:"AIVideoSummaryJob,omitempty" type:"Struct"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIVideoSummaryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoSummaryJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoSummaryJobResponseBody) GetAIVideoSummaryJob() *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob {
	return s.AIVideoSummaryJob
}

func (s *SubmitAIVideoSummaryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIVideoSummaryJobResponseBody) SetAIVideoSummaryJob(v *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) *SubmitAIVideoSummaryJobResponseBody {
	s.AIVideoSummaryJob = v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBody) SetRequestId(v string) *SubmitAIVideoSummaryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) SetCode(v string) *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob {
	s.Code = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) SetCreationTime(v string) *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) SetData(v string) *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob {
	s.Data = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) SetJobId(v string) *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) SetMediaId(v string) *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) SetMessage(v string) *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob {
	s.Message = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) SetStatus(v string) *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob {
	s.Status = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponseBodyAIVideoSummaryJob) Validate() error {
	return dara.Validate(s)
}

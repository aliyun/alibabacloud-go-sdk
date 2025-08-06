// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCategoryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCategoryJob(v *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) *SubmitAIVideoCategoryJobResponseBody
	GetAIVideoCategoryJob() *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob
	SetRequestId(v string) *SubmitAIVideoCategoryJobResponseBody
	GetRequestId() *string
}

type SubmitAIVideoCategoryJobResponseBody struct {
	AIVideoCategoryJob *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob `json:"AIVideoCategoryJob,omitempty" xml:"AIVideoCategoryJob,omitempty" type:"Struct"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIVideoCategoryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCategoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCategoryJobResponseBody) GetAIVideoCategoryJob() *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob {
	return s.AIVideoCategoryJob
}

func (s *SubmitAIVideoCategoryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIVideoCategoryJobResponseBody) SetAIVideoCategoryJob(v *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) *SubmitAIVideoCategoryJobResponseBody {
	s.AIVideoCategoryJob = v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBody) SetRequestId(v string) *SubmitAIVideoCategoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) SetCode(v string) *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob {
	s.Code = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) SetCreationTime(v string) *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) SetData(v string) *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob {
	s.Data = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) SetJobId(v string) *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) SetMediaId(v string) *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) SetMessage(v string) *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob {
	s.Message = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) SetStatus(v string) *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob {
	s.Status = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponseBodyAIVideoCategoryJob) Validate() error {
	return dara.Validate(s)
}

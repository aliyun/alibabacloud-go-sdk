// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCensorJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCensorJob(v *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) *SubmitAIVideoCensorJobResponseBody
	GetAIVideoCensorJob() *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob
	SetRequestId(v string) *SubmitAIVideoCensorJobResponseBody
	GetRequestId() *string
}

type SubmitAIVideoCensorJobResponseBody struct {
	AIVideoCensorJob *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob `json:"AIVideoCensorJob,omitempty" xml:"AIVideoCensorJob,omitempty" type:"Struct"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIVideoCensorJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCensorJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCensorJobResponseBody) GetAIVideoCensorJob() *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob {
	return s.AIVideoCensorJob
}

func (s *SubmitAIVideoCensorJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIVideoCensorJobResponseBody) SetAIVideoCensorJob(v *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) *SubmitAIVideoCensorJobResponseBody {
	s.AIVideoCensorJob = v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBody) SetRequestId(v string) *SubmitAIVideoCensorJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) SetCode(v string) *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob {
	s.Code = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) SetCreationTime(v string) *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) SetData(v string) *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob {
	s.Data = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) SetJobId(v string) *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) SetMediaId(v string) *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) SetMessage(v string) *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob {
	s.Message = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) SetStatus(v string) *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob {
	s.Status = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponseBodyAIVideoCensorJob) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCoverJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCoverJob(v *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) *SubmitAIVideoCoverJobResponseBody
	GetAIVideoCoverJob() *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob
	SetRequestId(v string) *SubmitAIVideoCoverJobResponseBody
	GetRequestId() *string
}

type SubmitAIVideoCoverJobResponseBody struct {
	AIVideoCoverJob *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob `json:"AIVideoCoverJob,omitempty" xml:"AIVideoCoverJob,omitempty" type:"Struct"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIVideoCoverJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCoverJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCoverJobResponseBody) GetAIVideoCoverJob() *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob {
	return s.AIVideoCoverJob
}

func (s *SubmitAIVideoCoverJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIVideoCoverJobResponseBody) SetAIVideoCoverJob(v *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) *SubmitAIVideoCoverJobResponseBody {
	s.AIVideoCoverJob = v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBody) SetRequestId(v string) *SubmitAIVideoCoverJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) GetData() *string {
	return s.Data
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) SetCode(v string) *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob {
	s.Code = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) SetCreationTime(v string) *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) SetData(v string) *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob {
	s.Data = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) SetJobId(v string) *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) SetMediaId(v string) *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) SetMessage(v string) *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob {
	s.Message = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) SetStatus(v string) *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob {
	s.Status = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponseBodyAIVideoCoverJob) Validate() error {
	return dara.Validate(s)
}

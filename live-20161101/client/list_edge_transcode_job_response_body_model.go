// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v *ListEdgeTranscodeJobResponseBodyJobList) *ListEdgeTranscodeJobResponseBody
	GetJobList() *ListEdgeTranscodeJobResponseBodyJobList
	SetRequestId(v string) *ListEdgeTranscodeJobResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEdgeTranscodeJobResponseBody
	GetTotalCount() *int32
}

type ListEdgeTranscodeJobResponseBody struct {
	JobList *ListEdgeTranscodeJobResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeJobResponseBody) GetJobList() *ListEdgeTranscodeJobResponseBodyJobList {
	return s.JobList
}

func (s *ListEdgeTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeTranscodeJobResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEdgeTranscodeJobResponseBody) SetJobList(v *ListEdgeTranscodeJobResponseBodyJobList) *ListEdgeTranscodeJobResponseBody {
	s.JobList = v
	return s
}

func (s *ListEdgeTranscodeJobResponseBody) SetRequestId(v string) *ListEdgeTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBody) SetTotalCount(v int32) *ListEdgeTranscodeJobResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBody) Validate() error {
	if s.JobList != nil {
		if err := s.JobList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEdgeTranscodeJobResponseBodyJobList struct {
	Job []*ListEdgeTranscodeJobResponseBodyJobListJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Repeated"`
}

func (s ListEdgeTranscodeJobResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeJobResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeJobResponseBodyJobList) GetJob() []*ListEdgeTranscodeJobResponseBodyJobListJob {
	return s.Job
}

func (s *ListEdgeTranscodeJobResponseBodyJobList) SetJob(v []*ListEdgeTranscodeJobResponseBodyJobListJob) *ListEdgeTranscodeJobResponseBodyJobList {
	s.Job = v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobList) Validate() error {
	if s.Job != nil {
		for _, item := range s.Job {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEdgeTranscodeJobResponseBodyJobListJob struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	LastStartAt  *string `json:"LastStartAt,omitempty" xml:"LastStartAt,omitempty"`
	LastStopAt   *string `json:"LastStopAt,omitempty" xml:"LastStopAt,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StreamInput  *string `json:"StreamInput,omitempty" xml:"StreamInput,omitempty"`
	StreamOutput *string `json:"StreamOutput,omitempty" xml:"StreamOutput,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEdgeTranscodeJobResponseBodyJobListJob) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeJobResponseBodyJobListJob) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetJobId() *string {
	return s.JobId
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetLastStartAt() *string {
	return s.LastStartAt
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetLastStopAt() *string {
	return s.LastStopAt
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetName() *string {
	return s.Name
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetStatus() *string {
	return s.Status
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetStreamInput() *string {
	return s.StreamInput
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetStreamOutput() *string {
	return s.StreamOutput
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) GetType() *string {
	return s.Type
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetCreateTime(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetJobId(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.JobId = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetLastStartAt(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.LastStartAt = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetLastStopAt(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.LastStopAt = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetName(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.Name = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetStatus(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.Status = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetStreamInput(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.StreamInput = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetStreamOutput(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.StreamOutput = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetTemplateId(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.TemplateId = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetTemplateName(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.TemplateName = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) SetType(v string) *ListEdgeTranscodeJobResponseBodyJobListJob {
	s.Type = &v
	return s
}

func (s *ListEdgeTranscodeJobResponseBodyJobListJob) Validate() error {
	return dara.Validate(s)
}

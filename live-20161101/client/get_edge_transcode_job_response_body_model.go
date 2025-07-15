// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetEdgeTranscodeJobResponseBodyJob) *GetEdgeTranscodeJobResponseBody
	GetJob() *GetEdgeTranscodeJobResponseBodyJob
	SetRequestId(v string) *GetEdgeTranscodeJobResponseBody
	GetRequestId() *string
}

type GetEdgeTranscodeJobResponseBody struct {
	// The details of the edge transcoding task.
	Job *GetEdgeTranscodeJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEdgeTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeTranscodeJobResponseBody) GetJob() *GetEdgeTranscodeJobResponseBodyJob {
	return s.Job
}

func (s *GetEdgeTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeTranscodeJobResponseBody) SetJob(v *GetEdgeTranscodeJobResponseBodyJob) *GetEdgeTranscodeJobResponseBody {
	s.Job = v
	return s
}

func (s *GetEdgeTranscodeJobResponseBody) SetRequestId(v string) *GetEdgeTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEdgeTranscodeJobResponseBodyJob struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2023-07-24T16:44:55Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The time when the task was last started.
	//
	// example:
	//
	// 2023-07-25T02:48:58Z
	LastStartAt *string `json:"LastStartAt,omitempty" xml:"LastStartAt,omitempty"`
	// The time when the task was last stopped.
	//
	// example:
	//
	// 2023-07-25T05:48:58Z
	LastStopAt *string `json:"LastStopAt,omitempty" xml:"LastStopAt,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// my_job
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- 0: not started
	//
	// 	- 1: in progress
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The URL of the input stream.
	//
	// example:
	//
	// rtmp://mydomain/app/stream1
	StreamInput *string `json:"StreamInput,omitempty" xml:"StreamInput,omitempty"`
	// The URL of the output stream.
	//
	// example:
	//
	// rtmp://testdomain/app/stream2
	StreamOutput *string `json:"StreamOutput,omitempty" xml:"StreamOutput,omitempty"`
	// The ID of the edge transcoding template used by the task.
	//
	// example:
	//
	// 9b1571b513cb44f7a1ba6ae561ff****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the edge transcoding template used by the task.
	//
	// example:
	//
	// my_template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of edge transcoding. Valid values:
	//
	// 	- common: standard transcoding and Narrowband HD™ 1.0 transcoding
	//
	// 	- nbhd-2: Narrowband HD™ 2.0 transcoding
	//
	// 	- ultra-hd: ultra-high definition transcoding
	//
	// example:
	//
	// common
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetEdgeTranscodeJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTranscodeJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetLastStartAt() *string {
	return s.LastStartAt
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetLastStopAt() *string {
	return s.LastStopAt
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetName() *string {
	return s.Name
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetStreamInput() *string {
	return s.StreamInput
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetStreamOutput() *string {
	return s.StreamOutput
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetEdgeTranscodeJobResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetCreateTime(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetJobId(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetLastStartAt(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.LastStartAt = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetLastStopAt(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.LastStopAt = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetName(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.Name = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetStatus(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetStreamInput(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.StreamInput = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetStreamOutput(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.StreamOutput = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetTemplateId(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.TemplateId = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetTemplateName(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.TemplateName = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) SetType(v string) *GetEdgeTranscodeJobResponseBodyJob {
	s.Type = &v
	return s
}

func (s *GetEdgeTranscodeJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}

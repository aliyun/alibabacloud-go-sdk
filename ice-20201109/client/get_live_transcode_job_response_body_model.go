// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetLiveTranscodeJobResponseBodyJob) *GetLiveTranscodeJobResponseBody
	GetJob() *GetLiveTranscodeJobResponseBodyJob
	SetRequestId(v string) *GetLiveTranscodeJobResponseBody
	GetRequestId() *string
}

type GetLiveTranscodeJobResponseBody struct {
	// The information about the transcoding job.
	Job *GetLiveTranscodeJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLiveTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeJobResponseBody) GetJob() *GetLiveTranscodeJobResponseBodyJob {
	return s.Job
}

func (s *GetLiveTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveTranscodeJobResponseBody) SetJob(v *GetLiveTranscodeJobResponseBodyJob) *GetLiveTranscodeJobResponseBody {
	s.Job = v
	return s
}

func (s *GetLiveTranscodeJobResponseBody) SetRequestId(v string) *GetLiveTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveTranscodeJobResponseBodyJob struct {
	// The time when the job was created.
	//
	// example:
	//
	// 2022-07-20T02:48:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the transcoding job.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the transcoding job.
	//
	// example:
	//
	// task1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the output stream.
	OutputStream *GetLiveTranscodeJobResponseBodyJobOutputStream `json:"OutputStream,omitempty" xml:"OutputStream,omitempty" type:"Struct"`
	// The start mode of the job.
	//
	// example:
	//
	// 0
	StartMode *int32 `json:"StartMode,omitempty" xml:"StartMode,omitempty"`
	// The state of the job.
	//
	// 	- 0: The job is not started.
	//
	// 	- 1: The job is in progress.
	//
	// 	- 2: The job is stopped.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the input stream.
	StreamInput *GetLiveTranscodeJobResponseBodyJobStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// b6491d5b3e514b7d895d14b5453ea119
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// basic
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the template.
	//
	// example:
	//
	// normal
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetLiveTranscodeJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetName() *string {
	return s.Name
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetOutputStream() *GetLiveTranscodeJobResponseBodyJobOutputStream {
	return s.OutputStream
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetStartMode() *int32 {
	return s.StartMode
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetStatus() *int32 {
	return s.Status
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetStreamInput() *GetLiveTranscodeJobResponseBodyJobStreamInput {
	return s.StreamInput
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetLiveTranscodeJobResponseBodyJob) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetCreateTime(v string) *GetLiveTranscodeJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetJobId(v string) *GetLiveTranscodeJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetName(v string) *GetLiveTranscodeJobResponseBodyJob {
	s.Name = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetOutputStream(v *GetLiveTranscodeJobResponseBodyJobOutputStream) *GetLiveTranscodeJobResponseBodyJob {
	s.OutputStream = v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetStartMode(v int32) *GetLiveTranscodeJobResponseBodyJob {
	s.StartMode = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetStatus(v int32) *GetLiveTranscodeJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetStreamInput(v *GetLiveTranscodeJobResponseBodyJobStreamInput) *GetLiveTranscodeJobResponseBodyJob {
	s.StreamInput = v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetTemplateId(v string) *GetLiveTranscodeJobResponseBodyJob {
	s.TemplateId = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetTemplateName(v string) *GetLiveTranscodeJobResponseBodyJob {
	s.TemplateName = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) SetTemplateType(v string) *GetLiveTranscodeJobResponseBodyJob {
	s.TemplateType = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJob) Validate() error {
	if s.OutputStream != nil {
		if err := s.OutputStream.Validate(); err != nil {
			return err
		}
	}
	if s.StreamInput != nil {
		if err := s.StreamInput.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveTranscodeJobResponseBodyJobOutputStream struct {
	// The information about the output stream.
	StreamInfos []*GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" type:"Repeated"`
}

func (s GetLiveTranscodeJobResponseBodyJobOutputStream) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeJobResponseBodyJobOutputStream) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeJobResponseBodyJobOutputStream) GetStreamInfos() []*GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos {
	return s.StreamInfos
}

func (s *GetLiveTranscodeJobResponseBodyJobOutputStream) SetStreamInfos(v []*GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos) *GetLiveTranscodeJobResponseBodyJobOutputStream {
	s.StreamInfos = v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJobOutputStream) Validate() error {
	if s.StreamInfos != nil {
		for _, item := range s.StreamInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos struct {
	// The URL of the output stream.
	//
	// example:
	//
	// rtmp://mydomain/app/mytranscode1
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The type of the output stream protocol. Only the RTMP protocol is supported.
	//
	// example:
	//
	// rtmp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos) GetType() *string {
	return s.Type
}

func (s *GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos) SetOutputUrl(v string) *GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos {
	s.OutputUrl = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos) SetType(v string) *GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos {
	s.Type = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJobOutputStreamStreamInfos) Validate() error {
	return dara.Validate(s)
}

type GetLiveTranscodeJobResponseBodyJobStreamInput struct {
	// The URL of the input stream.
	//
	// example:
	//
	// rtmp://mydomain/app/stream1
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// The type of the input stream.
	//
	// example:
	//
	// rtmp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLiveTranscodeJobResponseBodyJobStreamInput) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeJobResponseBodyJobStreamInput) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeJobResponseBodyJobStreamInput) GetInputUrl() *string {
	return s.InputUrl
}

func (s *GetLiveTranscodeJobResponseBodyJobStreamInput) GetType() *string {
	return s.Type
}

func (s *GetLiveTranscodeJobResponseBodyJobStreamInput) SetInputUrl(v string) *GetLiveTranscodeJobResponseBodyJobStreamInput {
	s.InputUrl = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJobStreamInput) SetType(v string) *GetLiveTranscodeJobResponseBodyJobStreamInput {
	s.Type = &v
	return s
}

func (s *GetLiveTranscodeJobResponseBodyJobStreamInput) Validate() error {
	return dara.Validate(s)
}

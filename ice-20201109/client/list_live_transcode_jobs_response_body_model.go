// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTranscodeJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v []*ListLiveTranscodeJobsResponseBodyJobList) *ListLiveTranscodeJobsResponseBody
	GetJobList() []*ListLiveTranscodeJobsResponseBodyJobList
	SetRequestId(v string) *ListLiveTranscodeJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLiveTranscodeJobsResponseBody
	GetTotalCount() *int32
}

type ListLiveTranscodeJobsResponseBody struct {
	// The list of transcoding jobs.
	JobList []*ListLiveTranscodeJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
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

func (s ListLiveTranscodeJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeJobsResponseBody) GetJobList() []*ListLiveTranscodeJobsResponseBodyJobList {
	return s.JobList
}

func (s *ListLiveTranscodeJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveTranscodeJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLiveTranscodeJobsResponseBody) SetJobList(v []*ListLiveTranscodeJobsResponseBodyJobList) *ListLiveTranscodeJobsResponseBody {
	s.JobList = v
	return s
}

func (s *ListLiveTranscodeJobsResponseBody) SetRequestId(v string) *ListLiveTranscodeJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBody) SetTotalCount(v int32) *ListLiveTranscodeJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBody) Validate() error {
	if s.JobList != nil {
		for _, item := range s.JobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveTranscodeJobsResponseBodyJobList struct {
	// The time when the job was created.
	//
	// example:
	//
	// 2022-07-20T02:48:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the transcoding job.
	//
	// example:
	//
	// mytask
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the output stream.
	OutputStream *ListLiveTranscodeJobsResponseBodyJobListOutputStream `json:"OutputStream,omitempty" xml:"OutputStream,omitempty" type:"Struct"`
	// The start mode of the job.
	//
	// example:
	//
	// 0
	StartMode *int32 `json:"StartMode,omitempty" xml:"StartMode,omitempty"`
	// The state of the job.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the input stream.
	StreamInput *ListLiveTranscodeJobsResponseBodyJobListStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The ID of the transcoding template used by the transcoding job.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287666****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the transcoding template used by the transcoding job.
	//
	// example:
	//
	// normal
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListLiveTranscodeJobsResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeJobsResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetName() *string {
	return s.Name
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetOutputStream() *ListLiveTranscodeJobsResponseBodyJobListOutputStream {
	return s.OutputStream
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetStartMode() *int32 {
	return s.StartMode
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetStatus() *int32 {
	return s.Status
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetStreamInput() *ListLiveTranscodeJobsResponseBodyJobListStreamInput {
	return s.StreamInput
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetCreateTime(v string) *ListLiveTranscodeJobsResponseBodyJobList {
	s.CreateTime = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetJobId(v string) *ListLiveTranscodeJobsResponseBodyJobList {
	s.JobId = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetName(v string) *ListLiveTranscodeJobsResponseBodyJobList {
	s.Name = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetOutputStream(v *ListLiveTranscodeJobsResponseBodyJobListOutputStream) *ListLiveTranscodeJobsResponseBodyJobList {
	s.OutputStream = v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetStartMode(v int32) *ListLiveTranscodeJobsResponseBodyJobList {
	s.StartMode = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetStatus(v int32) *ListLiveTranscodeJobsResponseBodyJobList {
	s.Status = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetStreamInput(v *ListLiveTranscodeJobsResponseBodyJobListStreamInput) *ListLiveTranscodeJobsResponseBodyJobList {
	s.StreamInput = v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetTemplateId(v string) *ListLiveTranscodeJobsResponseBodyJobList {
	s.TemplateId = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetTemplateName(v string) *ListLiveTranscodeJobsResponseBodyJobList {
	s.TemplateName = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) SetTemplateType(v string) *ListLiveTranscodeJobsResponseBodyJobList {
	s.TemplateType = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobList) Validate() error {
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

type ListLiveTranscodeJobsResponseBodyJobListOutputStream struct {
	// The list of stream URLs.
	StreamInfos []*ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" type:"Repeated"`
}

func (s ListLiveTranscodeJobsResponseBodyJobListOutputStream) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeJobsResponseBodyJobListOutputStream) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeJobsResponseBodyJobListOutputStream) GetStreamInfos() []*ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos {
	return s.StreamInfos
}

func (s *ListLiveTranscodeJobsResponseBodyJobListOutputStream) SetStreamInfos(v []*ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos) *ListLiveTranscodeJobsResponseBodyJobListOutputStream {
	s.StreamInfos = v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobListOutputStream) Validate() error {
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

type ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos struct {
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

func (s ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos) GetType() *string {
	return s.Type
}

func (s *ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos) SetOutputUrl(v string) *ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos {
	s.OutputUrl = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos) SetType(v string) *ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos {
	s.Type = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobListOutputStreamStreamInfos) Validate() error {
	return dara.Validate(s)
}

type ListLiveTranscodeJobsResponseBodyJobListStreamInput struct {
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

func (s ListLiveTranscodeJobsResponseBodyJobListStreamInput) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeJobsResponseBodyJobListStreamInput) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeJobsResponseBodyJobListStreamInput) GetInputUrl() *string {
	return s.InputUrl
}

func (s *ListLiveTranscodeJobsResponseBodyJobListStreamInput) GetType() *string {
	return s.Type
}

func (s *ListLiveTranscodeJobsResponseBodyJobListStreamInput) SetInputUrl(v string) *ListLiveTranscodeJobsResponseBodyJobListStreamInput {
	s.InputUrl = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobListStreamInput) SetType(v string) *ListLiveTranscodeJobsResponseBodyJobListStreamInput {
	s.Type = &v
	return s
}

func (s *ListLiveTranscodeJobsResponseBodyJobListStreamInput) Validate() error {
	return dara.Validate(s)
}

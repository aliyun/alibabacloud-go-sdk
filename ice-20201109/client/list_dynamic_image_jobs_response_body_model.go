// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicImageJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListDynamicImageJobsResponseBodyJobs) *ListDynamicImageJobsResponseBody
	GetJobs() []*ListDynamicImageJobsResponseBodyJobs
	SetNextPageToken(v string) *ListDynamicImageJobsResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListDynamicImageJobsResponseBody
	GetRequestId() *string
}

type ListDynamicImageJobsResponseBody struct {
	// The list of jobs.
	Jobs []*ListDynamicImageJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDynamicImageJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicImageJobsResponseBody) GetJobs() []*ListDynamicImageJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListDynamicImageJobsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDynamicImageJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDynamicImageJobsResponseBody) SetJobs(v []*ListDynamicImageJobsResponseBodyJobs) *ListDynamicImageJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListDynamicImageJobsResponseBody) SetNextPageToken(v string) *ListDynamicImageJobsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListDynamicImageJobsResponseBody) SetRequestId(v string) *ListDynamicImageJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDynamicImageJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDynamicImageJobsResponseBodyJobs struct {
	// The time when the job was created.
	//
	// example:
	//
	// 2022-07-12T16:17:54Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job was complete.
	//
	// example:
	//
	// 2022-07-12T16:30:54Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input of the job.
	Input *ListDynamicImageJobsResponseBodyJobsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2022-07-12T16:30:54Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// SampleJob
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the job.
	Output *ListDynamicImageJobsResponseBodyJobsOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The ID of the MPS queue to which the job was submitted.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The state of the job.
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job failed.
	//
	// 	- **Init**: The job is submitted.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 2022-07-12T16:17:54Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The request trigger source.
	//
	// Valid values:
	//
	// 	- Console
	//
	// 	- Workflow
	//
	// 	- API
	//
	// example:
	//
	// API
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
}

func (s ListDynamicImageJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetInput() *ListDynamicImageJobsResponseBodyJobsInput {
	return s.Input
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetName() *string {
	return s.Name
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetOutput() *ListDynamicImageJobsResponseBodyJobsOutput {
	return s.Output
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListDynamicImageJobsResponseBodyJobs) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetCreateTime(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetFinishTime(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.FinishTime = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetInput(v *ListDynamicImageJobsResponseBodyJobsInput) *ListDynamicImageJobsResponseBodyJobs {
	s.Input = v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetJobId(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetModifiedTime(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.ModifiedTime = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetName(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetOutput(v *ListDynamicImageJobsResponseBodyJobsOutput) *ListDynamicImageJobsResponseBodyJobs {
	s.Output = v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetPipelineId(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.PipelineId = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetStatus(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetSubmitTime(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.SubmitTime = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetTemplateId(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.TemplateId = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) SetTriggerSource(v string) *ListDynamicImageJobsResponseBodyJobs {
	s.TriggerSource = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobs) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDynamicImageJobsResponseBodyJobsInput struct {
	// The input file. The file can be an OSS object or a media asset. The URL of an OSS object can be in one of the following formats:
	//
	// 1.  OSS://bucket/object
	//
	// 2.  http(s)://bucket.oss-[regionId].aliyuncs.com/object In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// 1.  OSS: an Object Storage Service (OSS) object.
	//
	// 2.  Media: a media asset.
	//
	// *
	//
	// *
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDynamicImageJobsResponseBodyJobsInput) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageJobsResponseBodyJobsInput) GoString() string {
	return s.String()
}

func (s *ListDynamicImageJobsResponseBodyJobsInput) GetMedia() *string {
	return s.Media
}

func (s *ListDynamicImageJobsResponseBodyJobsInput) GetType() *string {
	return s.Type
}

func (s *ListDynamicImageJobsResponseBodyJobsInput) SetMedia(v string) *ListDynamicImageJobsResponseBodyJobsInput {
	s.Media = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobsInput) SetType(v string) *ListDynamicImageJobsResponseBodyJobsInput {
	s.Type = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobsInput) Validate() error {
	return dara.Validate(s)
}

type ListDynamicImageJobsResponseBodyJobsOutput struct {
	// The input file. If Type is set to OSS, the URL of an OSS object is returned. If Type is set to Media, the ID of a media asset is returned. The URL of an OSS object can be in one of the following formats:
	//
	// 1.  OSS://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output file. Valid values:
	//
	// 1.  OSS: an OSS object.
	//
	// 2.  Media: a media asset.
	//
	// *
	//
	// *
	//
	// example:
	//
	// Media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDynamicImageJobsResponseBodyJobsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageJobsResponseBodyJobsOutput) GoString() string {
	return s.String()
}

func (s *ListDynamicImageJobsResponseBodyJobsOutput) GetMedia() *string {
	return s.Media
}

func (s *ListDynamicImageJobsResponseBodyJobsOutput) GetType() *string {
	return s.Type
}

func (s *ListDynamicImageJobsResponseBodyJobsOutput) SetMedia(v string) *ListDynamicImageJobsResponseBodyJobsOutput {
	s.Media = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobsOutput) SetType(v string) *ListDynamicImageJobsResponseBodyJobsOutput {
	s.Type = &v
	return s
}

func (s *ListDynamicImageJobsResponseBodyJobsOutput) Validate() error {
	return dara.Validate(s)
}

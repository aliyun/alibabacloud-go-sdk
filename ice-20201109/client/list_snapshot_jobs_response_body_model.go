// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListSnapshotJobsResponseBodyJobs) *ListSnapshotJobsResponseBody
	GetJobs() []*ListSnapshotJobsResponseBodyJobs
	SetNextPageToken(v string) *ListSnapshotJobsResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListSnapshotJobsResponseBody
	GetRequestId() *string
}

type ListSnapshotJobsResponseBody struct {
	// The list of jobs.
	Jobs []*ListSnapshotJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSnapshotJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotJobsResponseBody) GetJobs() []*ListSnapshotJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListSnapshotJobsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListSnapshotJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSnapshotJobsResponseBody) SetJobs(v []*ListSnapshotJobsResponseBodyJobs) *ListSnapshotJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListSnapshotJobsResponseBody) SetNextPageToken(v string) *ListSnapshotJobsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListSnapshotJobsResponseBody) SetRequestId(v string) *ListSnapshotJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotJobsResponseBody) Validate() error {
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

type ListSnapshotJobsResponseBodyJobs struct {
	// Indicates whether the snapshots were captured in asynchronous mode.
	//
	// example:
	//
	// true
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The number of snapshots.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
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
	Input *ListSnapshotJobsResponseBodyJobsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
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
	Output *ListSnapshotJobsResponseBodyJobsOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The ID of the MPS queue to which the job was submitted.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
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
	// ****20b48fb04483915d4f2cd8ac****
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
	// The type of the job.
	//
	// Valid values:
	//
	// 	- WebVtt
	//
	// 	- Sprite
	//
	// 	- Normal
	//
	// example:
	//
	// Sprite
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSnapshotJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListSnapshotJobsResponseBodyJobs) GetAsync() *bool {
	return s.Async
}

func (s *ListSnapshotJobsResponseBodyJobs) GetCount() *int32 {
	return s.Count
}

func (s *ListSnapshotJobsResponseBodyJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSnapshotJobsResponseBodyJobs) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListSnapshotJobsResponseBodyJobs) GetInput() *ListSnapshotJobsResponseBodyJobsInput {
	return s.Input
}

func (s *ListSnapshotJobsResponseBodyJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListSnapshotJobsResponseBodyJobs) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListSnapshotJobsResponseBodyJobs) GetName() *string {
	return s.Name
}

func (s *ListSnapshotJobsResponseBodyJobs) GetOutput() *ListSnapshotJobsResponseBodyJobsOutput {
	return s.Output
}

func (s *ListSnapshotJobsResponseBodyJobs) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListSnapshotJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListSnapshotJobsResponseBodyJobs) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListSnapshotJobsResponseBodyJobs) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListSnapshotJobsResponseBodyJobs) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *ListSnapshotJobsResponseBodyJobs) GetType() *string {
	return s.Type
}

func (s *ListSnapshotJobsResponseBodyJobs) SetAsync(v bool) *ListSnapshotJobsResponseBodyJobs {
	s.Async = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetCount(v int32) *ListSnapshotJobsResponseBodyJobs {
	s.Count = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetCreateTime(v string) *ListSnapshotJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetFinishTime(v string) *ListSnapshotJobsResponseBodyJobs {
	s.FinishTime = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetInput(v *ListSnapshotJobsResponseBodyJobsInput) *ListSnapshotJobsResponseBodyJobs {
	s.Input = v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetJobId(v string) *ListSnapshotJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetModifiedTime(v string) *ListSnapshotJobsResponseBodyJobs {
	s.ModifiedTime = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetName(v string) *ListSnapshotJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetOutput(v *ListSnapshotJobsResponseBodyJobsOutput) *ListSnapshotJobsResponseBodyJobs {
	s.Output = v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetPipelineId(v string) *ListSnapshotJobsResponseBodyJobs {
	s.PipelineId = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetStatus(v string) *ListSnapshotJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetSubmitTime(v string) *ListSnapshotJobsResponseBodyJobs {
	s.SubmitTime = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetTemplateId(v string) *ListSnapshotJobsResponseBodyJobs {
	s.TemplateId = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetTriggerSource(v string) *ListSnapshotJobsResponseBodyJobs {
	s.TriggerSource = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) SetType(v string) *ListSnapshotJobsResponseBodyJobs {
	s.Type = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobs) Validate() error {
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

type ListSnapshotJobsResponseBodyJobsInput struct {
	// The input file. If Type is set to OSS, the URL of an OSS object is returned. If Type is set to Media, the ID of a media asset is returned. The URL of an OSS object can be in one of the following formats: 1. OSS://bucket/object 2. http(s)://bucket.oss-[RegionId].aliyuncs.com/object In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS.
	//
	// example:
	//
	// oss://bucket/object.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// 1.  OSS: an Object Storage Service (OSS) object.
	//
	// 2.  Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSnapshotJobsResponseBodyJobsInput) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotJobsResponseBodyJobsInput) GoString() string {
	return s.String()
}

func (s *ListSnapshotJobsResponseBodyJobsInput) GetMedia() *string {
	return s.Media
}

func (s *ListSnapshotJobsResponseBodyJobsInput) GetType() *string {
	return s.Type
}

func (s *ListSnapshotJobsResponseBodyJobsInput) SetMedia(v string) *ListSnapshotJobsResponseBodyJobsInput {
	s.Media = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobsInput) SetType(v string) *ListSnapshotJobsResponseBodyJobsInput {
	s.Type = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobsInput) Validate() error {
	return dara.Validate(s)
}

type ListSnapshotJobsResponseBodyJobsOutput struct {
	// The output file. If Type is set to OSS, the URL of an OSS object is returned. If Type is set to Media, the ID of a media asset is returned. The URL of an OSS object can be in one of the following formats:
	//
	// 1.  OSS://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	//
	// In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS. If multiple static snapshots were captured, the object must contain the "{Count}" placeholder. In the case of a sprite, the object must contain the "{TileCount}" placeholder. The suffix of the WebVTT snapshot objects must be ".vtt".
	//
	// example:
	//
	// http://test-bucket.oss-cn-shanghai.aliyuncs.com/output-{Count}.jpg
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output file. Valid values:
	//
	// 1.  OSS: an OSS object.
	//
	// 2.  Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSnapshotJobsResponseBodyJobsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotJobsResponseBodyJobsOutput) GoString() string {
	return s.String()
}

func (s *ListSnapshotJobsResponseBodyJobsOutput) GetMedia() *string {
	return s.Media
}

func (s *ListSnapshotJobsResponseBodyJobsOutput) GetType() *string {
	return s.Type
}

func (s *ListSnapshotJobsResponseBodyJobsOutput) SetMedia(v string) *ListSnapshotJobsResponseBodyJobsOutput {
	s.Media = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobsOutput) SetType(v string) *ListSnapshotJobsResponseBodyJobsOutput {
	s.Type = &v
	return s
}

func (s *ListSnapshotJobsResponseBodyJobsOutput) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSnapshotJobResponseBody
	GetRequestId() *string
	SetSnapshotJob(v *GetSnapshotJobResponseBodySnapshotJob) *GetSnapshotJobResponseBody
	GetSnapshotJob() *GetSnapshotJobResponseBodySnapshotJob
}

type GetSnapshotJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the snapshot job.
	SnapshotJob *GetSnapshotJobResponseBodySnapshotJob `json:"SnapshotJob,omitempty" xml:"SnapshotJob,omitempty" type:"Struct"`
}

func (s GetSnapshotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSnapshotJobResponseBody) GetSnapshotJob() *GetSnapshotJobResponseBodySnapshotJob {
	return s.SnapshotJob
}

func (s *GetSnapshotJobResponseBody) SetRequestId(v string) *GetSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSnapshotJobResponseBody) SetSnapshotJob(v *GetSnapshotJobResponseBodySnapshotJob) *GetSnapshotJobResponseBody {
	s.SnapshotJob = v
	return s
}

func (s *GetSnapshotJobResponseBody) Validate() error {
	if s.SnapshotJob != nil {
		if err := s.SnapshotJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSnapshotJobResponseBodySnapshotJob struct {
	// Indicates whether the snapshots were captured in asynchronous mode. Default value: true.
	//
	// example:
	//
	// true
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// Error codes
	//
	// example:
	//
	// ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of snapshots.
	//
	// example:
	//
	// 8
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
	Input *GetSnapshotJobResponseBodySnapshotJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// The specified resource for "Pipeline" could not be found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2022-07-12T16:30:54Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the job.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the job.
	Output *GetSnapshotJobResponseBodySnapshotJobOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The ID of the MPS queue to which the job was submitted.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The state of the job.
	//
	// Valid values:
	//
	// 	- Init: The job is submitted.
	//
	// 	- Success: The job is successful.
	//
	// 	- Fail: The job failed.
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
	// The snapshot template configuration.
	//
	// example:
	//
	// {"Type":"Normal","FrameType":"normal","Time":0,"Count":10}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
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
	// Snapshot types
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
	// The user-defined parameters.
	//
	// example:
	//
	// {"test parameter": "test value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetSnapshotJobResponseBodySnapshotJob) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotJobResponseBodySnapshotJob) GoString() string {
	return s.String()
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetAsync() *bool {
	return s.Async
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetCode() *string {
	return s.Code
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetCount() *int32 {
	return s.Count
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetInput() *GetSnapshotJobResponseBodySnapshotJobInput {
	return s.Input
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetJobId() *string {
	return s.JobId
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetMessage() *string {
	return s.Message
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetName() *string {
	return s.Name
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetOutput() *GetSnapshotJobResponseBodySnapshotJobOutput {
	return s.Output
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetStatus() *string {
	return s.Status
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetType() *string {
	return s.Type
}

func (s *GetSnapshotJobResponseBodySnapshotJob) GetUserData() *string {
	return s.UserData
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetAsync(v bool) *GetSnapshotJobResponseBodySnapshotJob {
	s.Async = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetCode(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.Code = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetCount(v int32) *GetSnapshotJobResponseBodySnapshotJob {
	s.Count = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetCreateTime(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.CreateTime = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetFinishTime(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.FinishTime = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetInput(v *GetSnapshotJobResponseBodySnapshotJobInput) *GetSnapshotJobResponseBodySnapshotJob {
	s.Input = v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetJobId(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.JobId = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetMessage(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.Message = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetModifiedTime(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.ModifiedTime = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetName(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.Name = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetOutput(v *GetSnapshotJobResponseBodySnapshotJobOutput) *GetSnapshotJobResponseBodySnapshotJob {
	s.Output = v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetPipelineId(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.PipelineId = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetStatus(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.Status = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetSubmitTime(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.SubmitTime = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetTemplateConfig(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.TemplateConfig = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetTemplateId(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.TemplateId = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetTriggerSource(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.TriggerSource = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetType(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.Type = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) SetUserData(v string) *GetSnapshotJobResponseBodySnapshotJob {
	s.UserData = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJob) Validate() error {
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

type GetSnapshotJobResponseBodySnapshotJobInput struct {
	// The input file. If Type is set to OSS, the URL of an OSS object is returned. If Type is set to Media, the ID of a media asset is returned. The URL of an OSS object can be in one of the following formats:
	//
	// 1.  oss://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS.
	//
	// example:
	//
	// oss://test-bucket/object.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The three key elements of OSS.
	OssFile *GetSnapshotJobResponseBodySnapshotJobInputOssFile `json:"OssFile,omitempty" xml:"OssFile,omitempty" type:"Struct"`
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

func (s GetSnapshotJobResponseBodySnapshotJobInput) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotJobResponseBodySnapshotJobInput) GoString() string {
	return s.String()
}

func (s *GetSnapshotJobResponseBodySnapshotJobInput) GetMedia() *string {
	return s.Media
}

func (s *GetSnapshotJobResponseBodySnapshotJobInput) GetOssFile() *GetSnapshotJobResponseBodySnapshotJobInputOssFile {
	return s.OssFile
}

func (s *GetSnapshotJobResponseBodySnapshotJobInput) GetType() *string {
	return s.Type
}

func (s *GetSnapshotJobResponseBodySnapshotJobInput) SetMedia(v string) *GetSnapshotJobResponseBodySnapshotJobInput {
	s.Media = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobInput) SetOssFile(v *GetSnapshotJobResponseBodySnapshotJobInputOssFile) *GetSnapshotJobResponseBodySnapshotJobInput {
	s.OssFile = v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobInput) SetType(v string) *GetSnapshotJobResponseBodySnapshotJobInput {
	s.Type = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobInput) Validate() error {
	if s.OssFile != nil {
		if err := s.OssFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSnapshotJobResponseBodySnapshotJobInputOssFile struct {
	// The OSS bucket.
	//
	// example:
	//
	// test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS location.
	//
	// example:
	//
	// oss-cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The OSS object.
	//
	// example:
	//
	// object.mp4
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s GetSnapshotJobResponseBodySnapshotJobInputOssFile) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotJobResponseBodySnapshotJobInputOssFile) GoString() string {
	return s.String()
}

func (s *GetSnapshotJobResponseBodySnapshotJobInputOssFile) GetBucket() *string {
	return s.Bucket
}

func (s *GetSnapshotJobResponseBodySnapshotJobInputOssFile) GetLocation() *string {
	return s.Location
}

func (s *GetSnapshotJobResponseBodySnapshotJobInputOssFile) GetObject() *string {
	return s.Object
}

func (s *GetSnapshotJobResponseBodySnapshotJobInputOssFile) SetBucket(v string) *GetSnapshotJobResponseBodySnapshotJobInputOssFile {
	s.Bucket = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobInputOssFile) SetLocation(v string) *GetSnapshotJobResponseBodySnapshotJobInputOssFile {
	s.Location = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobInputOssFile) SetObject(v string) *GetSnapshotJobResponseBodySnapshotJobInputOssFile {
	s.Object = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobInputOssFile) Validate() error {
	return dara.Validate(s)
}

type GetSnapshotJobResponseBodySnapshotJobOutput struct {
	// The output file. If Type is set to OSS, the URL of an OSS object is returned. If Type is set to Media, the ID of a media asset is returned. The URL of an OSS object can be in one of the following formats:
	//
	// 1.  oss://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	//
	// In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS. If multiple static snapshots were captured, the object must contain the "{Count}" placeholder. In the case of a sprite, the object must contain the "{TileCount}" placeholder. The suffix of the WebVTT snapshot objects must be ".vtt".
	//
	// example:
	//
	// http://test-bucket.oss-cn-shanghai.aliyuncs.com/output-{Count}.jpg
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The three key elements of OSS.
	OssFile *GetSnapshotJobResponseBodySnapshotJobOutputOssFile `json:"OssFile,omitempty" xml:"OssFile,omitempty" type:"Struct"`
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

func (s GetSnapshotJobResponseBodySnapshotJobOutput) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotJobResponseBodySnapshotJobOutput) GoString() string {
	return s.String()
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutput) GetMedia() *string {
	return s.Media
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutput) GetOssFile() *GetSnapshotJobResponseBodySnapshotJobOutputOssFile {
	return s.OssFile
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutput) GetType() *string {
	return s.Type
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutput) SetMedia(v string) *GetSnapshotJobResponseBodySnapshotJobOutput {
	s.Media = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutput) SetOssFile(v *GetSnapshotJobResponseBodySnapshotJobOutputOssFile) *GetSnapshotJobResponseBodySnapshotJobOutput {
	s.OssFile = v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutput) SetType(v string) *GetSnapshotJobResponseBodySnapshotJobOutput {
	s.Type = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutput) Validate() error {
	if s.OssFile != nil {
		if err := s.OssFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSnapshotJobResponseBodySnapshotJobOutputOssFile struct {
	// The OSS bucket.
	//
	// example:
	//
	// test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS location.
	//
	// example:
	//
	// oss-cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The OSS object.
	//
	// example:
	//
	// output-{Count}.jpg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s GetSnapshotJobResponseBodySnapshotJobOutputOssFile) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotJobResponseBodySnapshotJobOutputOssFile) GoString() string {
	return s.String()
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutputOssFile) GetBucket() *string {
	return s.Bucket
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutputOssFile) GetLocation() *string {
	return s.Location
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutputOssFile) GetObject() *string {
	return s.Object
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutputOssFile) SetBucket(v string) *GetSnapshotJobResponseBodySnapshotJobOutputOssFile {
	s.Bucket = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutputOssFile) SetLocation(v string) *GetSnapshotJobResponseBodySnapshotJobOutputOssFile {
	s.Location = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutputOssFile) SetObject(v string) *GetSnapshotJobResponseBodySnapshotJobOutputOssFile {
	s.Object = &v
	return s
}

func (s *GetSnapshotJobResponseBodySnapshotJobOutputOssFile) Validate() error {
	return dara.Validate(s)
}

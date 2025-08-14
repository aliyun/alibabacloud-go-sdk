// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDynamicImageJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicImageJob(v *GetDynamicImageJobResponseBodyDynamicImageJob) *GetDynamicImageJobResponseBody
	GetDynamicImageJob() *GetDynamicImageJobResponseBodyDynamicImageJob
	SetRequestId(v string) *GetDynamicImageJobResponseBody
	GetRequestId() *string
}

type GetDynamicImageJobResponseBody struct {
	// The information about the snapshot job.
	DynamicImageJob *GetDynamicImageJobResponseBodyDynamicImageJob `json:"DynamicImageJob,omitempty" xml:"DynamicImageJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******36-3C1E-4417-BDB2-1E034F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDynamicImageJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDynamicImageJobResponseBody) GetDynamicImageJob() *GetDynamicImageJobResponseBodyDynamicImageJob {
	return s.DynamicImageJob
}

func (s *GetDynamicImageJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDynamicImageJobResponseBody) SetDynamicImageJob(v *GetDynamicImageJobResponseBodyDynamicImageJob) *GetDynamicImageJobResponseBody {
	s.DynamicImageJob = v
	return s
}

func (s *GetDynamicImageJobResponseBody) SetRequestId(v string) *GetDynamicImageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDynamicImageJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDynamicImageJobResponseBodyDynamicImageJob struct {
	// Error codes
	//
	// example:
	//
	// ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	Input *GetDynamicImageJobResponseBodyDynamicImageJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
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
	// The specified resource for "CustomTemplate" could not be found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Output *GetDynamicImageJobResponseBodyDynamicImageJobOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The URL of the output animated image.
	//
	// example:
	//
	// http://test-bucket.oss-cn-shanghai.aliyuncs.com/output.gif
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
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
	// The animation template configuration.
	//
	// example:
	//
	// {"Format":"gif","Fps":5,"Height":1080,"Width":1920}
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
	// The user-defined data.
	//
	// example:
	//
	// {"sampleParam": "sampleValue"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetDynamicImageJobResponseBodyDynamicImageJob) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicImageJobResponseBodyDynamicImageJob) GoString() string {
	return s.String()
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetCode() *string {
	return s.Code
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetInput() *GetDynamicImageJobResponseBodyDynamicImageJobInput {
	return s.Input
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetJobId() *string {
	return s.JobId
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetMessage() *string {
	return s.Message
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetName() *string {
	return s.Name
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetOutput() *GetDynamicImageJobResponseBodyDynamicImageJobOutput {
	return s.Output
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetStatus() *string {
	return s.Status
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) GetUserData() *string {
	return s.UserData
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetCode(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.Code = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetCreateTime(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.CreateTime = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetFinishTime(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.FinishTime = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetInput(v *GetDynamicImageJobResponseBodyDynamicImageJobInput) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.Input = v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetJobId(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.JobId = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetMessage(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.Message = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetModifiedTime(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.ModifiedTime = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetName(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.Name = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetOutput(v *GetDynamicImageJobResponseBodyDynamicImageJobOutput) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.Output = v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetOutputUrl(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.OutputUrl = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetPipelineId(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.PipelineId = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetStatus(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.Status = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetSubmitTime(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.SubmitTime = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetTemplateConfig(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.TemplateConfig = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetTemplateId(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.TemplateId = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetTriggerSource(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.TriggerSource = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) SetUserData(v string) *GetDynamicImageJobResponseBodyDynamicImageJob {
	s.UserData = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJob) Validate() error {
	return dara.Validate(s)
}

type GetDynamicImageJobResponseBodyDynamicImageJobInput struct {
	// The input file. If Type is set to OSS, the URL of an OSS object is returned. If Type is set to Media, the ID of a media asset is returned. The URL of an OSS object can be in one of the following formats:
	//
	// 1.  OSS://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	//
	// In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS.
	//
	// example:
	//
	// oss://test-bucket/sample-input.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The three key elements of OSS.
	OssFile *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile `json:"OssFile,omitempty" xml:"OssFile,omitempty" type:"Struct"`
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

func (s GetDynamicImageJobResponseBodyDynamicImageJobInput) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicImageJobResponseBodyDynamicImageJobInput) GoString() string {
	return s.String()
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInput) GetMedia() *string {
	return s.Media
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInput) GetOssFile() *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile {
	return s.OssFile
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInput) GetType() *string {
	return s.Type
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInput) SetMedia(v string) *GetDynamicImageJobResponseBodyDynamicImageJobInput {
	s.Media = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInput) SetOssFile(v *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) *GetDynamicImageJobResponseBodyDynamicImageJobInput {
	s.OssFile = v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInput) SetType(v string) *GetDynamicImageJobResponseBodyDynamicImageJobInput {
	s.Type = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInput) Validate() error {
	return dara.Validate(s)
}

type GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile struct {
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
	// sample-input.mp4
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) GoString() string {
	return s.String()
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) GetBucket() *string {
	return s.Bucket
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) GetLocation() *string {
	return s.Location
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) GetObject() *string {
	return s.Object
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) SetBucket(v string) *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile {
	s.Bucket = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) SetLocation(v string) *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile {
	s.Location = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) SetObject(v string) *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile {
	s.Object = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobInputOssFile) Validate() error {
	return dara.Validate(s)
}

type GetDynamicImageJobResponseBodyDynamicImageJobOutput struct {
	// The input file. If Type is set to OSS, the URL of an OSS object is returned. If Type is set to Media, the ID of a media asset is returned. The URL of an OSS object can be in one of the following formats:
	//
	// 1.  OSS://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The three key elements of OSS.
	OssFile *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile `json:"OssFile,omitempty" xml:"OssFile,omitempty" type:"Struct"`
	// The type of the input file. Valid values: OSS: an OSS object. Media: a media asset.
	//
	// example:
	//
	// Media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDynamicImageJobResponseBodyDynamicImageJobOutput) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicImageJobResponseBodyDynamicImageJobOutput) GoString() string {
	return s.String()
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutput) GetMedia() *string {
	return s.Media
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutput) GetOssFile() *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile {
	return s.OssFile
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutput) GetType() *string {
	return s.Type
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutput) SetMedia(v string) *GetDynamicImageJobResponseBodyDynamicImageJobOutput {
	s.Media = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutput) SetOssFile(v *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) *GetDynamicImageJobResponseBodyDynamicImageJobOutput {
	s.OssFile = v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutput) SetType(v string) *GetDynamicImageJobResponseBodyDynamicImageJobOutput {
	s.Type = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutput) Validate() error {
	return dara.Validate(s)
}

type GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile struct {
	// The OSS bucket.
	//
	// example:
	//
	// sample-bucket
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
	// path/to/object
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) GoString() string {
	return s.String()
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) GetBucket() *string {
	return s.Bucket
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) GetLocation() *string {
	return s.Location
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) GetObject() *string {
	return s.Object
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) SetBucket(v string) *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile {
	s.Bucket = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) SetLocation(v string) *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile {
	s.Location = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) SetObject(v string) *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile {
	s.Object = &v
	return s
}

func (s *GetDynamicImageJobResponseBodyDynamicImageJobOutputOssFile) Validate() error {
	return dara.Validate(s)
}

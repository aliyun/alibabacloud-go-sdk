// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveRecordJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordJob(v *GetLiveRecordJobResponseBodyRecordJob) *GetLiveRecordJobResponseBody
	GetRecordJob() *GetLiveRecordJobResponseBodyRecordJob
	SetRequestId(v string) *GetLiveRecordJobResponseBody
	GetRequestId() *string
}

type GetLiveRecordJobResponseBody struct {
	// The details of the recording job.
	RecordJob *GetLiveRecordJobResponseBodyRecordJob `json:"RecordJob,omitempty" xml:"RecordJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B57A046C-CE33-5FBB-B57A-D2B89ACF6907
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLiveRecordJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRecordJobResponseBody) GetRecordJob() *GetLiveRecordJobResponseBodyRecordJob {
	return s.RecordJob
}

func (s *GetLiveRecordJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveRecordJobResponseBody) SetRecordJob(v *GetLiveRecordJobResponseBodyRecordJob) *GetLiveRecordJobResponseBody {
	s.RecordJob = v
	return s
}

func (s *GetLiveRecordJobResponseBody) SetRequestId(v string) *GetLiveRecordJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRecordJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLiveRecordJobResponseBodyRecordJob struct {
	// The time when the job was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-07-20T02:48:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the recording job.
	//
	// example:
	//
	// ab0e3e76-1e9d-11ed-ba64-0c42a1b73d66
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the recording job.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// https://example.com/imsnotify
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The storage address of the recording.
	RecordOutput *GetLiveRecordJobResponseBodyRecordJobRecordOutput `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty" type:"Struct"`
	// The state of the recording job.
	//
	// Valid values:
	//
	// 	- paused: The job is paused.
	//
	// 	- initial: The job is not started.
	//
	// 	- started: The job is in progress.
	//
	// example:
	//
	// paused
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The URL of the live stream.
	StreamInput *GetLiveRecordJobResponseBodyRecordJobStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The ID of the recording template.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the recording template.
	//
	// example:
	//
	// test template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetLiveRecordJobResponseBodyRecordJob) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordJobResponseBodyRecordJob) GoString() string {
	return s.String()
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetName() *string {
	return s.Name
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetRecordOutput() *GetLiveRecordJobResponseBodyRecordJobRecordOutput {
	return s.RecordOutput
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetStatus() *string {
	return s.Status
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetStreamInput() *GetLiveRecordJobResponseBodyRecordJobStreamInput {
	return s.StreamInput
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveRecordJobResponseBodyRecordJob) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetCreateTime(v string) *GetLiveRecordJobResponseBodyRecordJob {
	s.CreateTime = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetJobId(v string) *GetLiveRecordJobResponseBodyRecordJob {
	s.JobId = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetName(v string) *GetLiveRecordJobResponseBodyRecordJob {
	s.Name = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetNotifyUrl(v string) *GetLiveRecordJobResponseBodyRecordJob {
	s.NotifyUrl = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetRecordOutput(v *GetLiveRecordJobResponseBodyRecordJobRecordOutput) *GetLiveRecordJobResponseBodyRecordJob {
	s.RecordOutput = v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetStatus(v string) *GetLiveRecordJobResponseBodyRecordJob {
	s.Status = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetStreamInput(v *GetLiveRecordJobResponseBodyRecordJobStreamInput) *GetLiveRecordJobResponseBodyRecordJob {
	s.StreamInput = v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetTemplateId(v string) *GetLiveRecordJobResponseBodyRecordJob {
	s.TemplateId = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) SetTemplateName(v string) *GetLiveRecordJobResponseBodyRecordJob {
	s.TemplateName = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJob) Validate() error {
	return dara.Validate(s)
}

type GetLiveRecordJobResponseBodyRecordJobRecordOutput struct {
	// The bucket name.
	//
	// example:
	//
	// imsbucket1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The endpoint of the storage service.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The type of the storage address.
	//
	// Valid values:
	//
	// 	- vod
	//
	// 	- oss
	//
	// example:
	//
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLiveRecordJobResponseBodyRecordJobRecordOutput) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordJobResponseBodyRecordJobRecordOutput) GoString() string {
	return s.String()
}

func (s *GetLiveRecordJobResponseBodyRecordJobRecordOutput) GetBucket() *string {
	return s.Bucket
}

func (s *GetLiveRecordJobResponseBodyRecordJobRecordOutput) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetLiveRecordJobResponseBodyRecordJobRecordOutput) GetType() *string {
	return s.Type
}

func (s *GetLiveRecordJobResponseBodyRecordJobRecordOutput) SetBucket(v string) *GetLiveRecordJobResponseBodyRecordJobRecordOutput {
	s.Bucket = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJobRecordOutput) SetEndpoint(v string) *GetLiveRecordJobResponseBodyRecordJobRecordOutput {
	s.Endpoint = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJobRecordOutput) SetType(v string) *GetLiveRecordJobResponseBodyRecordJobRecordOutput {
	s.Type = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJobRecordOutput) Validate() error {
	return dara.Validate(s)
}

type GetLiveRecordJobResponseBodyRecordJobStreamInput struct {
	// The type of the live stream. The value can only be rtmp.
	//
	// example:
	//
	// rtmp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the live stream.
	//
	// example:
	//
	// rtmp://example.com/app/stream
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetLiveRecordJobResponseBodyRecordJobStreamInput) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordJobResponseBodyRecordJobStreamInput) GoString() string {
	return s.String()
}

func (s *GetLiveRecordJobResponseBodyRecordJobStreamInput) GetType() *string {
	return s.Type
}

func (s *GetLiveRecordJobResponseBodyRecordJobStreamInput) GetUrl() *string {
	return s.Url
}

func (s *GetLiveRecordJobResponseBodyRecordJobStreamInput) SetType(v string) *GetLiveRecordJobResponseBodyRecordJobStreamInput {
	s.Type = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJobStreamInput) SetUrl(v string) *GetLiveRecordJobResponseBodyRecordJobStreamInput {
	s.Url = &v
	return s
}

func (s *GetLiveRecordJobResponseBodyRecordJobStreamInput) Validate() error {
	return dara.Validate(s)
}

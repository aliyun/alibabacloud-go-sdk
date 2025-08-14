// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveSnapshotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *GetLiveSnapshotJobResponseBody
	GetCallbackUrl() *string
	SetCreateTime(v string) *GetLiveSnapshotJobResponseBody
	GetCreateTime() *string
	SetJobId(v string) *GetLiveSnapshotJobResponseBody
	GetJobId() *string
	SetJobName(v string) *GetLiveSnapshotJobResponseBody
	GetJobName() *string
	SetLastModified(v string) *GetLiveSnapshotJobResponseBody
	GetLastModified() *string
	SetOverwriteFormat(v string) *GetLiveSnapshotJobResponseBody
	GetOverwriteFormat() *string
	SetRequestId(v string) *GetLiveSnapshotJobResponseBody
	GetRequestId() *string
	SetSequenceFormat(v string) *GetLiveSnapshotJobResponseBody
	GetSequenceFormat() *string
	SetSnapshotOutput(v *GetLiveSnapshotJobResponseBodySnapshotOutput) *GetLiveSnapshotJobResponseBody
	GetSnapshotOutput() *GetLiveSnapshotJobResponseBodySnapshotOutput
	SetStatus(v string) *GetLiveSnapshotJobResponseBody
	GetStatus() *string
	SetStreamInput(v *GetLiveSnapshotJobResponseBodyStreamInput) *GetLiveSnapshotJobResponseBody
	GetStreamInput() *GetLiveSnapshotJobResponseBodyStreamInput
	SetTemplateId(v string) *GetLiveSnapshotJobResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *GetLiveSnapshotJobResponseBody
	GetTemplateName() *string
	SetTimeInterval(v int32) *GetLiveSnapshotJobResponseBody
	GetTimeInterval() *int32
}

type GetLiveSnapshotJobResponseBody struct {
	// The snapshot callback URL.
	//
	// example:
	//
	// http://www.aliyun.com/snapshot/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The time when the file was created.
	//
	// example:
	//
	// 2022-02-02T22:22:22Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the job.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The time when the file was last modified.
	//
	// example:
	//
	// 2022-02-02T22:22:22Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The naming format of the snapshot captured in overwrite mode.
	//
	// example:
	//
	// snapshot/{JobId}.jpg
	OverwriteFormat *string `json:"OverwriteFormat,omitempty" xml:"OverwriteFormat,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The naming format of the snapshot captured in time series mode.
	//
	// example:
	//
	// snapshot/{JobId}/{UnixTimestamp}.jpg
	SequenceFormat *string `json:"SequenceFormat,omitempty" xml:"SequenceFormat,omitempty"`
	// The output information.
	SnapshotOutput *GetLiveSnapshotJobResponseBodySnapshotOutput `json:"SnapshotOutput,omitempty" xml:"SnapshotOutput,omitempty" type:"Struct"`
	// The state of the job.
	//
	// Valid values:
	//
	// 	- init: The job is not started.
	//
	// 	- paused: The job is paused.
	//
	// 	- started: The job is in progress.
	//
	// example:
	//
	// started
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The input information.
	StreamInput *GetLiveSnapshotJobResponseBodyStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287666****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The interval between two adjacent snapshots.
	//
	// example:
	//
	// 5
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s GetLiveSnapshotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveSnapshotJobResponseBody) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *GetLiveSnapshotJobResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLiveSnapshotJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveSnapshotJobResponseBody) GetJobName() *string {
	return s.JobName
}

func (s *GetLiveSnapshotJobResponseBody) GetLastModified() *string {
	return s.LastModified
}

func (s *GetLiveSnapshotJobResponseBody) GetOverwriteFormat() *string {
	return s.OverwriteFormat
}

func (s *GetLiveSnapshotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveSnapshotJobResponseBody) GetSequenceFormat() *string {
	return s.SequenceFormat
}

func (s *GetLiveSnapshotJobResponseBody) GetSnapshotOutput() *GetLiveSnapshotJobResponseBodySnapshotOutput {
	return s.SnapshotOutput
}

func (s *GetLiveSnapshotJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetLiveSnapshotJobResponseBody) GetStreamInput() *GetLiveSnapshotJobResponseBodyStreamInput {
	return s.StreamInput
}

func (s *GetLiveSnapshotJobResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveSnapshotJobResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetLiveSnapshotJobResponseBody) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *GetLiveSnapshotJobResponseBody) SetCallbackUrl(v string) *GetLiveSnapshotJobResponseBody {
	s.CallbackUrl = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetCreateTime(v string) *GetLiveSnapshotJobResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetJobId(v string) *GetLiveSnapshotJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetJobName(v string) *GetLiveSnapshotJobResponseBody {
	s.JobName = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetLastModified(v string) *GetLiveSnapshotJobResponseBody {
	s.LastModified = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetOverwriteFormat(v string) *GetLiveSnapshotJobResponseBody {
	s.OverwriteFormat = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetRequestId(v string) *GetLiveSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetSequenceFormat(v string) *GetLiveSnapshotJobResponseBody {
	s.SequenceFormat = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetSnapshotOutput(v *GetLiveSnapshotJobResponseBodySnapshotOutput) *GetLiveSnapshotJobResponseBody {
	s.SnapshotOutput = v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetStatus(v string) *GetLiveSnapshotJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetStreamInput(v *GetLiveSnapshotJobResponseBodyStreamInput) *GetLiveSnapshotJobResponseBody {
	s.StreamInput = v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetTemplateId(v string) *GetLiveSnapshotJobResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetTemplateName(v string) *GetLiveSnapshotJobResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) SetTimeInterval(v int32) *GetLiveSnapshotJobResponseBody {
	s.TimeInterval = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLiveSnapshotJobResponseBodySnapshotOutput struct {
	// The bucket of the output endpoint. If the storage type is set to oss, the OSS bucket is returned.
	//
	// example:
	//
	// testbucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The output endpoint. If the storage type is set to oss, the Object Storage Service (OSS) domain name is returned.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The storage type. The value can only be oss.
	//
	// example:
	//
	// oss
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s GetLiveSnapshotJobResponseBodySnapshotOutput) String() string {
	return dara.Prettify(s)
}

func (s GetLiveSnapshotJobResponseBodySnapshotOutput) GoString() string {
	return s.String()
}

func (s *GetLiveSnapshotJobResponseBodySnapshotOutput) GetBucket() *string {
	return s.Bucket
}

func (s *GetLiveSnapshotJobResponseBodySnapshotOutput) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetLiveSnapshotJobResponseBodySnapshotOutput) GetStorageType() *string {
	return s.StorageType
}

func (s *GetLiveSnapshotJobResponseBodySnapshotOutput) SetBucket(v string) *GetLiveSnapshotJobResponseBodySnapshotOutput {
	s.Bucket = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBodySnapshotOutput) SetEndpoint(v string) *GetLiveSnapshotJobResponseBodySnapshotOutput {
	s.Endpoint = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBodySnapshotOutput) SetStorageType(v string) *GetLiveSnapshotJobResponseBodySnapshotOutput {
	s.StorageType = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBodySnapshotOutput) Validate() error {
	return dara.Validate(s)
}

type GetLiveSnapshotJobResponseBodyStreamInput struct {
	// The type of the input stream. The value can only be rtmp.
	//
	// example:
	//
	// rtmp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the input stream.
	//
	// example:
	//
	// rtmp://www.aliyun.com/stream
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetLiveSnapshotJobResponseBodyStreamInput) String() string {
	return dara.Prettify(s)
}

func (s GetLiveSnapshotJobResponseBodyStreamInput) GoString() string {
	return s.String()
}

func (s *GetLiveSnapshotJobResponseBodyStreamInput) GetType() *string {
	return s.Type
}

func (s *GetLiveSnapshotJobResponseBodyStreamInput) GetUrl() *string {
	return s.Url
}

func (s *GetLiveSnapshotJobResponseBodyStreamInput) SetType(v string) *GetLiveSnapshotJobResponseBodyStreamInput {
	s.Type = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBodyStreamInput) SetUrl(v string) *GetLiveSnapshotJobResponseBodyStreamInput {
	s.Url = &v
	return s
}

func (s *GetLiveSnapshotJobResponseBodyStreamInput) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *SubmitLiveSnapshotJobRequest
	GetCallbackUrl() *string
	SetJobName(v string) *SubmitLiveSnapshotJobRequest
	GetJobName() *string
	SetSnapshotOutput(v *SubmitLiveSnapshotJobRequestSnapshotOutput) *SubmitLiveSnapshotJobRequest
	GetSnapshotOutput() *SubmitLiveSnapshotJobRequestSnapshotOutput
	SetStreamInput(v *SubmitLiveSnapshotJobRequestStreamInput) *SubmitLiveSnapshotJobRequest
	GetStreamInput() *SubmitLiveSnapshotJobRequestStreamInput
	SetTemplateId(v string) *SubmitLiveSnapshotJobRequest
	GetTemplateId() *string
}

type SubmitLiveSnapshotJobRequest struct {
	// The snapshot callback URL.
	//
	// 	- It cannot exceed 255 characters in length.
	//
	// 	- Both HTTP and HTTPS URLs are supported.
	//
	// example:
	//
	// http://www.aliyun.com/snapshot/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The name of the job.
	//
	// 	- It cannot exceed 128 characters in length.
	//
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The information about the output snapshot.
	//
	// This parameter is required.
	SnapshotOutput *SubmitLiveSnapshotJobRequestSnapshotOutput `json:"SnapshotOutput,omitempty" xml:"SnapshotOutput,omitempty" type:"Struct"`
	// The information about the input stream.
	//
	// This parameter is required.
	StreamInput *SubmitLiveSnapshotJobRequestStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitLiveSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveSnapshotJobRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *SubmitLiveSnapshotJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *SubmitLiveSnapshotJobRequest) GetSnapshotOutput() *SubmitLiveSnapshotJobRequestSnapshotOutput {
	return s.SnapshotOutput
}

func (s *SubmitLiveSnapshotJobRequest) GetStreamInput() *SubmitLiveSnapshotJobRequestStreamInput {
	return s.StreamInput
}

func (s *SubmitLiveSnapshotJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitLiveSnapshotJobRequest) SetCallbackUrl(v string) *SubmitLiveSnapshotJobRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SubmitLiveSnapshotJobRequest) SetJobName(v string) *SubmitLiveSnapshotJobRequest {
	s.JobName = &v
	return s
}

func (s *SubmitLiveSnapshotJobRequest) SetSnapshotOutput(v *SubmitLiveSnapshotJobRequestSnapshotOutput) *SubmitLiveSnapshotJobRequest {
	s.SnapshotOutput = v
	return s
}

func (s *SubmitLiveSnapshotJobRequest) SetStreamInput(v *SubmitLiveSnapshotJobRequestStreamInput) *SubmitLiveSnapshotJobRequest {
	s.StreamInput = v
	return s
}

func (s *SubmitLiveSnapshotJobRequest) SetTemplateId(v string) *SubmitLiveSnapshotJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitLiveSnapshotJobRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitLiveSnapshotJobRequestSnapshotOutput struct {
	// The bucket of the snapshot output endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// testbucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The output endpoint of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The storage type of the snapshot. The value can only be oss.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s SubmitLiveSnapshotJobRequestSnapshotOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveSnapshotJobRequestSnapshotOutput) GoString() string {
	return s.String()
}

func (s *SubmitLiveSnapshotJobRequestSnapshotOutput) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitLiveSnapshotJobRequestSnapshotOutput) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SubmitLiveSnapshotJobRequestSnapshotOutput) GetStorageType() *string {
	return s.StorageType
}

func (s *SubmitLiveSnapshotJobRequestSnapshotOutput) SetBucket(v string) *SubmitLiveSnapshotJobRequestSnapshotOutput {
	s.Bucket = &v
	return s
}

func (s *SubmitLiveSnapshotJobRequestSnapshotOutput) SetEndpoint(v string) *SubmitLiveSnapshotJobRequestSnapshotOutput {
	s.Endpoint = &v
	return s
}

func (s *SubmitLiveSnapshotJobRequestSnapshotOutput) SetStorageType(v string) *SubmitLiveSnapshotJobRequestSnapshotOutput {
	s.StorageType = &v
	return s
}

func (s *SubmitLiveSnapshotJobRequestSnapshotOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitLiveSnapshotJobRequestStreamInput struct {
	// The type of the input stream. The value can only be rtmp.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the input stream.
	//
	// 	- It cannot exceed 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://www.aliyun.com/stream
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitLiveSnapshotJobRequestStreamInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveSnapshotJobRequestStreamInput) GoString() string {
	return s.String()
}

func (s *SubmitLiveSnapshotJobRequestStreamInput) GetType() *string {
	return s.Type
}

func (s *SubmitLiveSnapshotJobRequestStreamInput) GetUrl() *string {
	return s.Url
}

func (s *SubmitLiveSnapshotJobRequestStreamInput) SetType(v string) *SubmitLiveSnapshotJobRequestStreamInput {
	s.Type = &v
	return s
}

func (s *SubmitLiveSnapshotJobRequestStreamInput) SetUrl(v string) *SubmitLiveSnapshotJobRequestStreamInput {
	s.Url = &v
	return s
}

func (s *SubmitLiveSnapshotJobRequestStreamInput) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveRecordJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SubmitLiveRecordJobRequest
	GetName() *string
	SetNotifyUrl(v string) *SubmitLiveRecordJobRequest
	GetNotifyUrl() *string
	SetRecordOutput(v *SubmitLiveRecordJobRequestRecordOutput) *SubmitLiveRecordJobRequest
	GetRecordOutput() *SubmitLiveRecordJobRequestRecordOutput
	SetStreamInput(v *SubmitLiveRecordJobRequestStreamInput) *SubmitLiveRecordJobRequest
	GetStreamInput() *SubmitLiveRecordJobRequestStreamInput
	SetTemplateId(v string) *SubmitLiveRecordJobRequest
	GetTemplateId() *string
}

type SubmitLiveRecordJobRequest struct {
	// The name of the recording job.
	//
	// This parameter is required.
	//
	// example:
	//
	// live stream record 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// https://example.com/imsnotify
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The storage address of the recording.
	//
	// This parameter is required.
	RecordOutput *SubmitLiveRecordJobRequestRecordOutput `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty" type:"Struct"`
	// The URL of the live stream.
	//
	// This parameter is required.
	StreamInput *SubmitLiveRecordJobRequestStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The ID of the recording template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitLiveRecordJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveRecordJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveRecordJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitLiveRecordJobRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitLiveRecordJobRequest) GetRecordOutput() *SubmitLiveRecordJobRequestRecordOutput {
	return s.RecordOutput
}

func (s *SubmitLiveRecordJobRequest) GetStreamInput() *SubmitLiveRecordJobRequestStreamInput {
	return s.StreamInput
}

func (s *SubmitLiveRecordJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitLiveRecordJobRequest) SetName(v string) *SubmitLiveRecordJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitLiveRecordJobRequest) SetNotifyUrl(v string) *SubmitLiveRecordJobRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitLiveRecordJobRequest) SetRecordOutput(v *SubmitLiveRecordJobRequestRecordOutput) *SubmitLiveRecordJobRequest {
	s.RecordOutput = v
	return s
}

func (s *SubmitLiveRecordJobRequest) SetStreamInput(v *SubmitLiveRecordJobRequestStreamInput) *SubmitLiveRecordJobRequest {
	s.StreamInput = v
	return s
}

func (s *SubmitLiveRecordJobRequest) SetTemplateId(v string) *SubmitLiveRecordJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitLiveRecordJobRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitLiveRecordJobRequestRecordOutput struct {
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
	// oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The type of the storage address.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitLiveRecordJobRequestRecordOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveRecordJobRequestRecordOutput) GoString() string {
	return s.String()
}

func (s *SubmitLiveRecordJobRequestRecordOutput) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitLiveRecordJobRequestRecordOutput) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SubmitLiveRecordJobRequestRecordOutput) GetType() *string {
	return s.Type
}

func (s *SubmitLiveRecordJobRequestRecordOutput) SetBucket(v string) *SubmitLiveRecordJobRequestRecordOutput {
	s.Bucket = &v
	return s
}

func (s *SubmitLiveRecordJobRequestRecordOutput) SetEndpoint(v string) *SubmitLiveRecordJobRequestRecordOutput {
	s.Endpoint = &v
	return s
}

func (s *SubmitLiveRecordJobRequestRecordOutput) SetType(v string) *SubmitLiveRecordJobRequestRecordOutput {
	s.Type = &v
	return s
}

func (s *SubmitLiveRecordJobRequestRecordOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitLiveRecordJobRequestStreamInput struct {
	// The type of the live stream URL. The value can only be rtmp.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the live stream.
	//
	// example:
	//
	// rtmp://example.com/live/stream1
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitLiveRecordJobRequestStreamInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveRecordJobRequestStreamInput) GoString() string {
	return s.String()
}

func (s *SubmitLiveRecordJobRequestStreamInput) GetType() *string {
	return s.Type
}

func (s *SubmitLiveRecordJobRequestStreamInput) GetUrl() *string {
	return s.Url
}

func (s *SubmitLiveRecordJobRequestStreamInput) SetType(v string) *SubmitLiveRecordJobRequestStreamInput {
	s.Type = &v
	return s
}

func (s *SubmitLiveRecordJobRequestStreamInput) SetUrl(v string) *SubmitLiveRecordJobRequestStreamInput {
	s.Url = &v
	return s
}

func (s *SubmitLiveRecordJobRequestStreamInput) Validate() error {
	return dara.Validate(s)
}

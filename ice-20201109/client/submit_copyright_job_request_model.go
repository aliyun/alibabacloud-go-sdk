// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubmitCopyrightJobRequest
	GetDescription() *string
	SetInput(v *SubmitCopyrightJobRequestInput) *SubmitCopyrightJobRequest
	GetInput() *SubmitCopyrightJobRequestInput
	SetLevel(v int64) *SubmitCopyrightJobRequest
	GetLevel() *int64
	SetMessage(v string) *SubmitCopyrightJobRequest
	GetMessage() *string
	SetOutput(v *SubmitCopyrightJobRequestOutput) *SubmitCopyrightJobRequest
	GetOutput() *SubmitCopyrightJobRequestOutput
	SetParams(v string) *SubmitCopyrightJobRequest
	GetParams() *string
	SetStartTime(v int64) *SubmitCopyrightJobRequest
	GetStartTime() *int64
	SetTotalTime(v int64) *SubmitCopyrightJobRequest
	GetTotalTime() *int64
	SetUserData(v string) *SubmitCopyrightJobRequest
	GetUserData() *string
}

type SubmitCopyrightJobRequest struct {
	// A description of the watermark job.
	//
	// example:
	//
	// Task description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input video file to be watermarked.
	//
	// > - The OSS object or media asset must be in the same region as the service call.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket","Location":"oss-cn-shanghai","Object":"example.mp4"}
	Input *SubmitCopyrightJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The watermark level, which specifies the embedding channel. Valid values are 0, 1, and 2, which correspond to the U, UV, and YUV channels, respectively.
	//
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The watermark content to embed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Copyright watermark test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The location of the output file.
	//
	// > - The OSS bucket must be in the same region as the service call.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket","Location":"oss-cn-shanghai","Object":"example_result.mp4"}
	Output *SubmitCopyrightJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The parameters for the watermark job, specified as a JSON string. The following parameter is supported:
	//
	// - `algoType`: The algorithm type. Defaults to `v1`.
	//
	//   - `v1`: For videos 3 minutes or longer.
	//
	//   - `v2`: For short videos.
	//
	// example:
	//
	// {"algoType":"v2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The start time of the watermark in seconds. Defaults to 0.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The end time of the watermark in seconds. If unspecified, the watermark is applied until the video ends.
	//
	// example:
	//
	// 10
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The user data. The value can be up to 1,024 bytes in length.
	//
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitCopyrightJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitCopyrightJobRequest) GetInput() *SubmitCopyrightJobRequestInput {
	return s.Input
}

func (s *SubmitCopyrightJobRequest) GetLevel() *int64 {
	return s.Level
}

func (s *SubmitCopyrightJobRequest) GetMessage() *string {
	return s.Message
}

func (s *SubmitCopyrightJobRequest) GetOutput() *SubmitCopyrightJobRequestOutput {
	return s.Output
}

func (s *SubmitCopyrightJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitCopyrightJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SubmitCopyrightJobRequest) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *SubmitCopyrightJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitCopyrightJobRequest) SetDescription(v string) *SubmitCopyrightJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetInput(v *SubmitCopyrightJobRequestInput) *SubmitCopyrightJobRequest {
	s.Input = v
	return s
}

func (s *SubmitCopyrightJobRequest) SetLevel(v int64) *SubmitCopyrightJobRequest {
	s.Level = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetMessage(v string) *SubmitCopyrightJobRequest {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetOutput(v *SubmitCopyrightJobRequestOutput) *SubmitCopyrightJobRequest {
	s.Output = v
	return s
}

func (s *SubmitCopyrightJobRequest) SetParams(v string) *SubmitCopyrightJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetStartTime(v int64) *SubmitCopyrightJobRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetTotalTime(v int64) *SubmitCopyrightJobRequest {
	s.TotalTime = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetUserData(v string) *SubmitCopyrightJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightJobRequest) Validate() error {
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

type SubmitCopyrightJobRequestInput struct {
	// The input file, specified as either an OSS URL or a media asset ID. The following formats are supported for OSS URLs:
	//
	// 1\\. `oss://bucket/object`
	//
	// 2\\. `http(s)://bucket.oss-[regionId].aliyuncs.com/object`
	//
	// In these formats, `bucket` specifies the name of an OSS bucket in the same region as the service, and `object` specifies the file path.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// 1. `OSS`: The URL of an OSS object.
	//
	// 2. `Media`: The media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitCopyrightJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitCopyrightJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitCopyrightJobRequestInput) SetMedia(v string) *SubmitCopyrightJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitCopyrightJobRequestInput) SetType(v string) *SubmitCopyrightJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitCopyrightJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitCopyrightJobRequestOutput struct {
	// The OSS URL for the output file. The following formats are supported:
	//
	// 1\\. `oss://bucket/object`
	//
	// 2\\. `http(s)://bucket.oss-[regionId].aliyuncs.com/object`<br>In these formats, `bucket` specifies the name of an OSS bucket in the same region as the service, and `object` specifies the file path.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/output.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output file. Valid value:
	//
	// 1. `OSS`: The URL of an OSS object.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitCopyrightJobRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobRequestOutput) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobRequestOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitCopyrightJobRequestOutput) GetType() *string {
	return s.Type
}

func (s *SubmitCopyrightJobRequestOutput) SetMedia(v string) *SubmitCopyrightJobRequestOutput {
	s.Media = &v
	return s
}

func (s *SubmitCopyrightJobRequestOutput) SetType(v string) *SubmitCopyrightJobRequestOutput {
	s.Type = &v
	return s
}

func (s *SubmitCopyrightJobRequestOutput) Validate() error {
	return dara.Validate(s)
}

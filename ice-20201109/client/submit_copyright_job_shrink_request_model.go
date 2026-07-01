// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubmitCopyrightJobShrinkRequest
	GetDescription() *string
	SetInputShrink(v string) *SubmitCopyrightJobShrinkRequest
	GetInputShrink() *string
	SetLevel(v int64) *SubmitCopyrightJobShrinkRequest
	GetLevel() *int64
	SetMessage(v string) *SubmitCopyrightJobShrinkRequest
	GetMessage() *string
	SetOutputShrink(v string) *SubmitCopyrightJobShrinkRequest
	GetOutputShrink() *string
	SetParams(v string) *SubmitCopyrightJobShrinkRequest
	GetParams() *string
	SetStartTime(v int64) *SubmitCopyrightJobShrinkRequest
	GetStartTime() *int64
	SetTotalTime(v int64) *SubmitCopyrightJobShrinkRequest
	GetTotalTime() *int64
	SetUserData(v string) *SubmitCopyrightJobShrinkRequest
	GetUserData() *string
}

type SubmitCopyrightJobShrinkRequest struct {
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
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
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

func (s SubmitCopyrightJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitCopyrightJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitCopyrightJobShrinkRequest) GetLevel() *int64 {
	return s.Level
}

func (s *SubmitCopyrightJobShrinkRequest) GetMessage() *string {
	return s.Message
}

func (s *SubmitCopyrightJobShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *SubmitCopyrightJobShrinkRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitCopyrightJobShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SubmitCopyrightJobShrinkRequest) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *SubmitCopyrightJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitCopyrightJobShrinkRequest) SetDescription(v string) *SubmitCopyrightJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) SetInputShrink(v string) *SubmitCopyrightJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) SetLevel(v int64) *SubmitCopyrightJobShrinkRequest {
	s.Level = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) SetMessage(v string) *SubmitCopyrightJobShrinkRequest {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) SetOutputShrink(v string) *SubmitCopyrightJobShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) SetParams(v string) *SubmitCopyrightJobShrinkRequest {
	s.Params = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) SetStartTime(v int64) *SubmitCopyrightJobShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) SetTotalTime(v int64) *SubmitCopyrightJobShrinkRequest {
	s.TotalTime = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) SetUserData(v string) *SubmitCopyrightJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

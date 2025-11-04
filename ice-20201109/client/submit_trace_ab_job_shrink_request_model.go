// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceAbJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCipherBase64ed(v string) *SubmitTraceAbJobShrinkRequest
	GetCipherBase64ed() *string
	SetInputShrink(v string) *SubmitTraceAbJobShrinkRequest
	GetInputShrink() *string
	SetLevel(v int64) *SubmitTraceAbJobShrinkRequest
	GetLevel() *int64
	SetOutputShrink(v string) *SubmitTraceAbJobShrinkRequest
	GetOutputShrink() *string
	SetStartTime(v int64) *SubmitTraceAbJobShrinkRequest
	GetStartTime() *int64
	SetTotalTime(v int64) *SubmitTraceAbJobShrinkRequest
	GetTotalTime() *int64
	SetUserData(v string) *SubmitTraceAbJobShrinkRequest
	GetUserData() *string
}

type SubmitTraceAbJobShrinkRequest struct {
	// The key that is encoded by using the Base64 algorithm.
	//
	// example:
	//
	// Qh6OdgIMcliQSI1fReOw****
	CipherBase64ed *string `json:"CipherBase64ed,omitempty" xml:"CipherBase64ed,omitempty"`
	// The source video file for A/B watermarking.
	//
	// > OSS object or media asset must reside in the same region as the IMS service region. This API supports only videos that last at least 3 minutes. If the video is too short, the call may fail, or no output may be returned.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The watermark level, which specifies the channel to embed watermarks. Valid values: 0 specifies the 0u channel, 1 specifies the 1uv channel, and 2 specifies the 2yuv channel.
	//
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The output directory path.
	//
	// This parameter is required.
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The start point of watermark embedding. Unit: seconds.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The duration of the watermark embedding. Unit: seconds.
	//
	// example:
	//
	// 360
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The custom data, which can be up to 1,024 bytes in size.
	//
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTraceAbJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceAbJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbJobShrinkRequest) GetCipherBase64ed() *string {
	return s.CipherBase64ed
}

func (s *SubmitTraceAbJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitTraceAbJobShrinkRequest) GetLevel() *int64 {
	return s.Level
}

func (s *SubmitTraceAbJobShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *SubmitTraceAbJobShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SubmitTraceAbJobShrinkRequest) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *SubmitTraceAbJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTraceAbJobShrinkRequest) SetCipherBase64ed(v string) *SubmitTraceAbJobShrinkRequest {
	s.CipherBase64ed = &v
	return s
}

func (s *SubmitTraceAbJobShrinkRequest) SetInputShrink(v string) *SubmitTraceAbJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitTraceAbJobShrinkRequest) SetLevel(v int64) *SubmitTraceAbJobShrinkRequest {
	s.Level = &v
	return s
}

func (s *SubmitTraceAbJobShrinkRequest) SetOutputShrink(v string) *SubmitTraceAbJobShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *SubmitTraceAbJobShrinkRequest) SetStartTime(v int64) *SubmitTraceAbJobShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitTraceAbJobShrinkRequest) SetTotalTime(v int64) *SubmitTraceAbJobShrinkRequest {
	s.TotalTime = &v
	return s
}

func (s *SubmitTraceAbJobShrinkRequest) SetUserData(v string) *SubmitTraceAbJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTraceAbJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

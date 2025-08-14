// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveTranscodeJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SubmitLiveTranscodeJobShrinkRequest
	GetName() *string
	SetStartMode(v int32) *SubmitLiveTranscodeJobShrinkRequest
	GetStartMode() *int32
	SetStreamInputShrink(v string) *SubmitLiveTranscodeJobShrinkRequest
	GetStreamInputShrink() *string
	SetTemplateId(v string) *SubmitLiveTranscodeJobShrinkRequest
	GetTemplateId() *string
	SetTimedConfigShrink(v string) *SubmitLiveTranscodeJobShrinkRequest
	GetTimedConfigShrink() *string
	SetTranscodeOutputShrink(v string) *SubmitLiveTranscodeJobShrinkRequest
	GetTranscodeOutputShrink() *string
}

type SubmitLiveTranscodeJobShrinkRequest struct {
	// The name of the transcoding job.
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The start mode of the transcoding job.
	//
	// 	- 0: The transcoding job immediately starts.
	//
	// 	- 1: The transcoding job starts at the scheduled time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	StartMode *int32 `json:"StartMode,omitempty" xml:"StartMode,omitempty"`
	// The information about the input stream.
	//
	// This parameter is required.
	StreamInputShrink *string `json:"StreamInput,omitempty" xml:"StreamInput,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The configuration of a timed transcoding job. This parameter is required if you set StartMode to 1.
	TimedConfigShrink *string `json:"TimedConfig,omitempty" xml:"TimedConfig,omitempty"`
	// The information about the transcoding output.
	//
	// This parameter is required.
	TranscodeOutputShrink *string `json:"TranscodeOutput,omitempty" xml:"TranscodeOutput,omitempty"`
}

func (s SubmitLiveTranscodeJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveTranscodeJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveTranscodeJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitLiveTranscodeJobShrinkRequest) GetStartMode() *int32 {
	return s.StartMode
}

func (s *SubmitLiveTranscodeJobShrinkRequest) GetStreamInputShrink() *string {
	return s.StreamInputShrink
}

func (s *SubmitLiveTranscodeJobShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitLiveTranscodeJobShrinkRequest) GetTimedConfigShrink() *string {
	return s.TimedConfigShrink
}

func (s *SubmitLiveTranscodeJobShrinkRequest) GetTranscodeOutputShrink() *string {
	return s.TranscodeOutputShrink
}

func (s *SubmitLiveTranscodeJobShrinkRequest) SetName(v string) *SubmitLiveTranscodeJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitLiveTranscodeJobShrinkRequest) SetStartMode(v int32) *SubmitLiveTranscodeJobShrinkRequest {
	s.StartMode = &v
	return s
}

func (s *SubmitLiveTranscodeJobShrinkRequest) SetStreamInputShrink(v string) *SubmitLiveTranscodeJobShrinkRequest {
	s.StreamInputShrink = &v
	return s
}

func (s *SubmitLiveTranscodeJobShrinkRequest) SetTemplateId(v string) *SubmitLiveTranscodeJobShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitLiveTranscodeJobShrinkRequest) SetTimedConfigShrink(v string) *SubmitLiveTranscodeJobShrinkRequest {
	s.TimedConfigShrink = &v
	return s
}

func (s *SubmitLiveTranscodeJobShrinkRequest) SetTranscodeOutputShrink(v string) *SubmitLiveTranscodeJobShrinkRequest {
	s.TranscodeOutputShrink = &v
	return s
}

func (s *SubmitLiveTranscodeJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

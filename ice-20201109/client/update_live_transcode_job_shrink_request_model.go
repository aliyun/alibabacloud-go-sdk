// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveTranscodeJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateLiveTranscodeJobShrinkRequest
	GetJobId() *string
	SetName(v string) *UpdateLiveTranscodeJobShrinkRequest
	GetName() *string
	SetStreamInputShrink(v string) *UpdateLiveTranscodeJobShrinkRequest
	GetStreamInputShrink() *string
	SetTimedConfigShrink(v string) *UpdateLiveTranscodeJobShrinkRequest
	GetTimedConfigShrink() *string
	SetTranscodeOutputShrink(v string) *UpdateLiveTranscodeJobShrinkRequest
	GetTranscodeOutputShrink() *string
}

type UpdateLiveTranscodeJobShrinkRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// mytest3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the input stream.
	StreamInputShrink *string `json:"StreamInput,omitempty" xml:"StreamInput,omitempty"`
	// The configuration of a timed transcoding job.
	TimedConfigShrink *string `json:"TimedConfig,omitempty" xml:"TimedConfig,omitempty"`
	// The information about the transcoding output.
	TranscodeOutputShrink *string `json:"TranscodeOutput,omitempty" xml:"TranscodeOutput,omitempty"`
}

func (s UpdateLiveTranscodeJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeJobShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateLiveTranscodeJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLiveTranscodeJobShrinkRequest) GetStreamInputShrink() *string {
	return s.StreamInputShrink
}

func (s *UpdateLiveTranscodeJobShrinkRequest) GetTimedConfigShrink() *string {
	return s.TimedConfigShrink
}

func (s *UpdateLiveTranscodeJobShrinkRequest) GetTranscodeOutputShrink() *string {
	return s.TranscodeOutputShrink
}

func (s *UpdateLiveTranscodeJobShrinkRequest) SetJobId(v string) *UpdateLiveTranscodeJobShrinkRequest {
	s.JobId = &v
	return s
}

func (s *UpdateLiveTranscodeJobShrinkRequest) SetName(v string) *UpdateLiveTranscodeJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateLiveTranscodeJobShrinkRequest) SetStreamInputShrink(v string) *UpdateLiveTranscodeJobShrinkRequest {
	s.StreamInputShrink = &v
	return s
}

func (s *UpdateLiveTranscodeJobShrinkRequest) SetTimedConfigShrink(v string) *UpdateLiveTranscodeJobShrinkRequest {
	s.TimedConfigShrink = &v
	return s
}

func (s *UpdateLiveTranscodeJobShrinkRequest) SetTranscodeOutputShrink(v string) *UpdateLiveTranscodeJobShrinkRequest {
	s.TranscodeOutputShrink = &v
	return s
}

func (s *UpdateLiveTranscodeJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

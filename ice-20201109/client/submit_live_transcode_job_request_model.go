// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SubmitLiveTranscodeJobRequest
	GetName() *string
	SetStartMode(v int32) *SubmitLiveTranscodeJobRequest
	GetStartMode() *int32
	SetStreamInput(v *SubmitLiveTranscodeJobRequestStreamInput) *SubmitLiveTranscodeJobRequest
	GetStreamInput() *SubmitLiveTranscodeJobRequestStreamInput
	SetTemplateId(v string) *SubmitLiveTranscodeJobRequest
	GetTemplateId() *string
	SetTimedConfig(v *SubmitLiveTranscodeJobRequestTimedConfig) *SubmitLiveTranscodeJobRequest
	GetTimedConfig() *SubmitLiveTranscodeJobRequestTimedConfig
	SetTranscodeOutput(v *SubmitLiveTranscodeJobRequestTranscodeOutput) *SubmitLiveTranscodeJobRequest
	GetTranscodeOutput() *SubmitLiveTranscodeJobRequestTranscodeOutput
}

type SubmitLiveTranscodeJobRequest struct {
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
	StreamInput *SubmitLiveTranscodeJobRequestStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The configuration of a timed transcoding job. This parameter is required if you set StartMode to 1.
	TimedConfig *SubmitLiveTranscodeJobRequestTimedConfig `json:"TimedConfig,omitempty" xml:"TimedConfig,omitempty" type:"Struct"`
	// The information about the transcoding output.
	//
	// This parameter is required.
	TranscodeOutput *SubmitLiveTranscodeJobRequestTranscodeOutput `json:"TranscodeOutput,omitempty" xml:"TranscodeOutput,omitempty" type:"Struct"`
}

func (s SubmitLiveTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveTranscodeJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitLiveTranscodeJobRequest) GetStartMode() *int32 {
	return s.StartMode
}

func (s *SubmitLiveTranscodeJobRequest) GetStreamInput() *SubmitLiveTranscodeJobRequestStreamInput {
	return s.StreamInput
}

func (s *SubmitLiveTranscodeJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitLiveTranscodeJobRequest) GetTimedConfig() *SubmitLiveTranscodeJobRequestTimedConfig {
	return s.TimedConfig
}

func (s *SubmitLiveTranscodeJobRequest) GetTranscodeOutput() *SubmitLiveTranscodeJobRequestTranscodeOutput {
	return s.TranscodeOutput
}

func (s *SubmitLiveTranscodeJobRequest) SetName(v string) *SubmitLiveTranscodeJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequest) SetStartMode(v int32) *SubmitLiveTranscodeJobRequest {
	s.StartMode = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequest) SetStreamInput(v *SubmitLiveTranscodeJobRequestStreamInput) *SubmitLiveTranscodeJobRequest {
	s.StreamInput = v
	return s
}

func (s *SubmitLiveTranscodeJobRequest) SetTemplateId(v string) *SubmitLiveTranscodeJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequest) SetTimedConfig(v *SubmitLiveTranscodeJobRequestTimedConfig) *SubmitLiveTranscodeJobRequest {
	s.TimedConfig = v
	return s
}

func (s *SubmitLiveTranscodeJobRequest) SetTranscodeOutput(v *SubmitLiveTranscodeJobRequestTranscodeOutput) *SubmitLiveTranscodeJobRequest {
	s.TranscodeOutput = v
	return s
}

func (s *SubmitLiveTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitLiveTranscodeJobRequestStreamInput struct {
	// The URL of the input stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://mydomain/app/stream1
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// The type of the input stream. The value can only be rtmp.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitLiveTranscodeJobRequestStreamInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveTranscodeJobRequestStreamInput) GoString() string {
	return s.String()
}

func (s *SubmitLiveTranscodeJobRequestStreamInput) GetInputUrl() *string {
	return s.InputUrl
}

func (s *SubmitLiveTranscodeJobRequestStreamInput) GetType() *string {
	return s.Type
}

func (s *SubmitLiveTranscodeJobRequestStreamInput) SetInputUrl(v string) *SubmitLiveTranscodeJobRequestStreamInput {
	s.InputUrl = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequestStreamInput) SetType(v string) *SubmitLiveTranscodeJobRequestStreamInput {
	s.Type = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequestStreamInput) Validate() error {
	return dara.Validate(s)
}

type SubmitLiveTranscodeJobRequestTimedConfig struct {
	// The stop time of the transcoding job. Note: The time span between the stop time and the current time cannot exceed seven days.
	//
	// example:
	//
	// 2022-07-20T08:20:32Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the transcoding job. Note: The time span between the start time and the current time cannot exceed seven days.
	//
	// example:
	//
	// 2022-02-21T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SubmitLiveTranscodeJobRequestTimedConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveTranscodeJobRequestTimedConfig) GoString() string {
	return s.String()
}

func (s *SubmitLiveTranscodeJobRequestTimedConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *SubmitLiveTranscodeJobRequestTimedConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitLiveTranscodeJobRequestTimedConfig) SetEndTime(v string) *SubmitLiveTranscodeJobRequestTimedConfig {
	s.EndTime = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequestTimedConfig) SetStartTime(v string) *SubmitLiveTranscodeJobRequestTimedConfig {
	s.StartTime = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequestTimedConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitLiveTranscodeJobRequestTranscodeOutput struct {
	// The streaming domain name of ApsaraVideo Live.
	//
	// example:
	//
	// mydomain
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the output stream. A value of LiveCenter indicates that the URL of the output stream is generated based on the domain name of ApsaraVideo Live. The value can only be LiveCenter.
	//
	// This parameter is required.
	//
	// example:
	//
	// LiveCenter
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitLiveTranscodeJobRequestTranscodeOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveTranscodeJobRequestTranscodeOutput) GoString() string {
	return s.String()
}

func (s *SubmitLiveTranscodeJobRequestTranscodeOutput) GetDomainName() *string {
	return s.DomainName
}

func (s *SubmitLiveTranscodeJobRequestTranscodeOutput) GetType() *string {
	return s.Type
}

func (s *SubmitLiveTranscodeJobRequestTranscodeOutput) SetDomainName(v string) *SubmitLiveTranscodeJobRequestTranscodeOutput {
	s.DomainName = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequestTranscodeOutput) SetType(v string) *SubmitLiveTranscodeJobRequestTranscodeOutput {
	s.Type = &v
	return s
}

func (s *SubmitLiveTranscodeJobRequestTranscodeOutput) Validate() error {
	return dara.Validate(s)
}

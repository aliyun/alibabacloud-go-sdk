// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateLiveTranscodeJobRequest
	GetJobId() *string
	SetName(v string) *UpdateLiveTranscodeJobRequest
	GetName() *string
	SetStreamInput(v *UpdateLiveTranscodeJobRequestStreamInput) *UpdateLiveTranscodeJobRequest
	GetStreamInput() *UpdateLiveTranscodeJobRequestStreamInput
	SetTimedConfig(v *UpdateLiveTranscodeJobRequestTimedConfig) *UpdateLiveTranscodeJobRequest
	GetTimedConfig() *UpdateLiveTranscodeJobRequestTimedConfig
	SetTranscodeOutput(v *UpdateLiveTranscodeJobRequestTranscodeOutput) *UpdateLiveTranscodeJobRequest
	GetTranscodeOutput() *UpdateLiveTranscodeJobRequestTranscodeOutput
}

type UpdateLiveTranscodeJobRequest struct {
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
	StreamInput *UpdateLiveTranscodeJobRequestStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The configuration of a timed transcoding job.
	TimedConfig *UpdateLiveTranscodeJobRequestTimedConfig `json:"TimedConfig,omitempty" xml:"TimedConfig,omitempty" type:"Struct"`
	// The information about the transcoding output.
	TranscodeOutput *UpdateLiveTranscodeJobRequestTranscodeOutput `json:"TranscodeOutput,omitempty" xml:"TranscodeOutput,omitempty" type:"Struct"`
}

func (s UpdateLiveTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateLiveTranscodeJobRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLiveTranscodeJobRequest) GetStreamInput() *UpdateLiveTranscodeJobRequestStreamInput {
	return s.StreamInput
}

func (s *UpdateLiveTranscodeJobRequest) GetTimedConfig() *UpdateLiveTranscodeJobRequestTimedConfig {
	return s.TimedConfig
}

func (s *UpdateLiveTranscodeJobRequest) GetTranscodeOutput() *UpdateLiveTranscodeJobRequestTranscodeOutput {
	return s.TranscodeOutput
}

func (s *UpdateLiveTranscodeJobRequest) SetJobId(v string) *UpdateLiveTranscodeJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateLiveTranscodeJobRequest) SetName(v string) *UpdateLiveTranscodeJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateLiveTranscodeJobRequest) SetStreamInput(v *UpdateLiveTranscodeJobRequestStreamInput) *UpdateLiveTranscodeJobRequest {
	s.StreamInput = v
	return s
}

func (s *UpdateLiveTranscodeJobRequest) SetTimedConfig(v *UpdateLiveTranscodeJobRequestTimedConfig) *UpdateLiveTranscodeJobRequest {
	s.TimedConfig = v
	return s
}

func (s *UpdateLiveTranscodeJobRequest) SetTranscodeOutput(v *UpdateLiveTranscodeJobRequestTranscodeOutput) *UpdateLiveTranscodeJobRequest {
	s.TranscodeOutput = v
	return s
}

func (s *UpdateLiveTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveTranscodeJobRequestStreamInput struct {
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

func (s UpdateLiveTranscodeJobRequestStreamInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeJobRequestStreamInput) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeJobRequestStreamInput) GetInputUrl() *string {
	return s.InputUrl
}

func (s *UpdateLiveTranscodeJobRequestStreamInput) GetType() *string {
	return s.Type
}

func (s *UpdateLiveTranscodeJobRequestStreamInput) SetInputUrl(v string) *UpdateLiveTranscodeJobRequestStreamInput {
	s.InputUrl = &v
	return s
}

func (s *UpdateLiveTranscodeJobRequestStreamInput) SetType(v string) *UpdateLiveTranscodeJobRequestStreamInput {
	s.Type = &v
	return s
}

func (s *UpdateLiveTranscodeJobRequestStreamInput) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveTranscodeJobRequestTimedConfig struct {
	// The stop time of the transcoding job. Note: The time span between the stop time and the current time cannot exceed seven days.
	//
	// example:
	//
	// 2022-08-05T06:08:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the transcoding job. Note: The time span between the start time and the current time cannot exceed seven days.
	//
	// example:
	//
	// 2022-06-19T02:16:41Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateLiveTranscodeJobRequestTimedConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeJobRequestTimedConfig) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeJobRequestTimedConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateLiveTranscodeJobRequestTimedConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateLiveTranscodeJobRequestTimedConfig) SetEndTime(v string) *UpdateLiveTranscodeJobRequestTimedConfig {
	s.EndTime = &v
	return s
}

func (s *UpdateLiveTranscodeJobRequestTimedConfig) SetStartTime(v string) *UpdateLiveTranscodeJobRequestTimedConfig {
	s.StartTime = &v
	return s
}

func (s *UpdateLiveTranscodeJobRequestTimedConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveTranscodeJobRequestTranscodeOutput struct {
	// The streaming domain name of ApsaraVideo Live.
	//
	// This parameter is required.
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

func (s UpdateLiveTranscodeJobRequestTranscodeOutput) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeJobRequestTranscodeOutput) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeJobRequestTranscodeOutput) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveTranscodeJobRequestTranscodeOutput) GetType() *string {
	return s.Type
}

func (s *UpdateLiveTranscodeJobRequestTranscodeOutput) SetDomainName(v string) *UpdateLiveTranscodeJobRequestTranscodeOutput {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveTranscodeJobRequestTranscodeOutput) SetType(v string) *UpdateLiveTranscodeJobRequestTranscodeOutput {
	s.Type = &v
	return s
}

func (s *UpdateLiveTranscodeJobRequestTranscodeOutput) Validate() error {
	return dara.Validate(s)
}

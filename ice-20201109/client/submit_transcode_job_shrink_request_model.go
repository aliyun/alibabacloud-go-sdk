// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranscodeJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitTranscodeJobShrinkRequest
	GetClientToken() *string
	SetInputGroupShrink(v string) *SubmitTranscodeJobShrinkRequest
	GetInputGroupShrink() *string
	SetName(v string) *SubmitTranscodeJobShrinkRequest
	GetName() *string
	SetOutputGroupShrink(v string) *SubmitTranscodeJobShrinkRequest
	GetOutputGroupShrink() *string
	SetScheduleConfigShrink(v string) *SubmitTranscodeJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetUserData(v string) *SubmitTranscodeJobShrinkRequest
	GetUserData() *string
}

type SubmitTranscodeJobShrinkRequest struct {
	// The idempotence key. Ensures request idempotence.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The input group for the job. A single input creates a transcoding job. Multiple inputs create a media merging job.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-name
	InputGroupShrink *string `json:"InputGroup,omitempty" xml:"InputGroup,omitempty"`
	// The job name.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output group for the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// user-data
	OutputGroupShrink *string `json:"OutputGroup,omitempty" xml:"OutputGroup,omitempty"`
	// The job scheduling information.
	//
	// example:
	//
	// job-name
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// Custom settings in JSON format. The length is limited to 512 bytes. Supports [custom webhook address configuration](https://help.aliyun.com/document_detail/451631.html).
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTranscodeJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitTranscodeJobShrinkRequest) GetInputGroupShrink() *string {
	return s.InputGroupShrink
}

func (s *SubmitTranscodeJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitTranscodeJobShrinkRequest) GetOutputGroupShrink() *string {
	return s.OutputGroupShrink
}

func (s *SubmitTranscodeJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitTranscodeJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTranscodeJobShrinkRequest) SetClientToken(v string) *SubmitTranscodeJobShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitTranscodeJobShrinkRequest) SetInputGroupShrink(v string) *SubmitTranscodeJobShrinkRequest {
	s.InputGroupShrink = &v
	return s
}

func (s *SubmitTranscodeJobShrinkRequest) SetName(v string) *SubmitTranscodeJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitTranscodeJobShrinkRequest) SetOutputGroupShrink(v string) *SubmitTranscodeJobShrinkRequest {
	s.OutputGroupShrink = &v
	return s
}

func (s *SubmitTranscodeJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitTranscodeJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitTranscodeJobShrinkRequest) SetUserData(v string) *SubmitTranscodeJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTranscodeJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

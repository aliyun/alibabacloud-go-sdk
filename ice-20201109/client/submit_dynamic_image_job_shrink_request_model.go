// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicImageJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *SubmitDynamicImageJobShrinkRequest
	GetInputShrink() *string
	SetName(v string) *SubmitDynamicImageJobShrinkRequest
	GetName() *string
	SetOutputShrink(v string) *SubmitDynamicImageJobShrinkRequest
	GetOutputShrink() *string
	SetScheduleConfigShrink(v string) *SubmitDynamicImageJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetTemplateConfigShrink(v string) *SubmitDynamicImageJobShrinkRequest
	GetTemplateConfigShrink() *string
	SetUserData(v string) *SubmitDynamicImageJobShrinkRequest
	GetUserData() *string
}

type SubmitDynamicImageJobShrinkRequest struct {
	// The input of the job.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// SampleJob
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the job.
	//
	// This parameter is required.
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The scheduling settings.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The snapshot template configuration.
	//
	// This parameter is required.
	TemplateConfigShrink *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// {"SampleKey": "SampleValue"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitDynamicImageJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitDynamicImageJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitDynamicImageJobShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *SubmitDynamicImageJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitDynamicImageJobShrinkRequest) GetTemplateConfigShrink() *string {
	return s.TemplateConfigShrink
}

func (s *SubmitDynamicImageJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitDynamicImageJobShrinkRequest) SetInputShrink(v string) *SubmitDynamicImageJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitDynamicImageJobShrinkRequest) SetName(v string) *SubmitDynamicImageJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitDynamicImageJobShrinkRequest) SetOutputShrink(v string) *SubmitDynamicImageJobShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *SubmitDynamicImageJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitDynamicImageJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitDynamicImageJobShrinkRequest) SetTemplateConfigShrink(v string) *SubmitDynamicImageJobShrinkRequest {
	s.TemplateConfigShrink = &v
	return s
}

func (s *SubmitDynamicImageJobShrinkRequest) SetUserData(v string) *SubmitDynamicImageJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitDynamicImageJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *SubmitSnapshotJobShrinkRequest
	GetInputShrink() *string
	SetName(v string) *SubmitSnapshotJobShrinkRequest
	GetName() *string
	SetOutputShrink(v string) *SubmitSnapshotJobShrinkRequest
	GetOutputShrink() *string
	SetScheduleConfigShrink(v string) *SubmitSnapshotJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetTemplateConfigShrink(v string) *SubmitSnapshotJobShrinkRequest
	GetTemplateConfigShrink() *string
	SetUserData(v string) *SubmitSnapshotJobShrinkRequest
	GetUserData() *string
}

type SubmitSnapshotJobShrinkRequest struct {
	// The snapshot input.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// SampleJob
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The snapshot output.
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
	// {"test parameter": "test value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSnapshotJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitSnapshotJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitSnapshotJobShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *SubmitSnapshotJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitSnapshotJobShrinkRequest) GetTemplateConfigShrink() *string {
	return s.TemplateConfigShrink
}

func (s *SubmitSnapshotJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSnapshotJobShrinkRequest) SetInputShrink(v string) *SubmitSnapshotJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetName(v string) *SubmitSnapshotJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetOutputShrink(v string) *SubmitSnapshotJobShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitSnapshotJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetTemplateConfigShrink(v string) *SubmitSnapshotJobShrinkRequest {
	s.TemplateConfigShrink = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetUserData(v string) *SubmitSnapshotJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

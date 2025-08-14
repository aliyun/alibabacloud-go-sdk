// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPackageJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputsShrink(v string) *SubmitPackageJobShrinkRequest
	GetInputsShrink() *string
	SetName(v string) *SubmitPackageJobShrinkRequest
	GetName() *string
	SetOutputShrink(v string) *SubmitPackageJobShrinkRequest
	GetOutputShrink() *string
	SetScheduleConfigShrink(v string) *SubmitPackageJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetUserData(v string) *SubmitPackageJobShrinkRequest
	GetUserData() *string
}

type SubmitPackageJobShrinkRequest struct {
	// The input of the job.
	//
	// This parameter is required.
	InputsShrink *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the job.
	//
	// This parameter is required.
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The scheduling settings.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// {"param": "value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitPackageJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPackageJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitPackageJobShrinkRequest) GetInputsShrink() *string {
	return s.InputsShrink
}

func (s *SubmitPackageJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitPackageJobShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *SubmitPackageJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitPackageJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitPackageJobShrinkRequest) SetInputsShrink(v string) *SubmitPackageJobShrinkRequest {
	s.InputsShrink = &v
	return s
}

func (s *SubmitPackageJobShrinkRequest) SetName(v string) *SubmitPackageJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitPackageJobShrinkRequest) SetOutputShrink(v string) *SubmitPackageJobShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *SubmitPackageJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitPackageJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitPackageJobShrinkRequest) SetUserData(v string) *SubmitPackageJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitPackageJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

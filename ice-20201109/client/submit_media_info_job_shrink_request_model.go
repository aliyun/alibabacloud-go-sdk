// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaInfoJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *SubmitMediaInfoJobShrinkRequest
	GetInputShrink() *string
	SetName(v string) *SubmitMediaInfoJobShrinkRequest
	GetName() *string
	SetScheduleConfigShrink(v string) *SubmitMediaInfoJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetUserData(v string) *SubmitMediaInfoJobShrinkRequest
	GetUserData() *string
}

type SubmitMediaInfoJobShrinkRequest struct {
	// The input of the job.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The job name.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The scheduling parameters.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The user data.
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaInfoJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitMediaInfoJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitMediaInfoJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitMediaInfoJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaInfoJobShrinkRequest) SetInputShrink(v string) *SubmitMediaInfoJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitMediaInfoJobShrinkRequest) SetName(v string) *SubmitMediaInfoJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitMediaInfoJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitMediaInfoJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitMediaInfoJobShrinkRequest) SetUserData(v string) *SubmitMediaInfoJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaInfoJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSyncMediaInfoJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *SubmitSyncMediaInfoJobShrinkRequest
	GetInputShrink() *string
	SetName(v string) *SubmitSyncMediaInfoJobShrinkRequest
	GetName() *string
	SetScheduleConfigShrink(v string) *SubmitSyncMediaInfoJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetUserData(v string) *SubmitSyncMediaInfoJobShrinkRequest
	GetUserData() *string
}

type SubmitSyncMediaInfoJobShrinkRequest struct {
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
	// The scheduling parameters. This parameter is optional.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The user data.
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSyncMediaInfoJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) SetInputShrink(v string) *SubmitSyncMediaInfoJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) SetName(v string) *SubmitSyncMediaInfoJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitSyncMediaInfoJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) SetUserData(v string) *SubmitSyncMediaInfoJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSyncMediaInfoJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

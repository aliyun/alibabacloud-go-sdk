// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDIJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *StartDIJobShrinkRequest
	GetDIJobId() *int64
	SetForceToRerun(v bool) *StartDIJobShrinkRequest
	GetForceToRerun() *bool
	SetRealtimeStartSettingsShrink(v string) *StartDIJobShrinkRequest
	GetRealtimeStartSettingsShrink() *string
}

type StartDIJobShrinkRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11743
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// Specifies whether to forcefully rerun all synchronization steps. If you do not configure this parameter, the system does not forcefully rerun the task.
	//
	// example:
	//
	// true
	ForceToRerun *bool `json:"ForceToRerun,omitempty" xml:"ForceToRerun,omitempty"`
	// The settings for the start.
	RealtimeStartSettingsShrink *string `json:"RealtimeStartSettings,omitempty" xml:"RealtimeStartSettings,omitempty"`
}

func (s StartDIJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartDIJobShrinkRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *StartDIJobShrinkRequest) GetForceToRerun() *bool {
	return s.ForceToRerun
}

func (s *StartDIJobShrinkRequest) GetRealtimeStartSettingsShrink() *string {
	return s.RealtimeStartSettingsShrink
}

func (s *StartDIJobShrinkRequest) SetDIJobId(v int64) *StartDIJobShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *StartDIJobShrinkRequest) SetForceToRerun(v bool) *StartDIJobShrinkRequest {
	s.ForceToRerun = &v
	return s
}

func (s *StartDIJobShrinkRequest) SetRealtimeStartSettingsShrink(v string) *StartDIJobShrinkRequest {
	s.RealtimeStartSettingsShrink = &v
	return s
}

func (s *StartDIJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

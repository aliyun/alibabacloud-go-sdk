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
	SetId(v int64) *StartDIJobShrinkRequest
	GetId() *int64
	SetRealtimeStartSettingsShrink(v string) *StartDIJobShrinkRequest
	GetRealtimeStartSettingsShrink() *string
}

type StartDIJobShrinkRequest struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 10000
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// Specifies whether to forcefully rerun all synchronization steps. If you do not configure this parameter, the system does not perform the forcible rerun operation.
	//
	// 	- If the system does not perform the forcible rerun operation, only the steps that are not run start to run.
	//
	// 	- If the system performs the forcible rerun operation, all steps start to rerun.
	//
	// example:
	//
	// false
	ForceToRerun *bool `json:"ForceToRerun,omitempty" xml:"ForceToRerun,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 10000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The settings for starting real-time synchronization.
	//
	//     {
	//
	//       "StartTime":1663765058
	//
	//     }
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

func (s *StartDIJobShrinkRequest) GetId() *int64 {
	return s.Id
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

func (s *StartDIJobShrinkRequest) SetId(v int64) *StartDIJobShrinkRequest {
	s.Id = &v
	return s
}

func (s *StartDIJobShrinkRequest) SetRealtimeStartSettingsShrink(v string) *StartDIJobShrinkRequest {
	s.RealtimeStartSettingsShrink = &v
	return s
}

func (s *StartDIJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

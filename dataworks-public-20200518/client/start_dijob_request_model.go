// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *StartDIJobRequest
	GetDIJobId() *int64
	SetForceToRerun(v bool) *StartDIJobRequest
	GetForceToRerun() *bool
	SetRealtimeStartSettings(v *StartDIJobRequestRealtimeStartSettings) *StartDIJobRequest
	GetRealtimeStartSettings() *StartDIJobRequestRealtimeStartSettings
}

type StartDIJobRequest struct {
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
	RealtimeStartSettings *StartDIJobRequestRealtimeStartSettings `json:"RealtimeStartSettings,omitempty" xml:"RealtimeStartSettings,omitempty" type:"Struct"`
}

func (s StartDIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDIJobRequest) GoString() string {
	return s.String()
}

func (s *StartDIJobRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *StartDIJobRequest) GetForceToRerun() *bool {
	return s.ForceToRerun
}

func (s *StartDIJobRequest) GetRealtimeStartSettings() *StartDIJobRequestRealtimeStartSettings {
	return s.RealtimeStartSettings
}

func (s *StartDIJobRequest) SetDIJobId(v int64) *StartDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *StartDIJobRequest) SetForceToRerun(v bool) *StartDIJobRequest {
	s.ForceToRerun = &v
	return s
}

func (s *StartDIJobRequest) SetRealtimeStartSettings(v *StartDIJobRequestRealtimeStartSettings) *StartDIJobRequest {
	s.RealtimeStartSettings = v
	return s
}

func (s *StartDIJobRequest) Validate() error {
	return dara.Validate(s)
}

type StartDIJobRequestRealtimeStartSettings struct {
	// The failover settings.
	FailoverSettings *StartDIJobRequestRealtimeStartSettingsFailoverSettings `json:"FailoverSettings,omitempty" xml:"FailoverSettings,omitempty" type:"Struct"`
	// The timestamp of the start offset. Unit: seconds. If you do not configure this parameter, the offset is not reset by default.
	//
	// example:
	//
	// 1671516776
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StartDIJobRequestRealtimeStartSettings) String() string {
	return dara.Prettify(s)
}

func (s StartDIJobRequestRealtimeStartSettings) GoString() string {
	return s.String()
}

func (s *StartDIJobRequestRealtimeStartSettings) GetFailoverSettings() *StartDIJobRequestRealtimeStartSettingsFailoverSettings {
	return s.FailoverSettings
}

func (s *StartDIJobRequestRealtimeStartSettings) GetStartTime() *int64 {
	return s.StartTime
}

func (s *StartDIJobRequestRealtimeStartSettings) SetFailoverSettings(v *StartDIJobRequestRealtimeStartSettingsFailoverSettings) *StartDIJobRequestRealtimeStartSettings {
	s.FailoverSettings = v
	return s
}

func (s *StartDIJobRequestRealtimeStartSettings) SetStartTime(v int64) *StartDIJobRequestRealtimeStartSettings {
	s.StartTime = &v
	return s
}

func (s *StartDIJobRequestRealtimeStartSettings) Validate() error {
	return dara.Validate(s)
}

type StartDIJobRequestRealtimeStartSettingsFailoverSettings struct {
	// The failover interval. Unit: minutes.
	//
	// example:
	//
	// 10
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The maximum number of failovers.
	//
	// example:
	//
	// 30
	UpperLimit *int64 `json:"UpperLimit,omitempty" xml:"UpperLimit,omitempty"`
}

func (s StartDIJobRequestRealtimeStartSettingsFailoverSettings) String() string {
	return dara.Prettify(s)
}

func (s StartDIJobRequestRealtimeStartSettingsFailoverSettings) GoString() string {
	return s.String()
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) GetInterval() *int64 {
	return s.Interval
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) GetUpperLimit() *int64 {
	return s.UpperLimit
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) SetInterval(v int64) *StartDIJobRequestRealtimeStartSettingsFailoverSettings {
	s.Interval = &v
	return s
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) SetUpperLimit(v int64) *StartDIJobRequestRealtimeStartSettingsFailoverSettings {
	s.UpperLimit = &v
	return s
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartupProbe interface {
	dara.Model
	String() string
	GoString() string
	SetFailureThreshold(v int32) *StartupProbe
	GetFailureThreshold() *int32
	SetInitialDelaySeconds(v int32) *StartupProbe
	GetInitialDelaySeconds() *int32
	SetPeriodSeconds(v int32) *StartupProbe
	GetPeriodSeconds() *int32
	SetProbeHandler(v *ProbeHandler) *StartupProbe
	GetProbeHandler() *ProbeHandler
	SetTimeoutSeconds(v int32) *StartupProbe
	GetTimeoutSeconds() *int32
}

type StartupProbe struct {
	FailureThreshold    *int32        `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	InitialDelaySeconds *int32        `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32        `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	ProbeHandler        *ProbeHandler `json:"ProbeHandler,omitempty" xml:"ProbeHandler,omitempty"`
	TimeoutSeconds      *int32        `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s StartupProbe) String() string {
	return dara.Prettify(s)
}

func (s StartupProbe) GoString() string {
	return s.String()
}

func (s *StartupProbe) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *StartupProbe) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *StartupProbe) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *StartupProbe) GetProbeHandler() *ProbeHandler {
	return s.ProbeHandler
}

func (s *StartupProbe) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *StartupProbe) SetFailureThreshold(v int32) *StartupProbe {
	s.FailureThreshold = &v
	return s
}

func (s *StartupProbe) SetInitialDelaySeconds(v int32) *StartupProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *StartupProbe) SetPeriodSeconds(v int32) *StartupProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *StartupProbe) SetProbeHandler(v *ProbeHandler) *StartupProbe {
	s.ProbeHandler = v
	return s
}

func (s *StartupProbe) SetTimeoutSeconds(v int32) *StartupProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *StartupProbe) Validate() error {
	if s.ProbeHandler != nil {
		if err := s.ProbeHandler.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobDiagnosisSymptoms interface {
	dara.Model
	String() string
	GoString() string
	SetAutopilot(v *JobDiagnosisSymptom) *JobDiagnosisSymptoms
	GetAutopilot() *JobDiagnosisSymptom
	SetOthers(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms
	GetOthers() []*JobDiagnosisSymptom
	SetRuntime(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms
	GetRuntime() []*JobDiagnosisSymptom
	SetStartup(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms
	GetStartup() []*JobDiagnosisSymptom
	SetState(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms
	GetState() []*JobDiagnosisSymptom
	SetTroubleshooting(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms
	GetTroubleshooting() []*JobDiagnosisSymptom
}

type JobDiagnosisSymptoms struct {
	Autopilot       *JobDiagnosisSymptom   `json:"autopilot,omitempty" xml:"autopilot,omitempty"`
	Others          []*JobDiagnosisSymptom `json:"others,omitempty" xml:"others,omitempty" type:"Repeated"`
	Runtime         []*JobDiagnosisSymptom `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Repeated"`
	Startup         []*JobDiagnosisSymptom `json:"startup,omitempty" xml:"startup,omitempty" type:"Repeated"`
	State           []*JobDiagnosisSymptom `json:"state,omitempty" xml:"state,omitempty" type:"Repeated"`
	Troubleshooting []*JobDiagnosisSymptom `json:"troubleshooting,omitempty" xml:"troubleshooting,omitempty" type:"Repeated"`
}

func (s JobDiagnosisSymptoms) String() string {
	return dara.Prettify(s)
}

func (s JobDiagnosisSymptoms) GoString() string {
	return s.String()
}

func (s *JobDiagnosisSymptoms) GetAutopilot() *JobDiagnosisSymptom {
	return s.Autopilot
}

func (s *JobDiagnosisSymptoms) GetOthers() []*JobDiagnosisSymptom {
	return s.Others
}

func (s *JobDiagnosisSymptoms) GetRuntime() []*JobDiagnosisSymptom {
	return s.Runtime
}

func (s *JobDiagnosisSymptoms) GetStartup() []*JobDiagnosisSymptom {
	return s.Startup
}

func (s *JobDiagnosisSymptoms) GetState() []*JobDiagnosisSymptom {
	return s.State
}

func (s *JobDiagnosisSymptoms) GetTroubleshooting() []*JobDiagnosisSymptom {
	return s.Troubleshooting
}

func (s *JobDiagnosisSymptoms) SetAutopilot(v *JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Autopilot = v
	return s
}

func (s *JobDiagnosisSymptoms) SetOthers(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Others = v
	return s
}

func (s *JobDiagnosisSymptoms) SetRuntime(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Runtime = v
	return s
}

func (s *JobDiagnosisSymptoms) SetStartup(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Startup = v
	return s
}

func (s *JobDiagnosisSymptoms) SetState(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.State = v
	return s
}

func (s *JobDiagnosisSymptoms) SetTroubleshooting(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Troubleshooting = v
	return s
}

func (s *JobDiagnosisSymptoms) Validate() error {
	if s.Autopilot != nil {
		if err := s.Autopilot.Validate(); err != nil {
			return err
		}
	}
	if s.Others != nil {
		for _, item := range s.Others {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Runtime != nil {
		for _, item := range s.Runtime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Startup != nil {
		for _, item := range s.Startup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.State != nil {
		for _, item := range s.State {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Troubleshooting != nil {
		for _, item := range s.Troubleshooting {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSignalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSignal(v string) *CreateSignalRequest
	GetSignal() *string
	SetTarget(v *SignalTarget) *CreateSignalRequest
	GetTarget() *SignalTarget
}

type CreateSignalRequest struct {
	// The signal code.
	//
	// example:
	//
	// SIGUSR1
	Signal *string `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The signal delivery scope.
	Target *SignalTarget `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s CreateSignalRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSignalRequest) GoString() string {
	return s.String()
}

func (s *CreateSignalRequest) GetSignal() *string {
	return s.Signal
}

func (s *CreateSignalRequest) GetTarget() *SignalTarget {
	return s.Target
}

func (s *CreateSignalRequest) SetSignal(v string) *CreateSignalRequest {
	s.Signal = &v
	return s
}

func (s *CreateSignalRequest) SetTarget(v *SignalTarget) *CreateSignalRequest {
	s.Target = v
	return s
}

func (s *CreateSignalRequest) Validate() error {
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

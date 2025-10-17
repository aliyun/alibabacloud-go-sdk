// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLifecycle interface {
	dara.Model
	String() string
	GoString() string
	SetPostStart(v *LifecyclePostStart) *Lifecycle
	GetPostStart() *LifecyclePostStart
	SetPreStop(v *LifecyclePreStop) *Lifecycle
	GetPreStop() *LifecyclePreStop
}

type Lifecycle struct {
	PostStart *LifecyclePostStart `json:"PostStart,omitempty" xml:"PostStart,omitempty" type:"Struct"`
	PreStop   *LifecyclePreStop   `json:"PreStop,omitempty" xml:"PreStop,omitempty" type:"Struct"`
}

func (s Lifecycle) String() string {
	return dara.Prettify(s)
}

func (s Lifecycle) GoString() string {
	return s.String()
}

func (s *Lifecycle) GetPostStart() *LifecyclePostStart {
	return s.PostStart
}

func (s *Lifecycle) GetPreStop() *LifecyclePreStop {
	return s.PreStop
}

func (s *Lifecycle) SetPostStart(v *LifecyclePostStart) *Lifecycle {
	s.PostStart = v
	return s
}

func (s *Lifecycle) SetPreStop(v *LifecyclePreStop) *Lifecycle {
	s.PreStop = v
	return s
}

func (s *Lifecycle) Validate() error {
	if s.PostStart != nil {
		if err := s.PostStart.Validate(); err != nil {
			return err
		}
	}
	if s.PreStop != nil {
		if err := s.PreStop.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LifecyclePostStart struct {
	Exec *LifecyclePostStartExec `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
}

func (s LifecyclePostStart) String() string {
	return dara.Prettify(s)
}

func (s LifecyclePostStart) GoString() string {
	return s.String()
}

func (s *LifecyclePostStart) GetExec() *LifecyclePostStartExec {
	return s.Exec
}

func (s *LifecyclePostStart) SetExec(v *LifecyclePostStartExec) *LifecyclePostStart {
	s.Exec = v
	return s
}

func (s *LifecyclePostStart) Validate() error {
	if s.Exec != nil {
		if err := s.Exec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LifecyclePostStartExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s LifecyclePostStartExec) String() string {
	return dara.Prettify(s)
}

func (s LifecyclePostStartExec) GoString() string {
	return s.String()
}

func (s *LifecyclePostStartExec) GetCommand() []*string {
	return s.Command
}

func (s *LifecyclePostStartExec) SetCommand(v []*string) *LifecyclePostStartExec {
	s.Command = v
	return s
}

func (s *LifecyclePostStartExec) Validate() error {
	return dara.Validate(s)
}

type LifecyclePreStop struct {
	Exec *LifecyclePreStopExec `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
}

func (s LifecyclePreStop) String() string {
	return dara.Prettify(s)
}

func (s LifecyclePreStop) GoString() string {
	return s.String()
}

func (s *LifecyclePreStop) GetExec() *LifecyclePreStopExec {
	return s.Exec
}

func (s *LifecyclePreStop) SetExec(v *LifecyclePreStopExec) *LifecyclePreStop {
	s.Exec = v
	return s
}

func (s *LifecyclePreStop) Validate() error {
	if s.Exec != nil {
		if err := s.Exec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LifecyclePreStopExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s LifecyclePreStopExec) String() string {
	return dara.Prettify(s)
}

func (s LifecyclePreStopExec) GoString() string {
	return s.String()
}

func (s *LifecyclePreStopExec) GetCommand() []*string {
	return s.Command
}

func (s *LifecyclePreStopExec) SetCommand(v []*string) *LifecyclePreStopExec {
	s.Command = v
	return s
}

func (s *LifecyclePreStopExec) Validate() error {
	return dara.Validate(s)
}

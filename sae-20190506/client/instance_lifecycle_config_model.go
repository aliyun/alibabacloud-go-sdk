// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceLifecycleConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPreFreeze(v *LifecycleHook) *InstanceLifecycleConfig
	GetPreFreeze() *LifecycleHook
	SetPreStop(v *LifecycleHook) *InstanceLifecycleConfig
	GetPreStop() *LifecycleHook
}

type InstanceLifecycleConfig struct {
	PreFreeze *LifecycleHook `json:"preFreeze,omitempty" xml:"preFreeze,omitempty"`
	PreStop   *LifecycleHook `json:"preStop,omitempty" xml:"preStop,omitempty"`
}

func (s InstanceLifecycleConfig) String() string {
	return dara.Prettify(s)
}

func (s InstanceLifecycleConfig) GoString() string {
	return s.String()
}

func (s *InstanceLifecycleConfig) GetPreFreeze() *LifecycleHook {
	return s.PreFreeze
}

func (s *InstanceLifecycleConfig) GetPreStop() *LifecycleHook {
	return s.PreStop
}

func (s *InstanceLifecycleConfig) SetPreFreeze(v *LifecycleHook) *InstanceLifecycleConfig {
	s.PreFreeze = v
	return s
}

func (s *InstanceLifecycleConfig) SetPreStop(v *LifecycleHook) *InstanceLifecycleConfig {
	s.PreStop = v
	return s
}

func (s *InstanceLifecycleConfig) Validate() error {
	if s.PreFreeze != nil {
		if err := s.PreFreeze.Validate(); err != nil {
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

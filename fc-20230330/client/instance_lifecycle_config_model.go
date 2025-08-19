// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceLifecycleConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInitializer(v *LifecycleHook) *InstanceLifecycleConfig
	GetInitializer() *LifecycleHook
	SetPreStop(v *LifecycleHook) *InstanceLifecycleConfig
	GetPreStop() *LifecycleHook
}

type InstanceLifecycleConfig struct {
	Initializer *LifecycleHook `json:"initializer,omitempty" xml:"initializer,omitempty"`
	PreStop     *LifecycleHook `json:"preStop,omitempty" xml:"preStop,omitempty"`
}

func (s InstanceLifecycleConfig) String() string {
	return dara.Prettify(s)
}

func (s InstanceLifecycleConfig) GoString() string {
	return s.String()
}

func (s *InstanceLifecycleConfig) GetInitializer() *LifecycleHook {
	return s.Initializer
}

func (s *InstanceLifecycleConfig) GetPreStop() *LifecycleHook {
	return s.PreStop
}

func (s *InstanceLifecycleConfig) SetInitializer(v *LifecycleHook) *InstanceLifecycleConfig {
	s.Initializer = v
	return s
}

func (s *InstanceLifecycleConfig) SetPreStop(v *LifecycleHook) *InstanceLifecycleConfig {
	s.PreStop = v
	return s
}

func (s *InstanceLifecycleConfig) Validate() error {
	return dara.Validate(s)
}

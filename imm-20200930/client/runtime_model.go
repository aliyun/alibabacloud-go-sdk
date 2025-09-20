// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRuntime interface {
	dara.Model
	String() string
	GoString() string
	SetHyperparameters(v *Hyperparameters) *Runtime
	GetHyperparameters() *Hyperparameters
	SetResource(v *Resource) *Runtime
	GetResource() *Resource
}

type Runtime struct {
	// This parameter is required.
	Hyperparameters *Hyperparameters `json:"Hyperparameters,omitempty" xml:"Hyperparameters,omitempty"`
	// This parameter is required.
	Resource *Resource `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s Runtime) String() string {
	return dara.Prettify(s)
}

func (s Runtime) GoString() string {
	return s.String()
}

func (s *Runtime) GetHyperparameters() *Hyperparameters {
	return s.Hyperparameters
}

func (s *Runtime) GetResource() *Resource {
	return s.Resource
}

func (s *Runtime) SetHyperparameters(v *Hyperparameters) *Runtime {
	s.Hyperparameters = v
	return s
}

func (s *Runtime) SetResource(v *Resource) *Runtime {
	s.Resource = v
	return s
}

func (s *Runtime) Validate() error {
	return dara.Validate(s)
}

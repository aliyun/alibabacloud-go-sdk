// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceDeploymentSpec interface {
	dara.Model
	String() string
	GoString() string
	SetBaseline(v *ServiceBaseline) *ServiceDeploymentSpec
	GetBaseline() *ServiceBaseline
	SetChanges(v *ServiceChanges) *ServiceDeploymentSpec
	GetChanges() *ServiceChanges
	SetSkipRemoveResources(v bool) *ServiceDeploymentSpec
	GetSkipRemoveResources() *bool
	SetTarget(v *ServiceBaseline) *ServiceDeploymentSpec
	GetTarget() *ServiceBaseline
}

type ServiceDeploymentSpec struct {
	Baseline *ServiceBaseline `json:"baseline,omitempty" xml:"baseline,omitempty"`
	Changes  *ServiceChanges  `json:"changes,omitempty" xml:"changes,omitempty"`
	// example:
	//
	// false
	SkipRemoveResources *bool            `json:"skipRemoveResources,omitempty" xml:"skipRemoveResources,omitempty"`
	Target              *ServiceBaseline `json:"target,omitempty" xml:"target,omitempty"`
}

func (s ServiceDeploymentSpec) String() string {
	return dara.Prettify(s)
}

func (s ServiceDeploymentSpec) GoString() string {
	return s.String()
}

func (s *ServiceDeploymentSpec) GetBaseline() *ServiceBaseline {
	return s.Baseline
}

func (s *ServiceDeploymentSpec) GetChanges() *ServiceChanges {
	return s.Changes
}

func (s *ServiceDeploymentSpec) GetSkipRemoveResources() *bool {
	return s.SkipRemoveResources
}

func (s *ServiceDeploymentSpec) GetTarget() *ServiceBaseline {
	return s.Target
}

func (s *ServiceDeploymentSpec) SetBaseline(v *ServiceBaseline) *ServiceDeploymentSpec {
	s.Baseline = v
	return s
}

func (s *ServiceDeploymentSpec) SetChanges(v *ServiceChanges) *ServiceDeploymentSpec {
	s.Changes = v
	return s
}

func (s *ServiceDeploymentSpec) SetSkipRemoveResources(v bool) *ServiceDeploymentSpec {
	s.SkipRemoveResources = &v
	return s
}

func (s *ServiceDeploymentSpec) SetTarget(v *ServiceBaseline) *ServiceDeploymentSpec {
	s.Target = v
	return s
}

func (s *ServiceDeploymentSpec) Validate() error {
	return dara.Validate(s)
}

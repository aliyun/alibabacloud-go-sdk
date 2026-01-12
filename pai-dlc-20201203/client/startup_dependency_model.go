// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartupDependency interface {
	dara.Model
	String() string
	GoString() string
	SetMinReplicas(v string) *StartupDependency
	GetMinReplicas() *string
	SetOnPhase(v string) *StartupDependency
	GetOnPhase() *string
	SetType(v string) *StartupDependency
	GetType() *string
}

type StartupDependency struct {
	MinReplicas *string `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	OnPhase     *string `json:"OnPhase,omitempty" xml:"OnPhase,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StartupDependency) String() string {
	return dara.Prettify(s)
}

func (s StartupDependency) GoString() string {
	return s.String()
}

func (s *StartupDependency) GetMinReplicas() *string {
	return s.MinReplicas
}

func (s *StartupDependency) GetOnPhase() *string {
	return s.OnPhase
}

func (s *StartupDependency) GetType() *string {
	return s.Type
}

func (s *StartupDependency) SetMinReplicas(v string) *StartupDependency {
	s.MinReplicas = &v
	return s
}

func (s *StartupDependency) SetOnPhase(v string) *StartupDependency {
	s.OnPhase = &v
	return s
}

func (s *StartupDependency) SetType(v string) *StartupDependency {
	s.Type = &v
	return s
}

func (s *StartupDependency) Validate() error {
	return dara.Validate(s)
}

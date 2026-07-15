// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignalTarget interface {
	dara.Model
	String() string
	GoString() string
	SetPodNames(v []*string) *SignalTarget
	GetPodNames() []*string
	SetRoles(v []*string) *SignalTarget
	GetRoles() []*string
	SetScope(v string) *SignalTarget
	GetScope() *string
}

type SignalTarget struct {
	PodNames []*string `json:"PodNames,omitempty" xml:"PodNames,omitempty" type:"Repeated"`
	Roles    []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	Scope    *string   `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s SignalTarget) String() string {
	return dara.Prettify(s)
}

func (s SignalTarget) GoString() string {
	return s.String()
}

func (s *SignalTarget) GetPodNames() []*string {
	return s.PodNames
}

func (s *SignalTarget) GetRoles() []*string {
	return s.Roles
}

func (s *SignalTarget) GetScope() *string {
	return s.Scope
}

func (s *SignalTarget) SetPodNames(v []*string) *SignalTarget {
	s.PodNames = v
	return s
}

func (s *SignalTarget) SetRoles(v []*string) *SignalTarget {
	s.Roles = v
	return s
}

func (s *SignalTarget) SetScope(v string) *SignalTarget {
	s.Scope = &v
	return s
}

func (s *SignalTarget) Validate() error {
	return dara.Validate(s)
}

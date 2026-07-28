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
	// The pod name. Required when Scope is set to pods.
	PodNames []*string `json:"PodNames,omitempty" xml:"PodNames,omitempty" type:"Repeated"`
	// The role information. Required when Scope is set to roles.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The send scope.
	//
	// example:
	//
	// pods
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
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

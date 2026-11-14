// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicWarmPoolConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityPolicy(v *WarmPoolCapacityPolicy) *PublicWarmPoolConfig
	GetCapacityPolicy() *WarmPoolCapacityPolicy
	SetDeleting(v bool) *PublicWarmPoolConfig
	GetDeleting() *bool
	SetTeamID(v string) *PublicWarmPoolConfig
	GetTeamID() *string
	SetTemplateID(v string) *PublicWarmPoolConfig
	GetTemplateID() *string
}

type PublicWarmPoolConfig struct {
	// The warm pool capacity policy, which includes the base target capacity and optional scheduling rules.
	CapacityPolicy *WarmPoolCapacityPolicy `json:"capacityPolicy,omitempty" xml:"capacityPolicy,omitempty"`
	// Indicates whether the warm pool is being deleted.
	Deleting *bool `json:"deleting,omitempty" xml:"deleting,omitempty"`
	// The team ID.
	//
	// example:
	//
	// team-example1234567890
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The template ID, which specifies the template associated with the warm pool.
	//
	// example:
	//
	// tpl-example1234567890
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s PublicWarmPoolConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicWarmPoolConfig) GoString() string {
	return s.String()
}

func (s *PublicWarmPoolConfig) GetCapacityPolicy() *WarmPoolCapacityPolicy {
	return s.CapacityPolicy
}

func (s *PublicWarmPoolConfig) GetDeleting() *bool {
	return s.Deleting
}

func (s *PublicWarmPoolConfig) GetTeamID() *string {
	return s.TeamID
}

func (s *PublicWarmPoolConfig) GetTemplateID() *string {
	return s.TemplateID
}

func (s *PublicWarmPoolConfig) SetCapacityPolicy(v *WarmPoolCapacityPolicy) *PublicWarmPoolConfig {
	s.CapacityPolicy = v
	return s
}

func (s *PublicWarmPoolConfig) SetDeleting(v bool) *PublicWarmPoolConfig {
	s.Deleting = &v
	return s
}

func (s *PublicWarmPoolConfig) SetTeamID(v string) *PublicWarmPoolConfig {
	s.TeamID = &v
	return s
}

func (s *PublicWarmPoolConfig) SetTemplateID(v string) *PublicWarmPoolConfig {
	s.TemplateID = &v
	return s
}

func (s *PublicWarmPoolConfig) Validate() error {
	if s.CapacityPolicy != nil {
		if err := s.CapacityPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

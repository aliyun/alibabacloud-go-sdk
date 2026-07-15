// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeccompProfile interface {
	dara.Model
	String() string
	GoString() string
	SetLocalhostProfile(v string) *SeccompProfile
	GetLocalhostProfile() *string
	SetType(v string) *SeccompProfile
	GetType() *string
}

type SeccompProfile struct {
	// The path of the Seccomp profile on the node. This parameter takes effect only when Type is set to Localhost.
	//
	// example:
	//
	// my-profiles/profile-allow.json
	LocalhostProfile *string `json:"LocalhostProfile,omitempty" xml:"LocalhostProfile,omitempty"`
	// The Seccomp configuration type. Valid values: Localhost, RuntimeDefault, Unconfined.
	//
	// example:
	//
	// Unconfined
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SeccompProfile) String() string {
	return dara.Prettify(s)
}

func (s SeccompProfile) GoString() string {
	return s.String()
}

func (s *SeccompProfile) GetLocalhostProfile() *string {
	return s.LocalhostProfile
}

func (s *SeccompProfile) GetType() *string {
	return s.Type
}

func (s *SeccompProfile) SetLocalhostProfile(v string) *SeccompProfile {
	s.LocalhostProfile = &v
	return s
}

func (s *SeccompProfile) SetType(v string) *SeccompProfile {
	s.Type = &v
	return s
}

func (s *SeccompProfile) Validate() error {
	return dara.Validate(s)
}

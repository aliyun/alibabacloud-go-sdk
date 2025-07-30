// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecurityContext interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilities(v *SecurityContextCapabilities) *SecurityContext
	GetCapabilities() *SecurityContextCapabilities
	SetPrivileged(v bool) *SecurityContext
	GetPrivileged() *bool
	SetRunAsGroup(v int64) *SecurityContext
	GetRunAsGroup() *int64
	SetRunAsUser(v int64) *SecurityContext
	GetRunAsUser() *int64
	SetSeccompProfile(v *SeccompProfile) *SecurityContext
	GetSeccompProfile() *SeccompProfile
}

type SecurityContext struct {
	Capabilities *SecurityContextCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty"`
	Privileged   *bool                        `json:"Privileged,omitempty" xml:"Privileged,omitempty"`
	// example:
	//
	// 1000
	RunAsGroup *int64 `json:"RunAsGroup,omitempty" xml:"RunAsGroup,omitempty"`
	// example:
	//
	// 1000
	RunAsUser      *int64          `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
	SeccompProfile *SeccompProfile `json:"SeccompProfile,omitempty" xml:"SeccompProfile,omitempty"`
}

func (s SecurityContext) String() string {
	return dara.Prettify(s)
}

func (s SecurityContext) GoString() string {
	return s.String()
}

func (s *SecurityContext) GetCapabilities() *SecurityContextCapabilities {
	return s.Capabilities
}

func (s *SecurityContext) GetPrivileged() *bool {
	return s.Privileged
}

func (s *SecurityContext) GetRunAsGroup() *int64 {
	return s.RunAsGroup
}

func (s *SecurityContext) GetRunAsUser() *int64 {
	return s.RunAsUser
}

func (s *SecurityContext) GetSeccompProfile() *SeccompProfile {
	return s.SeccompProfile
}

func (s *SecurityContext) SetCapabilities(v *SecurityContextCapabilities) *SecurityContext {
	s.Capabilities = v
	return s
}

func (s *SecurityContext) SetPrivileged(v bool) *SecurityContext {
	s.Privileged = &v
	return s
}

func (s *SecurityContext) SetRunAsGroup(v int64) *SecurityContext {
	s.RunAsGroup = &v
	return s
}

func (s *SecurityContext) SetRunAsUser(v int64) *SecurityContext {
	s.RunAsUser = &v
	return s
}

func (s *SecurityContext) SetSeccompProfile(v *SeccompProfile) *SecurityContext {
	s.SeccompProfile = v
	return s
}

func (s *SecurityContext) Validate() error {
	return dara.Validate(s)
}

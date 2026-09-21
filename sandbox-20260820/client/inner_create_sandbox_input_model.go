// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInnerCreateSandboxInput interface {
	dara.Model
	String() string
	GoString() string
	SetAllowInternetAccess(v bool) *InnerCreateSandboxInput
	GetAllowInternetAccess() *bool
	SetAutoPause(v bool) *InnerCreateSandboxInput
	GetAutoPause() *bool
	SetAutoResume(v bool) *InnerCreateSandboxInput
	GetAutoResume() *bool
	SetEnvVars(v map[string]*string) *InnerCreateSandboxInput
	GetEnvVars() map[string]*string
	SetMetadata(v map[string]*string) *InnerCreateSandboxInput
	GetMetadata() map[string]*string
	SetNetwork(v *E2BNetwork) *InnerCreateSandboxInput
	GetNetwork() *E2BNetwork
	SetRuntime(v *InnerSandboxRuntimeConfig) *InnerCreateSandboxInput
	GetRuntime() *InnerSandboxRuntimeConfig
	SetSecure(v bool) *InnerCreateSandboxInput
	GetSecure() *bool
	SetTeamID(v string) *InnerCreateSandboxInput
	GetTeamID() *string
	SetTemplateID(v string) *InnerCreateSandboxInput
	GetTemplateID() *string
	SetTimeout(v int32) *InnerCreateSandboxInput
	GetTimeout() *int32
	SetVolumeMounts(v *InnerCreateSandboxVolumeMounts) *InnerCreateSandboxInput
	GetVolumeMounts() *InnerCreateSandboxVolumeMounts
}

type InnerCreateSandboxInput struct {
	// example:
	//
	// true
	AllowInternetAccess *bool `json:"allowInternetAccess,omitempty" xml:"allowInternetAccess,omitempty"`
	// example:
	//
	// true
	AutoPause *bool `json:"autoPause,omitempty" xml:"autoPause,omitempty"`
	// example:
	//
	// true
	AutoResume *bool                      `json:"autoResume,omitempty" xml:"autoResume,omitempty"`
	EnvVars    map[string]*string         `json:"envVars,omitempty" xml:"envVars,omitempty"`
	Metadata   map[string]*string         `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Network    *E2BNetwork                `json:"network,omitempty" xml:"network,omitempty"`
	Runtime    *InnerSandboxRuntimeConfig `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// example:
	//
	// true
	Secure *bool `json:"secure,omitempty" xml:"secure,omitempty"`
	// example:
	//
	// 76eeecaa-****
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// example:
	//
	// v9vjyzw64bsu83vw0dgs
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
	// example:
	//
	// 180
	Timeout      *int32                          `json:"timeout,omitempty" xml:"timeout,omitempty"`
	VolumeMounts *InnerCreateSandboxVolumeMounts `json:"volumeMounts,omitempty" xml:"volumeMounts,omitempty"`
}

func (s InnerCreateSandboxInput) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxInput) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxInput) GetAllowInternetAccess() *bool {
	return s.AllowInternetAccess
}

func (s *InnerCreateSandboxInput) GetAutoPause() *bool {
	return s.AutoPause
}

func (s *InnerCreateSandboxInput) GetAutoResume() *bool {
	return s.AutoResume
}

func (s *InnerCreateSandboxInput) GetEnvVars() map[string]*string {
	return s.EnvVars
}

func (s *InnerCreateSandboxInput) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *InnerCreateSandboxInput) GetNetwork() *E2BNetwork {
	return s.Network
}

func (s *InnerCreateSandboxInput) GetRuntime() *InnerSandboxRuntimeConfig {
	return s.Runtime
}

func (s *InnerCreateSandboxInput) GetSecure() *bool {
	return s.Secure
}

func (s *InnerCreateSandboxInput) GetTeamID() *string {
	return s.TeamID
}

func (s *InnerCreateSandboxInput) GetTemplateID() *string {
	return s.TemplateID
}

func (s *InnerCreateSandboxInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *InnerCreateSandboxInput) GetVolumeMounts() *InnerCreateSandboxVolumeMounts {
	return s.VolumeMounts
}

func (s *InnerCreateSandboxInput) SetAllowInternetAccess(v bool) *InnerCreateSandboxInput {
	s.AllowInternetAccess = &v
	return s
}

func (s *InnerCreateSandboxInput) SetAutoPause(v bool) *InnerCreateSandboxInput {
	s.AutoPause = &v
	return s
}

func (s *InnerCreateSandboxInput) SetAutoResume(v bool) *InnerCreateSandboxInput {
	s.AutoResume = &v
	return s
}

func (s *InnerCreateSandboxInput) SetEnvVars(v map[string]*string) *InnerCreateSandboxInput {
	s.EnvVars = v
	return s
}

func (s *InnerCreateSandboxInput) SetMetadata(v map[string]*string) *InnerCreateSandboxInput {
	s.Metadata = v
	return s
}

func (s *InnerCreateSandboxInput) SetNetwork(v *E2BNetwork) *InnerCreateSandboxInput {
	s.Network = v
	return s
}

func (s *InnerCreateSandboxInput) SetRuntime(v *InnerSandboxRuntimeConfig) *InnerCreateSandboxInput {
	s.Runtime = v
	return s
}

func (s *InnerCreateSandboxInput) SetSecure(v bool) *InnerCreateSandboxInput {
	s.Secure = &v
	return s
}

func (s *InnerCreateSandboxInput) SetTeamID(v string) *InnerCreateSandboxInput {
	s.TeamID = &v
	return s
}

func (s *InnerCreateSandboxInput) SetTemplateID(v string) *InnerCreateSandboxInput {
	s.TemplateID = &v
	return s
}

func (s *InnerCreateSandboxInput) SetTimeout(v int32) *InnerCreateSandboxInput {
	s.Timeout = &v
	return s
}

func (s *InnerCreateSandboxInput) SetVolumeMounts(v *InnerCreateSandboxVolumeMounts) *InnerCreateSandboxInput {
	s.VolumeMounts = v
	return s
}

func (s *InnerCreateSandboxInput) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	if s.VolumeMounts != nil {
		if err := s.VolumeMounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

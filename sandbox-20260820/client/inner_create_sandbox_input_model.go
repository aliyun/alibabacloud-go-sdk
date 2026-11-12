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
	SetResourceTags(v []*InnerCreateSandboxInputResourceTags) *InnerCreateSandboxInput
	GetResourceTags() []*InnerCreateSandboxInputResourceTags
	SetRuntime(v *InnerSandboxRuntimeConfig) *InnerCreateSandboxInput
	GetRuntime() *InnerSandboxRuntimeConfig
	SetSecure(v bool) *InnerCreateSandboxInput
	GetSecure() *bool
	SetSnapshotID(v string) *InnerCreateSandboxInput
	GetSnapshotID() *string
	SetSnapshotName(v string) *InnerCreateSandboxInput
	GetSnapshotName() *string
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
	// Specifies whether the sandbox is allowed to access external networks.
	//
	// example:
	//
	// true
	AllowInternetAccess *bool `json:"allowInternetAccess,omitempty" xml:"allowInternetAccess,omitempty"`
	// Specifies whether to automatically pause the sandbox after a timeout.
	//
	// example:
	//
	// true
	AutoPause *bool `json:"autoPause,omitempty" xml:"autoPause,omitempty"`
	// Specifies whether to automatically resume the sandbox.
	//
	// example:
	//
	// true
	AutoResume *bool `json:"autoResume,omitempty" xml:"autoResume,omitempty"`
	// The environment variable configuration.
	//
	// example:
	//
	// {"APP_MODE":"demo"}
	EnvVars map[string]*string `json:"envVars" xml:"envVars"`
	// The metadata configuration.
	Metadata map[string]*string `json:"metadata" xml:"metadata"`
	// The network configuration.
	Network *E2BNetwork `json:"network,omitempty" xml:"network,omitempty"`
	// The resource tags.
	ResourceTags []*InnerCreateSandboxInputResourceTags `json:"resourceTags" xml:"resourceTags" type:"Repeated"`
	// The runtime configuration.
	Runtime *InnerSandboxRuntimeConfig `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Specifies whether authentication is required to access the sandbox envd.
	//
	// example:
	//
	// true
	Secure *bool `json:"secure,omitempty" xml:"secure,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// 0d66****b374
	SnapshotID *string `json:"snapshotID,omitempty" xml:"snapshotID,omitempty"`
	// The snapshot name.
	//
	// example:
	//
	// name
	SnapshotName *string `json:"snapshotName,omitempty" xml:"snapshotName,omitempty"`
	// The unique identifier of the team.
	//
	// example:
	//
	// 76eeecaa-****
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The unique identifier of the template.
	//
	// example:
	//
	// v9vjyzw64bsu83vw0dgs
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
	// The timeout period in seconds.
	//
	// example:
	//
	// 180
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The volume mount configuration.
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

func (s *InnerCreateSandboxInput) GetResourceTags() []*InnerCreateSandboxInputResourceTags {
	return s.ResourceTags
}

func (s *InnerCreateSandboxInput) GetRuntime() *InnerSandboxRuntimeConfig {
	return s.Runtime
}

func (s *InnerCreateSandboxInput) GetSecure() *bool {
	return s.Secure
}

func (s *InnerCreateSandboxInput) GetSnapshotID() *string {
	return s.SnapshotID
}

func (s *InnerCreateSandboxInput) GetSnapshotName() *string {
	return s.SnapshotName
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

func (s *InnerCreateSandboxInput) SetResourceTags(v []*InnerCreateSandboxInputResourceTags) *InnerCreateSandboxInput {
	s.ResourceTags = v
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

func (s *InnerCreateSandboxInput) SetSnapshotID(v string) *InnerCreateSandboxInput {
	s.SnapshotID = &v
	return s
}

func (s *InnerCreateSandboxInput) SetSnapshotName(v string) *InnerCreateSandboxInput {
	s.SnapshotName = &v
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
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type InnerCreateSandboxInputResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// project
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// demo
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InnerCreateSandboxInputResourceTags) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxInputResourceTags) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxInputResourceTags) GetKey() *string {
	return s.Key
}

func (s *InnerCreateSandboxInputResourceTags) GetValue() *string {
	return s.Value
}

func (s *InnerCreateSandboxInputResourceTags) SetKey(v string) *InnerCreateSandboxInputResourceTags {
	s.Key = &v
	return s
}

func (s *InnerCreateSandboxInputResourceTags) SetValue(v string) *InnerCreateSandboxInputResourceTags {
	s.Value = &v
	return s
}

func (s *InnerCreateSandboxInputResourceTags) Validate() error {
	return dara.Validate(s)
}

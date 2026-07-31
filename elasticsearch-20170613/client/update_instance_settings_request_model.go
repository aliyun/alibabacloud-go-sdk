// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEsConfig(v map[string]*string) *UpdateInstanceSettingsRequest
	GetEsConfig() map[string]*string
	SetClientToken(v string) *UpdateInstanceSettingsRequest
	GetClientToken() *string
	SetForce(v bool) *UpdateInstanceSettingsRequest
	GetForce() *bool
	SetUpdateStrategy(v string) *UpdateInstanceSettingsRequest
	GetUpdateStrategy() *string
}

type UpdateInstanceSettingsRequest struct {
	// The YML file configuration of the instance.
	EsConfig map[string]*string `json:"esConfig,omitempty" xml:"esConfig,omitempty"`
	// A unique token used to ensure the idempotence of the request. The client generates this value. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to forcibly apply the change.
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
	// The change strategy for Elasticsearch (for example, the change method used during index updates, cluster upgrades, or service deployments). Valid values:
	//
	// - blue_green: blue-green change. Achieves seamless switchover by running two identical environments (blue and green) in parallel.
	//
	// - normal: in-place change. Performs changes directly in the current environment (for example, upgrades or scaling) without requiring additional resources.
	//
	// - intelligent: intelligent change. The system automatically analyzes the change type and environment state, and dynamically selects the optimal change method (blue-green change or in-place change).
	//
	// example:
	//
	// normal
	UpdateStrategy *string `json:"updateStrategy,omitempty" xml:"updateStrategy,omitempty"`
}

func (s UpdateInstanceSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSettingsRequest) GetEsConfig() map[string]*string {
	return s.EsConfig
}

func (s *UpdateInstanceSettingsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateInstanceSettingsRequest) GetForce() *bool {
	return s.Force
}

func (s *UpdateInstanceSettingsRequest) GetUpdateStrategy() *string {
	return s.UpdateStrategy
}

func (s *UpdateInstanceSettingsRequest) SetEsConfig(v map[string]*string) *UpdateInstanceSettingsRequest {
	s.EsConfig = v
	return s
}

func (s *UpdateInstanceSettingsRequest) SetClientToken(v string) *UpdateInstanceSettingsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceSettingsRequest) SetForce(v bool) *UpdateInstanceSettingsRequest {
	s.Force = &v
	return s
}

func (s *UpdateInstanceSettingsRequest) SetUpdateStrategy(v string) *UpdateInstanceSettingsRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *UpdateInstanceSettingsRequest) Validate() error {
	return dara.Validate(s)
}

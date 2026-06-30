// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *UpdatePolarClawSkillShrinkRequest
	GetApiKey() *string
	SetApplicationId(v string) *UpdatePolarClawSkillShrinkRequest
	GetApplicationId() *string
	SetEnabled(v bool) *UpdatePolarClawSkillShrinkRequest
	GetEnabled() *bool
	SetEnvShrink(v string) *UpdatePolarClawSkillShrinkRequest
	GetEnvShrink() *string
	SetSkillKey(v string) *UpdatePolarClawSkillShrinkRequest
	GetSkillKey() *string
}

type UpdatePolarClawSkillShrinkRequest struct {
	// The Skill API key. An empty string indicates that the key is deleted.
	//
	// example:
	//
	// my-api
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Specifies whether to enable the Skill.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The environment variables as a key-value map. A null value indicates that the variable is deleted.
	//
	// example:
	//
	// {"NETA_TOKEN":"my-token"}
	EnvShrink *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The Skill identifier key.
	//
	// This parameter is required.
	//
	// example:
	//
	// alibacloud-rds-copilot
	SkillKey *string `json:"SkillKey,omitempty" xml:"SkillKey,omitempty"`
}

func (s UpdatePolarClawSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawSkillShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdatePolarClawSkillShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawSkillShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePolarClawSkillShrinkRequest) GetEnvShrink() *string {
	return s.EnvShrink
}

func (s *UpdatePolarClawSkillShrinkRequest) GetSkillKey() *string {
	return s.SkillKey
}

func (s *UpdatePolarClawSkillShrinkRequest) SetApiKey(v string) *UpdatePolarClawSkillShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *UpdatePolarClawSkillShrinkRequest) SetApplicationId(v string) *UpdatePolarClawSkillShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawSkillShrinkRequest) SetEnabled(v bool) *UpdatePolarClawSkillShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdatePolarClawSkillShrinkRequest) SetEnvShrink(v string) *UpdatePolarClawSkillShrinkRequest {
	s.EnvShrink = &v
	return s
}

func (s *UpdatePolarClawSkillShrinkRequest) SetSkillKey(v string) *UpdatePolarClawSkillShrinkRequest {
	s.SkillKey = &v
	return s
}

func (s *UpdatePolarClawSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}

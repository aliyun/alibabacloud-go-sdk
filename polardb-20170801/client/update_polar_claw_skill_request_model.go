// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *UpdatePolarClawSkillRequest
	GetApiKey() *string
	SetApplicationId(v string) *UpdatePolarClawSkillRequest
	GetApplicationId() *string
	SetEnabled(v bool) *UpdatePolarClawSkillRequest
	GetEnabled() *bool
	SetEnv(v map[string]*string) *UpdatePolarClawSkillRequest
	GetEnv() map[string]*string
	SetSkillKey(v string) *UpdatePolarClawSkillRequest
	GetSkillKey() *string
}

type UpdatePolarClawSkillRequest struct {
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
	Env map[string]*string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The Skill identifier key.
	//
	// This parameter is required.
	//
	// example:
	//
	// alibacloud-rds-copilot
	SkillKey *string `json:"SkillKey,omitempty" xml:"SkillKey,omitempty"`
}

func (s UpdatePolarClawSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawSkillRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdatePolarClawSkillRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawSkillRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePolarClawSkillRequest) GetEnv() map[string]*string {
	return s.Env
}

func (s *UpdatePolarClawSkillRequest) GetSkillKey() *string {
	return s.SkillKey
}

func (s *UpdatePolarClawSkillRequest) SetApiKey(v string) *UpdatePolarClawSkillRequest {
	s.ApiKey = &v
	return s
}

func (s *UpdatePolarClawSkillRequest) SetApplicationId(v string) *UpdatePolarClawSkillRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawSkillRequest) SetEnabled(v bool) *UpdatePolarClawSkillRequest {
	s.Enabled = &v
	return s
}

func (s *UpdatePolarClawSkillRequest) SetEnv(v map[string]*string) *UpdatePolarClawSkillRequest {
	s.Env = v
	return s
}

func (s *UpdatePolarClawSkillRequest) SetSkillKey(v string) *UpdatePolarClawSkillRequest {
	s.SkillKey = &v
	return s
}

func (s *UpdatePolarClawSkillRequest) Validate() error {
	return dara.Validate(s)
}

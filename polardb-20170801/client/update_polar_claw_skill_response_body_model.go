// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdatePolarClawSkillResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UpdatePolarClawSkillResponseBody
	GetCode() *int32
	SetConfig(v *UpdatePolarClawSkillResponseBodyConfig) *UpdatePolarClawSkillResponseBody
	GetConfig() *UpdatePolarClawSkillResponseBodyConfig
	SetMessage(v string) *UpdatePolarClawSkillResponseBody
	GetMessage() *string
	SetOk(v bool) *UpdatePolarClawSkillResponseBody
	GetOk() *bool
	SetRequestId(v string) *UpdatePolarClawSkillResponseBody
	GetRequestId() *string
	SetSkillKey(v string) *UpdatePolarClawSkillResponseBody
	GetSkillKey() *string
}

type UpdatePolarClawSkillResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The updated Skill configuration. Sensitive values are masked.
	Config *UpdatePolarClawSkillResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Skill identifier key.
	//
	// example:
	//
	// alibacloud-rds-copilot
	SkillKey *string `json:"SkillKey,omitempty" xml:"SkillKey,omitempty"`
}

func (s UpdatePolarClawSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawSkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawSkillResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawSkillResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePolarClawSkillResponseBody) GetConfig() *UpdatePolarClawSkillResponseBodyConfig {
	return s.Config
}

func (s *UpdatePolarClawSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolarClawSkillResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *UpdatePolarClawSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolarClawSkillResponseBody) GetSkillKey() *string {
	return s.SkillKey
}

func (s *UpdatePolarClawSkillResponseBody) SetApplicationId(v string) *UpdatePolarClawSkillResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawSkillResponseBody) SetCode(v int32) *UpdatePolarClawSkillResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolarClawSkillResponseBody) SetConfig(v *UpdatePolarClawSkillResponseBodyConfig) *UpdatePolarClawSkillResponseBody {
	s.Config = v
	return s
}

func (s *UpdatePolarClawSkillResponseBody) SetMessage(v string) *UpdatePolarClawSkillResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolarClawSkillResponseBody) SetOk(v bool) *UpdatePolarClawSkillResponseBody {
	s.Ok = &v
	return s
}

func (s *UpdatePolarClawSkillResponseBody) SetRequestId(v string) *UpdatePolarClawSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolarClawSkillResponseBody) SetSkillKey(v string) *UpdatePolarClawSkillResponseBody {
	s.SkillKey = &v
	return s
}

func (s *UpdatePolarClawSkillResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolarClawSkillResponseBodyConfig struct {
	// Specifies whether to enable the Skill.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The environment variable configuration.
	Env map[string]*string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s UpdatePolarClawSkillResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawSkillResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawSkillResponseBodyConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePolarClawSkillResponseBodyConfig) GetEnv() map[string]*string {
	return s.Env
}

func (s *UpdatePolarClawSkillResponseBodyConfig) SetEnabled(v bool) *UpdatePolarClawSkillResponseBodyConfig {
	s.Enabled = &v
	return s
}

func (s *UpdatePolarClawSkillResponseBodyConfig) SetEnv(v map[string]*string) *UpdatePolarClawSkillResponseBodyConfig {
	s.Env = v
	return s
}

func (s *UpdatePolarClawSkillResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}

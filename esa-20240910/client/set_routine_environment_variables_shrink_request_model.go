// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRoutineEnvironmentVariablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *SetRoutineEnvironmentVariablesShrinkRequest
	GetEnv() *string
	SetEnvironmentVariablesShrink(v string) *SetRoutineEnvironmentVariablesShrinkRequest
	GetEnvironmentVariablesShrink() *string
	SetName(v string) *SetRoutineEnvironmentVariablesShrinkRequest
	GetName() *string
}

type SetRoutineEnvironmentVariablesShrinkRequest struct {
	// The environment name. Valid values:
	//
	// - `staging`: staging environment.
	//
	// - `production`: production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// production
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The dictionary of environment variables. The key is the environment variable name, and the value is the environment variable value.
	//
	// This parameter is required.
	//
	// example:
	//
	// "EnvironmentVariables": {
	//
	//         "Env_Key_1": {
	//
	//             "Type": "plain_text",
	//
	//             "Value": "value"
	//
	//         },
	//
	//         "PASSWORD": {
	//
	//             "Type": "secret_text",
	//
	//             "Value": "secret-password"
	//
	//         }
	//
	//     }
	EnvironmentVariablesShrink *string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The function name.
	//
	// This parameter is required.
	//
	// example:
	//
	// er_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetRoutineEnvironmentVariablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRoutineEnvironmentVariablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetRoutineEnvironmentVariablesShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *SetRoutineEnvironmentVariablesShrinkRequest) GetEnvironmentVariablesShrink() *string {
	return s.EnvironmentVariablesShrink
}

func (s *SetRoutineEnvironmentVariablesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SetRoutineEnvironmentVariablesShrinkRequest) SetEnv(v string) *SetRoutineEnvironmentVariablesShrinkRequest {
	s.Env = &v
	return s
}

func (s *SetRoutineEnvironmentVariablesShrinkRequest) SetEnvironmentVariablesShrink(v string) *SetRoutineEnvironmentVariablesShrinkRequest {
	s.EnvironmentVariablesShrink = &v
	return s
}

func (s *SetRoutineEnvironmentVariablesShrinkRequest) SetName(v string) *SetRoutineEnvironmentVariablesShrinkRequest {
	s.Name = &v
	return s
}

func (s *SetRoutineEnvironmentVariablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

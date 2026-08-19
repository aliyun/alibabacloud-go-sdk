// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRoutineEnvironmentVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *SetRoutineEnvironmentVariablesRequest
	GetEnv() *string
	SetEnvironmentVariables(v map[string]*EnvironmentVariablesValue) *SetRoutineEnvironmentVariablesRequest
	GetEnvironmentVariables() map[string]*EnvironmentVariablesValue
	SetName(v string) *SetRoutineEnvironmentVariablesRequest
	GetName() *string
}

type SetRoutineEnvironmentVariablesRequest struct {
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
	EnvironmentVariables map[string]*EnvironmentVariablesValue `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The function name.
	//
	// This parameter is required.
	//
	// example:
	//
	// er_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetRoutineEnvironmentVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRoutineEnvironmentVariablesRequest) GoString() string {
	return s.String()
}

func (s *SetRoutineEnvironmentVariablesRequest) GetEnv() *string {
	return s.Env
}

func (s *SetRoutineEnvironmentVariablesRequest) GetEnvironmentVariables() map[string]*EnvironmentVariablesValue {
	return s.EnvironmentVariables
}

func (s *SetRoutineEnvironmentVariablesRequest) GetName() *string {
	return s.Name
}

func (s *SetRoutineEnvironmentVariablesRequest) SetEnv(v string) *SetRoutineEnvironmentVariablesRequest {
	s.Env = &v
	return s
}

func (s *SetRoutineEnvironmentVariablesRequest) SetEnvironmentVariables(v map[string]*EnvironmentVariablesValue) *SetRoutineEnvironmentVariablesRequest {
	s.EnvironmentVariables = v
	return s
}

func (s *SetRoutineEnvironmentVariablesRequest) SetName(v string) *SetRoutineEnvironmentVariablesRequest {
	s.Name = &v
	return s
}

func (s *SetRoutineEnvironmentVariablesRequest) Validate() error {
	return dara.Validate(s)
}

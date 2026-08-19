// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineEnvironmentVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *DeleteRoutineEnvironmentVariablesRequest
	GetEnv() *string
	SetEnvironmentVariableKeys(v []*string) *DeleteRoutineEnvironmentVariablesRequest
	GetEnvironmentVariableKeys() []*string
	SetName(v string) *DeleteRoutineEnvironmentVariablesRequest
	GetName() *string
}

type DeleteRoutineEnvironmentVariablesRequest struct {
	// The environment name.
	//
	// Valid values:
	//
	// - `staging`: staging environment
	//
	// - `production`: production environment
	//
	// This parameter is required.
	//
	// example:
	//
	// production
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The list of environment variable keys to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["LOG_LEVEL"]
	EnvironmentVariableKeys []*string `json:"EnvironmentVariableKeys,omitempty" xml:"EnvironmentVariableKeys,omitempty" type:"Repeated"`
	// The name of the Routine function.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-routine
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteRoutineEnvironmentVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineEnvironmentVariablesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineEnvironmentVariablesRequest) GetEnv() *string {
	return s.Env
}

func (s *DeleteRoutineEnvironmentVariablesRequest) GetEnvironmentVariableKeys() []*string {
	return s.EnvironmentVariableKeys
}

func (s *DeleteRoutineEnvironmentVariablesRequest) GetName() *string {
	return s.Name
}

func (s *DeleteRoutineEnvironmentVariablesRequest) SetEnv(v string) *DeleteRoutineEnvironmentVariablesRequest {
	s.Env = &v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesRequest) SetEnvironmentVariableKeys(v []*string) *DeleteRoutineEnvironmentVariablesRequest {
	s.EnvironmentVariableKeys = v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesRequest) SetName(v string) *DeleteRoutineEnvironmentVariablesRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineEnvironmentVariablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *DeleteRoutineEnvironmentVariablesShrinkRequest
	GetEnv() *string
	SetEnvironmentVariableKeysShrink(v string) *DeleteRoutineEnvironmentVariablesShrinkRequest
	GetEnvironmentVariableKeysShrink() *string
	SetName(v string) *DeleteRoutineEnvironmentVariablesShrinkRequest
	GetName() *string
}

type DeleteRoutineEnvironmentVariablesShrinkRequest struct {
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
	EnvironmentVariableKeysShrink *string `json:"EnvironmentVariableKeys,omitempty" xml:"EnvironmentVariableKeys,omitempty"`
	// The name of the Routine function.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-routine
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteRoutineEnvironmentVariablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineEnvironmentVariablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineEnvironmentVariablesShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *DeleteRoutineEnvironmentVariablesShrinkRequest) GetEnvironmentVariableKeysShrink() *string {
	return s.EnvironmentVariableKeysShrink
}

func (s *DeleteRoutineEnvironmentVariablesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *DeleteRoutineEnvironmentVariablesShrinkRequest) SetEnv(v string) *DeleteRoutineEnvironmentVariablesShrinkRequest {
	s.Env = &v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesShrinkRequest) SetEnvironmentVariableKeysShrink(v string) *DeleteRoutineEnvironmentVariablesShrinkRequest {
	s.EnvironmentVariableKeysShrink = &v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesShrinkRequest) SetName(v string) *DeleteRoutineEnvironmentVariablesShrinkRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

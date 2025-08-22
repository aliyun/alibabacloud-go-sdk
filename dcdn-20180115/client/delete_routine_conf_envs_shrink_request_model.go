// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineConfEnvsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvsShrink(v string) *DeleteRoutineConfEnvsShrinkRequest
	GetEnvsShrink() *string
	SetName(v string) *DeleteRoutineConfEnvsShrinkRequest
	GetName() *string
}

type DeleteRoutineConfEnvsShrinkRequest struct {
	// The custom canary release environments that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["presetCanaryZheJiang"]
	EnvsShrink *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteRoutineConfEnvsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineConfEnvsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineConfEnvsShrinkRequest) GetEnvsShrink() *string {
	return s.EnvsShrink
}

func (s *DeleteRoutineConfEnvsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *DeleteRoutineConfEnvsShrinkRequest) SetEnvsShrink(v string) *DeleteRoutineConfEnvsShrinkRequest {
	s.EnvsShrink = &v
	return s
}

func (s *DeleteRoutineConfEnvsShrinkRequest) SetName(v string) *DeleteRoutineConfEnvsShrinkRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineConfEnvsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

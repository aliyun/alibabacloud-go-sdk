// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineConfEnvsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvs(v map[string]interface{}) *DeleteRoutineConfEnvsRequest
	GetEnvs() map[string]interface{}
	SetName(v string) *DeleteRoutineConfEnvsRequest
	GetName() *string
}

type DeleteRoutineConfEnvsRequest struct {
	// The custom canary release environments that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["presetCanaryZheJiang"]
	Envs map[string]interface{} `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteRoutineConfEnvsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineConfEnvsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineConfEnvsRequest) GetEnvs() map[string]interface{} {
	return s.Envs
}

func (s *DeleteRoutineConfEnvsRequest) GetName() *string {
	return s.Name
}

func (s *DeleteRoutineConfEnvsRequest) SetEnvs(v map[string]interface{}) *DeleteRoutineConfEnvsRequest {
	s.Envs = v
	return s
}

func (s *DeleteRoutineConfEnvsRequest) SetName(v string) *DeleteRoutineConfEnvsRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineConfEnvsRequest) Validate() error {
	return dara.Validate(s)
}

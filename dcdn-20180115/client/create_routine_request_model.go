// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRoutineRequest
	GetDescription() *string
	SetEnvConf(v map[string]interface{}) *CreateRoutineRequest
	GetEnvConf() map[string]interface{}
	SetName(v string) *CreateRoutineRequest
	GetName() *string
}

type CreateRoutineRequest struct {
	// The description of the routine.
	//
	// example:
	//
	// the description of this routine
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configurations of the specified environment.
	//
	// example:
	//
	// {"staging":{"SpecName":"50ms"},"production":{"SpecName":"50ms"}}
	EnvConf map[string]interface{} `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRoutineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRoutineRequest) GetEnvConf() map[string]interface{} {
	return s.EnvConf
}

func (s *CreateRoutineRequest) GetName() *string {
	return s.Name
}

func (s *CreateRoutineRequest) SetDescription(v string) *CreateRoutineRequest {
	s.Description = &v
	return s
}

func (s *CreateRoutineRequest) SetEnvConf(v map[string]interface{}) *CreateRoutineRequest {
	s.EnvConf = v
	return s
}

func (s *CreateRoutineRequest) SetName(v string) *CreateRoutineRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineRequest) Validate() error {
	return dara.Validate(s)
}

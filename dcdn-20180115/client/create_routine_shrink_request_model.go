// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRoutineShrinkRequest
	GetDescription() *string
	SetEnvConfShrink(v string) *CreateRoutineShrinkRequest
	GetEnvConfShrink() *string
	SetName(v string) *CreateRoutineShrinkRequest
	GetName() *string
}

type CreateRoutineShrinkRequest struct {
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
	EnvConfShrink *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRoutineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRoutineShrinkRequest) GetEnvConfShrink() *string {
	return s.EnvConfShrink
}

func (s *CreateRoutineShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateRoutineShrinkRequest) SetDescription(v string) *CreateRoutineShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRoutineShrinkRequest) SetEnvConfShrink(v string) *CreateRoutineShrinkRequest {
	s.EnvConfShrink = &v
	return s
}

func (s *CreateRoutineShrinkRequest) SetName(v string) *CreateRoutineShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineCodeDeploymentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersionsShrink(v string) *CreateRoutineCodeDeploymentShrinkRequest
	GetCodeVersionsShrink() *string
	SetEnv(v string) *CreateRoutineCodeDeploymentShrinkRequest
	GetEnv() *string
	SetName(v string) *CreateRoutineCodeDeploymentShrinkRequest
	GetName() *string
	SetStrategy(v string) *CreateRoutineCodeDeploymentShrinkRequest
	GetStrategy() *string
}

type CreateRoutineCodeDeploymentShrinkRequest struct {
	// The list of percentage-based canary release version configurations. A maximum of two versions are supported, and the total percentage must equal 100.
	//
	// This parameter is required.
	CodeVersionsShrink *string `json:"CodeVersions,omitempty" xml:"CodeVersions,omitempty"`
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
	// staging
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The name of the Edge Function Routine.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The deployment strategy.
	//
	// Valid values:
	//
	// - `percentage`: percentage mode
	//
	// This parameter is required.
	//
	// example:
	//
	// percentage
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s CreateRoutineCodeDeploymentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineCodeDeploymentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) GetCodeVersionsShrink() *string {
	return s.CodeVersionsShrink
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) SetCodeVersionsShrink(v string) *CreateRoutineCodeDeploymentShrinkRequest {
	s.CodeVersionsShrink = &v
	return s
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) SetEnv(v string) *CreateRoutineCodeDeploymentShrinkRequest {
	s.Env = &v
	return s
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) SetName(v string) *CreateRoutineCodeDeploymentShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) SetStrategy(v string) *CreateRoutineCodeDeploymentShrinkRequest {
	s.Strategy = &v
	return s
}

func (s *CreateRoutineCodeDeploymentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineCodeDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersions(v []*CreateRoutineCodeDeploymentRequestCodeVersions) *CreateRoutineCodeDeploymentRequest
	GetCodeVersions() []*CreateRoutineCodeDeploymentRequestCodeVersions
	SetEnv(v string) *CreateRoutineCodeDeploymentRequest
	GetEnv() *string
	SetName(v string) *CreateRoutineCodeDeploymentRequest
	GetName() *string
	SetStrategy(v string) *CreateRoutineCodeDeploymentRequest
	GetStrategy() *string
}

type CreateRoutineCodeDeploymentRequest struct {
	// The list of percentage-based canary release version configurations. A maximum of two versions are supported, and the total percentage must equal 100.
	//
	// This parameter is required.
	CodeVersions []*CreateRoutineCodeDeploymentRequestCodeVersions `json:"CodeVersions,omitempty" xml:"CodeVersions,omitempty" type:"Repeated"`
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

func (s CreateRoutineCodeDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineCodeDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineCodeDeploymentRequest) GetCodeVersions() []*CreateRoutineCodeDeploymentRequestCodeVersions {
	return s.CodeVersions
}

func (s *CreateRoutineCodeDeploymentRequest) GetEnv() *string {
	return s.Env
}

func (s *CreateRoutineCodeDeploymentRequest) GetName() *string {
	return s.Name
}

func (s *CreateRoutineCodeDeploymentRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateRoutineCodeDeploymentRequest) SetCodeVersions(v []*CreateRoutineCodeDeploymentRequestCodeVersions) *CreateRoutineCodeDeploymentRequest {
	s.CodeVersions = v
	return s
}

func (s *CreateRoutineCodeDeploymentRequest) SetEnv(v string) *CreateRoutineCodeDeploymentRequest {
	s.Env = &v
	return s
}

func (s *CreateRoutineCodeDeploymentRequest) SetName(v string) *CreateRoutineCodeDeploymentRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineCodeDeploymentRequest) SetStrategy(v string) *CreateRoutineCodeDeploymentRequest {
	s.Strategy = &v
	return s
}

func (s *CreateRoutineCodeDeploymentRequest) Validate() error {
	if s.CodeVersions != nil {
		for _, item := range s.CodeVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRoutineCodeDeploymentRequestCodeVersions struct {
	// The code version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1723599747213377175
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The canary release percentage for the code version. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Percentage *int64 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s CreateRoutineCodeDeploymentRequestCodeVersions) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineCodeDeploymentRequestCodeVersions) GoString() string {
	return s.String()
}

func (s *CreateRoutineCodeDeploymentRequestCodeVersions) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *CreateRoutineCodeDeploymentRequestCodeVersions) GetPercentage() *int64 {
	return s.Percentage
}

func (s *CreateRoutineCodeDeploymentRequestCodeVersions) SetCodeVersion(v string) *CreateRoutineCodeDeploymentRequestCodeVersions {
	s.CodeVersion = &v
	return s
}

func (s *CreateRoutineCodeDeploymentRequestCodeVersions) SetPercentage(v int64) *CreateRoutineCodeDeploymentRequestCodeVersions {
	s.Percentage = &v
	return s
}

func (s *CreateRoutineCodeDeploymentRequestCodeVersions) Validate() error {
	return dara.Validate(s)
}

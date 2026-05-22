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
	// The configuration list of phased release version numbers. A maximum of two versions are supported, and the sum of the total proportions is equal to 100.
	//
	// This parameter is required.
	CodeVersions []*CreateRoutineCodeDeploymentRequestCodeVersions `json:"CodeVersions,omitempty" xml:"CodeVersions,omitempty" type:"Repeated"`
	// The name of the environment. Only supports test environment `staging` or production environment `production`.
	//
	// This parameter is required.
	//
	// example:
	//
	// staging
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The function name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The deployment policy. Valid value: percentage.
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
	// The version of the code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1723599747213377175
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The phased release ratio of the code version. Valid values: 1 to 100.
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

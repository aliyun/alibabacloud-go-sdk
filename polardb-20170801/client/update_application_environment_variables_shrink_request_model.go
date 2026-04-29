// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationEnvironmentVariablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationEnvironmentVariablesShrinkRequest
	GetApplicationId() *string
	SetRestart(v bool) *UpdateApplicationEnvironmentVariablesShrinkRequest
	GetRestart() *bool
	SetVariablesShrink(v string) *UpdateApplicationEnvironmentVariablesShrinkRequest
	GetVariablesShrink() *string
}

type UpdateApplicationEnvironmentVariablesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// example:
	//
	// {
	//
	//     "ENV_TEST": "test-value"
	//
	// }
	VariablesShrink *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s UpdateApplicationEnvironmentVariablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationEnvironmentVariablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationEnvironmentVariablesShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationEnvironmentVariablesShrinkRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdateApplicationEnvironmentVariablesShrinkRequest) GetVariablesShrink() *string {
	return s.VariablesShrink
}

func (s *UpdateApplicationEnvironmentVariablesShrinkRequest) SetApplicationId(v string) *UpdateApplicationEnvironmentVariablesShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesShrinkRequest) SetRestart(v bool) *UpdateApplicationEnvironmentVariablesShrinkRequest {
	s.Restart = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesShrinkRequest) SetVariablesShrink(v string) *UpdateApplicationEnvironmentVariablesShrinkRequest {
	s.VariablesShrink = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

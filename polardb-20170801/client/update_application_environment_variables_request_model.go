// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationEnvironmentVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationEnvironmentVariablesRequest
	GetApplicationId() *string
	SetRestart(v bool) *UpdateApplicationEnvironmentVariablesRequest
	GetRestart() *bool
	SetVariables(v map[string]interface{}) *UpdateApplicationEnvironmentVariablesRequest
	GetVariables() map[string]interface{}
}

type UpdateApplicationEnvironmentVariablesRequest struct {
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
	Variables map[string]interface{} `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s UpdateApplicationEnvironmentVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationEnvironmentVariablesRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationEnvironmentVariablesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationEnvironmentVariablesRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdateApplicationEnvironmentVariablesRequest) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *UpdateApplicationEnvironmentVariablesRequest) SetApplicationId(v string) *UpdateApplicationEnvironmentVariablesRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesRequest) SetRestart(v bool) *UpdateApplicationEnvironmentVariablesRequest {
	s.Restart = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesRequest) SetVariables(v map[string]interface{}) *UpdateApplicationEnvironmentVariablesRequest {
	s.Variables = v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesRequest) Validate() error {
	return dara.Validate(s)
}

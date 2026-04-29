// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationEnvironmentVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemoveApplicationEnvironmentVariablesRequest
	GetApplicationId() *string
	SetRestart(v bool) *RemoveApplicationEnvironmentVariablesRequest
	GetRestart() *bool
	SetVariableNames(v []*string) *RemoveApplicationEnvironmentVariablesRequest
	GetVariableNames() []*string
}

type RemoveApplicationEnvironmentVariablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-********************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// This parameter is required.
	VariableNames []*string `json:"VariableNames,omitempty" xml:"VariableNames,omitempty" type:"Repeated"`
}

func (s RemoveApplicationEnvironmentVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationEnvironmentVariablesRequest) GoString() string {
	return s.String()
}

func (s *RemoveApplicationEnvironmentVariablesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemoveApplicationEnvironmentVariablesRequest) GetRestart() *bool {
	return s.Restart
}

func (s *RemoveApplicationEnvironmentVariablesRequest) GetVariableNames() []*string {
	return s.VariableNames
}

func (s *RemoveApplicationEnvironmentVariablesRequest) SetApplicationId(v string) *RemoveApplicationEnvironmentVariablesRequest {
	s.ApplicationId = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesRequest) SetRestart(v bool) *RemoveApplicationEnvironmentVariablesRequest {
	s.Restart = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesRequest) SetVariableNames(v []*string) *RemoveApplicationEnvironmentVariablesRequest {
	s.VariableNames = v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationEnvironmentVariablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemoveApplicationEnvironmentVariablesShrinkRequest
	GetApplicationId() *string
	SetRestart(v bool) *RemoveApplicationEnvironmentVariablesShrinkRequest
	GetRestart() *bool
	SetVariableNamesShrink(v string) *RemoveApplicationEnvironmentVariablesShrinkRequest
	GetVariableNamesShrink() *string
}

type RemoveApplicationEnvironmentVariablesShrinkRequest struct {
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
	VariableNamesShrink *string `json:"VariableNames,omitempty" xml:"VariableNames,omitempty"`
}

func (s RemoveApplicationEnvironmentVariablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationEnvironmentVariablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveApplicationEnvironmentVariablesShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemoveApplicationEnvironmentVariablesShrinkRequest) GetRestart() *bool {
	return s.Restart
}

func (s *RemoveApplicationEnvironmentVariablesShrinkRequest) GetVariableNamesShrink() *string {
	return s.VariableNamesShrink
}

func (s *RemoveApplicationEnvironmentVariablesShrinkRequest) SetApplicationId(v string) *RemoveApplicationEnvironmentVariablesShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesShrinkRequest) SetRestart(v bool) *RemoveApplicationEnvironmentVariablesShrinkRequest {
	s.Restart = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesShrinkRequest) SetVariableNamesShrink(v string) *RemoveApplicationEnvironmentVariablesShrinkRequest {
	s.VariableNamesShrink = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

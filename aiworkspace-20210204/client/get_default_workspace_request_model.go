// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetDefaultWorkspaceRequest
	GetVerbose() *bool
}

type GetDefaultWorkspaceRequest struct {
	// Specifies whether to show the details of the default workspace. The details include the conditions of the workspace in different phases. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetDefaultWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetDefaultWorkspaceRequest) SetVerbose(v bool) *GetDefaultWorkspaceRequest {
	s.Verbose = &v
	return s
}

func (s *GetDefaultWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}

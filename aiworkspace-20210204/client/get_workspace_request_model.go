// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetWorkspaceRequest
	GetVerbose() *bool
}

type GetWorkspaceRequest struct {
	// Specifies whether to return additional information, such as the workspace owner. Valid values:
	//
	// - false (default): Does not return additional information.
	//
	// - true: Returns additional information.
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetWorkspaceRequest) SetVerbose(v bool) *GetWorkspaceRequest {
	s.Verbose = &v
	return s
}

func (s *GetWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}

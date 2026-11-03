// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtraVerboseFields(v string) *GetWorkspaceRequest
	GetExtraVerboseFields() *string
	SetVerbose(v bool) *GetWorkspaceRequest
	GetVerbose() *bool
}

type GetWorkspaceRequest struct {
	// A comma-separated string used to specify additional information to return. This field does not take effect.
	//
	// example:
	//
	// NotSupport
	ExtraVerboseFields *string `json:"ExtraVerboseFields,omitempty" xml:"ExtraVerboseFields,omitempty"`
	// Specifies whether to display supplementary information such as the workspace owner. Valid values:
	//
	// - false: Does not display supplementary information.
	//
	// - true: Displays supplementary information.
	//
	// Default value: false.
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

func (s *GetWorkspaceRequest) GetExtraVerboseFields() *string {
	return s.ExtraVerboseFields
}

func (s *GetWorkspaceRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetWorkspaceRequest) SetExtraVerboseFields(v string) *GetWorkspaceRequest {
	s.ExtraVerboseFields = &v
	return s
}

func (s *GetWorkspaceRequest) SetVerbose(v bool) *GetWorkspaceRequest {
	s.Verbose = &v
	return s
}

func (s *GetWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}

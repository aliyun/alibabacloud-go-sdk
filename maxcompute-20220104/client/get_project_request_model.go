// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetProjectRequest
	GetVerbose() *bool
}

type GetProjectRequest struct {
	// Specifies whether to use additional information.
	//
	// example:
	//
	// true
	Verbose *bool `json:"verbose,omitempty" xml:"verbose,omitempty"`
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetProjectRequest) SetVerbose(v bool) *GetProjectRequest {
	s.Verbose = &v
	return s
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}

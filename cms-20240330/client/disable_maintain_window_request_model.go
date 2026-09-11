// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMaintainWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspace(v string) *DisableMaintainWindowRequest
	GetWorkspace() *string
}

type DisableMaintainWindowRequest struct {
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DisableMaintainWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableMaintainWindowRequest) GoString() string {
	return s.String()
}

func (s *DisableMaintainWindowRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *DisableMaintainWindowRequest) SetWorkspace(v string) *DisableMaintainWindowRequest {
	s.Workspace = &v
	return s
}

func (s *DisableMaintainWindowRequest) Validate() error {
	return dara.Validate(s)
}

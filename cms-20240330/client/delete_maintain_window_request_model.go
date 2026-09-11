// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaintainWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspace(v string) *DeleteMaintainWindowRequest
	GetWorkspace() *string
}

type DeleteMaintainWindowRequest struct {
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteMaintainWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaintainWindowRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaintainWindowRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteMaintainWindowRequest) SetWorkspace(v string) *DeleteMaintainWindowRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteMaintainWindowRequest) Validate() error {
	return dara.Validate(s)
}

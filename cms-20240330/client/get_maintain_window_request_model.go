// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaintainWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspace(v string) *GetMaintainWindowRequest
	GetWorkspace() *string
}

type GetMaintainWindowRequest struct {
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetMaintainWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMaintainWindowRequest) GoString() string {
	return s.String()
}

func (s *GetMaintainWindowRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetMaintainWindowRequest) SetWorkspace(v string) *GetMaintainWindowRequest {
	s.Workspace = &v
	return s
}

func (s *GetMaintainWindowRequest) Validate() error {
	return dara.Validate(s)
}

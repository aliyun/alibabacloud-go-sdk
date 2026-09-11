// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaintainWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MaintainWindowForModify) *UpdateMaintainWindowRequest
	GetBody() *MaintainWindowForModify
	SetWorkspace(v string) *UpdateMaintainWindowRequest
	GetWorkspace() *string
}

type UpdateMaintainWindowRequest struct {
	Body *MaintainWindowForModify `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateMaintainWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaintainWindowRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaintainWindowRequest) GetBody() *MaintainWindowForModify {
	return s.Body
}

func (s *UpdateMaintainWindowRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateMaintainWindowRequest) SetBody(v *MaintainWindowForModify) *UpdateMaintainWindowRequest {
	s.Body = v
	return s
}

func (s *UpdateMaintainWindowRequest) SetWorkspace(v string) *UpdateMaintainWindowRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateMaintainWindowRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

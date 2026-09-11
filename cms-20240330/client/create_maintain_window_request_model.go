// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaintainWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MaintainWindowForModify) *CreateMaintainWindowRequest
	GetBody() *MaintainWindowForModify
	SetWorkspace(v string) *CreateMaintainWindowRequest
	GetWorkspace() *string
}

type CreateMaintainWindowRequest struct {
	Body      *MaintainWindowForModify `json:"body,omitempty" xml:"body,omitempty"`
	Workspace *string                  `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateMaintainWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMaintainWindowRequest) GoString() string {
	return s.String()
}

func (s *CreateMaintainWindowRequest) GetBody() *MaintainWindowForModify {
	return s.Body
}

func (s *CreateMaintainWindowRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateMaintainWindowRequest) SetBody(v *MaintainWindowForModify) *CreateMaintainWindowRequest {
	s.Body = v
	return s
}

func (s *CreateMaintainWindowRequest) SetWorkspace(v string) *CreateMaintainWindowRequest {
	s.Workspace = &v
	return s
}

func (s *CreateMaintainWindowRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

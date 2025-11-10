// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNotifyStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *NotifyStrategyForModify) *UpdateNotifyStrategyRequest
	GetBody() *NotifyStrategyForModify
	SetWorkspace(v string) *UpdateNotifyStrategyRequest
	GetWorkspace() *string
}

type UpdateNotifyStrategyRequest struct {
	Body *NotifyStrategyForModify `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateNotifyStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNotifyStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateNotifyStrategyRequest) GetBody() *NotifyStrategyForModify {
	return s.Body
}

func (s *UpdateNotifyStrategyRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateNotifyStrategyRequest) SetBody(v *NotifyStrategyForModify) *UpdateNotifyStrategyRequest {
	s.Body = v
	return s
}

func (s *UpdateNotifyStrategyRequest) SetWorkspace(v string) *UpdateNotifyStrategyRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateNotifyStrategyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

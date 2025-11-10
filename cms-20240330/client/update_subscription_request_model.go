// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SubscriptionForModify) *UpdateSubscriptionRequest
	GetBody() *SubscriptionForModify
	SetWorkspace(v string) *UpdateSubscriptionRequest
	GetWorkspace() *string
}

type UpdateSubscriptionRequest struct {
	Body *SubscriptionForModify `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequest) GetBody() *SubscriptionForModify {
	return s.Body
}

func (s *UpdateSubscriptionRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateSubscriptionRequest) SetBody(v *SubscriptionForModify) *UpdateSubscriptionRequest {
	s.Body = v
	return s
}

func (s *UpdateSubscriptionRequest) SetWorkspace(v string) *UpdateSubscriptionRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateSubscriptionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

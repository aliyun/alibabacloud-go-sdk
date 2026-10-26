// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SubscriptionForModify) *CreateSubscriptionRequest
	GetBody() *SubscriptionForModify
	SetWorkspace(v string) *CreateSubscriptionRequest
	GetWorkspace() *string
}

type CreateSubscriptionRequest struct {
	Body *SubscriptionForModify `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequest) GetBody() *SubscriptionForModify {
	return s.Body
}

func (s *CreateSubscriptionRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateSubscriptionRequest) SetBody(v *SubscriptionForModify) *CreateSubscriptionRequest {
	s.Body = v
	return s
}

func (s *CreateSubscriptionRequest) SetWorkspace(v string) *CreateSubscriptionRequest {
	s.Workspace = &v
	return s
}

func (s *CreateSubscriptionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

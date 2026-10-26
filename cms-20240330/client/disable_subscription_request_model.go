// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspace(v string) *DisableSubscriptionRequest
	GetWorkspace() *string
}

type DisableSubscriptionRequest struct {
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DisableSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *DisableSubscriptionRequest) SetWorkspace(v string) *DisableSubscriptionRequest {
	s.Workspace = &v
	return s
}

func (s *DisableSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}

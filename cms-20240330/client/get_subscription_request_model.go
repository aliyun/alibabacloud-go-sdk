// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspace(v string) *GetSubscriptionRequest
	GetWorkspace() *string
}

type GetSubscriptionRequest struct {
	// The name of the workspace.
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetSubscriptionRequest) SetWorkspace(v string) *GetSubscriptionRequest {
	s.Workspace = &v
	return s
}

func (s *GetSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}

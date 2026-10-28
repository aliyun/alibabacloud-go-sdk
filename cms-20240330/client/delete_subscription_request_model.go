// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspace(v string) *DeleteSubscriptionRequest
	GetWorkspace() *string
}

type DeleteSubscriptionRequest struct {
	// The name of the workspace.
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteSubscriptionRequest) SetWorkspace(v string) *DeleteSubscriptionRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateWorkspaceShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateWorkspaceShrinkRequest
	GetClientToken() *string
}

type UpdateWorkspaceShrinkRequest struct {
	// The request body for updating a workspace.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client idempotency token.
	//
	// example:
	//
	// workspace-update-20260805-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateWorkspaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateWorkspaceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateWorkspaceShrinkRequest) SetBodyShrink(v string) *UpdateWorkspaceShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateWorkspaceShrinkRequest) SetClientToken(v string) *UpdateWorkspaceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWorkspaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateWorkspaceShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateWorkspaceShrinkRequest
	GetClientToken() *string
}

type CreateWorkspaceShrinkRequest struct {
	// The request body for creating a workspace.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client idempotency token.
	//
	// example:
	//
	// workspace-create-20260805-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateWorkspaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateWorkspaceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWorkspaceShrinkRequest) SetBodyShrink(v string) *CreateWorkspaceShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateWorkspaceShrinkRequest) SetClientToken(v string) *CreateWorkspaceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkspaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateAgentSpecShrinkRequest
	GetBodyShrink() *string
}

type UpdateAgentSpecShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentSpecShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateAgentSpecShrinkRequest) SetBodyShrink(v string) *UpdateAgentSpecShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateAgentSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}

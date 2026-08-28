// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpecVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateAgentSpecVersionShrinkRequest
	GetBodyShrink() *string
}

type CreateAgentSpecVersionShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentSpecVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecVersionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateAgentSpecVersionShrinkRequest) SetBodyShrink(v string) *CreateAgentSpecVersionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateAgentSpecVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}

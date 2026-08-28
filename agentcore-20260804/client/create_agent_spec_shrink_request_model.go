// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateAgentSpecShrinkRequest
	GetBodyShrink() *string
}

type CreateAgentSpecShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateAgentSpecShrinkRequest) SetBodyShrink(v string) *CreateAgentSpecShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateAgentSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}

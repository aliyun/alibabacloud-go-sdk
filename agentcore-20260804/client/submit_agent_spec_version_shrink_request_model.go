// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAgentSpecVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *SubmitAgentSpecVersionShrinkRequest
	GetBodyShrink() *string
}

type SubmitAgentSpecVersionShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAgentSpecVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAgentSpecVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitAgentSpecVersionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *SubmitAgentSpecVersionShrinkRequest) SetBodyShrink(v string) *SubmitAgentSpecVersionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *SubmitAgentSpecVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}

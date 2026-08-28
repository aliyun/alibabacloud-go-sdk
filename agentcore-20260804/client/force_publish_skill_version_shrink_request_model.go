// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForcePublishSkillVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *ForcePublishSkillVersionShrinkRequest
	GetBodyShrink() *string
}

type ForcePublishSkillVersionShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForcePublishSkillVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ForcePublishSkillVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ForcePublishSkillVersionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ForcePublishSkillVersionShrinkRequest) SetBodyShrink(v string) *ForcePublishSkillVersionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ForcePublishSkillVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishSkillVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *PublishSkillVersionShrinkRequest
	GetBodyShrink() *string
}

type PublishSkillVersionShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishSkillVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishSkillVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *PublishSkillVersionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *PublishSkillVersionShrinkRequest) SetBodyShrink(v string) *PublishSkillVersionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *PublishSkillVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}

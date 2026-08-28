// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedraftSkillVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *RedraftSkillVersionShrinkRequest
	GetBodyShrink() *string
}

type RedraftSkillVersionShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedraftSkillVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RedraftSkillVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *RedraftSkillVersionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *RedraftSkillVersionShrinkRequest) SetBodyShrink(v string) *RedraftSkillVersionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *RedraftSkillVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillLabelsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateSkillLabelsShrinkRequest
	GetBodyShrink() *string
}

type UpdateSkillLabelsShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillLabelsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillLabelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillLabelsShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateSkillLabelsShrinkRequest) SetBodyShrink(v string) *UpdateSkillLabelsShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateSkillLabelsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

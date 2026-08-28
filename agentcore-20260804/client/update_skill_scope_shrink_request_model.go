// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillScopeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateSkillScopeShrinkRequest
	GetBodyShrink() *string
}

type UpdateSkillScopeShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillScopeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillScopeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillScopeShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateSkillScopeShrinkRequest) SetBodyShrink(v string) *UpdateSkillScopeShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateSkillScopeShrinkRequest) Validate() error {
	return dara.Validate(s)
}

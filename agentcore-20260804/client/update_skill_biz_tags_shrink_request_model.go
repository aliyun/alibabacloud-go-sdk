// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillBizTagsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateSkillBizTagsShrinkRequest
	GetBodyShrink() *string
}

type UpdateSkillBizTagsShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillBizTagsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillBizTagsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillBizTagsShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateSkillBizTagsShrinkRequest) SetBodyShrink(v string) *UpdateSkillBizTagsShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateSkillBizTagsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSkillVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *SubmitSkillVersionShrinkRequest
	GetBodyShrink() *string
}

type SubmitSkillVersionShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSkillVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSkillVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSkillVersionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *SubmitSkillVersionShrinkRequest) SetBodyShrink(v string) *SubmitSkillVersionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *SubmitSkillVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}

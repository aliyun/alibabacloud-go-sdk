// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillDraftShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateSkillDraftShrinkRequest
	GetBodyShrink() *string
}

type CreateSkillDraftShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillDraftShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillDraftShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillDraftShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateSkillDraftShrinkRequest) SetBodyShrink(v string) *CreateSkillDraftShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateSkillDraftShrinkRequest) Validate() error {
	return dara.Validate(s)
}

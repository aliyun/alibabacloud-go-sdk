// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *OnlineSkillShrinkRequest
	GetBodyShrink() *string
}

type OnlineSkillShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *OnlineSkillShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *OnlineSkillShrinkRequest) SetBodyShrink(v string) *OnlineSkillShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *OnlineSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}

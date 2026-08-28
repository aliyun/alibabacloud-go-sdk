// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *OfflineSkillShrinkRequest
	GetBodyShrink() *string
}

type OfflineSkillShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *OfflineSkillShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *OfflineSkillShrinkRequest) SetBodyShrink(v string) *OfflineSkillShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *OfflineSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}

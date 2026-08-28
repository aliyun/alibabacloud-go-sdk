// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSkillViaOssShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UploadSkillViaOssShrinkRequest
	GetBodyShrink() *string
}

type UploadSkillViaOssShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadSkillViaOssShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadSkillViaOssShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadSkillViaOssShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UploadSkillViaOssShrinkRequest) SetBodyShrink(v string) *UploadSkillViaOssShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UploadSkillViaOssShrinkRequest) Validate() error {
	return dara.Validate(s)
}

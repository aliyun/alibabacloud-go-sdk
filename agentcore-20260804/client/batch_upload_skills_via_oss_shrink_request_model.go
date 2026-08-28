// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUploadSkillsViaOssShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *BatchUploadSkillsViaOssShrinkRequest
	GetBodyShrink() *string
}

type BatchUploadSkillsViaOssShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUploadSkillsViaOssShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUploadSkillsViaOssShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadSkillsViaOssShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *BatchUploadSkillsViaOssShrinkRequest) SetBodyShrink(v string) *BatchUploadSkillsViaOssShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *BatchUploadSkillsViaOssShrinkRequest) Validate() error {
	return dara.Validate(s)
}

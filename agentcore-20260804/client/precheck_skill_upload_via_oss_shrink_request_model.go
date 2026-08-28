// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckSkillUploadViaOssShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *PrecheckSkillUploadViaOssShrinkRequest
	GetBodyShrink() *string
}

type PrecheckSkillUploadViaOssShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrecheckSkillUploadViaOssShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PrecheckSkillUploadViaOssShrinkRequest) GoString() string {
	return s.String()
}

func (s *PrecheckSkillUploadViaOssShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *PrecheckSkillUploadViaOssShrinkRequest) SetBodyShrink(v string) *PrecheckSkillUploadViaOssShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *PrecheckSkillUploadViaOssShrinkRequest) Validate() error {
	return dara.Validate(s)
}

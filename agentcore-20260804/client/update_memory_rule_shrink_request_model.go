// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateMemoryRuleShrinkRequest
	GetBodyShrink() *string
}

type UpdateMemoryRuleShrinkRequest struct {
	// The JSON request body.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRuleShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateMemoryRuleShrinkRequest) SetBodyShrink(v string) *UpdateMemoryRuleShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateMemoryRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}

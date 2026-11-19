// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateMemoryRuleShrinkRequest
	GetBodyShrink() *string
}

type CreateMemoryRuleShrinkRequest struct {
	// The JSON request body.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryRuleShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateMemoryRuleShrinkRequest) SetBodyShrink(v string) *CreateMemoryRuleShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateMemoryRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}

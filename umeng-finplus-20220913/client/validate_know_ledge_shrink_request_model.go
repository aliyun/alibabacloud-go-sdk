// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateKnowLedgeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *ValidateKnowLedgeShrinkRequest
	GetBodyShrink() *string
}

type ValidateKnowLedgeShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateKnowLedgeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateKnowLedgeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ValidateKnowLedgeShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ValidateKnowLedgeShrinkRequest) SetBodyShrink(v string) *ValidateKnowLedgeShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ValidateKnowLedgeShrinkRequest) Validate() error {
	return dara.Validate(s)
}

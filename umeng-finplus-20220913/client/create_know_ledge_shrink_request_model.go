// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowLedgeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateKnowLedgeShrinkRequest
	GetBodyShrink() *string
}

type CreateKnowLedgeShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowLedgeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowLedgeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowLedgeShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateKnowLedgeShrinkRequest) SetBodyShrink(v string) *CreateKnowLedgeShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateKnowLedgeShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *ResetUserPasswordShrinkRequest
	GetBodyShrink() *string
}

type ResetUserPasswordShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetUserPasswordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ResetUserPasswordShrinkRequest) SetBodyShrink(v string) *ResetUserPasswordShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ResetUserPasswordShrinkRequest) Validate() error {
	return dara.Validate(s)
}

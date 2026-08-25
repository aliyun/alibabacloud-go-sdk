// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *DebugModelShrinkRequest
	GetBodyShrink() *string
}

type DebugModelShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *DebugModelShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *DebugModelShrinkRequest) SetBodyShrink(v string) *DebugModelShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *DebugModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}

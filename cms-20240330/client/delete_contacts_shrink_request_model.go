// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIdsShrink(v string) *DeleteContactsShrinkRequest
	GetContactIdsShrink() *string
}

type DeleteContactsShrinkRequest struct {
	ContactIdsShrink *string `json:"contactIds,omitempty" xml:"contactIds,omitempty"`
}

func (s DeleteContactsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactsShrinkRequest) GetContactIdsShrink() *string {
	return s.ContactIdsShrink
}

func (s *DeleteContactsShrinkRequest) SetContactIdsShrink(v string) *DeleteContactsShrinkRequest {
	s.ContactIdsShrink = &v
	return s
}

func (s *DeleteContactsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

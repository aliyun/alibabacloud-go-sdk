// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIdsShrink(v string) *DeleteContactGroupsShrinkRequest
	GetContactGroupIdsShrink() *string
}

type DeleteContactGroupsShrinkRequest struct {
	// This parameter is required.
	ContactGroupIdsShrink *string `json:"contactGroupIds,omitempty" xml:"contactGroupIds,omitempty"`
}

func (s DeleteContactGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupsShrinkRequest) GetContactGroupIdsShrink() *string {
	return s.ContactGroupIdsShrink
}

func (s *DeleteContactGroupsShrinkRequest) SetContactGroupIdsShrink(v string) *DeleteContactGroupsShrinkRequest {
	s.ContactGroupIdsShrink = &v
	return s
}

func (s *DeleteContactGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHaVipsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHaVipIdsShrink(v string) *DeleteHaVipsShrinkRequest
	GetHaVipIdsShrink() *string
}

type DeleteHaVipsShrinkRequest struct {
	// This parameter is required.
	HaVipIdsShrink *string `json:"HaVipIds,omitempty" xml:"HaVipIds,omitempty"`
}

func (s DeleteHaVipsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHaVipsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteHaVipsShrinkRequest) GetHaVipIdsShrink() *string {
	return s.HaVipIdsShrink
}

func (s *DeleteHaVipsShrinkRequest) SetHaVipIdsShrink(v string) *DeleteHaVipsShrinkRequest {
	s.HaVipIdsShrink = &v
	return s
}

func (s *DeleteHaVipsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

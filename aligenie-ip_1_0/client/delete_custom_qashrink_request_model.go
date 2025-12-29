// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomQAShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomQAIdsShrink(v string) *DeleteCustomQAShrinkRequest
	GetCustomQAIdsShrink() *string
	SetHotelId(v string) *DeleteCustomQAShrinkRequest
	GetHotelId() *string
}

type DeleteCustomQAShrinkRequest struct {
	CustomQAIdsShrink *string `json:"CustomQAIds,omitempty" xml:"CustomQAIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteCustomQAShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomQAShrinkRequest) GetCustomQAIdsShrink() *string {
	return s.CustomQAIdsShrink
}

func (s *DeleteCustomQAShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *DeleteCustomQAShrinkRequest) SetCustomQAIdsShrink(v string) *DeleteCustomQAShrinkRequest {
	s.CustomQAIdsShrink = &v
	return s
}

func (s *DeleteCustomQAShrinkRequest) SetHotelId(v string) *DeleteCustomQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteCustomQAShrinkRequest) Validate() error {
	return dara.Validate(s)
}

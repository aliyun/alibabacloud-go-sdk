// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCartoonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *DeleteCartoonRequest
	GetHotelId() *string
}

type DeleteCartoonRequest struct {
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteCartoonRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCartoonRequest) GoString() string {
	return s.String()
}

func (s *DeleteCartoonRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *DeleteCartoonRequest) SetHotelId(v string) *DeleteCartoonRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteCartoonRequest) Validate() error {
	return dara.Validate(s)
}

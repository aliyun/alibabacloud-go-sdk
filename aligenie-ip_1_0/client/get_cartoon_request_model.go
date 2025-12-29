// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCartoonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *GetCartoonRequest
	GetHotelId() *string
}

type GetCartoonRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s GetCartoonRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCartoonRequest) GoString() string {
	return s.String()
}

func (s *GetCartoonRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *GetCartoonRequest) SetHotelId(v string) *GetCartoonRequest {
	s.HotelId = &v
	return s
}

func (s *GetCartoonRequest) Validate() error {
	return dara.Validate(s)
}

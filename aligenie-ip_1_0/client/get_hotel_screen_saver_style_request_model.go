// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverStyleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *GetHotelScreenSaverStyleRequest
	GetHotelId() *string
}

type GetHotelScreenSaverStyleRequest struct {
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s GetHotelScreenSaverStyleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverStyleRequest) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelScreenSaverStyleRequest) SetHotelId(v string) *GetHotelScreenSaverStyleRequest {
	s.HotelId = &v
	return s
}

func (s *GetHotelScreenSaverStyleRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWelcomeTextAndMusicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *GetWelcomeTextAndMusicRequest
	GetHotelId() *string
}

type GetWelcomeTextAndMusicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s GetWelcomeTextAndMusicRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWelcomeTextAndMusicRequest) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *GetWelcomeTextAndMusicRequest) SetHotelId(v string) *GetWelcomeTextAndMusicRequest {
	s.HotelId = &v
	return s
}

func (s *GetWelcomeTextAndMusicRequest) Validate() error {
	return dara.Validate(s)
}

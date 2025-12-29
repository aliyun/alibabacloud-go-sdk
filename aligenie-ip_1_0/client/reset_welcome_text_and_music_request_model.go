// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetWelcomeTextAndMusicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ResetWelcomeTextAndMusicRequest
	GetHotelId() *string
}

type ResetWelcomeTextAndMusicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ResetWelcomeTextAndMusicRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetWelcomeTextAndMusicRequest) GoString() string {
	return s.String()
}

func (s *ResetWelcomeTextAndMusicRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ResetWelcomeTextAndMusicRequest) SetHotelId(v string) *ResetWelcomeTextAndMusicRequest {
	s.HotelId = &v
	return s
}

func (s *ResetWelcomeTextAndMusicRequest) Validate() error {
	return dara.Validate(s)
}

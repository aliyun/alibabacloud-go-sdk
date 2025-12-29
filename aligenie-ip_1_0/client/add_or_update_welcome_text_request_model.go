// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateWelcomeTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *AddOrUpdateWelcomeTextRequest
	GetHotelId() *string
	SetMusicUrl(v string) *AddOrUpdateWelcomeTextRequest
	GetMusicUrl() *string
	SetWelcomeText(v string) *AddOrUpdateWelcomeTextRequest
	GetWelcomeText() *string
}

type AddOrUpdateWelcomeTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://ailabsaicloudservice.alicdn.com/tmp/a.wav
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	// This parameter is required.
	WelcomeText *string `json:"WelcomeText,omitempty" xml:"WelcomeText,omitempty"`
}

func (s AddOrUpdateWelcomeTextRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateWelcomeTextRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateWelcomeTextRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddOrUpdateWelcomeTextRequest) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *AddOrUpdateWelcomeTextRequest) GetWelcomeText() *string {
	return s.WelcomeText
}

func (s *AddOrUpdateWelcomeTextRequest) SetHotelId(v string) *AddOrUpdateWelcomeTextRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateWelcomeTextRequest) SetMusicUrl(v string) *AddOrUpdateWelcomeTextRequest {
	s.MusicUrl = &v
	return s
}

func (s *AddOrUpdateWelcomeTextRequest) SetWelcomeText(v string) *AddOrUpdateWelcomeTextRequest {
	s.WelcomeText = &v
	return s
}

func (s *AddOrUpdateWelcomeTextRequest) Validate() error {
	return dara.Validate(s)
}

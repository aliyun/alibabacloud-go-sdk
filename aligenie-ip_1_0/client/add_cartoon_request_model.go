// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCartoonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *AddCartoonRequest
	GetHotelId() *string
	SetStartVideoMd5(v string) *AddCartoonRequest
	GetStartVideoMd5() *string
	SetStartVideoUrl(v string) *AddCartoonRequest
	GetStartVideoUrl() *string
}

type AddCartoonRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0***eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40c804***697
	StartVideoMd5 *string `json:"StartVideoMd5,omitempty" xml:"StartVideoMd5,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://***.mp4
	StartVideoUrl *string `json:"StartVideoUrl,omitempty" xml:"StartVideoUrl,omitempty"`
}

func (s AddCartoonRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCartoonRequest) GoString() string {
	return s.String()
}

func (s *AddCartoonRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddCartoonRequest) GetStartVideoMd5() *string {
	return s.StartVideoMd5
}

func (s *AddCartoonRequest) GetStartVideoUrl() *string {
	return s.StartVideoUrl
}

func (s *AddCartoonRequest) SetHotelId(v string) *AddCartoonRequest {
	s.HotelId = &v
	return s
}

func (s *AddCartoonRequest) SetStartVideoMd5(v string) *AddCartoonRequest {
	s.StartVideoMd5 = &v
	return s
}

func (s *AddCartoonRequest) SetStartVideoUrl(v string) *AddCartoonRequest {
	s.StartVideoUrl = &v
	return s
}

func (s *AddCartoonRequest) Validate() error {
	return dara.Validate(s)
}

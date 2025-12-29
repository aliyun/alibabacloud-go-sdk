// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveHotelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *RemoveHotelRequest
	GetAppKey() *string
	SetHotelId(v string) *RemoveHotelRequest
	GetHotelId() *string
	SetTbOpenId(v string) *RemoveHotelRequest
	GetTbOpenId() *string
}

type RemoveHotelRequest struct {
	// appkey
	//
	// This parameter is required.
	//
	// example:
	//
	// 30193305
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEVK***UE3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s RemoveHotelRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveHotelRequest) GoString() string {
	return s.String()
}

func (s *RemoveHotelRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *RemoveHotelRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *RemoveHotelRequest) GetTbOpenId() *string {
	return s.TbOpenId
}

func (s *RemoveHotelRequest) SetAppKey(v string) *RemoveHotelRequest {
	s.AppKey = &v
	return s
}

func (s *RemoveHotelRequest) SetHotelId(v string) *RemoveHotelRequest {
	s.HotelId = &v
	return s
}

func (s *RemoveHotelRequest) SetTbOpenId(v string) *RemoveHotelRequest {
	s.TbOpenId = &v
	return s
}

func (s *RemoveHotelRequest) Validate() error {
	return dara.Validate(s)
}

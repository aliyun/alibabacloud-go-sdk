// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicInfoQARequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *GetBasicInfoQARequest
	GetHotelId() *string
}

type GetBasicInfoQARequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s GetBasicInfoQARequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicInfoQARequest) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQARequest) GetHotelId() *string {
	return s.HotelId
}

func (s *GetBasicInfoQARequest) SetHotelId(v string) *GetBasicInfoQARequest {
	s.HotelId = &v
	return s
}

func (s *GetBasicInfoQARequest) Validate() error {
	return dara.Validate(s)
}

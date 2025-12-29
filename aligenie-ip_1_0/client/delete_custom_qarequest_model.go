// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomQARequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomQAIds(v []*string) *DeleteCustomQARequest
	GetCustomQAIds() []*string
	SetHotelId(v string) *DeleteCustomQARequest
	GetHotelId() *string
}

type DeleteCustomQARequest struct {
	CustomQAIds []*string `json:"CustomQAIds,omitempty" xml:"CustomQAIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteCustomQARequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomQARequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomQARequest) GetCustomQAIds() []*string {
	return s.CustomQAIds
}

func (s *DeleteCustomQARequest) GetHotelId() *string {
	return s.HotelId
}

func (s *DeleteCustomQARequest) SetCustomQAIds(v []*string) *DeleteCustomQARequest {
	s.CustomQAIds = v
	return s
}

func (s *DeleteCustomQARequest) SetHotelId(v string) *DeleteCustomQARequest {
	s.HotelId = &v
	return s
}

func (s *DeleteCustomQARequest) Validate() error {
	return dara.Validate(s)
}

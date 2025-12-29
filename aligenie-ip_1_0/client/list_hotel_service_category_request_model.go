// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelServiceCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayload(v *ListHotelServiceCategoryRequestPayload) *ListHotelServiceCategoryRequest
	GetPayload() *ListHotelServiceCategoryRequestPayload
}

type ListHotelServiceCategoryRequest struct {
	// This parameter is required.
	Payload *ListHotelServiceCategoryRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
}

func (s ListHotelServiceCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelServiceCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryRequest) GetPayload() *ListHotelServiceCategoryRequestPayload {
	return s.Payload
}

func (s *ListHotelServiceCategoryRequest) SetPayload(v *ListHotelServiceCategoryRequestPayload) *ListHotelServiceCategoryRequest {
	s.Payload = v
	return s
}

func (s *ListHotelServiceCategoryRequest) Validate() error {
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelServiceCategoryRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// HOTEL_SERVICE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelServiceCategoryRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s ListHotelServiceCategoryRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryRequestPayload) GetType() *string {
	return s.Type
}

func (s *ListHotelServiceCategoryRequestPayload) SetType(v string) *ListHotelServiceCategoryRequestPayload {
	s.Type = &v
	return s
}

func (s *ListHotelServiceCategoryRequestPayload) Validate() error {
	return dara.Validate(s)
}

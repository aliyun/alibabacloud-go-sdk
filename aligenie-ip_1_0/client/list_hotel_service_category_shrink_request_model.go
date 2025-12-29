// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelServiceCategoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayloadShrink(v string) *ListHotelServiceCategoryShrinkRequest
	GetPayloadShrink() *string
}

type ListHotelServiceCategoryShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s ListHotelServiceCategoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelServiceCategoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *ListHotelServiceCategoryShrinkRequest) SetPayloadShrink(v string) *ListHotelServiceCategoryShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListHotelServiceCategoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}

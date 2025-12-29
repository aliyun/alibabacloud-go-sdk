// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelOrderDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayloadShrink(v string) *GetHotelOrderDetailShrinkRequest
	GetPayloadShrink() *string
}

type GetHotelOrderDetailShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s GetHotelOrderDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelOrderDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *GetHotelOrderDetailShrinkRequest) SetPayloadShrink(v string) *GetHotelOrderDetailShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetHotelOrderDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}

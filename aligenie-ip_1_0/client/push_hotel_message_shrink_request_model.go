// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushHotelMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPushHotelMessageReqShrink(v string) *PushHotelMessageShrinkRequest
	GetPushHotelMessageReqShrink() *string
}

type PushHotelMessageShrinkRequest struct {
	// pushHotelMessageReq
	//
	// This parameter is required.
	PushHotelMessageReqShrink *string `json:"PushHotelMessageReq,omitempty" xml:"PushHotelMessageReq,omitempty"`
}

func (s PushHotelMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushHotelMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushHotelMessageShrinkRequest) GetPushHotelMessageReqShrink() *string {
	return s.PushHotelMessageReqShrink
}

func (s *PushHotelMessageShrinkRequest) SetPushHotelMessageReqShrink(v string) *PushHotelMessageShrinkRequest {
	s.PushHotelMessageReqShrink = &v
	return s
}

func (s *PushHotelMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}

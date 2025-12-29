// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushHotelMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPushHotelMessageReq(v *PushHotelMessageRequestPushHotelMessageReq) *PushHotelMessageRequest
	GetPushHotelMessageReq() *PushHotelMessageRequestPushHotelMessageReq
}

type PushHotelMessageRequest struct {
	// pushHotelMessageReq
	//
	// This parameter is required.
	PushHotelMessageReq *PushHotelMessageRequestPushHotelMessageReq `json:"PushHotelMessageReq,omitempty" xml:"PushHotelMessageReq,omitempty" type:"Struct"`
}

func (s PushHotelMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s PushHotelMessageRequest) GoString() string {
	return s.String()
}

func (s *PushHotelMessageRequest) GetPushHotelMessageReq() *PushHotelMessageRequestPushHotelMessageReq {
	return s.PushHotelMessageReq
}

func (s *PushHotelMessageRequest) SetPushHotelMessageReq(v *PushHotelMessageRequestPushHotelMessageReq) *PushHotelMessageRequest {
	s.PushHotelMessageReq = v
	return s
}

func (s *PushHotelMessageRequest) Validate() error {
	if s.PushHotelMessageReq != nil {
		if err := s.PushHotelMessageReq.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushHotelMessageRequestPushHotelMessageReq struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId  *string            `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	ParamMap map[string]*string `json:"ParamMap,omitempty" xml:"ParamMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s PushHotelMessageRequestPushHotelMessageReq) String() string {
	return dara.Prettify(s)
}

func (s PushHotelMessageRequestPushHotelMessageReq) GoString() string {
	return s.String()
}

func (s *PushHotelMessageRequestPushHotelMessageReq) GetHotelId() *string {
	return s.HotelId
}

func (s *PushHotelMessageRequestPushHotelMessageReq) GetParamMap() map[string]*string {
	return s.ParamMap
}

func (s *PushHotelMessageRequestPushHotelMessageReq) GetRoomNo() *string {
	return s.RoomNo
}

func (s *PushHotelMessageRequestPushHotelMessageReq) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *PushHotelMessageRequestPushHotelMessageReq) SetHotelId(v string) *PushHotelMessageRequestPushHotelMessageReq {
	s.HotelId = &v
	return s
}

func (s *PushHotelMessageRequestPushHotelMessageReq) SetParamMap(v map[string]*string) *PushHotelMessageRequestPushHotelMessageReq {
	s.ParamMap = v
	return s
}

func (s *PushHotelMessageRequestPushHotelMessageReq) SetRoomNo(v string) *PushHotelMessageRequestPushHotelMessageReq {
	s.RoomNo = &v
	return s
}

func (s *PushHotelMessageRequestPushHotelMessageReq) SetTemplateId(v int64) *PushHotelMessageRequestPushHotelMessageReq {
	s.TemplateId = &v
	return s
}

func (s *PushHotelMessageRequestPushHotelMessageReq) Validate() error {
	return dara.Validate(s)
}

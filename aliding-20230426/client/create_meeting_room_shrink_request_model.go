// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCycleReservation(v bool) *CreateMeetingRoomShrinkRequest
	GetEnableCycleReservation() *bool
	SetGroupId(v int64) *CreateMeetingRoomShrinkRequest
	GetGroupId() *int64
	SetIsvRoomId(v string) *CreateMeetingRoomShrinkRequest
	GetIsvRoomId() *string
	SetReservationAuthorityShrink(v string) *CreateMeetingRoomShrinkRequest
	GetReservationAuthorityShrink() *string
	SetRoomCapacity(v int32) *CreateMeetingRoomShrinkRequest
	GetRoomCapacity() *int32
	SetRoomLabelIdsShrink(v string) *CreateMeetingRoomShrinkRequest
	GetRoomLabelIdsShrink() *string
	SetRoomLocationShrink(v string) *CreateMeetingRoomShrinkRequest
	GetRoomLocationShrink() *string
	SetRoomName(v string) *CreateMeetingRoomShrinkRequest
	GetRoomName() *string
	SetRoomPicture(v string) *CreateMeetingRoomShrinkRequest
	GetRoomPicture() *string
	SetRoomStatus(v int32) *CreateMeetingRoomShrinkRequest
	GetRoomStatus() *int32
	SetTenantContextShrink(v string) *CreateMeetingRoomShrinkRequest
	GetTenantContextShrink() *string
}

type CreateMeetingRoomShrinkRequest struct {
	EnableCycleReservation *bool `json:"EnableCycleReservation,omitempty" xml:"EnableCycleReservation,omitempty"`
	// example:
	//
	// 4644
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId                  *string `json:"IsvRoomId,omitempty" xml:"IsvRoomId,omitempty"`
	ReservationAuthorityShrink *string `json:"ReservationAuthority,omitempty" xml:"ReservationAuthority,omitempty"`
	// example:
	//
	// 100
	RoomCapacity       *int32  `json:"RoomCapacity,omitempty" xml:"RoomCapacity,omitempty"`
	RoomLabelIdsShrink *string `json:"RoomLabelIds,omitempty" xml:"RoomLabelIds,omitempty"`
	RoomLocationShrink *string `json:"RoomLocation,omitempty" xml:"RoomLocation,omitempty"`
	RoomName           *string `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADPxxxxx.jpg
	RoomPicture *string `json:"RoomPicture,omitempty" xml:"RoomPicture,omitempty"`
	// example:
	//
	// 1
	RoomStatus          *int32  `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s CreateMeetingRoomShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomShrinkRequest) GetEnableCycleReservation() *bool {
	return s.EnableCycleReservation
}

func (s *CreateMeetingRoomShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateMeetingRoomShrinkRequest) GetIsvRoomId() *string {
	return s.IsvRoomId
}

func (s *CreateMeetingRoomShrinkRequest) GetReservationAuthorityShrink() *string {
	return s.ReservationAuthorityShrink
}

func (s *CreateMeetingRoomShrinkRequest) GetRoomCapacity() *int32 {
	return s.RoomCapacity
}

func (s *CreateMeetingRoomShrinkRequest) GetRoomLabelIdsShrink() *string {
	return s.RoomLabelIdsShrink
}

func (s *CreateMeetingRoomShrinkRequest) GetRoomLocationShrink() *string {
	return s.RoomLocationShrink
}

func (s *CreateMeetingRoomShrinkRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *CreateMeetingRoomShrinkRequest) GetRoomPicture() *string {
	return s.RoomPicture
}

func (s *CreateMeetingRoomShrinkRequest) GetRoomStatus() *int32 {
	return s.RoomStatus
}

func (s *CreateMeetingRoomShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateMeetingRoomShrinkRequest) SetEnableCycleReservation(v bool) *CreateMeetingRoomShrinkRequest {
	s.EnableCycleReservation = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetGroupId(v int64) *CreateMeetingRoomShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetIsvRoomId(v string) *CreateMeetingRoomShrinkRequest {
	s.IsvRoomId = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetReservationAuthorityShrink(v string) *CreateMeetingRoomShrinkRequest {
	s.ReservationAuthorityShrink = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetRoomCapacity(v int32) *CreateMeetingRoomShrinkRequest {
	s.RoomCapacity = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetRoomLabelIdsShrink(v string) *CreateMeetingRoomShrinkRequest {
	s.RoomLabelIdsShrink = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetRoomLocationShrink(v string) *CreateMeetingRoomShrinkRequest {
	s.RoomLocationShrink = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetRoomName(v string) *CreateMeetingRoomShrinkRequest {
	s.RoomName = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetRoomPicture(v string) *CreateMeetingRoomShrinkRequest {
	s.RoomPicture = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetRoomStatus(v int32) *CreateMeetingRoomShrinkRequest {
	s.RoomStatus = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) SetTenantContextShrink(v string) *CreateMeetingRoomShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateMeetingRoomShrinkRequest) Validate() error {
	return dara.Validate(s)
}

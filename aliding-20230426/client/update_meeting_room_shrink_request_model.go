// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCycleReservation(v bool) *UpdateMeetingRoomShrinkRequest
	GetEnableCycleReservation() *bool
	SetGroupId(v int64) *UpdateMeetingRoomShrinkRequest
	GetGroupId() *int64
	SetIsvRoomId(v string) *UpdateMeetingRoomShrinkRequest
	GetIsvRoomId() *string
	SetReservationAuthorityShrink(v string) *UpdateMeetingRoomShrinkRequest
	GetReservationAuthorityShrink() *string
	SetRoomCapacity(v int32) *UpdateMeetingRoomShrinkRequest
	GetRoomCapacity() *int32
	SetRoomId(v string) *UpdateMeetingRoomShrinkRequest
	GetRoomId() *string
	SetRoomLabelIdsShrink(v string) *UpdateMeetingRoomShrinkRequest
	GetRoomLabelIdsShrink() *string
	SetRoomLocationShrink(v string) *UpdateMeetingRoomShrinkRequest
	GetRoomLocationShrink() *string
	SetRoomName(v string) *UpdateMeetingRoomShrinkRequest
	GetRoomName() *string
	SetRoomPicture(v string) *UpdateMeetingRoomShrinkRequest
	GetRoomPicture() *string
	SetRoomStatus(v int32) *UpdateMeetingRoomShrinkRequest
	GetRoomStatus() *int32
	SetTenantContextShrink(v string) *UpdateMeetingRoomShrinkRequest
	GetTenantContextShrink() *string
}

type UpdateMeetingRoomShrinkRequest struct {
	EnableCycleReservation *bool `json:"EnableCycleReservation,omitempty" xml:"EnableCycleReservation,omitempty"`
	// example:
	//
	// 0
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId                  *string `json:"IsvRoomId,omitempty" xml:"IsvRoomId,omitempty"`
	ReservationAuthorityShrink *string `json:"ReservationAuthority,omitempty" xml:"ReservationAuthority,omitempty"`
	// example:
	//
	// 100
	RoomCapacity *int32 `json:"RoomCapacity,omitempty" xml:"RoomCapacity,omitempty"`
	// example:
	//
	// 0ffbxxxxx
	RoomId             *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
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

func (s UpdateMeetingRoomShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomShrinkRequest) GetEnableCycleReservation() *bool {
	return s.EnableCycleReservation
}

func (s *UpdateMeetingRoomShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateMeetingRoomShrinkRequest) GetIsvRoomId() *string {
	return s.IsvRoomId
}

func (s *UpdateMeetingRoomShrinkRequest) GetReservationAuthorityShrink() *string {
	return s.ReservationAuthorityShrink
}

func (s *UpdateMeetingRoomShrinkRequest) GetRoomCapacity() *int32 {
	return s.RoomCapacity
}

func (s *UpdateMeetingRoomShrinkRequest) GetRoomId() *string {
	return s.RoomId
}

func (s *UpdateMeetingRoomShrinkRequest) GetRoomLabelIdsShrink() *string {
	return s.RoomLabelIdsShrink
}

func (s *UpdateMeetingRoomShrinkRequest) GetRoomLocationShrink() *string {
	return s.RoomLocationShrink
}

func (s *UpdateMeetingRoomShrinkRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *UpdateMeetingRoomShrinkRequest) GetRoomPicture() *string {
	return s.RoomPicture
}

func (s *UpdateMeetingRoomShrinkRequest) GetRoomStatus() *int32 {
	return s.RoomStatus
}

func (s *UpdateMeetingRoomShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateMeetingRoomShrinkRequest) SetEnableCycleReservation(v bool) *UpdateMeetingRoomShrinkRequest {
	s.EnableCycleReservation = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetGroupId(v int64) *UpdateMeetingRoomShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetIsvRoomId(v string) *UpdateMeetingRoomShrinkRequest {
	s.IsvRoomId = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetReservationAuthorityShrink(v string) *UpdateMeetingRoomShrinkRequest {
	s.ReservationAuthorityShrink = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetRoomCapacity(v int32) *UpdateMeetingRoomShrinkRequest {
	s.RoomCapacity = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetRoomId(v string) *UpdateMeetingRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetRoomLabelIdsShrink(v string) *UpdateMeetingRoomShrinkRequest {
	s.RoomLabelIdsShrink = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetRoomLocationShrink(v string) *UpdateMeetingRoomShrinkRequest {
	s.RoomLocationShrink = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetRoomName(v string) *UpdateMeetingRoomShrinkRequest {
	s.RoomName = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetRoomPicture(v string) *UpdateMeetingRoomShrinkRequest {
	s.RoomPicture = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetRoomStatus(v int32) *UpdateMeetingRoomShrinkRequest {
	s.RoomStatus = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) SetTenantContextShrink(v string) *UpdateMeetingRoomShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateMeetingRoomShrinkRequest) Validate() error {
	return dara.Validate(s)
}

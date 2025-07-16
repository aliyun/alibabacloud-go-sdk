// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoomId(v string) *DeleteMeetingRoomShrinkRequest
	GetRoomId() *string
	SetTenantContextShrink(v string) *DeleteMeetingRoomShrinkRequest
	GetTenantContextShrink() *string
}

type DeleteMeetingRoomShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0ffb7xxxxx
	RoomId              *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DeleteMeetingRoomShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomShrinkRequest) GetRoomId() *string {
	return s.RoomId
}

func (s *DeleteMeetingRoomShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteMeetingRoomShrinkRequest) SetRoomId(v string) *DeleteMeetingRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *DeleteMeetingRoomShrinkRequest) SetTenantContextShrink(v string) *DeleteMeetingRoomShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteMeetingRoomShrinkRequest) Validate() error {
	return dara.Validate(s)
}

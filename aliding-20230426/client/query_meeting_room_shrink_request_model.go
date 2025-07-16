// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoomId(v string) *QueryMeetingRoomShrinkRequest
	GetRoomId() *string
	SetTenantContextShrink(v string) *QueryMeetingRoomShrinkRequest
	GetTenantContextShrink() *string
}

type QueryMeetingRoomShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0ffb7xxxxx
	RoomId              *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryMeetingRoomShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomShrinkRequest) GetRoomId() *string {
	return s.RoomId
}

func (s *QueryMeetingRoomShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryMeetingRoomShrinkRequest) SetRoomId(v string) *QueryMeetingRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *QueryMeetingRoomShrinkRequest) SetTenantContextShrink(v string) *QueryMeetingRoomShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryMeetingRoomShrinkRequest) Validate() error {
	return dara.Validate(s)
}

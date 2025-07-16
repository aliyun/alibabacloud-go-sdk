// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteMeetingRoomGroupShrinkRequest
	GetGroupId() *string
	SetTenantContextShrink(v string) *DeleteMeetingRoomGroupShrinkRequest
	GetTenantContextShrink() *string
}

type DeleteMeetingRoomGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 172
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DeleteMeetingRoomGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteMeetingRoomGroupShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteMeetingRoomGroupShrinkRequest) SetGroupId(v string) *DeleteMeetingRoomGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMeetingRoomGroupShrinkRequest) SetTenantContextShrink(v string) *DeleteMeetingRoomGroupShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteMeetingRoomGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

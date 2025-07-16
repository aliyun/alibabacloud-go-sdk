// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *UpdateMeetingRoomGroupShrinkRequest
	GetGroupId() *string
	SetGroupName(v string) *UpdateMeetingRoomGroupShrinkRequest
	GetGroupName() *string
	SetTenantContextShrink(v string) *UpdateMeetingRoomGroupShrinkRequest
	GetTenantContextShrink() *string
}

type UpdateMeetingRoomGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 172
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s UpdateMeetingRoomGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateMeetingRoomGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateMeetingRoomGroupShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateMeetingRoomGroupShrinkRequest) SetGroupId(v string) *UpdateMeetingRoomGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMeetingRoomGroupShrinkRequest) SetGroupName(v string) *UpdateMeetingRoomGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateMeetingRoomGroupShrinkRequest) SetTenantContextShrink(v string) *UpdateMeetingRoomGroupShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateMeetingRoomGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

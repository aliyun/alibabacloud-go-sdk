// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *CreateMeetingRoomGroupShrinkRequest
	GetGroupName() *string
	SetParentGroupId(v int64) *CreateMeetingRoomGroupShrinkRequest
	GetParentGroupId() *int64
	SetTenantContextShrink(v string) *CreateMeetingRoomGroupShrinkRequest
	GetTenantContextShrink() *string
}

type CreateMeetingRoomGroupShrinkRequest struct {
	// example:
	//
	// 测试分组
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 172L
	ParentGroupId       *int64  `json:"ParentGroupId,omitempty" xml:"ParentGroupId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s CreateMeetingRoomGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateMeetingRoomGroupShrinkRequest) GetParentGroupId() *int64 {
	return s.ParentGroupId
}

func (s *CreateMeetingRoomGroupShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateMeetingRoomGroupShrinkRequest) SetGroupName(v string) *CreateMeetingRoomGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateMeetingRoomGroupShrinkRequest) SetParentGroupId(v int64) *CreateMeetingRoomGroupShrinkRequest {
	s.ParentGroupId = &v
	return s
}

func (s *CreateMeetingRoomGroupShrinkRequest) SetTenantContextShrink(v string) *CreateMeetingRoomGroupShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateMeetingRoomGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

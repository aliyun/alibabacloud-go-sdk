// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *QueryMeetingRoomGroupShrinkRequest
	GetGroupId() *string
	SetTenantContextShrink(v string) *QueryMeetingRoomGroupShrinkRequest
	GetTenantContextShrink() *string
}

type QueryMeetingRoomGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 172
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryMeetingRoomGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryMeetingRoomGroupShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryMeetingRoomGroupShrinkRequest) SetGroupId(v string) *QueryMeetingRoomGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomGroupShrinkRequest) SetTenantContextShrink(v string) *QueryMeetingRoomGroupShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryMeetingRoomGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

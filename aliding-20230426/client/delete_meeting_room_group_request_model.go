// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteMeetingRoomGroupRequest
	GetGroupId() *string
	SetTenantContext(v *DeleteMeetingRoomGroupRequestTenantContext) *DeleteMeetingRoomGroupRequest
	GetTenantContext() *DeleteMeetingRoomGroupRequestTenantContext
}

type DeleteMeetingRoomGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 172
	GroupId       *string                                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	TenantContext *DeleteMeetingRoomGroupRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DeleteMeetingRoomGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteMeetingRoomGroupRequest) GetTenantContext() *DeleteMeetingRoomGroupRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteMeetingRoomGroupRequest) SetGroupId(v string) *DeleteMeetingRoomGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMeetingRoomGroupRequest) SetTenantContext(v *DeleteMeetingRoomGroupRequestTenantContext) *DeleteMeetingRoomGroupRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteMeetingRoomGroupRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteMeetingRoomGroupRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteMeetingRoomGroupRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomGroupRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMeetingRoomGroupRequestTenantContext) SetTenantId(v string) *DeleteMeetingRoomGroupRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteMeetingRoomGroupRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoomId(v string) *DeleteMeetingRoomRequest
	GetRoomId() *string
	SetTenantContext(v *DeleteMeetingRoomRequestTenantContext) *DeleteMeetingRoomRequest
	GetTenantContext() *DeleteMeetingRoomRequestTenantContext
}

type DeleteMeetingRoomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0ffb7xxxxx
	RoomId        *string                                `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	TenantContext *DeleteMeetingRoomRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DeleteMeetingRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomRequest) GetRoomId() *string {
	return s.RoomId
}

func (s *DeleteMeetingRoomRequest) GetTenantContext() *DeleteMeetingRoomRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteMeetingRoomRequest) SetRoomId(v string) *DeleteMeetingRoomRequest {
	s.RoomId = &v
	return s
}

func (s *DeleteMeetingRoomRequest) SetTenantContext(v *DeleteMeetingRoomRequestTenantContext) *DeleteMeetingRoomRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteMeetingRoomRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteMeetingRoomRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteMeetingRoomRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMeetingRoomRequestTenantContext) SetTenantId(v string) *DeleteMeetingRoomRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteMeetingRoomRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

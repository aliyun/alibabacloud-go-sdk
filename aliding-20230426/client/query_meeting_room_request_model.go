// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoomId(v string) *QueryMeetingRoomRequest
	GetRoomId() *string
	SetTenantContext(v *QueryMeetingRoomRequestTenantContext) *QueryMeetingRoomRequest
	GetTenantContext() *QueryMeetingRoomRequestTenantContext
}

type QueryMeetingRoomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0ffb7xxxxx
	RoomId        *string                               `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	TenantContext *QueryMeetingRoomRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomRequest) GetRoomId() *string {
	return s.RoomId
}

func (s *QueryMeetingRoomRequest) GetTenantContext() *QueryMeetingRoomRequestTenantContext {
	return s.TenantContext
}

func (s *QueryMeetingRoomRequest) SetRoomId(v string) *QueryMeetingRoomRequest {
	s.RoomId = &v
	return s
}

func (s *QueryMeetingRoomRequest) SetTenantContext(v *QueryMeetingRoomRequestTenantContext) *QueryMeetingRoomRequest {
	s.TenantContext = v
	return s
}

func (s *QueryMeetingRoomRequest) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryMeetingRoomRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMeetingRoomRequestTenantContext) SetTenantId(v string) *QueryMeetingRoomRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryMeetingRoomRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

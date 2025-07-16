// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *QueryMeetingRoomGroupRequest
	GetGroupId() *string
	SetTenantContext(v *QueryMeetingRoomGroupRequestTenantContext) *QueryMeetingRoomGroupRequest
	GetTenantContext() *QueryMeetingRoomGroupRequestTenantContext
}

type QueryMeetingRoomGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 172
	GroupId       *string                                    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	TenantContext *QueryMeetingRoomGroupRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryMeetingRoomGroupRequest) GetTenantContext() *QueryMeetingRoomGroupRequestTenantContext {
	return s.TenantContext
}

func (s *QueryMeetingRoomGroupRequest) SetGroupId(v string) *QueryMeetingRoomGroupRequest {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomGroupRequest) SetTenantContext(v *QueryMeetingRoomGroupRequestTenantContext) *QueryMeetingRoomGroupRequest {
	s.TenantContext = v
	return s
}

func (s *QueryMeetingRoomGroupRequest) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomGroupRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryMeetingRoomGroupRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMeetingRoomGroupRequestTenantContext) SetTenantId(v string) *QueryMeetingRoomGroupRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryMeetingRoomGroupRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

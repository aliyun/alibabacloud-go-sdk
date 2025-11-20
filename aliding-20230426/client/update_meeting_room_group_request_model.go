// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *UpdateMeetingRoomGroupRequest
	GetGroupId() *string
	SetGroupName(v string) *UpdateMeetingRoomGroupRequest
	GetGroupName() *string
	SetTenantContext(v *UpdateMeetingRoomGroupRequestTenantContext) *UpdateMeetingRoomGroupRequest
	GetTenantContext() *UpdateMeetingRoomGroupRequestTenantContext
}

type UpdateMeetingRoomGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 172
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName     *string                                     `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	TenantContext *UpdateMeetingRoomGroupRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UpdateMeetingRoomGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateMeetingRoomGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateMeetingRoomGroupRequest) GetTenantContext() *UpdateMeetingRoomGroupRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateMeetingRoomGroupRequest) SetGroupId(v string) *UpdateMeetingRoomGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMeetingRoomGroupRequest) SetGroupName(v string) *UpdateMeetingRoomGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateMeetingRoomGroupRequest) SetTenantContext(v *UpdateMeetingRoomGroupRequestTenantContext) *UpdateMeetingRoomGroupRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateMeetingRoomGroupRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMeetingRoomGroupRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateMeetingRoomGroupRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomGroupRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMeetingRoomGroupRequestTenantContext) SetTenantId(v string) *UpdateMeetingRoomGroupRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateMeetingRoomGroupRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

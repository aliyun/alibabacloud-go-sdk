// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *CreateMeetingRoomGroupRequest
	GetGroupName() *string
	SetParentGroupId(v int64) *CreateMeetingRoomGroupRequest
	GetParentGroupId() *int64
	SetTenantContext(v *CreateMeetingRoomGroupRequestTenantContext) *CreateMeetingRoomGroupRequest
	GetTenantContext() *CreateMeetingRoomGroupRequestTenantContext
}

type CreateMeetingRoomGroupRequest struct {
	// example:
	//
	// 测试分组
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 172L
	ParentGroupId *int64                                      `json:"ParentGroupId,omitempty" xml:"ParentGroupId,omitempty"`
	TenantContext *CreateMeetingRoomGroupRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s CreateMeetingRoomGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateMeetingRoomGroupRequest) GetParentGroupId() *int64 {
	return s.ParentGroupId
}

func (s *CreateMeetingRoomGroupRequest) GetTenantContext() *CreateMeetingRoomGroupRequestTenantContext {
	return s.TenantContext
}

func (s *CreateMeetingRoomGroupRequest) SetGroupName(v string) *CreateMeetingRoomGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateMeetingRoomGroupRequest) SetParentGroupId(v int64) *CreateMeetingRoomGroupRequest {
	s.ParentGroupId = &v
	return s
}

func (s *CreateMeetingRoomGroupRequest) SetTenantContext(v *CreateMeetingRoomGroupRequestTenantContext) *CreateMeetingRoomGroupRequest {
	s.TenantContext = v
	return s
}

func (s *CreateMeetingRoomGroupRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMeetingRoomGroupRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateMeetingRoomGroupRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomGroupRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMeetingRoomGroupRequestTenantContext) SetTenantId(v string) *CreateMeetingRoomGroupRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateMeetingRoomGroupRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

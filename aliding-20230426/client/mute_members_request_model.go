// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *MuteMembersRequestTenantContext) *MuteMembersRequest
	GetTenantContext() *MuteMembersRequestTenantContext
	SetUserIds(v []*string) *MuteMembersRequest
	GetUserIds() []*string
	SetConferenceId(v string) *MuteMembersRequest
	GetConferenceId() *string
	SetMuteAction(v string) *MuteMembersRequest
	GetMuteAction() *string
}

type MuteMembersRequest struct {
	TenantContext *MuteMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// ["012345"]
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mute
	MuteAction *string `json:"muteAction,omitempty" xml:"muteAction,omitempty"`
}

func (s MuteMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s MuteMembersRequest) GoString() string {
	return s.String()
}

func (s *MuteMembersRequest) GetTenantContext() *MuteMembersRequestTenantContext {
	return s.TenantContext
}

func (s *MuteMembersRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *MuteMembersRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *MuteMembersRequest) GetMuteAction() *string {
	return s.MuteAction
}

func (s *MuteMembersRequest) SetTenantContext(v *MuteMembersRequestTenantContext) *MuteMembersRequest {
	s.TenantContext = v
	return s
}

func (s *MuteMembersRequest) SetUserIds(v []*string) *MuteMembersRequest {
	s.UserIds = v
	return s
}

func (s *MuteMembersRequest) SetConferenceId(v string) *MuteMembersRequest {
	s.ConferenceId = &v
	return s
}

func (s *MuteMembersRequest) SetMuteAction(v string) *MuteMembersRequest {
	s.MuteAction = &v
	return s
}

func (s *MuteMembersRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MuteMembersRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MuteMembersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s MuteMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *MuteMembersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *MuteMembersRequestTenantContext) SetTenantId(v string) *MuteMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *MuteMembersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

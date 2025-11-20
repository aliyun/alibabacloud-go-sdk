// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInviteeList(v []*InviteUsersRequestInviteeList) *InviteUsersRequest
	GetInviteeList() []*InviteUsersRequestInviteeList
	SetTenantContext(v *InviteUsersRequestTenantContext) *InviteUsersRequest
	GetTenantContext() *InviteUsersRequestTenantContext
	SetConferenceId(v string) *InviteUsersRequest
	GetConferenceId() *string
	SetPhoneInviteeList(v []*InviteUsersRequestPhoneInviteeList) *InviteUsersRequest
	GetPhoneInviteeList() []*InviteUsersRequestPhoneInviteeList
}

type InviteUsersRequest struct {
	InviteeList   []*InviteUsersRequestInviteeList `json:"InviteeList,omitempty" xml:"InviteeList,omitempty" type:"Repeated"`
	TenantContext *InviteUsersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId     *string                               `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	PhoneInviteeList []*InviteUsersRequestPhoneInviteeList `json:"phoneInviteeList,omitempty" xml:"phoneInviteeList,omitempty" type:"Repeated"`
}

func (s InviteUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersRequest) GoString() string {
	return s.String()
}

func (s *InviteUsersRequest) GetInviteeList() []*InviteUsersRequestInviteeList {
	return s.InviteeList
}

func (s *InviteUsersRequest) GetTenantContext() *InviteUsersRequestTenantContext {
	return s.TenantContext
}

func (s *InviteUsersRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *InviteUsersRequest) GetPhoneInviteeList() []*InviteUsersRequestPhoneInviteeList {
	return s.PhoneInviteeList
}

func (s *InviteUsersRequest) SetInviteeList(v []*InviteUsersRequestInviteeList) *InviteUsersRequest {
	s.InviteeList = v
	return s
}

func (s *InviteUsersRequest) SetTenantContext(v *InviteUsersRequestTenantContext) *InviteUsersRequest {
	s.TenantContext = v
	return s
}

func (s *InviteUsersRequest) SetConferenceId(v string) *InviteUsersRequest {
	s.ConferenceId = &v
	return s
}

func (s *InviteUsersRequest) SetPhoneInviteeList(v []*InviteUsersRequestPhoneInviteeList) *InviteUsersRequest {
	s.PhoneInviteeList = v
	return s
}

func (s *InviteUsersRequest) Validate() error {
	if s.InviteeList != nil {
		for _, item := range s.InviteeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	if s.PhoneInviteeList != nil {
		for _, item := range s.PhoneInviteeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InviteUsersRequestInviteeList struct {
	// This parameter is required.
	//
	// example:
	//
	// 测试用户
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InviteUsersRequestInviteeList) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersRequestInviteeList) GoString() string {
	return s.String()
}

func (s *InviteUsersRequestInviteeList) GetNick() *string {
	return s.Nick
}

func (s *InviteUsersRequestInviteeList) GetUserId() *string {
	return s.UserId
}

func (s *InviteUsersRequestInviteeList) SetNick(v string) *InviteUsersRequestInviteeList {
	s.Nick = &v
	return s
}

func (s *InviteUsersRequestInviteeList) SetUserId(v string) *InviteUsersRequestInviteeList {
	s.UserId = &v
	return s
}

func (s *InviteUsersRequestInviteeList) Validate() error {
	return dara.Validate(s)
}

type InviteUsersRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InviteUsersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InviteUsersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *InviteUsersRequestTenantContext) SetTenantId(v string) *InviteUsersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *InviteUsersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

type InviteUsersRequestPhoneInviteeList struct {
	InviteClient *bool   `json:"InviteClient,omitempty" xml:"InviteClient,omitempty"`
	Nick         *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	PhoneNumber  *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	StatusCode   *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s InviteUsersRequestPhoneInviteeList) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersRequestPhoneInviteeList) GoString() string {
	return s.String()
}

func (s *InviteUsersRequestPhoneInviteeList) GetInviteClient() *bool {
	return s.InviteClient
}

func (s *InviteUsersRequestPhoneInviteeList) GetNick() *string {
	return s.Nick
}

func (s *InviteUsersRequestPhoneInviteeList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *InviteUsersRequestPhoneInviteeList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *InviteUsersRequestPhoneInviteeList) SetInviteClient(v bool) *InviteUsersRequestPhoneInviteeList {
	s.InviteClient = &v
	return s
}

func (s *InviteUsersRequestPhoneInviteeList) SetNick(v string) *InviteUsersRequestPhoneInviteeList {
	s.Nick = &v
	return s
}

func (s *InviteUsersRequestPhoneInviteeList) SetPhoneNumber(v string) *InviteUsersRequestPhoneInviteeList {
	s.PhoneNumber = &v
	return s
}

func (s *InviteUsersRequestPhoneInviteeList) SetStatusCode(v string) *InviteUsersRequestPhoneInviteeList {
	s.StatusCode = &v
	return s
}

func (s *InviteUsersRequestPhoneInviteeList) Validate() error {
	return dara.Validate(s)
}

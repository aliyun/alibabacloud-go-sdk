// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserInvitationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *CreateUserInvitationCodeRequest
	GetEndUserId() *string
	SetEndUserIds(v []*string) *CreateUserInvitationCodeRequest
	GetEndUserIds() []*string
	SetTerminalGroupId(v string) *CreateUserInvitationCodeRequest
	GetTerminalGroupId() *string
	SetValidityMinutes(v int32) *CreateUserInvitationCodeRequest
	GetValidityMinutes() *int32
}

type CreateUserInvitationCodeRequest struct {
	// The logon identifier (EndUserId) of a single user, used for compatibility with single-user creation. You must specify either this parameter or EndUserIds, but not both. If neither is specified, PARAMETER_MISSING is returned. If both are specified, PARAM_ERROR is returned.
	//
	// example:
	//
	// user-001
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The array of user logon identifiers (EndUserIds). You must specify either this parameter or EndUserId. You can specify 1 to 100 unique, non-empty strings at a time, and a code is generated for each user. If any user does not exist, USER_NOT_FOUND is returned. If the array is empty, or contains duplicate items, empty items, or non-string items, PARAM_ERROR is returned.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The ID of the terminal group. The group must belong to the current account and must not be deleted. After the corresponding terminal uses the management code to complete management, it is added to this group. If the group does not exist, is deleted, or does not belong to the current account, TERMINAL_GROUP_NOT_FOUND is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// tg-example
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	// The validity period of the management code, in minutes. A value of -1 indicates that the code is permanently valid. For values between 60 and 1440, only integer multiples of 60 (1 to 24 hours) are allowed. For values greater than 1440, only integer multiples of 1440 are allowed, up to a maximum of 12960 (2 to 9 days). Other values return PARAM_ERROR. The server calculates the validity period based on this duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1440
	ValidityMinutes *int32 `json:"ValidityMinutes,omitempty" xml:"ValidityMinutes,omitempty"`
}

func (s CreateUserInvitationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserInvitationCodeRequest) GoString() string {
	return s.String()
}

func (s *CreateUserInvitationCodeRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CreateUserInvitationCodeRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *CreateUserInvitationCodeRequest) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *CreateUserInvitationCodeRequest) GetValidityMinutes() *int32 {
	return s.ValidityMinutes
}

func (s *CreateUserInvitationCodeRequest) SetEndUserId(v string) *CreateUserInvitationCodeRequest {
	s.EndUserId = &v
	return s
}

func (s *CreateUserInvitationCodeRequest) SetEndUserIds(v []*string) *CreateUserInvitationCodeRequest {
	s.EndUserIds = v
	return s
}

func (s *CreateUserInvitationCodeRequest) SetTerminalGroupId(v string) *CreateUserInvitationCodeRequest {
	s.TerminalGroupId = &v
	return s
}

func (s *CreateUserInvitationCodeRequest) SetValidityMinutes(v int32) *CreateUserInvitationCodeRequest {
	s.ValidityMinutes = &v
	return s
}

func (s *CreateUserInvitationCodeRequest) Validate() error {
	return dara.Validate(s)
}

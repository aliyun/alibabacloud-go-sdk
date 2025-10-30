// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTenantMembersBySourceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommand(v *AddTenantMembersBySourceUserRequestAddCommand) *AddTenantMembersBySourceUserRequest
	GetAddCommand() *AddTenantMembersBySourceUserRequestAddCommand
	SetOpTenantId(v int64) *AddTenantMembersBySourceUserRequest
	GetOpTenantId() *int64
}

type AddTenantMembersBySourceUserRequest struct {
	AddCommand *AddTenantMembersBySourceUserRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddTenantMembersBySourceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersBySourceUserRequest) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserRequest) GetAddCommand() *AddTenantMembersBySourceUserRequestAddCommand {
	return s.AddCommand
}

func (s *AddTenantMembersBySourceUserRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddTenantMembersBySourceUserRequest) SetAddCommand(v *AddTenantMembersBySourceUserRequestAddCommand) *AddTenantMembersBySourceUserRequest {
	s.AddCommand = v
	return s
}

func (s *AddTenantMembersBySourceUserRequest) SetOpTenantId(v int64) *AddTenantMembersBySourceUserRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequest) Validate() error {
	if s.AddCommand != nil {
		if err := s.AddCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTenantMembersBySourceUserRequestAddCommand struct {
	SourceUserList []*AddTenantMembersBySourceUserRequestAddCommandSourceUserList `json:"SourceUserList,omitempty" xml:"SourceUserList,omitempty" type:"Repeated"`
}

func (s AddTenantMembersBySourceUserRequestAddCommand) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersBySourceUserRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserRequestAddCommand) GetSourceUserList() []*AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	return s.SourceUserList
}

func (s *AddTenantMembersBySourceUserRequestAddCommand) SetSourceUserList(v []*AddTenantMembersBySourceUserRequestAddCommandSourceUserList) *AddTenantMembersBySourceUserRequestAddCommand {
	s.SourceUserList = v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommand) Validate() error {
	if s.SourceUserList != nil {
		for _, item := range s.SourceUserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddTenantMembersBySourceUserRequestAddCommandSourceUserList struct {
	// example:
	//
	// 123@xx.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 123@dingding
	DingNumber  *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123@xx.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 13888888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// example:
	//
	// 2323131
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s AddTenantMembersBySourceUserRequestAddCommandSourceUserList) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersBySourceUserRequestAddCommandSourceUserList) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) GetAccountName() *string {
	return s.AccountName
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) GetDingNumber() *string {
	return s.DingNumber
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) GetMail() *string {
	return s.Mail
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) GetSourceId() *string {
	return s.SourceId
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetAccountName(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.AccountName = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetDingNumber(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.DingNumber = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetDisplayName(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.DisplayName = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetMail(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.Mail = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetMobilePhone(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.MobilePhone = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetSourceId(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.SourceId = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) Validate() error {
	return dara.Validate(s)
}

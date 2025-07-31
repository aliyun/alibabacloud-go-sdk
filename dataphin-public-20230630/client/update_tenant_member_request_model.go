// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateTenantMemberRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateTenantMemberRequestUpdateCommand) *UpdateTenantMemberRequest
	GetUpdateCommand() *UpdateTenantMemberRequestUpdateCommand
}

type UpdateTenantMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateTenantMemberRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateTenantMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateTenantMemberRequest) GetUpdateCommand() *UpdateTenantMemberRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateTenantMemberRequest) SetOpTenantId(v int64) *UpdateTenantMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateTenantMemberRequest) SetUpdateCommand(v *UpdateTenantMemberRequestUpdateCommand) *UpdateTenantMemberRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateTenantMemberRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTenantMemberRequestUpdateCommand struct {
	// This parameter is required.
	MemberList []*UpdateTenantMemberRequestUpdateCommandMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s UpdateTenantMemberRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantMemberRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberRequestUpdateCommand) GetMemberList() []*UpdateTenantMemberRequestUpdateCommandMemberList {
	return s.MemberList
}

func (s *UpdateTenantMemberRequestUpdateCommand) SetMemberList(v []*UpdateTenantMemberRequestUpdateCommandMemberList) *UpdateTenantMemberRequestUpdateCommand {
	s.MemberList = v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}

type UpdateTenantMemberRequestUpdateCommandMemberList struct {
	// example:
	//
	// 123@dingding
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// example:
	//
	// 123@xx.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 13888888888
	MobilePhone *string   `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	RoleList    []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// example:
	//
	// 2331
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateTenantMemberRequestUpdateCommandMemberList) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantMemberRequestUpdateCommandMemberList) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) GetDingNumber() *string {
	return s.DingNumber
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) GetMail() *string {
	return s.Mail
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) GetRoleList() []*string {
	return s.RoleList
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) GetUserId() *string {
	return s.UserId
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetDingNumber(v string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.DingNumber = &v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetMail(v string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.Mail = &v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetMobilePhone(v string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.MobilePhone = &v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetRoleList(v []*string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.RoleList = v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetUserId(v string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.UserId = &v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataServiceProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommand(v *AddDataServiceProjectMemberRequestAddCommand) *AddDataServiceProjectMemberRequest
	GetAddCommand() *AddDataServiceProjectMemberRequestAddCommand
	SetOpTenantId(v int64) *AddDataServiceProjectMemberRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *AddDataServiceProjectMemberRequest
	GetProjectId() *int32
}

type AddDataServiceProjectMemberRequest struct {
	// This parameter is required.
	AddCommand *AddDataServiceProjectMemberRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AddDataServiceProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *AddDataServiceProjectMemberRequest) GetAddCommand() *AddDataServiceProjectMemberRequestAddCommand {
	return s.AddCommand
}

func (s *AddDataServiceProjectMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddDataServiceProjectMemberRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *AddDataServiceProjectMemberRequest) SetAddCommand(v *AddDataServiceProjectMemberRequestAddCommand) *AddDataServiceProjectMemberRequest {
	s.AddCommand = v
	return s
}

func (s *AddDataServiceProjectMemberRequest) SetOpTenantId(v int64) *AddDataServiceProjectMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddDataServiceProjectMemberRequest) SetProjectId(v int32) *AddDataServiceProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *AddDataServiceProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}

type AddDataServiceProjectMemberRequestAddCommand struct {
	// This parameter is required.
	ProjectMemberList []*AddDataServiceProjectMemberRequestAddCommandProjectMemberList `json:"ProjectMemberList,omitempty" xml:"ProjectMemberList,omitempty" type:"Repeated"`
}

func (s AddDataServiceProjectMemberRequestAddCommand) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceProjectMemberRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddDataServiceProjectMemberRequestAddCommand) GetProjectMemberList() []*AddDataServiceProjectMemberRequestAddCommandProjectMemberList {
	return s.ProjectMemberList
}

func (s *AddDataServiceProjectMemberRequestAddCommand) SetProjectMemberList(v []*AddDataServiceProjectMemberRequestAddCommandProjectMemberList) *AddDataServiceProjectMemberRequestAddCommand {
	s.ProjectMemberList = v
	return s
}

func (s *AddDataServiceProjectMemberRequestAddCommand) Validate() error {
	return dara.Validate(s)
}

type AddDataServiceProjectMemberRequestAddCommandProjectMemberList struct {
	// This parameter is required.
	//
	// example:
	//
	// xx@aliyuncs.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30012011
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddDataServiceProjectMemberRequestAddCommandProjectMemberList) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceProjectMemberRequestAddCommandProjectMemberList) GoString() string {
	return s.String()
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) GetAccountName() *string {
	return s.AccountName
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) GetRole() *int32 {
	return s.Role
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) GetUserId() *string {
	return s.UserId
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) SetAccountName(v string) *AddDataServiceProjectMemberRequestAddCommandProjectMemberList {
	s.AccountName = &v
	return s
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) SetDisplayName(v string) *AddDataServiceProjectMemberRequestAddCommandProjectMemberList {
	s.DisplayName = &v
	return s
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) SetRole(v int32) *AddDataServiceProjectMemberRequestAddCommandProjectMemberList {
	s.Role = &v
	return s
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) SetUserId(v string) *AddDataServiceProjectMemberRequestAddCommandProjectMemberList {
	s.UserId = &v
	return s
}

func (s *AddDataServiceProjectMemberRequestAddCommandProjectMemberList) Validate() error {
	return dara.Validate(s)
}

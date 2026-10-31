// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyResourcePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyCommand(v *ApplyResourcePermissionRequestApplyCommand) *ApplyResourcePermissionRequest
	GetApplyCommand() *ApplyResourcePermissionRequestApplyCommand
	SetOpTenantId(v int64) *ApplyResourcePermissionRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ApplyResourcePermissionRequest
	GetOpUserId() *string
}

type ApplyResourcePermissionRequest struct {
	// The permission application request.
	//
	// This parameter is required.
	ApplyCommand *ApplyResourcePermissionRequestApplyCommand `json:"ApplyCommand,omitempty" xml:"ApplyCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ApplyResourcePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *ApplyResourcePermissionRequest) GetApplyCommand() *ApplyResourcePermissionRequestApplyCommand {
	return s.ApplyCommand
}

func (s *ApplyResourcePermissionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ApplyResourcePermissionRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ApplyResourcePermissionRequest) SetApplyCommand(v *ApplyResourcePermissionRequestApplyCommand) *ApplyResourcePermissionRequest {
	s.ApplyCommand = v
	return s
}

func (s *ApplyResourcePermissionRequest) SetOpTenantId(v int64) *ApplyResourcePermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *ApplyResourcePermissionRequest) SetOpUserId(v string) *ApplyResourcePermissionRequest {
	s.OpUserId = &v
	return s
}

func (s *ApplyResourcePermissionRequest) Validate() error {
	if s.ApplyCommand != nil {
		if err := s.ApplyCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyResourcePermissionRequestApplyCommand struct {
	// The list of users to whom permissions are granted.
	//
	// This parameter is required.
	GrantToUsers []*ApplyResourcePermissionRequestApplyCommandGrantToUsers `json:"GrantToUsers,omitempty" xml:"GrantToUsers,omitempty" type:"Repeated"`
	// The permission granularity. Valid values: TABLE (table-level permissions) and COLUMN (column-level permissions).
	//
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The list of operation types.
	//
	// This parameter is required.
	OperateList []*string `json:"OperateList,omitempty" xml:"OperateList,omitempty" type:"Repeated"`
	// The reason for the application. This parameter is required and cannot be empty or contain only whitespace characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The list of resources for which permissions are applied.
	//
	// This parameter is required.
	ResourceList []*ApplyResourcePermissionRequestApplyCommandResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHYSICAL_TABLE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ApplyResourcePermissionRequestApplyCommand) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourcePermissionRequestApplyCommand) GoString() string {
	return s.String()
}

func (s *ApplyResourcePermissionRequestApplyCommand) GetGrantToUsers() []*ApplyResourcePermissionRequestApplyCommandGrantToUsers {
	return s.GrantToUsers
}

func (s *ApplyResourcePermissionRequestApplyCommand) GetGranularity() *string {
	return s.Granularity
}

func (s *ApplyResourcePermissionRequestApplyCommand) GetOperateList() []*string {
	return s.OperateList
}

func (s *ApplyResourcePermissionRequestApplyCommand) GetReason() *string {
	return s.Reason
}

func (s *ApplyResourcePermissionRequestApplyCommand) GetResourceList() []*ApplyResourcePermissionRequestApplyCommandResourceList {
	return s.ResourceList
}

func (s *ApplyResourcePermissionRequestApplyCommand) GetResourceType() *string {
	return s.ResourceType
}

func (s *ApplyResourcePermissionRequestApplyCommand) SetGrantToUsers(v []*ApplyResourcePermissionRequestApplyCommandGrantToUsers) *ApplyResourcePermissionRequestApplyCommand {
	s.GrantToUsers = v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommand) SetGranularity(v string) *ApplyResourcePermissionRequestApplyCommand {
	s.Granularity = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommand) SetOperateList(v []*string) *ApplyResourcePermissionRequestApplyCommand {
	s.OperateList = v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommand) SetReason(v string) *ApplyResourcePermissionRequestApplyCommand {
	s.Reason = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommand) SetResourceList(v []*ApplyResourcePermissionRequestApplyCommandResourceList) *ApplyResourcePermissionRequestApplyCommand {
	s.ResourceList = v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommand) SetResourceType(v string) *ApplyResourcePermissionRequestApplyCommand {
	s.ResourceType = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommand) Validate() error {
	if s.GrantToUsers != nil {
		for _, item := range s.GrantToUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyResourcePermissionRequestApplyCommandGrantToUsers struct {
	// The account type.
	//
	// This parameter is required.
	//
	// example:
	//
	// PERSONAL
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The expiration date of the permission. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2026-12-31
	PeriodEnd *string `json:"PeriodEnd,omitempty" xml:"PeriodEnd,omitempty"`
	// The type of the permission validity period.
	//
	// example:
	//
	// DAYS_30
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The project ID of the production account. This parameter is required when the account type is PRODUCE.
	//
	// example:
	//
	// 7298904510790528
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300001186
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ApplyResourcePermissionRequestApplyCommandGrantToUsers) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourcePermissionRequestApplyCommandGrantToUsers) GoString() string {
	return s.String()
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) GetAccountType() *string {
	return s.AccountType
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) GetPeriodType() *string {
	return s.PeriodType
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) GetProjectId() *string {
	return s.ProjectId
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) GetUserId() *string {
	return s.UserId
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) SetAccountType(v string) *ApplyResourcePermissionRequestApplyCommandGrantToUsers {
	s.AccountType = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) SetPeriodEnd(v string) *ApplyResourcePermissionRequestApplyCommandGrantToUsers {
	s.PeriodEnd = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) SetPeriodType(v string) *ApplyResourcePermissionRequestApplyCommandGrantToUsers {
	s.PeriodType = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) SetProjectId(v string) *ApplyResourcePermissionRequestApplyCommandGrantToUsers {
	s.ProjectId = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) SetUserId(v string) *ApplyResourcePermissionRequestApplyCommandGrantToUsers {
	s.UserId = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommandGrantToUsers) Validate() error {
	return dara.Validate(s)
}

type ApplyResourcePermissionRequestApplyCommandResourceList struct {
	// The list of column names specified for column-level permissions. This parameter is required when the granularity is COLUMN.
	ColumnList []*string `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	// The resource ID, which is the GUID of the table metadata.
	//
	// This parameter is required.
	//
	// example:
	//
	// dp_ds_table_xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ApplyResourcePermissionRequestApplyCommandResourceList) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourcePermissionRequestApplyCommandResourceList) GoString() string {
	return s.String()
}

func (s *ApplyResourcePermissionRequestApplyCommandResourceList) GetColumnList() []*string {
	return s.ColumnList
}

func (s *ApplyResourcePermissionRequestApplyCommandResourceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *ApplyResourcePermissionRequestApplyCommandResourceList) SetColumnList(v []*string) *ApplyResourcePermissionRequestApplyCommandResourceList {
	s.ColumnList = v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommandResourceList) SetResourceId(v string) *ApplyResourcePermissionRequestApplyCommandResourceList {
	s.ResourceId = &v
	return s
}

func (s *ApplyResourcePermissionRequestApplyCommandResourceList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardSetRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateStandardSetRequestUpdateCommand) *UpdateStandardSetRequest
	GetUpdateCommand() *UpdateStandardSetRequestUpdateCommand
}

type UpdateStandardSetRequest struct {
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// Update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateStandardSetRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateStandardSetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardSetRequest) GetUpdateCommand() *UpdateStandardSetRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateStandardSetRequest) SetOpTenantId(v int64) *UpdateStandardSetRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardSetRequest) SetUpdateCommand(v *UpdateStandardSetRequestUpdateCommand) *UpdateStandardSetRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateStandardSetRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardSetRequestUpdateCommand struct {
	// Publishing approval configuration.
	ApprovalConfig *UpdateStandardSetRequestUpdateCommandApprovalConfig `json:"ApprovalConfig,omitempty" xml:"ApprovalConfig,omitempty" type:"Struct"`
	// Standard set code.
	//
	// This parameter is required.
	//
	// example:
	//
	// CITY
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Default standard template ID.
	//
	// example:
	//
	// 1001
	DefaultStandardTemplateId *int64 `json:"DefaultStandardTemplateId,omitempty" xml:"DefaultStandardTemplateId,omitempty"`
	// Standard set description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Parent directory.
	DirectoryReference *UpdateStandardSetRequestUpdateCommandDirectoryReference `json:"DirectoryReference,omitempty" xml:"DirectoryReference,omitempty" type:"Struct"`
	// Maintainers.
	MaintainerList []*string `json:"MaintainerList,omitempty" xml:"MaintainerList,omitempty" type:"Repeated"`
	// Member group list.
	MemberGroupList []*string `json:"MemberGroupList,omitempty" xml:"MemberGroupList,omitempty" type:"Repeated"`
	// Member list.
	MemberList []*string `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// Standard set name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unpublishing approval configuration.
	OfflineApprovalConfig *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig `json:"OfflineApprovalConfig,omitempty" xml:"OfflineApprovalConfig,omitempty" type:"Struct"`
	// Standard set ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	StandardSetId *int64 `json:"StandardSetId,omitempty" xml:"StandardSetId,omitempty"`
	// Visibility configuration.
	VisibilityConfig *UpdateStandardSetRequestUpdateCommandVisibilityConfig `json:"VisibilityConfig,omitempty" xml:"VisibilityConfig,omitempty" type:"Struct"`
}

func (s UpdateStandardSetRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetRequestUpdateCommand) GetApprovalConfig() *UpdateStandardSetRequestUpdateCommandApprovalConfig {
	return s.ApprovalConfig
}

func (s *UpdateStandardSetRequestUpdateCommand) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardSetRequestUpdateCommand) GetDefaultStandardTemplateId() *int64 {
	return s.DefaultStandardTemplateId
}

func (s *UpdateStandardSetRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardSetRequestUpdateCommand) GetDirectoryReference() *UpdateStandardSetRequestUpdateCommandDirectoryReference {
	return s.DirectoryReference
}

func (s *UpdateStandardSetRequestUpdateCommand) GetMaintainerList() []*string {
	return s.MaintainerList
}

func (s *UpdateStandardSetRequestUpdateCommand) GetMemberGroupList() []*string {
	return s.MemberGroupList
}

func (s *UpdateStandardSetRequestUpdateCommand) GetMemberList() []*string {
	return s.MemberList
}

func (s *UpdateStandardSetRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateStandardSetRequestUpdateCommand) GetOfflineApprovalConfig() *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig {
	return s.OfflineApprovalConfig
}

func (s *UpdateStandardSetRequestUpdateCommand) GetStandardSetId() *int64 {
	return s.StandardSetId
}

func (s *UpdateStandardSetRequestUpdateCommand) GetVisibilityConfig() *UpdateStandardSetRequestUpdateCommandVisibilityConfig {
	return s.VisibilityConfig
}

func (s *UpdateStandardSetRequestUpdateCommand) SetApprovalConfig(v *UpdateStandardSetRequestUpdateCommandApprovalConfig) *UpdateStandardSetRequestUpdateCommand {
	s.ApprovalConfig = v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetCode(v string) *UpdateStandardSetRequestUpdateCommand {
	s.Code = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetDefaultStandardTemplateId(v int64) *UpdateStandardSetRequestUpdateCommand {
	s.DefaultStandardTemplateId = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetDescription(v string) *UpdateStandardSetRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetDirectoryReference(v *UpdateStandardSetRequestUpdateCommandDirectoryReference) *UpdateStandardSetRequestUpdateCommand {
	s.DirectoryReference = v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetMaintainerList(v []*string) *UpdateStandardSetRequestUpdateCommand {
	s.MaintainerList = v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetMemberGroupList(v []*string) *UpdateStandardSetRequestUpdateCommand {
	s.MemberGroupList = v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetMemberList(v []*string) *UpdateStandardSetRequestUpdateCommand {
	s.MemberList = v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetName(v string) *UpdateStandardSetRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetOfflineApprovalConfig(v *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) *UpdateStandardSetRequestUpdateCommand {
	s.OfflineApprovalConfig = v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetStandardSetId(v int64) *UpdateStandardSetRequestUpdateCommand {
	s.StandardSetId = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) SetVisibilityConfig(v *UpdateStandardSetRequestUpdateCommandVisibilityConfig) *UpdateStandardSetRequestUpdateCommand {
	s.VisibilityConfig = v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommand) Validate() error {
	if s.ApprovalConfig != nil {
		if err := s.ApprovalConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DirectoryReference != nil {
		if err := s.DirectoryReference.Validate(); err != nil {
			return err
		}
	}
	if s.OfflineApprovalConfig != nil {
		if err := s.OfflineApprovalConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VisibilityConfig != nil {
		if err := s.VisibilityConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardSetRequestUpdateCommandApprovalConfig struct {
	// Approval process type. Valid values: BY_DEFAULT (default approval type) and BY_TEMPLATE (template-based approval).
	//
	// This parameter is required.
	//
	// example:
	//
	// BY_DEFAULT
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// Specifies whether to enable approval.
	//
	// This parameter is required.
	EnableApproval *bool `json:"EnableApproval,omitempty" xml:"EnableApproval,omitempty"`
	// Specifies whether to submit approvals in batch.
	//
	// This parameter is required.
	IsSubmitInBatch *bool `json:"IsSubmitInBatch,omitempty" xml:"IsSubmitInBatch,omitempty"`
	// Approval template ID. This parameter takes effect only when the approval process type is set to BY_TEMPLATE.
	//
	// example:
	//
	// 1121
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateStandardSetRequestUpdateCommandApprovalConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetRequestUpdateCommandApprovalConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) GetEnableApproval() *bool {
	return s.EnableApproval
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) GetIsSubmitInBatch() *bool {
	return s.IsSubmitInBatch
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) SetApprovalType(v string) *UpdateStandardSetRequestUpdateCommandApprovalConfig {
	s.ApprovalType = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) SetEnableApproval(v bool) *UpdateStandardSetRequestUpdateCommandApprovalConfig {
	s.EnableApproval = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) SetIsSubmitInBatch(v bool) *UpdateStandardSetRequestUpdateCommandApprovalConfig {
	s.IsSubmitInBatch = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) SetTemplateId(v int64) *UpdateStandardSetRequestUpdateCommandApprovalConfig {
	s.TemplateId = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandApprovalConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardSetRequestUpdateCommandDirectoryReference struct {
	// Directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /dir1
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
}

func (s UpdateStandardSetRequestUpdateCommandDirectoryReference) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetRequestUpdateCommandDirectoryReference) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetRequestUpdateCommandDirectoryReference) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateStandardSetRequestUpdateCommandDirectoryReference) SetDirectory(v string) *UpdateStandardSetRequestUpdateCommandDirectoryReference {
	s.Directory = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandDirectoryReference) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig struct {
	// Approval process type. Valid values: BY_DEFAULT (default approval type) and BY_TEMPLATE (template-based approval).
	//
	// This parameter is required.
	//
	// example:
	//
	// BY_DEFAULT
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// Specifies whether to enable approval.
	//
	// This parameter is required.
	EnableApproval *bool `json:"EnableApproval,omitempty" xml:"EnableApproval,omitempty"`
	// Specifies whether to submit approvals in batch.
	//
	// This parameter is required.
	IsSubmitInBatch *bool `json:"IsSubmitInBatch,omitempty" xml:"IsSubmitInBatch,omitempty"`
	// Approval template ID. This parameter takes effect only when the approval process type is set to BY_TEMPLATE.
	//
	// example:
	//
	// 1121
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) GetEnableApproval() *bool {
	return s.EnableApproval
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) GetIsSubmitInBatch() *bool {
	return s.IsSubmitInBatch
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) SetApprovalType(v string) *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig {
	s.ApprovalType = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) SetEnableApproval(v bool) *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig {
	s.EnableApproval = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) SetIsSubmitInBatch(v bool) *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig {
	s.IsSubmitInBatch = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) SetTemplateId(v int64) *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig {
	s.TemplateId = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandOfflineApprovalConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardSetRequestUpdateCommandVisibilityConfig struct {
	// List of specified visible users. This parameter takes effect only when the visibility type is set to SPECIFIED.
	SpecifiedUserList []*string `json:"SpecifiedUserList,omitempty" xml:"SpecifiedUserList,omitempty" type:"Repeated"`
	// Visibility type. Valid values: PUBLIC (public access), PRIVATE (private access, visible only to standard set members and administrators), and SPECIFIED (visible to specified users).
	//
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateStandardSetRequestUpdateCommandVisibilityConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetRequestUpdateCommandVisibilityConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetRequestUpdateCommandVisibilityConfig) GetSpecifiedUserList() []*string {
	return s.SpecifiedUserList
}

func (s *UpdateStandardSetRequestUpdateCommandVisibilityConfig) GetType() *string {
	return s.Type
}

func (s *UpdateStandardSetRequestUpdateCommandVisibilityConfig) SetSpecifiedUserList(v []*string) *UpdateStandardSetRequestUpdateCommandVisibilityConfig {
	s.SpecifiedUserList = v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandVisibilityConfig) SetType(v string) *UpdateStandardSetRequestUpdateCommandVisibilityConfig {
	s.Type = &v
	return s
}

func (s *UpdateStandardSetRequestUpdateCommandVisibilityConfig) Validate() error {
	return dara.Validate(s)
}

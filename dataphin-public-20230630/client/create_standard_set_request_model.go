// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateStandardSetRequestCreateCommand) *CreateStandardSetRequest
	GetCreateCommand() *CreateStandardSetRequestCreateCommand
	SetOpTenantId(v int64) *CreateStandardSetRequest
	GetOpTenantId() *int64
}

type CreateStandardSetRequest struct {
	// This parameter is required.
	CreateCommand *CreateStandardSetRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardSetRequest) GetCreateCommand() *CreateStandardSetRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateStandardSetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardSetRequest) SetCreateCommand(v *CreateStandardSetRequestCreateCommand) *CreateStandardSetRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateStandardSetRequest) SetOpTenantId(v int64) *CreateStandardSetRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardSetRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardSetRequestCreateCommand struct {
	ApprovalConfig *CreateStandardSetRequestCreateCommandApprovalConfig `json:"ApprovalConfig,omitempty" xml:"ApprovalConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// CITY
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1001
	DefaultStandardTemplateId *int64 `json:"DefaultStandardTemplateId,omitempty" xml:"DefaultStandardTemplateId,omitempty"`
	// example:
	//
	// test
	Description        *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryReference *CreateStandardSetRequestCreateCommandDirectoryReference `json:"DirectoryReference,omitempty" xml:"DirectoryReference,omitempty" type:"Struct"`
	MaintainerList     []*string                                                `json:"MaintainerList,omitempty" xml:"MaintainerList,omitempty" type:"Repeated"`
	MemberGroupList    []*string                                                `json:"MemberGroupList,omitempty" xml:"MemberGroupList,omitempty" type:"Repeated"`
	MemberList         []*string                                                `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name                  *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	OfflineApprovalConfig *CreateStandardSetRequestCreateCommandOfflineApprovalConfig `json:"OfflineApprovalConfig,omitempty" xml:"OfflineApprovalConfig,omitempty" type:"Struct"`
	VisibilityConfig      *CreateStandardSetRequestCreateCommandVisibilityConfig      `json:"VisibilityConfig,omitempty" xml:"VisibilityConfig,omitempty" type:"Struct"`
}

func (s CreateStandardSetRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateStandardSetRequestCreateCommand) GetApprovalConfig() *CreateStandardSetRequestCreateCommandApprovalConfig {
	return s.ApprovalConfig
}

func (s *CreateStandardSetRequestCreateCommand) GetCode() *string {
	return s.Code
}

func (s *CreateStandardSetRequestCreateCommand) GetDefaultStandardTemplateId() *int64 {
	return s.DefaultStandardTemplateId
}

func (s *CreateStandardSetRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardSetRequestCreateCommand) GetDirectoryReference() *CreateStandardSetRequestCreateCommandDirectoryReference {
	return s.DirectoryReference
}

func (s *CreateStandardSetRequestCreateCommand) GetMaintainerList() []*string {
	return s.MaintainerList
}

func (s *CreateStandardSetRequestCreateCommand) GetMemberGroupList() []*string {
	return s.MemberGroupList
}

func (s *CreateStandardSetRequestCreateCommand) GetMemberList() []*string {
	return s.MemberList
}

func (s *CreateStandardSetRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateStandardSetRequestCreateCommand) GetOfflineApprovalConfig() *CreateStandardSetRequestCreateCommandOfflineApprovalConfig {
	return s.OfflineApprovalConfig
}

func (s *CreateStandardSetRequestCreateCommand) GetVisibilityConfig() *CreateStandardSetRequestCreateCommandVisibilityConfig {
	return s.VisibilityConfig
}

func (s *CreateStandardSetRequestCreateCommand) SetApprovalConfig(v *CreateStandardSetRequestCreateCommandApprovalConfig) *CreateStandardSetRequestCreateCommand {
	s.ApprovalConfig = v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetCode(v string) *CreateStandardSetRequestCreateCommand {
	s.Code = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetDefaultStandardTemplateId(v int64) *CreateStandardSetRequestCreateCommand {
	s.DefaultStandardTemplateId = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetDescription(v string) *CreateStandardSetRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetDirectoryReference(v *CreateStandardSetRequestCreateCommandDirectoryReference) *CreateStandardSetRequestCreateCommand {
	s.DirectoryReference = v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetMaintainerList(v []*string) *CreateStandardSetRequestCreateCommand {
	s.MaintainerList = v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetMemberGroupList(v []*string) *CreateStandardSetRequestCreateCommand {
	s.MemberGroupList = v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetMemberList(v []*string) *CreateStandardSetRequestCreateCommand {
	s.MemberList = v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetName(v string) *CreateStandardSetRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetOfflineApprovalConfig(v *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) *CreateStandardSetRequestCreateCommand {
	s.OfflineApprovalConfig = v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) SetVisibilityConfig(v *CreateStandardSetRequestCreateCommandVisibilityConfig) *CreateStandardSetRequestCreateCommand {
	s.VisibilityConfig = v
	return s
}

func (s *CreateStandardSetRequestCreateCommand) Validate() error {
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

type CreateStandardSetRequestCreateCommandApprovalConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// BY_DEFAULT
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// This parameter is required.
	EnableApproval *bool `json:"EnableApproval,omitempty" xml:"EnableApproval,omitempty"`
	// This parameter is required.
	IsSubmitInBatch *bool `json:"IsSubmitInBatch,omitempty" xml:"IsSubmitInBatch,omitempty"`
	// example:
	//
	// 1121
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateStandardSetRequestCreateCommandApprovalConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetRequestCreateCommandApprovalConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) GetEnableApproval() *bool {
	return s.EnableApproval
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) GetIsSubmitInBatch() *bool {
	return s.IsSubmitInBatch
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) SetApprovalType(v string) *CreateStandardSetRequestCreateCommandApprovalConfig {
	s.ApprovalType = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) SetEnableApproval(v bool) *CreateStandardSetRequestCreateCommandApprovalConfig {
	s.EnableApproval = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) SetIsSubmitInBatch(v bool) *CreateStandardSetRequestCreateCommandApprovalConfig {
	s.IsSubmitInBatch = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) SetTemplateId(v int64) *CreateStandardSetRequestCreateCommandApprovalConfig {
	s.TemplateId = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandApprovalConfig) Validate() error {
	return dara.Validate(s)
}

type CreateStandardSetRequestCreateCommandDirectoryReference struct {
	// This parameter is required.
	//
	// example:
	//
	// /dir1
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
}

func (s CreateStandardSetRequestCreateCommandDirectoryReference) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetRequestCreateCommandDirectoryReference) GoString() string {
	return s.String()
}

func (s *CreateStandardSetRequestCreateCommandDirectoryReference) GetDirectory() *string {
	return s.Directory
}

func (s *CreateStandardSetRequestCreateCommandDirectoryReference) SetDirectory(v string) *CreateStandardSetRequestCreateCommandDirectoryReference {
	s.Directory = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandDirectoryReference) Validate() error {
	return dara.Validate(s)
}

type CreateStandardSetRequestCreateCommandOfflineApprovalConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// BY_DEFAULT
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// This parameter is required.
	EnableApproval *bool `json:"EnableApproval,omitempty" xml:"EnableApproval,omitempty"`
	// This parameter is required.
	IsSubmitInBatch *bool `json:"IsSubmitInBatch,omitempty" xml:"IsSubmitInBatch,omitempty"`
	// example:
	//
	// 1121
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateStandardSetRequestCreateCommandOfflineApprovalConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetRequestCreateCommandOfflineApprovalConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) GetEnableApproval() *bool {
	return s.EnableApproval
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) GetIsSubmitInBatch() *bool {
	return s.IsSubmitInBatch
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) SetApprovalType(v string) *CreateStandardSetRequestCreateCommandOfflineApprovalConfig {
	s.ApprovalType = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) SetEnableApproval(v bool) *CreateStandardSetRequestCreateCommandOfflineApprovalConfig {
	s.EnableApproval = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) SetIsSubmitInBatch(v bool) *CreateStandardSetRequestCreateCommandOfflineApprovalConfig {
	s.IsSubmitInBatch = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) SetTemplateId(v int64) *CreateStandardSetRequestCreateCommandOfflineApprovalConfig {
	s.TemplateId = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandOfflineApprovalConfig) Validate() error {
	return dara.Validate(s)
}

type CreateStandardSetRequestCreateCommandVisibilityConfig struct {
	SpecifiedUserList []*string `json:"SpecifiedUserList,omitempty" xml:"SpecifiedUserList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateStandardSetRequestCreateCommandVisibilityConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetRequestCreateCommandVisibilityConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardSetRequestCreateCommandVisibilityConfig) GetSpecifiedUserList() []*string {
	return s.SpecifiedUserList
}

func (s *CreateStandardSetRequestCreateCommandVisibilityConfig) GetType() *string {
	return s.Type
}

func (s *CreateStandardSetRequestCreateCommandVisibilityConfig) SetSpecifiedUserList(v []*string) *CreateStandardSetRequestCreateCommandVisibilityConfig {
	s.SpecifiedUserList = v
	return s
}

func (s *CreateStandardSetRequestCreateCommandVisibilityConfig) SetType(v string) *CreateStandardSetRequestCreateCommandVisibilityConfig {
	s.Type = &v
	return s
}

func (s *CreateStandardSetRequestCreateCommandVisibilityConfig) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateUdfRequestCreateCommand) *CreateUdfRequest
	GetCreateCommand() *CreateUdfRequestCreateCommand
	SetOpTenantId(v int64) *CreateUdfRequest
	GetOpTenantId() *int64
}

type CreateUdfRequest struct {
	// This parameter is required.
	CreateCommand *CreateUdfRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateUdfRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfRequest) GoString() string {
	return s.String()
}

func (s *CreateUdfRequest) GetCreateCommand() *CreateUdfRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateUdfRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateUdfRequest) SetCreateCommand(v *CreateUdfRequestCreateCommand) *CreateUdfRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateUdfRequest) SetOpTenantId(v int64) *CreateUdfRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateUdfRequest) Validate() error {
	return dara.Validate(s)
}

type CreateUdfRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// com.lydaas.SleepTest
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// udf_sleep(100)
	CommandHelp *string `json:"CommandHelp,omitempty" xml:"CommandHelp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	ComputeEngineType *string `json:"ComputeEngineType,omitempty" xml:"ComputeEngineType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my_udf
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 711833
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RefResourceIdList []*int64 `json:"RefResourceIdList,omitempty" xml:"RefResourceIdList,omitempty" type:"Repeated"`
}

func (s CreateUdfRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateUdfRequestCreateCommand) GetCategory() *int32 {
	return s.Category
}

func (s *CreateUdfRequestCreateCommand) GetClassName() *string {
	return s.ClassName
}

func (s *CreateUdfRequestCreateCommand) GetCommandHelp() *string {
	return s.CommandHelp
}

func (s *CreateUdfRequestCreateCommand) GetComment() *string {
	return s.Comment
}

func (s *CreateUdfRequestCreateCommand) GetComputeEngineType() *string {
	return s.ComputeEngineType
}

func (s *CreateUdfRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateUdfRequestCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateUdfRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateUdfRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateUdfRequestCreateCommand) GetRefResourceIdList() []*int64 {
	return s.RefResourceIdList
}

func (s *CreateUdfRequestCreateCommand) SetCategory(v int32) *CreateUdfRequestCreateCommand {
	s.Category = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetClassName(v string) *CreateUdfRequestCreateCommand {
	s.ClassName = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetCommandHelp(v string) *CreateUdfRequestCreateCommand {
	s.CommandHelp = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetComment(v string) *CreateUdfRequestCreateCommand {
	s.Comment = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetComputeEngineType(v string) *CreateUdfRequestCreateCommand {
	s.ComputeEngineType = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetDescription(v string) *CreateUdfRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetDirectory(v string) *CreateUdfRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetName(v string) *CreateUdfRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetProjectId(v int64) *CreateUdfRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateUdfRequestCreateCommand) SetRefResourceIdList(v []*int64) *CreateUdfRequestCreateCommand {
	s.RefResourceIdList = v
	return s
}

func (s *CreateUdfRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateUdfRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateUdfRequestUpdateCommand) *UpdateUdfRequest
	GetUpdateCommand() *UpdateUdfRequestUpdateCommand
}

type UpdateUdfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateUdfRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateUdfRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfRequest) GoString() string {
	return s.String()
}

func (s *UpdateUdfRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateUdfRequest) GetUpdateCommand() *UpdateUdfRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateUdfRequest) SetOpTenantId(v int64) *UpdateUdfRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateUdfRequest) SetUpdateCommand(v *UpdateUdfRequestUpdateCommand) *UpdateUdfRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateUdfRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUdfRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
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
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 711833
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RefResourceIdList []*int64 `json:"RefResourceIdList,omitempty" xml:"RefResourceIdList,omitempty" type:"Repeated"`
}

func (s UpdateUdfRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateUdfRequestUpdateCommand) GetCategory() *int32 {
	return s.Category
}

func (s *UpdateUdfRequestUpdateCommand) GetClassName() *string {
	return s.ClassName
}

func (s *UpdateUdfRequestUpdateCommand) GetCommandHelp() *string {
	return s.CommandHelp
}

func (s *UpdateUdfRequestUpdateCommand) GetComment() *string {
	return s.Comment
}

func (s *UpdateUdfRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateUdfRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateUdfRequestUpdateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateUdfRequestUpdateCommand) GetRefResourceIdList() []*int64 {
	return s.RefResourceIdList
}

func (s *UpdateUdfRequestUpdateCommand) SetCategory(v int32) *UpdateUdfRequestUpdateCommand {
	s.Category = &v
	return s
}

func (s *UpdateUdfRequestUpdateCommand) SetClassName(v string) *UpdateUdfRequestUpdateCommand {
	s.ClassName = &v
	return s
}

func (s *UpdateUdfRequestUpdateCommand) SetCommandHelp(v string) *UpdateUdfRequestUpdateCommand {
	s.CommandHelp = &v
	return s
}

func (s *UpdateUdfRequestUpdateCommand) SetComment(v string) *UpdateUdfRequestUpdateCommand {
	s.Comment = &v
	return s
}

func (s *UpdateUdfRequestUpdateCommand) SetDescription(v string) *UpdateUdfRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateUdfRequestUpdateCommand) SetId(v int64) *UpdateUdfRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateUdfRequestUpdateCommand) SetProjectId(v int64) *UpdateUdfRequestUpdateCommand {
	s.ProjectId = &v
	return s
}

func (s *UpdateUdfRequestUpdateCommand) SetRefResourceIdList(v []*int64) *UpdateUdfRequestUpdateCommand {
	s.RefResourceIdList = v
	return s
}

func (s *UpdateUdfRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}

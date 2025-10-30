// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTaskUdfLineagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBatchTaskUdfLineagesRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateBatchTaskUdfLineagesRequestUpdateCommand) *UpdateBatchTaskUdfLineagesRequest
	GetUpdateCommand() *UpdateBatchTaskUdfLineagesRequestUpdateCommand
}

type UpdateBatchTaskUdfLineagesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateBatchTaskUdfLineagesRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateBatchTaskUdfLineagesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskUdfLineagesRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskUdfLineagesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBatchTaskUdfLineagesRequest) GetUpdateCommand() *UpdateBatchTaskUdfLineagesRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateBatchTaskUdfLineagesRequest) SetOpTenantId(v int64) *UpdateBatchTaskUdfLineagesRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequest) SetUpdateCommand(v *UpdateBatchTaskUdfLineagesRequestUpdateCommand) *UpdateBatchTaskUdfLineagesRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBatchTaskUdfLineagesRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	LineageGroupList []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList `json:"LineageGroupList,omitempty" xml:"LineageGroupList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateBatchTaskUdfLineagesRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskUdfLineagesRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommand) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommand) GetLineageGroupList() []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList {
	return s.LineageGroupList
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommand) SetFileId(v int64) *UpdateBatchTaskUdfLineagesRequestUpdateCommand {
	s.FileId = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommand) SetLineageGroupList(v []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList) *UpdateBatchTaskUdfLineagesRequestUpdateCommand {
	s.LineageGroupList = v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommand) SetProjectId(v int64) *UpdateBatchTaskUdfLineagesRequestUpdateCommand {
	s.ProjectId = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommand) Validate() error {
	if s.LineageGroupList != nil {
		for _, item := range s.LineageGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList struct {
	// This parameter is required.
	InputLineageList []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList `json:"InputLineageList,omitempty" xml:"InputLineageList,omitempty" type:"Repeated"`
	// This parameter is required.
	OutputLineageList []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList `json:"OutputLineageList,omitempty" xml:"OutputLineageList,omitempty" type:"Repeated"`
}

func (s UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList) GetInputLineageList() []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList {
	return s.InputLineageList
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList) GetOutputLineageList() []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList {
	return s.OutputLineageList
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList) SetInputLineageList(v []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList {
	s.InputLineageList = v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList) SetOutputLineageList(v []*UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList {
	s.OutputLineageList = v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupList) Validate() error {
	if s.InputLineageList != nil {
		for _, item := range s.InputLineageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputLineageList != nil {
		for _, item := range s.OutputLineageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList struct {
	// This parameter is required.
	ColumnList []*string `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// dev
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	FullTable *bool `json:"FullTable,omitempty" xml:"FullTable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) GetColumnList() []*string {
	return s.ColumnList
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) GetEnv() *string {
	return s.Env
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) GetFullTable() *bool {
	return s.FullTable
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) GetName() *string {
	return s.Name
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) SetColumnList(v []*string) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList {
	s.ColumnList = v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) SetEnv(v string) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList {
	s.Env = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) SetFullTable(v bool) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList {
	s.FullTable = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) SetName(v string) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList {
	s.Name = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListInputLineageList) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList struct {
	// This parameter is required.
	ColumnList []*string `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// dev
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	FullTable *bool `json:"FullTable,omitempty" xml:"FullTable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) GetColumnList() []*string {
	return s.ColumnList
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) GetEnv() *string {
	return s.Env
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) GetFullTable() *bool {
	return s.FullTable
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) GetName() *string {
	return s.Name
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) SetColumnList(v []*string) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList {
	s.ColumnList = v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) SetEnv(v string) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList {
	s.Env = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) SetFullTable(v bool) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList {
	s.FullTable = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) SetName(v string) *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList {
	s.Name = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesRequestUpdateCommandLineageGroupListOutputLineageList) Validate() error {
	return dara.Validate(s)
}

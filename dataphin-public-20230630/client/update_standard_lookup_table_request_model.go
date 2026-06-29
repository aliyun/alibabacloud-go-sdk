// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardLookupTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardLookupTableRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateStandardLookupTableRequestUpdateCommand) *UpdateStandardLookupTableRequest
	GetUpdateCommand() *UpdateStandardLookupTableRequestUpdateCommand
}

type UpdateStandardLookupTableRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateStandardLookupTableRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateStandardLookupTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardLookupTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardLookupTableRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardLookupTableRequest) GetUpdateCommand() *UpdateStandardLookupTableRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateStandardLookupTableRequest) SetOpTenantId(v int64) *UpdateStandardLookupTableRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardLookupTableRequest) SetUpdateCommand(v *UpdateStandardLookupTableRequestUpdateCommand) *UpdateStandardLookupTableRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateStandardLookupTableRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardLookupTableRequestUpdateCommand struct {
	// The lookup table code.
	//
	// This parameter is required.
	//
	// example:
	//
	// CITY
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The lookup table description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory to which the lookup table belongs.
	DirectoryReference *UpdateStandardLookupTableRequestUpdateCommandDirectoryReference `json:"DirectoryReference,omitempty" xml:"DirectoryReference,omitempty" type:"Struct"`
	// The lookup table ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of lookup table values.
	LookupTableValueList []*UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList `json:"LookupTableValueList,omitempty" xml:"LookupTableValueList,omitempty" type:"Repeated"`
	// The lookup table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 城市码表
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the lookup table owner. Default value: the user ID of the caller.
	//
	// example:
	//
	// 30012021
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s UpdateStandardLookupTableRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardLookupTableRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) GetDirectoryReference() *UpdateStandardLookupTableRequestUpdateCommandDirectoryReference {
	return s.DirectoryReference
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) GetLookupTableValueList() []*UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList {
	return s.LookupTableValueList
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) GetOwner() *string {
	return s.Owner
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) SetCode(v string) *UpdateStandardLookupTableRequestUpdateCommand {
	s.Code = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) SetDescription(v string) *UpdateStandardLookupTableRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) SetDirectoryReference(v *UpdateStandardLookupTableRequestUpdateCommandDirectoryReference) *UpdateStandardLookupTableRequestUpdateCommand {
	s.DirectoryReference = v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) SetId(v int64) *UpdateStandardLookupTableRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) SetLookupTableValueList(v []*UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) *UpdateStandardLookupTableRequestUpdateCommand {
	s.LookupTableValueList = v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) SetName(v string) *UpdateStandardLookupTableRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) SetOwner(v string) *UpdateStandardLookupTableRequestUpdateCommand {
	s.Owner = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommand) Validate() error {
	if s.DirectoryReference != nil {
		if err := s.DirectoryReference.Validate(); err != nil {
			return err
		}
	}
	if s.LookupTableValueList != nil {
		for _, item := range s.LookupTableValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateStandardLookupTableRequestUpdateCommandDirectoryReference struct {
	// The directory to which the lookup table belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// /dir1/dir2
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
}

func (s UpdateStandardLookupTableRequestUpdateCommandDirectoryReference) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardLookupTableRequestUpdateCommandDirectoryReference) GoString() string {
	return s.String()
}

func (s *UpdateStandardLookupTableRequestUpdateCommandDirectoryReference) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateStandardLookupTableRequestUpdateCommandDirectoryReference) SetDirectory(v string) *UpdateStandardLookupTableRequestUpdateCommandDirectoryReference {
	s.Directory = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommandDirectoryReference) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList struct {
	// The code description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The English name of the code.
	//
	// example:
	//
	// HZ
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	// The code name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The code value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 杭州
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) GoString() string {
	return s.String()
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) GetEnglishName() *string {
	return s.EnglishName
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) GetName() *string {
	return s.Name
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) GetValue() *string {
	return s.Value
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) SetDescription(v string) *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList {
	s.Description = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) SetEnglishName(v string) *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList {
	s.EnglishName = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) SetName(v string) *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList {
	s.Name = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) SetValue(v string) *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList {
	s.Value = &v
	return s
}

func (s *UpdateStandardLookupTableRequestUpdateCommandLookupTableValueList) Validate() error {
	return dara.Validate(s)
}

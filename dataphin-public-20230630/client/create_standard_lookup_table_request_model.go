// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardLookupTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateStandardLookupTableRequestCreateCommand) *CreateStandardLookupTableRequest
	GetCreateCommand() *CreateStandardLookupTableRequestCreateCommand
	SetOpTenantId(v int64) *CreateStandardLookupTableRequest
	GetOpTenantId() *int64
}

type CreateStandardLookupTableRequest struct {
	// This parameter is required.
	CreateCommand *CreateStandardLookupTableRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardLookupTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardLookupTableRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardLookupTableRequest) GetCreateCommand() *CreateStandardLookupTableRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateStandardLookupTableRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardLookupTableRequest) SetCreateCommand(v *CreateStandardLookupTableRequestCreateCommand) *CreateStandardLookupTableRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateStandardLookupTableRequest) SetOpTenantId(v int64) *CreateStandardLookupTableRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardLookupTableRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardLookupTableRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// CITY
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test
	Description          *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryReference   *CreateStandardLookupTableRequestCreateCommandDirectoryReference     `json:"DirectoryReference,omitempty" xml:"DirectoryReference,omitempty" type:"Struct"`
	LookupTableValueList []*CreateStandardLookupTableRequestCreateCommandLookupTableValueList `json:"LookupTableValueList,omitempty" xml:"LookupTableValueList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 城市码表
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012021
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s CreateStandardLookupTableRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardLookupTableRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateStandardLookupTableRequestCreateCommand) GetCode() *string {
	return s.Code
}

func (s *CreateStandardLookupTableRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardLookupTableRequestCreateCommand) GetDirectoryReference() *CreateStandardLookupTableRequestCreateCommandDirectoryReference {
	return s.DirectoryReference
}

func (s *CreateStandardLookupTableRequestCreateCommand) GetLookupTableValueList() []*CreateStandardLookupTableRequestCreateCommandLookupTableValueList {
	return s.LookupTableValueList
}

func (s *CreateStandardLookupTableRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateStandardLookupTableRequestCreateCommand) GetOwner() *string {
	return s.Owner
}

func (s *CreateStandardLookupTableRequestCreateCommand) SetCode(v string) *CreateStandardLookupTableRequestCreateCommand {
	s.Code = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommand) SetDescription(v string) *CreateStandardLookupTableRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommand) SetDirectoryReference(v *CreateStandardLookupTableRequestCreateCommandDirectoryReference) *CreateStandardLookupTableRequestCreateCommand {
	s.DirectoryReference = v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommand) SetLookupTableValueList(v []*CreateStandardLookupTableRequestCreateCommandLookupTableValueList) *CreateStandardLookupTableRequestCreateCommand {
	s.LookupTableValueList = v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommand) SetName(v string) *CreateStandardLookupTableRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommand) SetOwner(v string) *CreateStandardLookupTableRequestCreateCommand {
	s.Owner = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommand) Validate() error {
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

type CreateStandardLookupTableRequestCreateCommandDirectoryReference struct {
	// This parameter is required.
	//
	// example:
	//
	// /dir1/dir2
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
}

func (s CreateStandardLookupTableRequestCreateCommandDirectoryReference) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardLookupTableRequestCreateCommandDirectoryReference) GoString() string {
	return s.String()
}

func (s *CreateStandardLookupTableRequestCreateCommandDirectoryReference) GetDirectory() *string {
	return s.Directory
}

func (s *CreateStandardLookupTableRequestCreateCommandDirectoryReference) SetDirectory(v string) *CreateStandardLookupTableRequestCreateCommandDirectoryReference {
	s.Directory = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommandDirectoryReference) Validate() error {
	return dara.Validate(s)
}

type CreateStandardLookupTableRequestCreateCommandLookupTableValueList struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// HZ
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 杭州
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStandardLookupTableRequestCreateCommandLookupTableValueList) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardLookupTableRequestCreateCommandLookupTableValueList) GoString() string {
	return s.String()
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) GetEnglishName() *string {
	return s.EnglishName
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) GetName() *string {
	return s.Name
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) GetValue() *string {
	return s.Value
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) SetDescription(v string) *CreateStandardLookupTableRequestCreateCommandLookupTableValueList {
	s.Description = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) SetEnglishName(v string) *CreateStandardLookupTableRequestCreateCommandLookupTableValueList {
	s.EnglishName = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) SetName(v string) *CreateStandardLookupTableRequestCreateCommandLookupTableValueList {
	s.Name = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) SetValue(v string) *CreateStandardLookupTableRequestCreateCommandLookupTableValueList {
	s.Value = &v
	return s
}

func (s *CreateStandardLookupTableRequestCreateCommandLookupTableValueList) Validate() error {
	return dara.Validate(s)
}

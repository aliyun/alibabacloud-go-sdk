// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityTemplateRequest
	GetOpTenantId() *int64
	SetUpsertCommand(v *UpsertQualityTemplateRequestUpsertCommand) *UpsertQualityTemplateRequest
	GetUpsertCommand() *UpsertQualityTemplateRequestUpsertCommand
}

type UpsertQualityTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommand *UpsertQualityTemplateRequestUpsertCommand `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty" type:"Struct"`
}

func (s UpsertQualityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityTemplateRequest) GetUpsertCommand() *UpsertQualityTemplateRequestUpsertCommand {
	return s.UpsertCommand
}

func (s *UpsertQualityTemplateRequest) SetOpTenantId(v int64) *UpsertQualityTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityTemplateRequest) SetUpsertCommand(v *UpsertQualityTemplateRequestUpsertCommand) *UpsertQualityTemplateRequest {
	s.UpsertCommand = v
	return s
}

func (s *UpsertQualityTemplateRequest) Validate() error {
	if s.UpsertCommand != nil {
		if err := s.UpsertCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityTemplateRequestUpsertCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// CONSISTENT
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// test
	Description      *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	FormPropertyList []*UpsertQualityTemplateRequestUpsertCommandFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30012011
	Owner                     *string   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SupportDataSourceTypeList []*string `json:"SupportDataSourceTypeList,omitempty" xml:"SupportDataSourceTypeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpsertQualityTemplateRequestUpsertCommand) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateRequestUpsertCommand) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetCatalog() *string {
	return s.Catalog
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetDescription() *string {
	return s.Description
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetFormPropertyList() []*UpsertQualityTemplateRequestUpsertCommandFormPropertyList {
	return s.FormPropertyList
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetId() *int64 {
	return s.Id
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetName() *string {
	return s.Name
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetOwner() *string {
	return s.Owner
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetSupportDataSourceTypeList() []*string {
	return s.SupportDataSourceTypeList
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetType() *string {
	return s.Type
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetCatalog(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Catalog = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetDescription(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Description = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetFormPropertyList(v []*UpsertQualityTemplateRequestUpsertCommandFormPropertyList) *UpsertQualityTemplateRequestUpsertCommand {
	s.FormPropertyList = v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetId(v int64) *UpsertQualityTemplateRequestUpsertCommand {
	s.Id = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetName(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Name = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetOwner(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Owner = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetSupportDataSourceTypeList(v []*string) *UpsertQualityTemplateRequestUpsertCommand {
	s.SupportDataSourceTypeList = v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetType(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Type = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) Validate() error {
	if s.FormPropertyList != nil {
		for _, item := range s.FormPropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpsertQualityTemplateRequestUpsertCommandFormPropertyList struct {
	// example:
	//
	// expression
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// col
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpsertQualityTemplateRequestUpsertCommandFormPropertyList) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateRequestUpsertCommandFormPropertyList) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) GetComponentType() *string {
	return s.ComponentType
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) GetName() *string {
	return s.Name
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) GetValue() *string {
	return s.Value
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) SetComponentType(v string) *UpsertQualityTemplateRequestUpsertCommandFormPropertyList {
	s.ComponentType = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) SetName(v string) *UpsertQualityTemplateRequestUpsertCommandFormPropertyList {
	s.Name = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) SetValue(v string) *UpsertQualityTemplateRequestUpsertCommandFormPropertyList {
	s.Value = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) Validate() error {
	return dara.Validate(s)
}

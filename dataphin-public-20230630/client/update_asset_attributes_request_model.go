// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateAssetAttributesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateAssetAttributesRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateAssetAttributesRequestUpdateCommand) *UpdateAssetAttributesRequest
	GetUpdateCommand() *UpdateAssetAttributesRequestUpdateCommand
}

type UpdateAssetAttributesRequest struct {
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
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateAssetAttributesRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateAssetAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateAssetAttributesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateAssetAttributesRequest) GetUpdateCommand() *UpdateAssetAttributesRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateAssetAttributesRequest) SetOpTenantId(v int64) *UpdateAssetAttributesRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateAssetAttributesRequest) SetOpUserId(v string) *UpdateAssetAttributesRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateAssetAttributesRequest) SetUpdateCommand(v *UpdateAssetAttributesRequestUpdateCommand) *UpdateAssetAttributesRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateAssetAttributesRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAssetAttributesRequestUpdateCommand struct {
	// The list of asset property updates. A maximum of 50 entries can be specified in a single request.
	//
	// This parameter is required.
	AssetAttributeUpdateList []*UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList `json:"AssetAttributeUpdateList,omitempty" xml:"AssetAttributeUpdateList,omitempty" type:"Repeated"`
}

func (s UpdateAssetAttributesRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesRequestUpdateCommand) GetAssetAttributeUpdateList() []*UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList {
	return s.AssetAttributeUpdateList
}

func (s *UpdateAssetAttributesRequestUpdateCommand) SetAssetAttributeUpdateList(v []*UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList) *UpdateAssetAttributesRequestUpdateCommand {
	s.AssetAttributeUpdateList = v
	return s
}

func (s *UpdateAssetAttributesRequestUpdateCommand) Validate() error {
	if s.AssetAttributeUpdateList != nil {
		for _, item := range s.AssetAttributeUpdateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList struct {
	// The list of properties to update.
	//
	// This parameter is required.
	AttributeList []*UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Repeated"`
	// The globally unique identifier (GUID) of the asset. You can obtain this value by calling operations such as ListCatalogAssets and GetTableColumnByTableGuids.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.project_a.table_orders
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList) GetAttributeList() []*UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList {
	return s.AttributeList
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList) GetGuid() *string {
	return s.Guid
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList) SetAttributeList(v []*UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList) *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList {
	s.AttributeList = v
	return s
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList) SetGuid(v string) *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList {
	s.Guid = &v
	return s
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateList) Validate() error {
	if s.AttributeList != nil {
		for _, item := range s.AttributeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList struct {
	// The property code. This value must match the AttributeCode returned by the GetAssetTypeAttributeCodes operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// data_level
	AttributeCode *string `json:"AttributeCode,omitempty" xml:"AttributeCode,omitempty"`
	// The list of property values. For a single-value property, pass one element. For a multi-value property, pass multiple elements. Pass an empty array [] to clear the property value.
	//
	// This parameter is required.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList) GetAttributeCode() *string {
	return s.AttributeCode
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList) GetValues() []*string {
	return s.Values
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList) SetAttributeCode(v string) *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList {
	s.AttributeCode = &v
	return s
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList) SetValues(v []*string) *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList {
	s.Values = v
	return s
}

func (s *UpdateAssetAttributesRequestUpdateCommandAssetAttributeUpdateListAttributeList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityOnResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCommand(v *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) *CheckDataSourceConnectivityOnResourceGroupRequest
	GetCheckCommand() *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand
	SetOpTenantId(v int64) *CheckDataSourceConnectivityOnResourceGroupRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CheckDataSourceConnectivityOnResourceGroupRequest
	GetOpUserId() *string
}

type CheckDataSourceConnectivityOnResourceGroupRequest struct {
	// The check object.
	//
	// This parameter is required.
	CheckCommand *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty" type:"Struct"`
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

func (s CheckDataSourceConnectivityOnResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityOnResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequest) GetCheckCommand() *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand {
	return s.CheckCommand
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequest) SetCheckCommand(v *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) *CheckDataSourceConnectivityOnResourceGroupRequest {
	s.CheckCommand = v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequest) SetOpTenantId(v int64) *CheckDataSourceConnectivityOnResourceGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequest) SetOpUserId(v string) *CheckDataSourceConnectivityOnResourceGroupRequest {
	s.OpUserId = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequest) Validate() error {
	if s.CheckCommand != nil {
		if err := s.CheckCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand struct {
	// The connection configuration items. If DataSourceId is specified, this parameter is optional and the saved configurations of the data source are used by default.
	ConfigItemList []*CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// The data source ID. This parameter is optional. If specified, the check result is associated with the data source and visible on the data source page, and the old check records of the data source on the target resource group are cleared first. If Type and ConfigItemList are not specified, the saved type and configurations of the data source are used for the check.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The ID of the schedule resource group. The connectivity check task is executed on this resource group.
	//
	// example:
	//
	// rg_269xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The data source type. If DataSourceId is specified, this parameter is optional and the saved type of the data source is used by default.
	//
	// example:
	//
	// MYSQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) GetConfigItemList() []*CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList {
	return s.ConfigItemList
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) GetType() *string {
	return s.Type
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) SetConfigItemList(v []*CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList) *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand {
	s.ConfigItemList = v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) SetDataSourceId(v string) *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand {
	s.DataSourceId = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) SetResourceGroupId(v string) *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) SetType(v string) *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand {
	s.Type = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommand) Validate() error {
	if s.ConfigItemList != nil {
		for _, item := range s.ConfigItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList struct {
	// The key of the configuration item.
	//
	// example:
	//
	// jdbc.url
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The configuration value.
	//
	// example:
	//
	// jdbc:mysql://host:port/database
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList) GetKey() *string {
	return s.Key
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList) GetValue() *string {
	return s.Value
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList) SetKey(v string) *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList {
	s.Key = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList) SetValue(v string) *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList {
	s.Value = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupRequestCheckCommandConfigItemList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataSourceConfigRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateDataSourceConfigRequestUpdateCommand) *UpdateDataSourceConfigRequest
	GetUpdateCommand() *UpdateDataSourceConfigRequestUpdateCommand
}

type UpdateDataSourceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateDataSourceConfigRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataSourceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataSourceConfigRequest) GetUpdateCommand() *UpdateDataSourceConfigRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDataSourceConfigRequest) SetOpTenantId(v int64) *UpdateDataSourceConfigRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataSourceConfigRequest) SetUpdateCommand(v *UpdateDataSourceConfigRequestUpdateCommand) *UpdateDataSourceConfigRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDataSourceConfigRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataSourceConfigRequestUpdateCommand struct {
	// This parameter is required.
	ConfigItemList []*UpdateDataSourceConfigRequestUpdateCommandConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 13231313
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateDataSourceConfigRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceConfigRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigRequestUpdateCommand) GetConfigItemList() []*UpdateDataSourceConfigRequestUpdateCommandConfigItemList {
	return s.ConfigItemList
}

func (s *UpdateDataSourceConfigRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataSourceConfigRequestUpdateCommand) SetConfigItemList(v []*UpdateDataSourceConfigRequestUpdateCommandConfigItemList) *UpdateDataSourceConfigRequestUpdateCommand {
	s.ConfigItemList = v
	return s
}

func (s *UpdateDataSourceConfigRequestUpdateCommand) SetId(v int64) *UpdateDataSourceConfigRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateDataSourceConfigRequestUpdateCommand) Validate() error {
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

type UpdateDataSourceConfigRequestUpdateCommandConfigItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDataSourceConfigRequestUpdateCommandConfigItemList) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceConfigRequestUpdateCommandConfigItemList) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigRequestUpdateCommandConfigItemList) GetKey() *string {
	return s.Key
}

func (s *UpdateDataSourceConfigRequestUpdateCommandConfigItemList) GetValue() *string {
	return s.Value
}

func (s *UpdateDataSourceConfigRequestUpdateCommandConfigItemList) SetKey(v string) *UpdateDataSourceConfigRequestUpdateCommandConfigItemList {
	s.Key = &v
	return s
}

func (s *UpdateDataSourceConfigRequestUpdateCommandConfigItemList) SetValue(v string) *UpdateDataSourceConfigRequestUpdateCommandConfigItemList {
	s.Value = &v
	return s
}

func (s *UpdateDataSourceConfigRequestUpdateCommandConfigItemList) Validate() error {
	return dara.Validate(s)
}

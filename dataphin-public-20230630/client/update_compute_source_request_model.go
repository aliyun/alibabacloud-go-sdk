// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateComputeSourceRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateComputeSourceRequestUpdateCommand) *UpdateComputeSourceRequest
	GetUpdateCommand() *UpdateComputeSourceRequestUpdateCommand
}

type UpdateComputeSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateComputeSourceRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateComputeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeSourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateComputeSourceRequest) GetUpdateCommand() *UpdateComputeSourceRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateComputeSourceRequest) SetOpTenantId(v int64) *UpdateComputeSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateComputeSourceRequest) SetUpdateCommand(v *UpdateComputeSourceRequestUpdateCommand) *UpdateComputeSourceRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateComputeSourceRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComputeSourceRequestUpdateCommand struct {
	// This parameter is required.
	ConfigList []*UpdateComputeSourceRequestUpdateCommandConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test1021
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MacCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateComputeSourceRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSourceRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateComputeSourceRequestUpdateCommand) GetConfigList() []*UpdateComputeSourceRequestUpdateCommandConfigList {
	return s.ConfigList
}

func (s *UpdateComputeSourceRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateComputeSourceRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateComputeSourceRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateComputeSourceRequestUpdateCommand) GetType() *string {
	return s.Type
}

func (s *UpdateComputeSourceRequestUpdateCommand) SetConfigList(v []*UpdateComputeSourceRequestUpdateCommandConfigList) *UpdateComputeSourceRequestUpdateCommand {
	s.ConfigList = v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommand) SetDescription(v string) *UpdateComputeSourceRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommand) SetId(v int64) *UpdateComputeSourceRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommand) SetName(v string) *UpdateComputeSourceRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommand) SetType(v string) *UpdateComputeSourceRequestUpdateCommand {
	s.Type = &v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommand) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateComputeSourceRequestUpdateCommandConfigList struct {
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

func (s UpdateComputeSourceRequestUpdateCommandConfigList) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSourceRequestUpdateCommandConfigList) GoString() string {
	return s.String()
}

func (s *UpdateComputeSourceRequestUpdateCommandConfigList) GetKey() *string {
	return s.Key
}

func (s *UpdateComputeSourceRequestUpdateCommandConfigList) GetValue() *string {
	return s.Value
}

func (s *UpdateComputeSourceRequestUpdateCommandConfigList) SetKey(v string) *UpdateComputeSourceRequestUpdateCommandConfigList {
	s.Key = &v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommandConfigList) SetValue(v string) *UpdateComputeSourceRequestUpdateCommandConfigList {
	s.Value = &v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommandConfigList) Validate() error {
	return dara.Validate(s)
}

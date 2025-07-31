// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateComputeSourceRequestCreateCommand) *CreateComputeSourceRequest
	GetCreateCommand() *CreateComputeSourceRequestCreateCommand
	SetOpTenantId(v int64) *CreateComputeSourceRequest
	GetOpTenantId() *int64
}

type CreateComputeSourceRequest struct {
	// This parameter is required.
	CreateCommand *CreateComputeSourceRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateComputeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeSourceRequest) GetCreateCommand() *CreateComputeSourceRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateComputeSourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateComputeSourceRequest) SetCreateCommand(v *CreateComputeSourceRequestCreateCommand) *CreateComputeSourceRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateComputeSourceRequest) SetOpTenantId(v int64) *CreateComputeSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateComputeSourceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateComputeSourceRequestCreateCommand struct {
	// This parameter is required.
	ConfigList []*CreateComputeSourceRequestCreateCommandConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test1011
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MacCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateComputeSourceRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeSourceRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateComputeSourceRequestCreateCommand) GetConfigList() []*CreateComputeSourceRequestCreateCommandConfigList {
	return s.ConfigList
}

func (s *CreateComputeSourceRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateComputeSourceRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateComputeSourceRequestCreateCommand) GetType() *string {
	return s.Type
}

func (s *CreateComputeSourceRequestCreateCommand) SetConfigList(v []*CreateComputeSourceRequestCreateCommandConfigList) *CreateComputeSourceRequestCreateCommand {
	s.ConfigList = v
	return s
}

func (s *CreateComputeSourceRequestCreateCommand) SetDescription(v string) *CreateComputeSourceRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateComputeSourceRequestCreateCommand) SetName(v string) *CreateComputeSourceRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateComputeSourceRequestCreateCommand) SetType(v string) *CreateComputeSourceRequestCreateCommand {
	s.Type = &v
	return s
}

func (s *CreateComputeSourceRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}

type CreateComputeSourceRequestCreateCommandConfigList struct {
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

func (s CreateComputeSourceRequestCreateCommandConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeSourceRequestCreateCommandConfigList) GoString() string {
	return s.String()
}

func (s *CreateComputeSourceRequestCreateCommandConfigList) GetKey() *string {
	return s.Key
}

func (s *CreateComputeSourceRequestCreateCommandConfigList) GetValue() *string {
	return s.Value
}

func (s *CreateComputeSourceRequestCreateCommandConfigList) SetKey(v string) *CreateComputeSourceRequestCreateCommandConfigList {
	s.Key = &v
	return s
}

func (s *CreateComputeSourceRequestCreateCommandConfigList) SetValue(v string) *CreateComputeSourceRequestCreateCommandConfigList {
	s.Value = &v
	return s
}

func (s *CreateComputeSourceRequestCreateCommandConfigList) Validate() error {
	return dara.Validate(s)
}

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
	SetOpUserId(v string) *CreateComputeSourceRequest
	GetOpUserId() *string
}

type CreateComputeSourceRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *CreateComputeSourceRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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

func (s *CreateComputeSourceRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateComputeSourceRequest) SetCreateCommand(v *CreateComputeSourceRequestCreateCommand) *CreateComputeSourceRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateComputeSourceRequest) SetOpTenantId(v int64) *CreateComputeSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateComputeSourceRequest) SetOpUserId(v string) *CreateComputeSourceRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateComputeSourceRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateComputeSourceRequestCreateCommand struct {
	// The ID of the associated cluster. This parameter takes effect only when CreateType is not specified or is set to COMPUTE_SOURCE, which creates a compute source that references a cluster. This parameter is mutually exclusive with CreateType=CLUSTER.
	//
	// example:
	//
	// 102311
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The connection configuration items.
	//
	// This parameter is required.
	ConfigList []*CreateComputeSourceRequestCreateCommandConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The type of entity to create. Valid values:
	//
	// - CLUSTER: Creates a cluster. ClusterId cannot be specified.
	//
	// - COMPUTE_SOURCE: Creates a compute source. This is the default value.
	//
	// example:
	//
	// CLUSTER
	CreateType *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the compute source.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1011
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the compute source.
	//
	// This parameter is required.
	//
	// example:
	//
	// MacCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the compute source type.
	//
	// example:
	//
	// CDH6
	TypeVersion *string `json:"TypeVersion,omitempty" xml:"TypeVersion,omitempty"`
}

func (s CreateComputeSourceRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeSourceRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateComputeSourceRequestCreateCommand) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *CreateComputeSourceRequestCreateCommand) GetConfigList() []*CreateComputeSourceRequestCreateCommandConfigList {
	return s.ConfigList
}

func (s *CreateComputeSourceRequestCreateCommand) GetCreateType() *string {
	return s.CreateType
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

func (s *CreateComputeSourceRequestCreateCommand) GetTypeVersion() *string {
	return s.TypeVersion
}

func (s *CreateComputeSourceRequestCreateCommand) SetClusterId(v int64) *CreateComputeSourceRequestCreateCommand {
	s.ClusterId = &v
	return s
}

func (s *CreateComputeSourceRequestCreateCommand) SetConfigList(v []*CreateComputeSourceRequestCreateCommandConfigList) *CreateComputeSourceRequestCreateCommand {
	s.ConfigList = v
	return s
}

func (s *CreateComputeSourceRequestCreateCommand) SetCreateType(v string) *CreateComputeSourceRequestCreateCommand {
	s.CreateType = &v
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

func (s *CreateComputeSourceRequestCreateCommand) SetTypeVersion(v string) *CreateComputeSourceRequestCreateCommand {
	s.TypeVersion = &v
	return s
}

func (s *CreateComputeSourceRequestCreateCommand) Validate() error {
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

type CreateComputeSourceRequestCreateCommandConfigList struct {
	// The configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the configuration item.
	//
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

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
	SetOpUserId(v string) *UpdateComputeSourceRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateComputeSourceRequestUpdateCommand) *UpdateComputeSourceRequest
	GetUpdateCommand() *UpdateComputeSourceRequestUpdateCommand
}

type UpdateComputeSourceRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The edit command.
	//
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

func (s *UpdateComputeSourceRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateComputeSourceRequest) GetUpdateCommand() *UpdateComputeSourceRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateComputeSourceRequest) SetOpTenantId(v int64) *UpdateComputeSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateComputeSourceRequest) SetOpUserId(v string) *UpdateComputeSourceRequest {
	s.OpUserId = &v
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
	// The associated cluster ID. This parameter takes effect only when CreateType is left empty or set to COMPUTE_SOURCE (to create a compute source that references a cluster). This parameter is mutually exclusive with CreateType=CLUSTER.
	//
	// example:
	//
	// 102311
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The connection configuration items.
	//
	// This parameter is required.
	ConfigList []*UpdateComputeSourceRequestUpdateCommandConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The creation entity type. Valid values:
	//
	// - CLUSTER: Creates the entity as a cluster. ClusterId cannot be specified.
	//
	// - COMPUTE_SOURCE: Creates the entity as a compute source. This is the default value.
	//
	// example:
	//
	// CLUSTER
	CreateType *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// The description.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The compute source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The compute source name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1021
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The compute source type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MacCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The compute source type version.
	//
	// example:
	//
	// CDH6
	TypeVersion *string `json:"TypeVersion,omitempty" xml:"TypeVersion,omitempty"`
}

func (s UpdateComputeSourceRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSourceRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateComputeSourceRequestUpdateCommand) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *UpdateComputeSourceRequestUpdateCommand) GetConfigList() []*UpdateComputeSourceRequestUpdateCommandConfigList {
	return s.ConfigList
}

func (s *UpdateComputeSourceRequestUpdateCommand) GetCreateType() *string {
	return s.CreateType
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

func (s *UpdateComputeSourceRequestUpdateCommand) GetTypeVersion() *string {
	return s.TypeVersion
}

func (s *UpdateComputeSourceRequestUpdateCommand) SetClusterId(v int64) *UpdateComputeSourceRequestUpdateCommand {
	s.ClusterId = &v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommand) SetConfigList(v []*UpdateComputeSourceRequestUpdateCommandConfigList) *UpdateComputeSourceRequestUpdateCommand {
	s.ConfigList = v
	return s
}

func (s *UpdateComputeSourceRequestUpdateCommand) SetCreateType(v string) *UpdateComputeSourceRequestUpdateCommand {
	s.CreateType = &v
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

func (s *UpdateComputeSourceRequestUpdateCommand) SetTypeVersion(v string) *UpdateComputeSourceRequestUpdateCommand {
	s.TypeVersion = &v
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
	// The configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The configuration item value.
	//
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

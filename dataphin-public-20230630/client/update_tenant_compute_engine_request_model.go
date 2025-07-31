// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantComputeEngineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateTenantComputeEngineRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateTenantComputeEngineRequestUpdateCommand) *UpdateTenantComputeEngineRequest
	GetUpdateCommand() *UpdateTenantComputeEngineRequestUpdateCommand
}

type UpdateTenantComputeEngineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateTenantComputeEngineRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateTenantComputeEngineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantComputeEngineRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantComputeEngineRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateTenantComputeEngineRequest) GetUpdateCommand() *UpdateTenantComputeEngineRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateTenantComputeEngineRequest) SetOpTenantId(v int64) *UpdateTenantComputeEngineRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateTenantComputeEngineRequest) SetUpdateCommand(v *UpdateTenantComputeEngineRequestUpdateCommand) *UpdateTenantComputeEngineRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateTenantComputeEngineRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTenantComputeEngineRequestUpdateCommand struct {
	// This parameter is required.
	ClusterUrlList []*string `json:"ClusterUrlList,omitempty" xml:"ClusterUrlList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MacCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateTenantComputeEngineRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantComputeEngineRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateTenantComputeEngineRequestUpdateCommand) GetClusterUrlList() []*string {
	return s.ClusterUrlList
}

func (s *UpdateTenantComputeEngineRequestUpdateCommand) GetType() *string {
	return s.Type
}

func (s *UpdateTenantComputeEngineRequestUpdateCommand) GetVersion() *string {
	return s.Version
}

func (s *UpdateTenantComputeEngineRequestUpdateCommand) SetClusterUrlList(v []*string) *UpdateTenantComputeEngineRequestUpdateCommand {
	s.ClusterUrlList = v
	return s
}

func (s *UpdateTenantComputeEngineRequestUpdateCommand) SetType(v string) *UpdateTenantComputeEngineRequestUpdateCommand {
	s.Type = &v
	return s
}

func (s *UpdateTenantComputeEngineRequestUpdateCommand) SetVersion(v string) *UpdateTenantComputeEngineRequestUpdateCommand {
	s.Version = &v
	return s
}

func (s *UpdateTenantComputeEngineRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}

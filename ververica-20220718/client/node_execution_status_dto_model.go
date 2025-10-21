// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeExecutionStatusDTO interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *NodeExecutionStatusDTO
	GetExecutionId() *string
	SetNamespace(v string) *NodeExecutionStatusDTO
	GetNamespace() *string
	SetResourceId(v string) *NodeExecutionStatusDTO
	GetResourceId() *string
	SetStatus(v string) *NodeExecutionStatusDTO
	GetStatus() *string
	SetType(v string) *NodeExecutionStatusDTO
	GetType() *string
	SetWorkspace(v string) *NodeExecutionStatusDTO
	GetWorkspace() *string
}

type NodeExecutionStatusDTO struct {
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
	Namespace   *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	ResourceId  *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
	Workspace   *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s NodeExecutionStatusDTO) String() string {
	return dara.Prettify(s)
}

func (s NodeExecutionStatusDTO) GoString() string {
	return s.String()
}

func (s *NodeExecutionStatusDTO) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *NodeExecutionStatusDTO) GetNamespace() *string {
	return s.Namespace
}

func (s *NodeExecutionStatusDTO) GetResourceId() *string {
	return s.ResourceId
}

func (s *NodeExecutionStatusDTO) GetStatus() *string {
	return s.Status
}

func (s *NodeExecutionStatusDTO) GetType() *string {
	return s.Type
}

func (s *NodeExecutionStatusDTO) GetWorkspace() *string {
	return s.Workspace
}

func (s *NodeExecutionStatusDTO) SetExecutionId(v string) *NodeExecutionStatusDTO {
	s.ExecutionId = &v
	return s
}

func (s *NodeExecutionStatusDTO) SetNamespace(v string) *NodeExecutionStatusDTO {
	s.Namespace = &v
	return s
}

func (s *NodeExecutionStatusDTO) SetResourceId(v string) *NodeExecutionStatusDTO {
	s.ResourceId = &v
	return s
}

func (s *NodeExecutionStatusDTO) SetStatus(v string) *NodeExecutionStatusDTO {
	s.Status = &v
	return s
}

func (s *NodeExecutionStatusDTO) SetType(v string) *NodeExecutionStatusDTO {
	s.Type = &v
	return s
}

func (s *NodeExecutionStatusDTO) SetWorkspace(v string) *NodeExecutionStatusDTO {
	s.Workspace = &v
	return s
}

func (s *NodeExecutionStatusDTO) Validate() error {
	return dara.Validate(s)
}

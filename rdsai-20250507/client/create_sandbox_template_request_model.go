// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSandboxTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultCpu(v string) *CreateSandboxTemplateRequest
	GetDefaultCpu() *string
	SetDefaultMemory(v string) *CreateSandboxTemplateRequest
	GetDefaultMemory() *string
	SetDescription(v string) *CreateSandboxTemplateRequest
	GetDescription() *string
	SetInstanceName(v string) *CreateSandboxTemplateRequest
	GetInstanceName() *string
	SetRegionId(v string) *CreateSandboxTemplateRequest
	GetRegionId() *string
	SetReplicas(v int64) *CreateSandboxTemplateRequest
	GetReplicas() *int64
	SetTemplateName(v string) *CreateSandboxTemplateRequest
	GetTemplateName() *string
}

type CreateSandboxTemplateRequest struct {
	// The number of CPUs for sandboxes created by using this template. Valid values: 1 to 4.
	//
	// example:
	//
	// 1
	DefaultCpu *string `json:"DefaultCpu,omitempty" xml:"DefaultCpu,omitempty"`
	// The memory size for sandboxes created by using this template. Unit: Gi. Valid values: 1Gi to 8Gi.
	//
	// example:
	//
	// 1Gi
	DefaultMemory *string `json:"DefaultMemory,omitempty" xml:"DefaultMemory,omitempty"`
	// The description of the sandbox template. The description must be unique within the VPC.
	//
	// example:
	//
	// code-interpreter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID of the AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of prewarmed sandboxes. Valid values: 1 to 1000.
	//
	// example:
	//
	// 1
	Replicas *int64 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The name of the sandbox template.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-interpreter
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateSandboxTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSandboxTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateSandboxTemplateRequest) GetDefaultCpu() *string {
	return s.DefaultCpu
}

func (s *CreateSandboxTemplateRequest) GetDefaultMemory() *string {
	return s.DefaultMemory
}

func (s *CreateSandboxTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSandboxTemplateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateSandboxTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSandboxTemplateRequest) GetReplicas() *int64 {
	return s.Replicas
}

func (s *CreateSandboxTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSandboxTemplateRequest) SetDefaultCpu(v string) *CreateSandboxTemplateRequest {
	s.DefaultCpu = &v
	return s
}

func (s *CreateSandboxTemplateRequest) SetDefaultMemory(v string) *CreateSandboxTemplateRequest {
	s.DefaultMemory = &v
	return s
}

func (s *CreateSandboxTemplateRequest) SetDescription(v string) *CreateSandboxTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateSandboxTemplateRequest) SetInstanceName(v string) *CreateSandboxTemplateRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateSandboxTemplateRequest) SetRegionId(v string) *CreateSandboxTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSandboxTemplateRequest) SetReplicas(v int64) *CreateSandboxTemplateRequest {
	s.Replicas = &v
	return s
}

func (s *CreateSandboxTemplateRequest) SetTemplateName(v string) *CreateSandboxTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSandboxTemplateRequest) Validate() error {
	return dara.Validate(s)
}

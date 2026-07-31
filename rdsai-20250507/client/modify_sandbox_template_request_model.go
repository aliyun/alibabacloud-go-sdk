// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySandboxTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultCpu(v string) *ModifySandboxTemplateRequest
	GetDefaultCpu() *string
	SetDefaultMemory(v string) *ModifySandboxTemplateRequest
	GetDefaultMemory() *string
	SetInstanceName(v string) *ModifySandboxTemplateRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifySandboxTemplateRequest
	GetRegionId() *string
	SetReplicas(v int64) *ModifySandboxTemplateRequest
	GetReplicas() *int64
	SetTemplateId(v string) *ModifySandboxTemplateRequest
	GetTemplateId() *string
}

type ModifySandboxTemplateRequest struct {
	// The number of CPUs for sandboxes created from this template. Valid values: 1 to 4.
	//
	// example:
	//
	// 1
	DefaultCpu *string `json:"DefaultCpu,omitempty" xml:"DefaultCpu,omitempty"`
	// The memory size for sandboxes created from this template. Unit: Gi. Valid values: 1Gi to 8Gi.
	//
	// example:
	//
	// 1Gi
	DefaultMemory *string `json:"DefaultMemory,omitempty" xml:"DefaultMemory,omitempty"`
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
	// The initial number of instances. Valid values: 1 to 1000.
	//
	// example:
	//
	// 2
	Replicas *int64 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The sandbox template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop-xxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifySandboxTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySandboxTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifySandboxTemplateRequest) GetDefaultCpu() *string {
	return s.DefaultCpu
}

func (s *ModifySandboxTemplateRequest) GetDefaultMemory() *string {
	return s.DefaultMemory
}

func (s *ModifySandboxTemplateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifySandboxTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySandboxTemplateRequest) GetReplicas() *int64 {
	return s.Replicas
}

func (s *ModifySandboxTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifySandboxTemplateRequest) SetDefaultCpu(v string) *ModifySandboxTemplateRequest {
	s.DefaultCpu = &v
	return s
}

func (s *ModifySandboxTemplateRequest) SetDefaultMemory(v string) *ModifySandboxTemplateRequest {
	s.DefaultMemory = &v
	return s
}

func (s *ModifySandboxTemplateRequest) SetInstanceName(v string) *ModifySandboxTemplateRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifySandboxTemplateRequest) SetRegionId(v string) *ModifySandboxTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySandboxTemplateRequest) SetReplicas(v int64) *ModifySandboxTemplateRequest {
	s.Replicas = &v
	return s
}

func (s *ModifySandboxTemplateRequest) SetTemplateId(v string) *ModifySandboxTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifySandboxTemplateRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSandboxTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultCpu(v string) *CreateSandboxTemplateShrinkRequest
	GetDefaultCpu() *string
	SetDefaultMemory(v string) *CreateSandboxTemplateShrinkRequest
	GetDefaultMemory() *string
	SetDescription(v string) *CreateSandboxTemplateShrinkRequest
	GetDescription() *string
	SetImage(v string) *CreateSandboxTemplateShrinkRequest
	GetImage() *string
	SetInstanceName(v string) *CreateSandboxTemplateShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *CreateSandboxTemplateShrinkRequest
	GetRegionId() *string
	SetReplicas(v int64) *CreateSandboxTemplateShrinkRequest
	GetReplicas() *int64
	SetTagsShrink(v string) *CreateSandboxTemplateShrinkRequest
	GetTagsShrink() *string
	SetTemplateName(v string) *CreateSandboxTemplateShrinkRequest
	GetTemplateName() *string
}

type CreateSandboxTemplateShrinkRequest struct {
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
	Image       *string `json:"Image,omitempty" xml:"Image,omitempty"`
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
	Replicas   *int64  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the sandbox template.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-interpreter
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateSandboxTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSandboxTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSandboxTemplateShrinkRequest) GetDefaultCpu() *string {
	return s.DefaultCpu
}

func (s *CreateSandboxTemplateShrinkRequest) GetDefaultMemory() *string {
	return s.DefaultMemory
}

func (s *CreateSandboxTemplateShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSandboxTemplateShrinkRequest) GetImage() *string {
	return s.Image
}

func (s *CreateSandboxTemplateShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateSandboxTemplateShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSandboxTemplateShrinkRequest) GetReplicas() *int64 {
	return s.Replicas
}

func (s *CreateSandboxTemplateShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateSandboxTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSandboxTemplateShrinkRequest) SetDefaultCpu(v string) *CreateSandboxTemplateShrinkRequest {
	s.DefaultCpu = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) SetDefaultMemory(v string) *CreateSandboxTemplateShrinkRequest {
	s.DefaultMemory = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) SetDescription(v string) *CreateSandboxTemplateShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) SetImage(v string) *CreateSandboxTemplateShrinkRequest {
	s.Image = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) SetInstanceName(v string) *CreateSandboxTemplateShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) SetRegionId(v string) *CreateSandboxTemplateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) SetReplicas(v int64) *CreateSandboxTemplateShrinkRequest {
	s.Replicas = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) SetTagsShrink(v string) *CreateSandboxTemplateShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) SetTemplateName(v string) *CreateSandboxTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSandboxTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}

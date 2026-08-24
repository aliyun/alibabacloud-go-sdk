// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySandboxTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultCpu(v string) *ModifySandboxTemplateShrinkRequest
	GetDefaultCpu() *string
	SetDefaultMemory(v string) *ModifySandboxTemplateShrinkRequest
	GetDefaultMemory() *string
	SetImage(v string) *ModifySandboxTemplateShrinkRequest
	GetImage() *string
	SetInstanceName(v string) *ModifySandboxTemplateShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifySandboxTemplateShrinkRequest
	GetRegionId() *string
	SetReplicas(v int64) *ModifySandboxTemplateShrinkRequest
	GetReplicas() *int64
	SetTagsShrink(v string) *ModifySandboxTemplateShrinkRequest
	GetTagsShrink() *string
	SetTemplateId(v string) *ModifySandboxTemplateShrinkRequest
	GetTemplateId() *string
}

type ModifySandboxTemplateShrinkRequest struct {
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
	Image         *string `json:"Image,omitempty" xml:"Image,omitempty"`
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
	Replicas   *int64  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The sandbox template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop-xxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifySandboxTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySandboxTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifySandboxTemplateShrinkRequest) GetDefaultCpu() *string {
	return s.DefaultCpu
}

func (s *ModifySandboxTemplateShrinkRequest) GetDefaultMemory() *string {
	return s.DefaultMemory
}

func (s *ModifySandboxTemplateShrinkRequest) GetImage() *string {
	return s.Image
}

func (s *ModifySandboxTemplateShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifySandboxTemplateShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySandboxTemplateShrinkRequest) GetReplicas() *int64 {
	return s.Replicas
}

func (s *ModifySandboxTemplateShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ModifySandboxTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifySandboxTemplateShrinkRequest) SetDefaultCpu(v string) *ModifySandboxTemplateShrinkRequest {
	s.DefaultCpu = &v
	return s
}

func (s *ModifySandboxTemplateShrinkRequest) SetDefaultMemory(v string) *ModifySandboxTemplateShrinkRequest {
	s.DefaultMemory = &v
	return s
}

func (s *ModifySandboxTemplateShrinkRequest) SetImage(v string) *ModifySandboxTemplateShrinkRequest {
	s.Image = &v
	return s
}

func (s *ModifySandboxTemplateShrinkRequest) SetInstanceName(v string) *ModifySandboxTemplateShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifySandboxTemplateShrinkRequest) SetRegionId(v string) *ModifySandboxTemplateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySandboxTemplateShrinkRequest) SetReplicas(v int64) *ModifySandboxTemplateShrinkRequest {
	s.Replicas = &v
	return s
}

func (s *ModifySandboxTemplateShrinkRequest) SetTagsShrink(v string) *ModifySandboxTemplateShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ModifySandboxTemplateShrinkRequest) SetTemplateId(v string) *ModifySandboxTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifySandboxTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}

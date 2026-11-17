// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEditQueueSpec interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *BatchEditQueueSpec
	GetDescription() *string
	SetEnvironments(v []*string) *BatchEditQueueSpec
	GetEnvironments() []*string
	SetGpuSpec(v []*string) *BatchEditQueueSpec
	GetGpuSpec() []*string
	SetResourceSpec(v *ResourceSpec) *BatchEditQueueSpec
	GetResourceSpec() *ResourceSpec
	SetWorkspaceQueueName(v string) *BatchEditQueueSpec
	GetWorkspaceQueueName() *string
}

type BatchEditQueueSpec struct {
	// The description of the queue. If this parameter is not specified, the description is not modified.
	//
	// example:
	//
	// Development and debugging queue
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The runtime environments. If this parameter is not specified, the runtime environments are not modified.
	Environments []*string `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// The list of GPU instance types. If this parameter is not specified, the GPU specifications are not modified.
	GpuSpec []*string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty" type:"Repeated"`
	// The resource specifications, which represent the absolute target value. If this parameter is not specified, only the description, environments, or GPU specifications are modified, and the quota is not recalculated.
	ResourceSpec *ResourceSpec `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// dev_queue
	WorkspaceQueueName *string `json:"workspaceQueueName,omitempty" xml:"workspaceQueueName,omitempty"`
}

func (s BatchEditQueueSpec) String() string {
	return dara.Prettify(s)
}

func (s BatchEditQueueSpec) GoString() string {
	return s.String()
}

func (s *BatchEditQueueSpec) GetDescription() *string {
	return s.Description
}

func (s *BatchEditQueueSpec) GetEnvironments() []*string {
	return s.Environments
}

func (s *BatchEditQueueSpec) GetGpuSpec() []*string {
	return s.GpuSpec
}

func (s *BatchEditQueueSpec) GetResourceSpec() *ResourceSpec {
	return s.ResourceSpec
}

func (s *BatchEditQueueSpec) GetWorkspaceQueueName() *string {
	return s.WorkspaceQueueName
}

func (s *BatchEditQueueSpec) SetDescription(v string) *BatchEditQueueSpec {
	s.Description = &v
	return s
}

func (s *BatchEditQueueSpec) SetEnvironments(v []*string) *BatchEditQueueSpec {
	s.Environments = v
	return s
}

func (s *BatchEditQueueSpec) SetGpuSpec(v []*string) *BatchEditQueueSpec {
	s.GpuSpec = v
	return s
}

func (s *BatchEditQueueSpec) SetResourceSpec(v *ResourceSpec) *BatchEditQueueSpec {
	s.ResourceSpec = v
	return s
}

func (s *BatchEditQueueSpec) SetWorkspaceQueueName(v string) *BatchEditQueueSpec {
	s.WorkspaceQueueName = &v
	return s
}

func (s *BatchEditQueueSpec) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

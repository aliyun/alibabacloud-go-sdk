// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocateStrategy(v string) *CreateQuotaRequest
	GetAllocateStrategy() *string
	SetClusterSpec(v *ClusterSpec) *CreateQuotaRequest
	GetClusterSpec() *ClusterSpec
	SetDescription(v string) *CreateQuotaRequest
	GetDescription() *string
	SetLabels(v []*Label) *CreateQuotaRequest
	GetLabels() []*Label
	SetMin(v *ResourceSpec) *CreateQuotaRequest
	GetMin() *ResourceSpec
	SetParentQuotaId(v string) *CreateQuotaRequest
	GetParentQuotaId() *string
	SetQueueStrategy(v string) *CreateQuotaRequest
	GetQueueStrategy() *string
	SetQuotaConfig(v *QuotaConfig) *CreateQuotaRequest
	GetQuotaConfig() *QuotaConfig
	SetQuotaName(v string) *CreateQuotaRequest
	GetQuotaName() *string
	SetResourceGroupIds(v []*string) *CreateQuotaRequest
	GetResourceGroupIds() []*string
	SetResourceType(v string) *CreateQuotaRequest
	GetResourceType() *string
}

type CreateQuotaRequest struct {
	// The allocation strategy for the quota. Only `ByNodeSpecs` is supported.
	//
	// example:
	//
	// ByNodeSpecs
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// The native cluster specification for the quota.
	ClusterSpec *ClusterSpec `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// this is a test quota
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tags for the quota.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The minimum resources for the quota. You can define this in one of the following ways:
	//
	// - `ResourceAmount`: Specifies the CPU, memory, and GPU details.
	//
	// - `NodeSpecs`: Specifies the node specification and the number of nodes.
	//
	// Constraints:
	//
	// - If this quota allocates resources from a dedicated resource group, you must use the `NodeSpecs` method.
	//
	// - If this quota allocates resources from a parent quota, both methods are allowed. However, all its child quotas must use the same method.
	//
	// - All GPU specifications within the quota must have the same GPU type.
	//
	// - For quotas with the resource type set to ECS or Lingjun, only the `NodeSpecs` method can be used.
	Min *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	// The ID of the parent quota.
	//
	// - If you do not specify this parameter, a root quota is created. Resources are allocated from a dedicated resource group.
	//
	// - If you specify this parameter, a child quota is created. Resources are allocated from the nodes that are bound to the root quota.
	//
	// example:
	//
	// quota1ci8g793pgm
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// The queuing strategy for the quota. Four strategies are supported:
	//
	// - `PaiStrategyIntelligent`: The intelligent strategy.
	//
	// - `PaiStrategyBalance`: The balance strategy.
	//
	// - `PaiStrategyRoundRobin`: The round-robin strategy.
	//
	// - `PaiStrategyStrictFIFO`: The FIFO strategy.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// Constraints for the `QuotaConfig` parameter:
	//
	// - This parameter is ignored if the resource type is ECS or Lingjun.
	//
	// - If the resource type is ACS, the specified VPC and ACS configurations are applied.
	QuotaConfig *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// test-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The IDs of the dedicated resource groups. The following constraints apply:
	//
	// - Only a root quota, for which `ParentQuotaId` is empty, can allocate nodes from a resource group.
	//
	// - The VPC configurations of the specified resource groups must be the same.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The resource type of the quota. Valid values: Lingjun, ECS, and ACS. Default value: ECS.
	//
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaRequest) GetAllocateStrategy() *string {
	return s.AllocateStrategy
}

func (s *CreateQuotaRequest) GetClusterSpec() *ClusterSpec {
	return s.ClusterSpec
}

func (s *CreateQuotaRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateQuotaRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateQuotaRequest) GetMin() *ResourceSpec {
	return s.Min
}

func (s *CreateQuotaRequest) GetParentQuotaId() *string {
	return s.ParentQuotaId
}

func (s *CreateQuotaRequest) GetQueueStrategy() *string {
	return s.QueueStrategy
}

func (s *CreateQuotaRequest) GetQuotaConfig() *QuotaConfig {
	return s.QuotaConfig
}

func (s *CreateQuotaRequest) GetQuotaName() *string {
	return s.QuotaName
}

func (s *CreateQuotaRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *CreateQuotaRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateQuotaRequest) SetAllocateStrategy(v string) *CreateQuotaRequest {
	s.AllocateStrategy = &v
	return s
}

func (s *CreateQuotaRequest) SetClusterSpec(v *ClusterSpec) *CreateQuotaRequest {
	s.ClusterSpec = v
	return s
}

func (s *CreateQuotaRequest) SetDescription(v string) *CreateQuotaRequest {
	s.Description = &v
	return s
}

func (s *CreateQuotaRequest) SetLabels(v []*Label) *CreateQuotaRequest {
	s.Labels = v
	return s
}

func (s *CreateQuotaRequest) SetMin(v *ResourceSpec) *CreateQuotaRequest {
	s.Min = v
	return s
}

func (s *CreateQuotaRequest) SetParentQuotaId(v string) *CreateQuotaRequest {
	s.ParentQuotaId = &v
	return s
}

func (s *CreateQuotaRequest) SetQueueStrategy(v string) *CreateQuotaRequest {
	s.QueueStrategy = &v
	return s
}

func (s *CreateQuotaRequest) SetQuotaConfig(v *QuotaConfig) *CreateQuotaRequest {
	s.QuotaConfig = v
	return s
}

func (s *CreateQuotaRequest) SetQuotaName(v string) *CreateQuotaRequest {
	s.QuotaName = &v
	return s
}

func (s *CreateQuotaRequest) SetResourceGroupIds(v []*string) *CreateQuotaRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *CreateQuotaRequest) SetResourceType(v string) *CreateQuotaRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateQuotaRequest) Validate() error {
	if s.ClusterSpec != nil {
		if err := s.ClusterSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Min != nil {
		if err := s.Min.Validate(); err != nil {
			return err
		}
	}
	if s.QuotaConfig != nil {
		if err := s.QuotaConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

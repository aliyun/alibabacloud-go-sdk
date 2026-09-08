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
	// The quota allocation strategy. Currently, only ByNodeSpecs is supported.
	//
	// example:
	//
	// ByNodeSpecs
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// The specifications of the native cluster for the resource quota.
	ClusterSpec *ClusterSpec `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The quota description.
	//
	// example:
	//
	// this is a test quota
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The quota labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The minimum quota configuration. Valid options:
	//
	// - ResourceAmount: specifies CPU, memory, or GPU details.
	//
	// - NodeSpecs: specifies the instance type and quantity.
	//
	// Constraints:
	//
	// - If the quota allocates resources from a dedicated resource group, only the NodeSpecs strategy is allowed.
	//
	// - If the quota allocates resources from a parent quota, both strategies are allowed, but all child quotas must use the same strategy.
	//
	// - All GPU specifications within a quota must use the same GPU type.
	//
	// - Resource quotas with the ECS or Lingjun resource type can only use the NodeSpecs strategy.
	Min *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	// The parent QuotaId:
	//
	// - If ParentQuotaId is empty, a root quota is created and machines are allocated from the dedicated resource group.
	//
	// - If ParentQuotaId is not empty, a child quota is created and resources are allocated from the nodes bound to the root quota.
	//
	// example:
	//
	// quota1ci8g793pgm
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// Four queuing policies are supported for quotas.
	//
	// - PaiStrategyIntelligent: intelligent policies.
	//
	// - PaiStrategyBalance: balanced policy.
	//
	// - PaiStrategyRoundRobin: resource-priority policy.
	//
	// - PaiStrategyStrictFIFO: FIFO policy.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// QuotaConfig configuration constraints:
	//
	// - This configuration does not take effect when the ECS or Lingjun resource type is used.
	//
	// - When the ACS resource type is used, the user VPC information and ACS configuration take effect.
	QuotaConfig *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	// The quota name.
	//
	// example:
	//
	// test-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The list of dedicated resource groups. Constraints:
	//
	// - Only root quotas (where ParentQuotaId is empty) can allocate machines from resource groups.
	//
	// - The VPC configurations in the specified resource groups must be consistent.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The quota resource type (Lingjun/ECS/ACS). Default value: ECS.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuota interface {
	dara.Model
	String() string
	GoString() string
	SetAllocateStrategy(v string) *Quota
	GetAllocateStrategy() *string
	SetCreatorId(v string) *Quota
	GetCreatorId() *string
	SetDescription(v string) *Quota
	GetDescription() *string
	SetGPUType(v string) *Quota
	GetGPUType() *string
	SetGmtCreatedTime(v string) *Quota
	GetGmtCreatedTime() *string
	SetGmtModifiedTime(v string) *Quota
	GetGmtModifiedTime() *string
	SetHyperZones(v []*string) *Quota
	GetHyperZones() []*string
	SetLabels(v []*Label) *Quota
	GetLabels() []*Label
	SetLatestOperationId(v string) *Quota
	GetLatestOperationId() *string
	SetMin(v *ResourceSpec) *Quota
	GetMin() *ResourceSpec
	SetParentQuotaId(v string) *Quota
	GetParentQuotaId() *string
	SetQueueStrategy(v string) *Quota
	GetQueueStrategy() *string
	SetQuotaCluster(v *QuotaCluster) *Quota
	GetQuotaCluster() *QuotaCluster
	SetQuotaConfig(v *QuotaConfig) *Quota
	GetQuotaConfig() *QuotaConfig
	SetQuotaDetails(v *QuotaDetails) *Quota
	GetQuotaDetails() *QuotaDetails
	SetQuotaId(v string) *Quota
	GetQuotaId() *string
	SetQuotaName(v string) *Quota
	GetQuotaName() *string
	SetReasonCode(v string) *Quota
	GetReasonCode() *string
	SetReasonMessage(v string) *Quota
	GetReasonMessage() *string
	SetResourceGroupIds(v []*string) *Quota
	GetResourceGroupIds() []*string
	SetResourceType(v string) *Quota
	GetResourceType() *string
	SetStatus(v string) *Quota
	GetStatus() *string
	SetSubQuotas(v []*QuotaIdName) *Quota
	GetSubQuotas() []*QuotaIdName
	SetVersion(v string) *Quota
	GetVersion() *string
	SetWorkspaces(v []*WorkspaceIdName) *Quota
	GetWorkspaces() []*WorkspaceIdName
}

type Quota struct {
	// The resource allocation strategy. Currently, only `ByNodeSpec` is supported.
	//
	// example:
	//
	// ByNodeSpec
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// The ID of the user who created the quota.
	//
	// example:
	//
	// 1884692****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The quota description.
	//
	// example:
	//
	// this is a test quota
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GPUType     *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The time when the quota was created.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// The time when the quota was last modified.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModifiedTime *string   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HyperZones      []*string `json:"HyperZones,omitempty" xml:"HyperZones,omitempty" type:"Repeated"`
	// The quota labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The ID of the most recent operation on the quota.
	//
	// example:
	//
	// operation****
	LatestOperationId *string `json:"LatestOperationId,omitempty" xml:"LatestOperationId,omitempty"`
	// The guaranteed resources.
	Min *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	// The parent quota ID.
	//
	// example:
	//
	// ""
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// The queuing strategy. Valid values:
	//
	// - `PaiStrategyIntelligent`: intelligent strategy.
	//
	// - `PaiStrategyBalance`: balanced strategy.
	//
	// - `PaiStrategyRoundRobin`: round-robin strategy.
	//
	// - `PaiStrategyStrictFIFO`: FIFO strategy.
	//
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string       `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	QuotaCluster  *QuotaCluster `json:"QuotaCluster,omitempty" xml:"QuotaCluster,omitempty"`
	// The quota configuration.
	QuotaConfig *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	// The quota details.
	QuotaDetails *QuotaDetails `json:"QuotaDetails,omitempty" xml:"QuotaDetails,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// quota12345
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The quota name.
	//
	// example:
	//
	// dlc-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The error code.
	//
	// example:
	//
	// “”
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// “”
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// A list of resource group IDs.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// - `Lingjun`
	//
	// - `ECS` (default)
	//
	// - `ACS`
	//
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The quota status.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of sub-quotas.
	SubQuotas []*QuotaIdName `json:"SubQuotas,omitempty" xml:"SubQuotas,omitempty" type:"Repeated"`
	// The version.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The workspaces that are associated with the quota.
	Workspaces []*WorkspaceIdName `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s Quota) String() string {
	return dara.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) GetAllocateStrategy() *string {
	return s.AllocateStrategy
}

func (s *Quota) GetCreatorId() *string {
	return s.CreatorId
}

func (s *Quota) GetDescription() *string {
	return s.Description
}

func (s *Quota) GetGPUType() *string {
	return s.GPUType
}

func (s *Quota) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *Quota) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Quota) GetHyperZones() []*string {
	return s.HyperZones
}

func (s *Quota) GetLabels() []*Label {
	return s.Labels
}

func (s *Quota) GetLatestOperationId() *string {
	return s.LatestOperationId
}

func (s *Quota) GetMin() *ResourceSpec {
	return s.Min
}

func (s *Quota) GetParentQuotaId() *string {
	return s.ParentQuotaId
}

func (s *Quota) GetQueueStrategy() *string {
	return s.QueueStrategy
}

func (s *Quota) GetQuotaCluster() *QuotaCluster {
	return s.QuotaCluster
}

func (s *Quota) GetQuotaConfig() *QuotaConfig {
	return s.QuotaConfig
}

func (s *Quota) GetQuotaDetails() *QuotaDetails {
	return s.QuotaDetails
}

func (s *Quota) GetQuotaId() *string {
	return s.QuotaId
}

func (s *Quota) GetQuotaName() *string {
	return s.QuotaName
}

func (s *Quota) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *Quota) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *Quota) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *Quota) GetResourceType() *string {
	return s.ResourceType
}

func (s *Quota) GetStatus() *string {
	return s.Status
}

func (s *Quota) GetSubQuotas() []*QuotaIdName {
	return s.SubQuotas
}

func (s *Quota) GetVersion() *string {
	return s.Version
}

func (s *Quota) GetWorkspaces() []*WorkspaceIdName {
	return s.Workspaces
}

func (s *Quota) SetAllocateStrategy(v string) *Quota {
	s.AllocateStrategy = &v
	return s
}

func (s *Quota) SetCreatorId(v string) *Quota {
	s.CreatorId = &v
	return s
}

func (s *Quota) SetDescription(v string) *Quota {
	s.Description = &v
	return s
}

func (s *Quota) SetGPUType(v string) *Quota {
	s.GPUType = &v
	return s
}

func (s *Quota) SetGmtCreatedTime(v string) *Quota {
	s.GmtCreatedTime = &v
	return s
}

func (s *Quota) SetGmtModifiedTime(v string) *Quota {
	s.GmtModifiedTime = &v
	return s
}

func (s *Quota) SetHyperZones(v []*string) *Quota {
	s.HyperZones = v
	return s
}

func (s *Quota) SetLabels(v []*Label) *Quota {
	s.Labels = v
	return s
}

func (s *Quota) SetLatestOperationId(v string) *Quota {
	s.LatestOperationId = &v
	return s
}

func (s *Quota) SetMin(v *ResourceSpec) *Quota {
	s.Min = v
	return s
}

func (s *Quota) SetParentQuotaId(v string) *Quota {
	s.ParentQuotaId = &v
	return s
}

func (s *Quota) SetQueueStrategy(v string) *Quota {
	s.QueueStrategy = &v
	return s
}

func (s *Quota) SetQuotaCluster(v *QuotaCluster) *Quota {
	s.QuotaCluster = v
	return s
}

func (s *Quota) SetQuotaConfig(v *QuotaConfig) *Quota {
	s.QuotaConfig = v
	return s
}

func (s *Quota) SetQuotaDetails(v *QuotaDetails) *Quota {
	s.QuotaDetails = v
	return s
}

func (s *Quota) SetQuotaId(v string) *Quota {
	s.QuotaId = &v
	return s
}

func (s *Quota) SetQuotaName(v string) *Quota {
	s.QuotaName = &v
	return s
}

func (s *Quota) SetReasonCode(v string) *Quota {
	s.ReasonCode = &v
	return s
}

func (s *Quota) SetReasonMessage(v string) *Quota {
	s.ReasonMessage = &v
	return s
}

func (s *Quota) SetResourceGroupIds(v []*string) *Quota {
	s.ResourceGroupIds = v
	return s
}

func (s *Quota) SetResourceType(v string) *Quota {
	s.ResourceType = &v
	return s
}

func (s *Quota) SetStatus(v string) *Quota {
	s.Status = &v
	return s
}

func (s *Quota) SetSubQuotas(v []*QuotaIdName) *Quota {
	s.SubQuotas = v
	return s
}

func (s *Quota) SetVersion(v string) *Quota {
	s.Version = &v
	return s
}

func (s *Quota) SetWorkspaces(v []*WorkspaceIdName) *Quota {
	s.Workspaces = v
	return s
}

func (s *Quota) Validate() error {
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
	if s.QuotaCluster != nil {
		if err := s.QuotaCluster.Validate(); err != nil {
			return err
		}
	}
	if s.QuotaConfig != nil {
		if err := s.QuotaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.QuotaDetails != nil {
		if err := s.QuotaDetails.Validate(); err != nil {
			return err
		}
	}
	if s.SubQuotas != nil {
		for _, item := range s.SubQuotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Workspaces != nil {
		for _, item := range s.Workspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

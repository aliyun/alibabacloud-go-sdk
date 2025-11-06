// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocateStrategy(v string) *GetQuotaResponseBody
	GetAllocateStrategy() *string
	SetCreatorId(v string) *GetQuotaResponseBody
	GetCreatorId() *string
	SetDescription(v string) *GetQuotaResponseBody
	GetDescription() *string
	SetGmtCreatedTime(v string) *GetQuotaResponseBody
	GetGmtCreatedTime() *string
	SetGmtModifiedTime(v string) *GetQuotaResponseBody
	GetGmtModifiedTime() *string
	SetHyperZones(v []*string) *GetQuotaResponseBody
	GetHyperZones() []*string
	SetLabels(v []*Label) *GetQuotaResponseBody
	GetLabels() []*Label
	SetLatestOperationId(v string) *GetQuotaResponseBody
	GetLatestOperationId() *string
	SetMin(v *ResourceSpec) *GetQuotaResponseBody
	GetMin() *ResourceSpec
	SetParentQuotaId(v string) *GetQuotaResponseBody
	GetParentQuotaId() *string
	SetQueueStrategy(v string) *GetQuotaResponseBody
	GetQueueStrategy() *string
	SetQuotaCluster(v *QuotaCluster) *GetQuotaResponseBody
	GetQuotaCluster() *QuotaCluster
	SetQuotaConfig(v *QuotaConfig) *GetQuotaResponseBody
	GetQuotaConfig() *QuotaConfig
	SetQuotaDetails(v *QuotaDetails) *GetQuotaResponseBody
	GetQuotaDetails() *QuotaDetails
	SetQuotaId(v string) *GetQuotaResponseBody
	GetQuotaId() *string
	SetQuotaName(v string) *GetQuotaResponseBody
	GetQuotaName() *string
	SetReasonCode(v string) *GetQuotaResponseBody
	GetReasonCode() *string
	SetReasonMessage(v string) *GetQuotaResponseBody
	GetReasonMessage() *string
	SetRequestId(v string) *GetQuotaResponseBody
	GetRequestId() *string
	SetResourceGroupIds(v []*string) *GetQuotaResponseBody
	GetResourceGroupIds() []*string
	SetResourceType(v string) *GetQuotaResponseBody
	GetResourceType() *string
	SetStatus(v string) *GetQuotaResponseBody
	GetStatus() *string
	SetSubQuotas(v []*QuotaIdName) *GetQuotaResponseBody
	GetSubQuotas() []*QuotaIdName
	SetVersion(v string) *GetQuotaResponseBody
	GetVersion() *string
	SetWorkspaces(v []*WorkspaceIdName) *GetQuotaResponseBody
	GetWorkspaces() []*WorkspaceIdName
}

type GetQuotaResponseBody struct {
	// example:
	//
	// ByNodeSpec
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// example:
	//
	// 18846926616
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// this is a test quota
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModifiedTime *string   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HyperZones      []*string `json:"HyperZones,omitempty" xml:"HyperZones,omitempty" type:"Repeated"`
	Labels          []*Label  `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// operation1234
	LatestOperationId *string       `json:"LatestOperationId,omitempty" xml:"LatestOperationId,omitempty"`
	Min               *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// quota1ci8g793pgm
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string       `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	QuotaCluster  *QuotaCluster `json:"QuotaCluster,omitempty" xml:"QuotaCluster,omitempty"`
	QuotaConfig   *QuotaConfig  `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	QuotaDetails  *QuotaDetails `json:"QuotaDetails,omitempty" xml:"QuotaDetails,omitempty"`
	// Quota Id
	//
	// example:
	//
	// quotajradxh43rgb
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// test-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// “”
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// “”
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Ready
	Status    *string        `json:"Status,omitempty" xml:"Status,omitempty"`
	SubQuotas []*QuotaIdName `json:"SubQuotas,omitempty" xml:"SubQuotas,omitempty" type:"Repeated"`
	// example:
	//
	// 1.0
	Version    *string            `json:"Version,omitempty" xml:"Version,omitempty"`
	Workspaces []*WorkspaceIdName `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s GetQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBody) GetAllocateStrategy() *string {
	return s.AllocateStrategy
}

func (s *GetQuotaResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetQuotaResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetQuotaResponseBody) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *GetQuotaResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetQuotaResponseBody) GetHyperZones() []*string {
	return s.HyperZones
}

func (s *GetQuotaResponseBody) GetLabels() []*Label {
	return s.Labels
}

func (s *GetQuotaResponseBody) GetLatestOperationId() *string {
	return s.LatestOperationId
}

func (s *GetQuotaResponseBody) GetMin() *ResourceSpec {
	return s.Min
}

func (s *GetQuotaResponseBody) GetParentQuotaId() *string {
	return s.ParentQuotaId
}

func (s *GetQuotaResponseBody) GetQueueStrategy() *string {
	return s.QueueStrategy
}

func (s *GetQuotaResponseBody) GetQuotaCluster() *QuotaCluster {
	return s.QuotaCluster
}

func (s *GetQuotaResponseBody) GetQuotaConfig() *QuotaConfig {
	return s.QuotaConfig
}

func (s *GetQuotaResponseBody) GetQuotaDetails() *QuotaDetails {
	return s.QuotaDetails
}

func (s *GetQuotaResponseBody) GetQuotaId() *string {
	return s.QuotaId
}

func (s *GetQuotaResponseBody) GetQuotaName() *string {
	return s.QuotaName
}

func (s *GetQuotaResponseBody) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetQuotaResponseBody) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaResponseBody) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *GetQuotaResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetQuotaResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaResponseBody) GetSubQuotas() []*QuotaIdName {
	return s.SubQuotas
}

func (s *GetQuotaResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetQuotaResponseBody) GetWorkspaces() []*WorkspaceIdName {
	return s.Workspaces
}

func (s *GetQuotaResponseBody) SetAllocateStrategy(v string) *GetQuotaResponseBody {
	s.AllocateStrategy = &v
	return s
}

func (s *GetQuotaResponseBody) SetCreatorId(v string) *GetQuotaResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBody) SetDescription(v string) *GetQuotaResponseBody {
	s.Description = &v
	return s
}

func (s *GetQuotaResponseBody) SetGmtCreatedTime(v string) *GetQuotaResponseBody {
	s.GmtCreatedTime = &v
	return s
}

func (s *GetQuotaResponseBody) SetGmtModifiedTime(v string) *GetQuotaResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetQuotaResponseBody) SetHyperZones(v []*string) *GetQuotaResponseBody {
	s.HyperZones = v
	return s
}

func (s *GetQuotaResponseBody) SetLabels(v []*Label) *GetQuotaResponseBody {
	s.Labels = v
	return s
}

func (s *GetQuotaResponseBody) SetLatestOperationId(v string) *GetQuotaResponseBody {
	s.LatestOperationId = &v
	return s
}

func (s *GetQuotaResponseBody) SetMin(v *ResourceSpec) *GetQuotaResponseBody {
	s.Min = v
	return s
}

func (s *GetQuotaResponseBody) SetParentQuotaId(v string) *GetQuotaResponseBody {
	s.ParentQuotaId = &v
	return s
}

func (s *GetQuotaResponseBody) SetQueueStrategy(v string) *GetQuotaResponseBody {
	s.QueueStrategy = &v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaCluster(v *QuotaCluster) *GetQuotaResponseBody {
	s.QuotaCluster = v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaConfig(v *QuotaConfig) *GetQuotaResponseBody {
	s.QuotaConfig = v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaDetails(v *QuotaDetails) *GetQuotaResponseBody {
	s.QuotaDetails = v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaId(v string) *GetQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaName(v string) *GetQuotaResponseBody {
	s.QuotaName = &v
	return s
}

func (s *GetQuotaResponseBody) SetReasonCode(v string) *GetQuotaResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetQuotaResponseBody) SetReasonMessage(v string) *GetQuotaResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetQuotaResponseBody) SetRequestId(v string) *GetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaResponseBody) SetResourceGroupIds(v []*string) *GetQuotaResponseBody {
	s.ResourceGroupIds = v
	return s
}

func (s *GetQuotaResponseBody) SetResourceType(v string) *GetQuotaResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetQuotaResponseBody) SetStatus(v string) *GetQuotaResponseBody {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBody) SetSubQuotas(v []*QuotaIdName) *GetQuotaResponseBody {
	s.SubQuotas = v
	return s
}

func (s *GetQuotaResponseBody) SetVersion(v string) *GetQuotaResponseBody {
	s.Version = &v
	return s
}

func (s *GetQuotaResponseBody) SetWorkspaces(v []*WorkspaceIdName) *GetQuotaResponseBody {
	s.Workspaces = v
	return s
}

func (s *GetQuotaResponseBody) Validate() error {
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

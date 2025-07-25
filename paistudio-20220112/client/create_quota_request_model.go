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
	// example:
	//
	// ByNodeSpecs
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// example:
	//
	// this is a test quota
	Description *string       `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      []*Label      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Min         *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// quota1ci8g793pgm
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string      `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	QuotaConfig   *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	// example:
	//
	// test-quota
	QuotaName        *string   `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreAggregateEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *IgnoreAggregateEvaluationResultsRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *IgnoreAggregateEvaluationResultsRequest
	GetConfigRuleId() *string
	SetIgnoreDate(v string) *IgnoreAggregateEvaluationResultsRequest
	GetIgnoreDate() *string
	SetReason(v string) *IgnoreAggregateEvaluationResultsRequest
	GetReason() *string
	SetResources(v []*IgnoreAggregateEvaluationResultsRequestResources) *IgnoreAggregateEvaluationResultsRequest
	GetResources() []*IgnoreAggregateEvaluationResultsRequestResources
}

type IgnoreAggregateEvaluationResultsRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-5b6c626622af008f****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the rule.
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7e72626622af0051****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The date from which the system automatically re-evaluates the ignored incompliant resources.
	//
	// >  If you leave this parameter empty, the system does not automatically re-evaluate the ignored incompliant resources. You must manually re-evaluate the ignored incompliant resources.
	//
	// example:
	//
	// 2022-06-01
	IgnoreDate *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	// The reason why you ignore the resource.
	//
	// example:
	//
	// The reason why you ignore the resource.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The resources to be ignored.
	//
	// This parameter is required.
	Resources []*IgnoreAggregateEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s IgnoreAggregateEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetIgnoreDate() *string {
	return s.IgnoreDate
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetReason() *string {
	return s.Reason
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetResources() []*IgnoreAggregateEvaluationResultsRequestResources {
	return s.Resources
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetAggregatorId(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetConfigRuleId(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetIgnoreDate(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetReason(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetResources(v []*IgnoreAggregateEvaluationResultsRequestResources) *IgnoreAggregateEvaluationResultsRequest {
	s.Resources = v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IgnoreAggregateEvaluationResultsRequestResources struct {
	// The ID of the region in which the resource resides.
	//
	// For more information about how to obtain the ID of a region, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the Alibaba Cloud account to which the resources belong.
	//
	// >  You must specify the ID of the current management account or a member account in the account group of the management account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120886317861****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The ID of the resource.
	//
	// For more information about how to query the ID of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-hp3a3b4ztyfm2plgm****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// For more information about how to query the type of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::SLB::LoadBalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s IgnoreAggregateEvaluationResultsRequestResources) String() string {
	return dara.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) GetRegion() *string {
	return s.Region
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetRegion(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceAccountId(v int64) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceId(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceType(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) Validate() error {
	return dara.Validate(s)
}
